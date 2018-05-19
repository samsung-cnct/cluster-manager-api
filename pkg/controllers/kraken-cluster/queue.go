package kraken_cluster

import (
	"fmt"
	"time"

	"github.com/golang/glog"

	ccapi "github.com/samsung-cnct/cluster-controller/pkg/apis/clustercontroller/v1alpha1"
	"github.com/samsung-cnct/cluster-controller/pkg/client/clientset/versioned"
	cmaapi "github.com/samsung-cnct/cluster-manager-api/pkg/apis/cma/v1alpha1"
	sdsClient "github.com/samsung-cnct/cluster-manager-api/pkg/generated/cma/client/clientset/versioned"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/k8sutil"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"

	"log"
)

type Controller struct {
	indexer  cache.Indexer
	queue    workqueue.RateLimitingInterface
	informer cache.Controller

	client *versioned.Clientset
}

func NewController(queue workqueue.RateLimitingInterface, indexer cache.Indexer, informer cache.Controller, client *versioned.Clientset) *Controller {
	return &Controller{
		informer: informer,
		indexer:  indexer,
		queue:    queue,
		client:   client,
	}
}

func (c *Controller) processNextItem() bool {
	// Wait until there is a new item in the working queue
	key, quit := c.queue.Get()
	if quit {
		return false
	}
	// Tell the queue that we are done with processing this key. This unblocks the key for other workers
	// This allows safe parallel processing because two KrakenClusters with the same key are never processed in
	// parallel.
	defer c.queue.Done(key)

	// Invoke the method containing the business logic
	err := c.syncToStdout(key.(string))
	// Handle the error if something went wrong during the execution of the business logic
	c.handleErr(err, key)
	return true
}

// syncToStdout is the business logic of the controller. In this controller it simply prints
// information about the KrakenClusters to stdout. In case an error happened, it has to simply return the error.
// The retry logic should not be part of the business logic.
func (c *Controller) syncToStdout(key string) error {
	obj, exists, err := c.indexer.GetByKey(key)
	if err != nil {
		glog.Errorf("Fetching object with key %s from store failed with %v", key, err)
		return err
	}

	if !exists {
		// Below we will warm up our cache with a KrakenCluster, so that we will see a delete for one KrakenCluster
		fmt.Printf("KrakenCluster %s does not exist anymore\n", key)
	} else {
		// Note that you also have to check the uid if you have a local controlled resource, which
		// is dependent on the actual instance, to detect that a KrakenCluster was recreated with the same name
		fmt.Printf("Sync/Add/Update for KrakenCluster %s\n", obj.(*ccapi.KrakenCluster).GetName())
		if (obj.(*ccapi.KrakenCluster).Status.State == ccapi.Created) && (obj.(*ccapi.KrakenCluster).Status.Kubeconfig != "") {
			c.updateSDSCluster(obj.(*ccapi.KrakenCluster))
		}

	}
	return nil
}

// handleErr checks if an error happened and makes sure we will retry later.
func (c *Controller) handleErr(err error, key interface{}) {
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

func (c *Controller) Run(threadiness int, stopCh chan struct{}) {
	defer runtime.HandleCrash()

	// Let the workers stop when we are done
	defer c.queue.ShutDown()
	glog.Info("Starting KrakenCluster controller")

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
	glog.Info("Stopping KrakenCluster controller")
}

func (c *Controller) runWorker() {
	for c.processNextItem() {
	}
}

func ListenToKrakenClusterChanges(config *rest.Config) {
	if config == nil {
		config = k8sutil.DefaultConfig
	}
	clientset := versioned.NewForConfigOrDie(config)

	// create the KrakenCluster watcher
	krakentClusterListWatcher := cache.NewListWatchFromClient(clientset.SamsungV1alpha1().RESTClient(), "KrakenClusters", "default", fields.Everything())

	// create the workqueue
	queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())

	// Bind the workqueue to a cache with the help of an informer. This way we make sure that
	// whenever the cache is updated, the KrakenCluster key is added to the workqueue.
	// Note that when we finally process the item from the workqueue, we might see a newer version
	// of the KrakenCluster than the version which was responsible for triggering the update.
	indexer, informer := cache.NewIndexerInformer(krakentClusterListWatcher, &ccapi.KrakenCluster{}, 0, cache.ResourceEventHandlerFuncs{
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

	controller := NewController(queue, indexer, informer, clientset)

	// Now let's start the controller
	stop := make(chan struct{})
	defer close(stop)
	go controller.Run(1, stop)

	// Wait forever
	select {}
}

func (c *Controller) updateSDSCluster(krakenCluster *ccapi.KrakenCluster) {
	clusterName := krakenCluster.GetName()
	log.Printf("I'm here for cluster %s", clusterName)
	client := sdsClient.NewForConfigOrDie(k8sutil.DefaultConfig)
	sdsCluster, err := client.CmaV1alpha1().SDSClusters(krakenCluster.Namespace).Get(clusterName, v1.GetOptions{})
	if err != nil {
		log.Printf("Failed to get SDSCluster for KrakenCluster %s, error was: %s", clusterName, err)
	}
	changes := false
	if sdsCluster.Status.ClusterBuilt == false {
		changes = true
		sdsCluster.Status.ClusterBuilt = true
	}
	if sdsCluster.Status.Phase == cmaapi.ClusterPhaseWaitingForCluster || sdsCluster.Status.Phase == cmaapi.ClusterPhaseNone || sdsCluster.Status.Phase == cmaapi.ClusterPhasePending {
		changes = true
		sdsCluster.Status.Phase = cmaapi.ClusterPhaseHaveCluster
	}
	if changes {
		_, err = client.CmaV1alpha1().SDSClusters(sdsCluster.Namespace).Update(sdsCluster)
		if err != nil {
			log.Printf("Could not update SDSCluster for KrakenCluster %s, error was: %s", clusterName, err)
		}
	}
}
