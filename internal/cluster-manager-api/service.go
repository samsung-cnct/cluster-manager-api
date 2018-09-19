package cluster_manager_api

import (
	"github.com/juju/loggo"
	"github.com/samsung-cnct/cluster-manager-api/internal/cluster-manager-api/aws"
	"github.com/samsung-cnct/cluster-manager-api/internal/cluster-manager-api/azure"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util"
)

var (
	logger loggo.Logger
)

type Server struct {
	azure azure.ClientInterface
	aws   aws.ClientInterface
}

func NewServerFromDefaults() (*Server, error) {
	azure, err := azure.CreateFromDefaults()
	if err != nil {
		return nil, err
	}
	aws, err := aws.CreateFromDefaults()
	if err != nil {
		azure.Close()
		return nil, err
	}

	return &Server{azure: azure, aws: aws}, nil
}

func SetLogger() {
	logger = util.GetModuleLogger("internal.cluster-manager-api", loggo.INFO)
}
