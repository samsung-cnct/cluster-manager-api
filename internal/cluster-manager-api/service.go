package cluster_manager_api

import (
	"github.com/juju/loggo"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util"
)

var (
	logger loggo.Logger
)

type Server struct{}

func SetLogger() {
	logger = util.GetModuleLogger("internal.cluster-manager-api", loggo.INFO)
}
