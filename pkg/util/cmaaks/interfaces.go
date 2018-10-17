package cmaaks

import (
	pb "github.com/samsung-cnct/cma-aks/pkg/generated/api"
	"google.golang.org/grpc"
)

type Client struct {
	conn   *grpc.ClientConn
	client pb.ClusterClient
}

type ClientInterface interface {
	CreateCluster(CreateClusterInput) (CreateClusterOutput, error)
	GetCluster(GetClusterInput) (GetClusterOutput, error)
	DeleteCluster(DeleteClusterInput) (DeleteClusterOutput, error)
	ListClusters(ListClusterInput) (ListClusterOutput, error)
	GetClusterUpgrades(input GetClusterUpgradesInput) (GetClusterUpgradesOutput, error)
	ClusterUpgrade(input ClusterUpgradeInput) (ClusterUpgradeOutput, error)
	GetClusterNodeCount(input GetClusterNodeCountInput) (GetClusterNodeCountOutput, error)
	ScaleCluster(input ScaleClusterInput) (ScaleClusterOutput, error)

	CreateNewClient(string, bool) error
	Close() error
}
