package sds_application

import (
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/fields"

	"io/ioutil"
	"os"

	"github.com/golang/glog"
	"github.com/juju/loggo"
	api "github.com/samsung-cnct/cluster-manager-api/pkg/apis/cma/v1alpha1"
	"github.com/samsung-cnct/cluster-manager-api/pkg/generated/cma/client/clientset/versioned"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/ccutil"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/cma"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/helmutil"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/k8sutil"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/workqueue"
)

var (
	logger loggo.Logger
)

type SDSApplicationController struct {
	indexer  cache.Indexer
	queue    workqueue.RateLimitingInterface
	informer cache.Controller

	client *versioned.Clientset
}

func NewSDSApplicationController(config *rest.Config) (output *SDSApplicationController) {
	if config == nil {
		config = k8sutil.DefaultConfig
	}
	client := versioned.NewForConfigOrDie(config)

	// create sdsapplication list/watcher
	sdsApplicationListWatcher := cache.NewListWatchFromClient(
		client.CmaV1alpha1().RESTClient(),
		api.SDSApplicationResourcePlural,
		"default",
		fields.Everything())

	// create the workqueue
	queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())

	// Bind the workqueue to a cache with the help of an informer. This way we make sure that
	// whenever the cache is updated, the SDSCluster key is added to the workqueue.
	// Note that when we finally process the item from the workqueue, we might see a newer version
	// of the SDSPackageManager than the version which was responsible for triggering the update.
	indexer, informer := cache.NewIndexerInformer(sdsApplicationListWatcher, &api.SDSApplication{}, 0, cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(obj)
			if err == nil {
				queue.Add(key)
			}
		},
		UpdateFunc: func(old interface{}, new interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(new)
			if err == nil {
				queue.Add(key)
			}
		},
		DeleteFunc: func(obj interface{}) {
			// IndexerInformer uses a delta queue, therefore for deletes we have to use this
			// key function.
			key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
			if err == nil {
				queue.Add(key)
			}
		},
	}, cache.Indexers{})

	output = &SDSApplicationController{
		informer: informer,
		indexer:  indexer,
		queue:    queue,
		client:   client,
	}
	output.SetLogger()
	return
}

func (c *SDSApplicationController) processNextItem() bool {
	// Wait until there is a new item in the working queue
	key, quit := c.queue.Get()
	if quit {
		return false
	}
	// Tell the queue that we are done with processing this key. This unblocks the key for other workers
	// This allows safe parallel processing because two SDSClusters with the same key are never processed in
	// parallel.
	defer c.queue.Done(key)

	// Invoke the method containing the business logic
	err := c.processItem(key.(string))
	// Handle the error if something went wrong during the execution of the business logic
	c.handleErr(err, key)
	return true
}

// processItem is the business logic of the controller.
func (c *SDSApplicationController) processItem(key string) error {
	obj, exists, err := c.indexer.GetByKey(key)
	if err != nil {
		glog.Errorf("Fetching object with key %s from store failed with %v", key, err)
		return err
	}

	if !exists {
		// Below we will warm up our cache with a SDSPackageManager, so that we will see a delete for one SDSPackageManager
		fmt.Printf("SDSApplication -->%s<-- does not exist anymore\n", key)
	} else {
		SDSApplication := obj.(*api.SDSApplication)
		clusterName := SDSApplication.GetName()
		fmt.Printf("SDSApplication -->%s<-- does exist (name=%s)!\n", key, clusterName)

		switch SDSApplication.Status.Phase {
		case api.ApplicationPhaseNone, api.ApplicationPhasePending:
			c.deployApplication(SDSApplication)
		case api.ApplicationPhaseInstalling:
			c.waitForApplication(SDSApplication)
		}
	}
	return nil
}

// handleErr checks if an error happened and makes sure we will retry later.
func (c *SDSApplicationController) handleErr(err error, key interface{}) {
	if err == nil {
		// Forget about the #AddRateLimited history of the key on every successful synchronization.
		// This ensures that future processing of updates for this key is not delayed because of
		// an outdated error history.
		c.queue.Forget(key)
		return
	}

	// This controller retries 5 times if something goes wrong. After that, it stops trying.
	if c.queue.NumRequeues(key) < 5 {
		glog.Infof("Error syncing SDSApplication %v: %v", key, err)

		// Re-enqueue the key rate limited. Based on the rate limiter on the
		// queue and the re-enqueue history, the key will be processed later again.
		c.queue.AddRateLimited(key)
		return
	}

	c.queue.Forget(key)
	// Report to an external entity that, even after several retries, we could not successfully process this key
	runtime.HandleError(err)
	glog.Infof("Dropping krakenCluster %q out of the queue: %v", key, err)
}

func (c *SDSApplicationController) Run(threadiness int, stopCh chan struct{}) {
	defer runtime.HandleCrash()

	// Let the workers stop when we are done
	defer c.queue.ShutDown()
	glog.Info("Starting SDSCluster controller")

	go c.informer.Run(stopCh)

	// Wait for all involved caches to be synced, before processing items from the queue is started
	if !cache.WaitForCacheSync(stopCh, c.informer.HasSynced) {
		runtime.HandleError(fmt.Errorf("Timed out waiting for caches to sync"))
		return
	}

	for i := 0; i < threadiness; i++ {
		go wait.Until(c.runWorker, time.Second, stopCh)
	}

	<-stopCh
	glog.Info("Stopping SDSCluster controller")
}

func (c *SDSApplicationController) runWorker() {
	for c.processNextItem() {
	}
}

func (c *SDSApplicationController) deployApplication(application *api.SDSApplication) (bool, error) {
	config, err := c.getRestConfigForRemoteCluster(application.Spec.PackageManager.Name, application.Namespace, nil)
	if err != nil {
		return false, err
	}
	packageManager, err := cma.GetSDSPackageManager(application.Spec.PackageManager.Name, "default", nil)

	k8sutil.CreateJob(helmutil.GenerateHelmInstallJob(application.Spec), packageManager.Spec.Namespace, config)

	application.Status.Phase = api.ApplicationPhaseInstalling
	_, err = c.client.CmaV1alpha1().SDSApplications(application.Namespace).Update(application)
	if err == nil {
		logger.Infof("Deployed helm install job for -->%s<--", application.Spec.Name)
	} else {
		logger.Infof("Could not update the status error was %s", err)
	}

	return true, nil
}

func retrieveClusterRestConfig(name string, namespace string, config *rest.Config) (*rest.Config, error) {
	cluster, err := ccutil.GetKrakenCluster(name, namespace, config)
	if err != nil {
		return nil, err
	}
	// Let's create a tempfile and line it up for removal
	file, err := ioutil.TempFile(os.TempDir(), "kraken-kubeconfig")
	defer os.Remove(file.Name())
	file.WriteString(cluster.Status.Kubeconfig)

	clusterConfig, err := clientcmd.BuildConfigFromFlags("", file.Name())
	if os.Getenv("CLUSTERMANAGERAPI_INSECURE_TLS") == "true" {
		clusterConfig.TLSClientConfig = rest.TLSClientConfig{Insecure: true}
	}

	if err != nil {
		logger.Errorf("Could not load kubeconfig for cluster -->%s<-- in namespace -->%s<--", name, namespace)
		return nil, err
	}
	return clusterConfig, nil
}

func (c *SDSApplicationController) getRestConfigForRemoteCluster(clusterName string, namespace string, config *rest.Config) (*rest.Config, error) {
	cluster, err := ccutil.GetKrakenCluster(clusterName, namespace, config)
	if err != nil {
		glog.Errorf("Failed to retrieve KrakenCluster CR -->%s<-- in namespace -->%s<--, error was: %s", clusterName, namespace, err)
		return nil, err
	}
	if cluster.Status.Kubeconfig == "" {
		glog.Errorf("Could not install tiller yet for cluster -->%s<-- cluster is not ready, status is -->%s<--", cluster.Name, cluster.Status.State)
		return nil, err
	}

	remoteConfig, err := retrieveClusterRestConfig(clusterName, namespace, config)
	if err != nil {
		glog.Errorf("Could not install tiller yet for cluster -->%s<-- cluster is not ready, error is: %v", clusterName, err)
		return nil, err
	}

	return remoteConfig, nil
}

func (c *SDSApplicationController) SetLogger() {
	logger = util.GetModuleLogger("pkg.controllers.sds_application", loggo.INFO)
}

func (c *SDSApplicationController) waitForApplication(application *api.SDSApplication) (result bool, err error) {
	config, err := c.getRestConfigForRemoteCluster(application.Spec.PackageManager.Name, application.Namespace, nil)
	if err != nil {
		return false, err
	}

	packageManager, err := cma.GetSDSPackageManager(application.Spec.PackageManager.Name, "default", nil)
	if err != nil {
		logger.Infof("Cannot retrieve package manager for application %s", application.Spec.Name)
		return false, err
	}

	clientset, _ := kubernetes.NewForConfig(config)
	timeout := 0
	for timeout < 2000 {
		job, err := clientset.BatchV1().Jobs(packageManager.Spec.Namespace).Get("app-install-"+application.Spec.Name, v1.GetOptions{})
		if err == nil {
			if job.Status.Succeeded > 0 {
				application.Status.Phase = api.ApplicationPhaseImplemented
				application.Status.Ready = true
				_, err = c.client.CmaV1alpha1().SDSApplications(application.Namespace).Update(application)
				if err == nil {
					logger.Infof("Helm installed app -->%s<--", application.Spec.Name)
					c.updateSDSCluster(application.Spec.PackageManager.Name)
				} else {
					logger.Infof("Could not update the status error was %s", err)
				}
				return true, nil
			}
		}
		time.Sleep(5 * time.Second)
		timeout++
	}
	return false, nil
}

func (c *SDSApplicationController) updateSDSCluster(clusterName string) (result bool, err error) {
	// TODO This is dubious, but for the PoC, good enough
	sdsCluster, err := cma.GetSDSCluster(clusterName, "default", nil)
	if err != nil {
		logger.Infof("Failed to get SDSCluster for SDSApplication %s, error was: ", clusterName, err)
		return false, err
	}

	changes := false
	if sdsCluster.Status.AppsInstalled == false {
		changes = true
		sdsCluster.Status.AppsInstalled = true
	}
	switch sdsCluster.Status.Phase {
	case api.ClusterPhaseHavePackageManager, api.ClusterPhaseDeployingApplications:
		changes = true
		sdsCluster.Status.Phase = api.ClusterPhaseReady
	}

	if changes {
		_, err = cma.UpdateSDSCluster(sdsCluster, sdsCluster.Namespace, nil)
		if err != nil {
			logger.Infof("Could not update SDSCluster for KrakenCluster %s, error was: ", sdsCluster.Name, err)
		}
	}

	return true, nil
}
