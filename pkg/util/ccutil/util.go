package ccutil

import (
	"github.com/juju/loggo"
	"github.com/samsung-cnct/cluster-controller/pkg/client/clientset/versioned"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/k8sutil"
	"k8s.io/client-go/rest"
)

var (
	logger loggo.Logger
)

func prepareRestClient(config *rest.Config) *versioned.Clientset {
	if config == nil {
		config = k8sutil.DefaultConfig
	}

	return versioned.NewForConfigOrDie(config)
}

func SetLogger() {
	logger = util.GetModuleLogger("pkg.util.ccutil", loggo.INFO)
}
