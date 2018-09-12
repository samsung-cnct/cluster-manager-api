package cmaaws


const (
	AWSProvider = "AWS"
)

type CreateClusterInput struct {
	Name string
	K8SVersion string
	AWS AWSSpec
	HighAvailability bool
	NetworkFabric string
}

type CreateClusterOutput struct {
	Cluster ClusterItem
}

type GetClusterInput struct {
	Name string
	Credentials Credentials
}

type GetClusterOutput struct {
	Cluster ClusterDetailItem
}

type DeleteClusterInput struct {
	Name string
	Credentials Credentials
}

type DeleteClusterOutput struct {
	Status string
}

type ListClusterInput struct {
	Credentials Credentials
}

type ListClusterOutput struct {
	Clusters []ClusterItem
}

type Credentials struct {
	SecretKeyID string
	SecretAccessKey string
	Region string
}

type AWSSpec struct {
	DataCenter DataCenter
	Credentials Credentials
	PreconfiguredItems PreconfiguredItems
	InstanceGroups []InstanceGroup
}

type DataCenter struct {
	Region string
	AvailabilityZones []string
}

type PreconfiguredItems struct {
	VPCID string
	SecurityGroupID string
	IAMRoleARN string
}

type ClusterItem struct {
	ID string
	Name string
	Status string
}

type ClusterDetailItem struct {
	ID string
	Name string
	Status string
	Kubeconfig string
}

type InstanceGroup struct {
	Type string
	MinQuantity int
	MaxQuantity int
}
