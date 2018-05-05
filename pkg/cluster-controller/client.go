package cluster_controller

import (
	"github.com/samsung-cnct/cluster-controller/pkg/client/clientset/versioned"
	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/kubernetes"

	api "github.com/samsung-cnct/cluster-controller/pkg/apis/clustercontroller/v1alpha1"

	"github.com/sirupsen/logrus"
	kwatch "k8s.io/apimachinery/pkg/watch"

	"time"
)

var initRetryWaitTime = 30 * time.Second

var pt *panicTimer

func init() {
	pt = newPanicTimer(time.Minute, "unexpected long blocking (> 1 Minute) when handling cluster event")
}

// Event object
type Event struct {
	Type   kwatch.EventType
	Object *api.KrakenCluster
}

// Controller object, contains configs and map of redis instances.
type Client struct {
	logger *logrus.Entry
	Config
}

// Config object
type Config struct {
	Namespace      string
	ClusterWide    bool
	ServiceAccount string
	KubeCli        kubernetes.Interface
	KubeExtCli     apiextensionsclient.Interface
	KrakenCRCli    versioned.Interface
	CreateCRD      bool
}

// New creates new client object
func New(cfg Config) *Client {
	return &Client{
		logger: logrus.WithField("pkg", "cluster_controller"),
		Config: cfg,
	}
}
