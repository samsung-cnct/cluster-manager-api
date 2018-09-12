package cmaaks

type AKSClientInterface interface {
	CreateCluster(CreateClusterInput) (CreateClusterOutput, error)
	GetCluster(GetClusterInput) (GetClusterOutput, error)
	DeleteCluster(DeleteClusterInput) (DeleteClusterOutput, error)
	ListClusters(ListClusterInput) (ListClusterOutput, error)
}

type AKSRealClientInterface interface {
	CreateNewClient(string, bool) error
	Close() error
	AKSClientInterface
}

