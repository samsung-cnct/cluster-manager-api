package cmak8sutil

import (
	"github.com/samsung-cnct/cma-operator/pkg/generated/cma/client/clientset/versioned/typed/cma/v1alpha1"
	"k8s.io/client-go/rest"
)

const (
	CallbackURLAnnotation = "callbackURL"
	RequestIDAnnotation   = "requestID"
)

type Client struct {
	kubeConfigLocation string
	config             *rest.Config

	applicationClient    v1alpha1.SDSApplicationInterface
	clusterClient        v1alpha1.SDSClusterInterface
	packageManagerClient v1alpha1.SDSPackageManagerInterface
}

type ClientInterface interface {
	CreateApplication(name string, application Application) error
	GetApplication(name string, clusterName string) (Application, error)
	UpdateOrCreateApplication(name string, application Application) error
	DeleteApplication(name string, clusterName string) error
	ChangeApplicationStatus(name string, clusterName string, status string) error

	CreateCluster(name string, cluster Cluster) error
	GetCluster(name string) (Cluster, error)
	UpdateOrCreateCluster(name string, cluster Cluster) error
	DeleteCluster(name string) error
	ChangeClusterStatus(name string, status string) error

	CreatePackageManager(name string, packageManager PackageManager) error
	GetPackageManager(name string, clusterName string) (PackageManager, error)
	UpdateOrCreatePackageManager(name string, packageManager PackageManager) error
	DeletePackageManager(name string, clusterName string) error
	ChangePackageManagerStatus(name string, clusterName string, status string) error

	CreateNewClients() error
	SetConfig(config *rest.Config)
	SetApplicationClient(client v1alpha1.SDSApplicationInterface)
	SetClusterClient(client v1alpha1.SDSClusterInterface)
	SetPackageManagerClient(client v1alpha1.SDSPackageManagerInterface)
}
