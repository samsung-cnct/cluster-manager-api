package cmassh

const (
	SshProvider = "Ssh"
)

type CreateClusterInput struct {
	Name              string
	K8SVersion        string
	ControlPlaneNodes []MachineSpec
	WorkerNodes       []MachineSpec
	APIEndpoint       string
	HighAvailability  bool
	NetworkFabric     string
}

type CreateClusterOutput struct {
	Cluster ClusterItem
}

type GetClusterInput struct {
	Name string
}

type GetClusterOutput struct {
	Cluster ClusterDetailItem
}

type DeleteClusterInput struct {
	Name string
}

type DeleteClusterOutput struct {
	Status string
}

type ListClusterInput struct{}

type ListClusterOutput struct {
	Clusters []ClusterItem
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

type MachineSpec struct {
	Username string
	Host     string
	Port     int
	Password string
	Labels   []KubernetesLabel
}

type KubernetesLabel struct {
	Name  string
	Value string
}

type AdjustClusterInput struct {
	Name        string
	AddNodes    []MachineSpec
	RemoveNodes []RemoveMachineSpec
}

type AdjustClusterOutput struct {
}

type RemoveMachineSpec struct {
	Host string
}

type GetClusterUpgradesInput struct {
	Name string
}

type GetClusterUpgradesOutput struct {
	Versions []string
}

type ClusterUpgradeInput struct {
	Name    string
	Version string
}

type ClusterUpgradeOutput struct{}
