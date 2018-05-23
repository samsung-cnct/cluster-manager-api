package layouts

import (
	"github.com/samsung-cnct/cluster-manager-api/pkg/apis/cma/v1alpha1"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/cma"
)

type Layout interface {
	GenerateSDSCluster(options cma.SDSClusterOptions) v1alpha1.SDSCluster
	GenerateSDSPackageManager(options cma.SDSPackageManagerOptions, cluster *v1alpha1.SDSCluster) v1alpha1.SDSPackageManager
	GenerateSDSApplications(cluster *v1alpha1.SDSCluster, packageManager *v1alpha1.SDSPackageManager) []v1alpha1.SDSApplication
}