package cmak8sutil

type Application struct {
	CallbackURL    string
	Cluster        string
	Chart          Chart
	Namespace      string
	PackageManager string
	RequestID      string
	TillerSettings TillerSettings
	Values         string
}

type Chart struct {
	Name         string
	Repository   ChartRepository
	ChartPayload []byte
	Version      string
}

type ChartRepository struct {
	Name string
	URL  string
}

type Cluster struct {
	CallbackURL string
	Provider    string
	RequestID   string
}

type PackageManager struct {
	AdminNamespaces []string
	CallbackURL     string
	Cluster         string
	ClusterWide     bool
	Image           string
	Name            string
	Namespace       string
	RequestID       string
	Version         string
}

type TillerSettings struct {
	Namespace string
	Version   string
}
