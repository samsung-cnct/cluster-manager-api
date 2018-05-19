package sds_cluster

import (
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/fields"

	"github.com/golang/glog"
	"github.com/juju/loggo"
	api "github.com/samsung-cnct/cluster-manager-api/pkg/apis/cma/v1alpha1"
	"github.com/samsung-cnct/cluster-manager-api/pkg/generated/cma/client/clientset/versioned"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/ccutil"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/cma"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/k8sutil"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

var (
	logger loggo.Logger
)

type SDSClusterController struct {
	indexer  cache.Indexer
	queue    workqueue.RateLimitingInterface
	informer cache.Controller

	client *versioned.Clientset
}

func NewSDSClusterController(config *rest.Config) *SDSClusterController {
	if config == nil {
		config = k8sutil.DefaultConfig
	}
	client := versioned.NewForConfigOrDie(config)

	// create sdscluster list/watcher
	sdsClusterListWatcher := cache.NewListWatchFromClient(
		client.CmaV1alpha1().RESTClient(),
		api.SDSClusterResourcePlural,
		"default",
		fields.Everything())

	// create the workqueue
	queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())

	// Bind the workqueue to a cache with the help of an informer. This way we make sure that
	// whenever the cache is updated, the SDSCluster key is added to the workqueue.
	// Note that when we finally process the item from the workqueue, we might see a newer version
	// of the SDSCluster than the version which was responsible for triggering the update.
	indexer, informer := cache.NewIndexerInformer(sdsClusterListWatcher, &api.SDSCluster{}, 0, cache.ResourceEventHandlerFuncs{
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

	output := &SDSClusterController{
		informer: informer,
		indexer:  indexer,
		queue:    queue,
		client:   client,
	}
	output.SetLogger()
	return output
}

func (c *SDSClusterController) processNextItem() bool {
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
	err := c.proxessItem(key.(string))
	// Handle the error if something went wrong during the execution of the business logic
	c.handleErr(err, key)
	return true
}

// processItem is the business logic of the controller.
func (c *SDSClusterController) proxessItem(key string) error {
	obj, exists, err := c.indexer.GetByKey(key)
	if err != nil {
		glog.Errorf("Fetching object with key %s from store failed with %v", key, err)
		return err
	}

	if !exists {
		// Below we will warm up our cache with a SDSCluster, so that we will see a delete for one SDSCluster
		fmt.Printf("SDSCluster %s does not exist anymore\n", key)
	} else {
		sdsCluster := obj.(*api.SDSCluster)
		clusterName := sdsCluster.GetName()
		fmt.Printf("SDSCluster %s does exist (name=%s)!\n", key, clusterName)

		switch sdsCluster.Status.Phase {
		case api.ClusterPhaseNone, api.ClusterPhasePending:
			c.deployKrakenCluster(sdsCluster)
		case api.ClusterPhaseHaveCluster:
			c.deployPackageManager(sdsCluster)
		case api.ClusterPhaseHavePackageManager:
			c.deployApplications(sdsCluster)
		}
		if sdsCluster.Status.ClusterBuilt != true {
			c.deployKrakenCluster(sdsCluster)
		} else if sdsCluster.Status.TillerInstalled != true {
			c.deployPackageManager(sdsCluster)
		} else if sdsCluster.Status.AppsInstalled != true {
			c.deployApplications(sdsCluster)
		}
	}
	return nil
}

// handleErr checks if an error happened and makes sure we will retry later.
func (c *SDSClusterController) handleErr(err error, key interface{}) {
	if err == nil {
		// Forget about the #AddRateLimited history of the key on every successful synchronization.
		// This ensures that future processing of updates for this key is not delayed because of
		// an outdated error history.
		c.queue.Forget(key)
		return
	}

	// This controller retries 5 times if something goes wrong. After that, it stops trying.
	if c.queue.NumRequeues(key) < 5 {
		glog.Infof("Error syncing krakenCluster %v: %v", key, err)

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

func (c *SDSClusterController) Run(threadiness int, stopCh chan struct{}) {
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

func (c *SDSClusterController) runWorker() {
	for c.processNextItem() {
	}
}

func (c *SDSClusterController) SetLogger() {
	logger = util.GetModuleLogger("pkg.controllers.sds_cluster", loggo.INFO)
}

func (c *SDSClusterController) deployKrakenCluster(sdsCluster *api.SDSCluster) {
	clusterName := sdsCluster.GetName()
	_, err := ccutil.CreateKrakenCluster(
		ccutil.GenerateKrakenCluster(
			ccutil.KrakenClusterOptions{
				Name:     clusterName,
				Provider: sdsCluster.Spec.Provider.Name,
				MaaS: ccutil.MaaSOptions{
					Endpoint: sdsCluster.Spec.Provider.MaaS.Endpoint,
					Username: sdsCluster.Spec.Provider.MaaS.Username,
					OAuthKey: sdsCluster.Spec.Provider.MaaS.OAuthKey,
				},
				AWS: ccutil.AWSOptions{
					Region:          sdsCluster.Spec.Provider.AWS.Region,
					SecretKeyId:     sdsCluster.Spec.Provider.AWS.SecretKeyId,
					SecretAccessKey: sdsCluster.Spec.Provider.AWS.SecretAccessKey,
				},
			},
		), "default", nil)
	if (err != nil) && (!k8sutil.IsResourceAlreadyExistsError(err)) {
		logger.Infof("Could not create kraken cluster")
		sdsCluster.Status.Phase = api.ClusterPhasePending
		_, err = c.client.CmaV1alpha1().SDSClusters(sdsCluster.Namespace).Update(sdsCluster)
		if err != nil {
			logger.Infof("Could not update SDSCluster to Pending status, error was: %s", err)
		}
		return
	}

	sdsCluster.Status.Phase = api.ClusterPhaseWaitingForCluster
	_, err = c.client.CmaV1alpha1().SDSClusters(sdsCluster.Namespace).Update(sdsCluster)
	if err != nil {
		logger.Infof("Failed to request a KrakenCluster CR, error was: %s", err)
	}
}

func (c *SDSClusterController) deployPackageManager(sdsCluster *api.SDSCluster) {
	clusterName := sdsCluster.GetName()
	// Cluster name shouldn't have to be the name of the package manager - need to fix
	options := cma.SDSPackageManagerOptions{
		Name:            clusterName,
		Namespace:       sdsCluster.Spec.PackageManager.Namespace,
		Version:         sdsCluster.Spec.PackageManager.Version,
		ClusterWide:     sdsCluster.Spec.PackageManager.Permissions.ClusterWide,
		AdminNamespaces: sdsCluster.Spec.PackageManager.Permissions.Namespaces,
	}
	_, err := cma.CreateSDSPackageManager(cma.GenerateSDSPackageManager(options), sdsCluster.Namespace, nil)
	if (err != nil) && (!k8sutil.IsResourceAlreadyExistsError(err)) {
		logger.Infof("Could not create SDSPackageManager for cluster %s", clusterName)
		sdsCluster.Status.Phase = api.ClusterPhaseHaveCluster
		_, err = c.client.CmaV1alpha1().SDSClusters(sdsCluster.Namespace).Update(sdsCluster)
		if err != nil {
			logger.Infof("Could not update SDSCluster to HaveCluster status, error was: %s", err)
		}
		return
	}

	sdsCluster.Status.Phase = api.ClusterPhaseDeployingPackageManager
	_, err = c.client.CmaV1alpha1().SDSClusters(sdsCluster.Namespace).Update(sdsCluster)
	if err != nil {
		logger.Infof("Failed to request a SDSPackageManager CR, error was: %s", err)
	}
}

func (c *SDSClusterController) deployApplications(sdsCluster *api.SDSCluster) {
	clusterName := sdsCluster.GetName()
	for _, application := range sdsCluster.Spec.Applications {
		_, err := cma.CreateSDSApplication(cma.GenerateSDSApplication(cma.SDSApplicationOptions{
			Name:           sdsCluster.Name + "-" + application.Name,
			Namespace:      application.Namespace,
			Values:         application.Values,
			PackageManager: application.PackageManager.Name,
			Chart: cma.Chart{
				Name:    application.Chart.Name,
				Version: application.Chart.Version,
				Repository: cma.ChartRepository{
					Name: application.Chart.Repository.Name,
					URL:  application.Chart.Repository.URL,
				},
			},
		}), "default", nil)
		if err != nil {
			logger.Infof("Error creating SDSApplication -->%s<-- for cluster -->%s<--, error was -->%s<--", application.Name, sdsCluster.Name, err)
			continue
		}
		logger.Infof("Created SDSApplication -->%s<-- for SDSCluster -->%s<--", application.Name, sdsCluster.Name)
	}

	sdsCluster.Status.Phase = api.ClusterPhaseDeployingApplications
	_, err := c.client.CmaV1alpha1().SDSClusters(sdsCluster.Namespace).Update(sdsCluster)
	if err != nil {
		logger.Infof("Failed to request a SDSPackageManager CR for cluster %s, error was: %s", clusterName, err)
	}
}
