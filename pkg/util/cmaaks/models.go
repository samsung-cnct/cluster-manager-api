package cmaaks

const (
	AKSProvider = "AKS"
)

type CreateClusterInput struct {
	Name             string
	K8SVersion       string
	Azure            AzureSpec
	HighAvailability bool
	NetworkFabric    string
}

type CreateClusterOutput struct {
	Cluster ClusterItem
}

type GetClusterInput struct {
	Name        string
	Credentials Credentials
}

type GetClusterOutput struct {
	Cluster ClusterDetailItem
}

type DeleteClusterInput struct {
	Name        string
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
	AppID          string
	Tenant         string
	Password       string
	SubscriptionID string
}

type AzureSpec struct {
	Location              string
	Credentials           Credentials
	ClusterServiceAccount ClusterServiceAccount
	InstanceGroups        []InstanceGroup
}

type ClusterItem struct {
	ID     string
	Name   string
	Status string
}

type ClusterDetailItem struct {
	ID         string
	Name       string
	Status     string
	Kubeconfig string
}

type ClusterServiceAccount struct {
	ClientID     string
	ClientSecret string
}

type InstanceGroup struct {
	Name        string
	Type        string
	MinQuantity int
	MaxQuantity int
}

type GetClusterUpgradesInput struct {
	Name        string
	Credentials Credentials
}

type GetClusterUpgradesOutput struct {
	Versions []string
}

type ClusterUpgradeInput struct {
	Name        string
	Credentials Credentials
	Version     string
}

type ClusterUpgradeOutput struct{}

type GetClusterNodeCountInput struct {
	Name        string
	Credentials Credentials
}

type GetClusterNodeCountOutput struct {
	Name  string
	Count int32
}

type ScaleClusterInput struct {
	Name        string
	Credentials Credentials
	NodePool    string
	Count       int32
}

type ScaleClusterOutput struct {}
