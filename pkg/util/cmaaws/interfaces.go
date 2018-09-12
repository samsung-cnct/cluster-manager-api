package cmaaws

type ClientInterface interface {
	CreateCluster(CreateClusterInput) (CreateClusterOutput, error)
	GetCluster(GetClusterInput) (GetClusterOutput, error)
	DeleteCluster(DeleteClusterInput) (DeleteClusterOutput, error)
	ListClusters(ListClusterInput) (ListClusterOutput, error)
}

type RealClientInterface interface {
	CreateNewClient(string, bool) error
	Close() error
	ClientInterface
}
