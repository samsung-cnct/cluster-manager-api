package cmavmware

type VMWareClientInterface interface {
	CreateCluster(CreateClusterInput) (CreateClusterOutput, error)
	GetCluster(GetClusterInput) (GetClusterOutput, error)
	DeleteCluster(DeleteClusterInput) (DeleteClusterOutput, error)
	ListClusters(ListClusterInput) (ListClusterOutput, error)
	AdjustCluster(AdjustClusterInput) (AdjustClusterOutput, error)
	CreateNewClient(string, bool) error
	Close() error
}
