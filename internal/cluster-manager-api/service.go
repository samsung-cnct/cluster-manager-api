package cluster_manager_api

import (
	"github.com/juju/loggo"
	"github.com/samsung-cnct/cluster-manager-api/internal/cluster-manager-api/aws"
	"github.com/samsung-cnct/cluster-manager-api/internal/cluster-manager-api/azure"
	"github.com/samsung-cnct/cluster-manager-api/internal/cluster-manager-api/ssh"
	"github.com/samsung-cnct/cluster-manager-api/internal/cluster-manager-api/vmware"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util"
	cmak8sutil "github.com/samsung-cnct/cluster-manager-api/pkg/util/k8sutil/cma"
)

var (
	logger loggo.Logger
)

type Server struct {
	azure  azure.ClientInterface
	aws    aws.ClientInterface
	vmware vmware.ClientInterface
	ssh    ssh.ClientInterface

	cmak8s cmak8sutil.ClientInterface
}

func NewServerFromDefaults() (*Server, error) {
	var awsClient aws.ClientInterface
	var azureClient azure.ClientInterface
	var vmwareClient vmware.ClientInterface
	var cmak8sClient cmak8sutil.ClientInterface
	var sshClient ssh.ClientInterface
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

	if ssh.IsEnabled() {
		sshClient, err = ssh.CreateFromDefaults()
		if err != nil {
			return nil, err
		}
	}

	cmak8sClient, err = cmak8sutil.CreateFromDefaults()
	if err != nil {
		return nil, err
	}

	return &Server{
		aws:    awsClient,
		azure:  azureClient,
		vmware: vmwareClient,
		cmak8s: cmak8sClient,
		ssh:    sshClient,
	}, nil
}

func SetLogger() {
	logger = util.GetModuleLogger("internal.cluster-manager-api", loggo.INFO)
}
