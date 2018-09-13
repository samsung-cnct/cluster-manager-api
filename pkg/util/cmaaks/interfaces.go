package cmaaks

type AKSClientInterface interface {
	CreateCluster(CreateClusterInput) (CreateClusterOutput, error)
	GetCluster(GetClusterInput) (GetClusterOutput, error)
	DeleteCluster(DeleteClusterInput) (DeleteClusterOutput, error)
	ListClusters(ListClusterInput) (ListClusterOutput, error)
	CreateNewClient(string, bool) error
	Close() error
}
