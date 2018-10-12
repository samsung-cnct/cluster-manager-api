package cmak8sutil

import "github.com/samsung-cnct/cma-operator/pkg/apis/cma/v1alpha1"

type Application struct {
	CallbackURL    string
	Cluster        string
	Chart          Chart
	Namespace      string
	PackageManager string
	RequestID      string
	TillerSettings TillerSettings
	Values         string
	Phase          v1alpha1.ApplicationPhase
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
	Phase       v1alpha1.Phase
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
	Phase           v1alpha1.PackageManagerPhase
}

type TillerSettings struct {
	Namespace string
	Version   string
}
