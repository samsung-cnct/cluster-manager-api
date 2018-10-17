package azure

import (
	pb "github.com/samsung-cnct/cluster-manager-api/pkg/generated/api"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/cmaaks"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/k8sutil/azure"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/k8sutil/cma"
)

type Client struct {
	cmaAKSClient cmaaks.ClientInterface
	secretClient azurek8sutil.ClientInterface
	cmaK8sClient cmak8sutil.ClientInterface
}

type ClientInterface interface {
	CreateCluster(in *pb.CreateClusterMsg) (*pb.CreateClusterReply, error)
	GetCluster(in *pb.GetClusterMsg) (*pb.GetClusterReply, error)
	GetClusterList(in *pb.GetClusterListMsg) (*pb.GetClusterListReply, error)
	DeleteCluster(in *pb.DeleteClusterMsg) (*pb.DeleteClusterReply, error)
	UpdateCredentials(in *pb.UpdateAzureCredentialsMsg) (*pb.UpdateAzureCredentialsReply, error)
	GetClusterUpgrades(in *pb.GetUpgradeClusterInformationMsg) (output *pb.GetUpgradeClusterInformationReply, err error)
	ClusterUpgrade(in *pb.UpgradeClusterMsg) (output *pb.UpgradeClusterReply, err error)
	AdjustCluster(in *pb.AdjustClusterMsg) (*pb.AdjustClusterReply, error)

	SetCMAAKSClient(client cmaaks.ClientInterface)
	SetSecretClient(client azurek8sutil.ClientInterface)
	SetCMAK8sClient(client cmak8sutil.ClientInterface)
	Close() error
}
