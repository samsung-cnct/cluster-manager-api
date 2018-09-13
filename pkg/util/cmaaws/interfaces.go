package cmaaws

import (
	pb "gitlab.com/mvenezia/cma-aws/pkg/generated/api"
)

type ClientInterface interface {
	CreateCluster(CreateClusterInput) (CreateClusterOutput, error)
	GetCluster(GetClusterInput) (GetClusterOutput, error)
	DeleteCluster(DeleteClusterInput) (DeleteClusterOutput, error)
	ListClusters(ListClusterInput) (ListClusterOutput, error)
	CreateNewClient(string, bool) error
	Close() error
	SetClient(client pb.ClusterClient)
}
