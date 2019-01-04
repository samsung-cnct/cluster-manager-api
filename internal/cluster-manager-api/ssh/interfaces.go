package ssh

import (
	pb "github.com/samsung-cnct/cluster-manager-api/pkg/generated/api"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/cmassh"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/k8sutil/cma"
)

type Client struct {
	cmaSshClient cmassh.ClientInterface
	cmaK8sClient    cmak8sutil.ClientInterface
}

type ClientInterface interface {
	CreateCluster(in *pb.CreateClusterMsg) (*pb.CreateClusterReply, error)
	GetCluster(in *pb.GetClusterMsg) (*pb.GetClusterReply, error)
	GetClusterList(in *pb.GetClusterListMsg) (*pb.GetClusterListReply, error)
	DeleteCluster(in *pb.DeleteClusterMsg) (*pb.DeleteClusterReply, error)
	GetClusterUpgrades(in *pb.GetUpgradeClusterInformationMsg) (output *pb.GetUpgradeClusterInformationReply, err error)
	ClusterUpgrade(in *pb.UpgradeClusterMsg) (output *pb.UpgradeClusterReply, err error)
	AdjustCluster(in *pb.AdjustClusterMsg) (*pb.AdjustClusterReply, error)

	SetCMASshClient(client cmassh.ClientInterface)
	SetCMAK8sClient(client cmak8sutil.ClientInterface)
	Close() error
}
