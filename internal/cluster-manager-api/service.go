package cluster_manager_api

import (
	"github.com/juju/loggo"
	"github.com/samsung-cnct/cluster-manager-api/internal/cluster-manager-api/aws"
	"github.com/samsung-cnct/cluster-manager-api/internal/cluster-manager-api/azure"
	"github.com/samsung-cnct/cluster-manager-api/internal/cluster-manager-api/vmware"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util"
)

var (
	logger loggo.Logger
)

type Server struct {
	azure  azure.ClientInterface
	aws    aws.ClientInterface
	vmware vmware.ClientInterface
}

func NewServerFromDefaults() (*Server, error) {
	var awsClient aws.ClientInterface
	var azureClient azure.ClientInterface
	var vmwareClient vmware.ClientInterface
	var err error

	if aws.IsEnabled() {
		awsClient, err = aws.CreateFromDefaults()
		if err != nil {
			return nil, err
		}
	}

	if azure.IsEnabled() {
		azureClient, err = azure.CreateFromDefaults()
		if err != nil {
			return nil, err
		}
	}

	if vmware.IsEnabled() {
		vmwareClient, err = vmware.CreateFromDefaults()
		if err != nil {
			return nil, err
		}
	}

	return &Server{aws: awsClient, azure: azureClient, vmware: vmwareClient}, nil
}

func SetLogger() {
	logger = util.GetModuleLogger("internal.cluster-manager-api", loggo.INFO)
}
