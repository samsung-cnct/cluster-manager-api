package cluster

import (
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/fields"

	"github.com/golang/glog"
	api "github.com/samsung-cnct/cluster-manager-api/pkg/apis/cma/v1alpha1"
	"github.com/samsung-cnct/cluster-manager-api/pkg/client/clientset/versioned"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/ccutil"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/k8sutil"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
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

	return &SDSClusterController{
		informer: informer,
		indexer:  indexer,
		queue:    queue,
		client:   client,
	}
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

		_, err := ccutil.CreateKrakenCluster(
			ccutil.GenerateKrakenCluster(
				ccutil.KrakenClusterOptions{Name: clusterName},
			), "default", nil)

		if err != nil {
			glog.Errorf("Failed to request a KrakenCluster CR", key, err)
			return err
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
