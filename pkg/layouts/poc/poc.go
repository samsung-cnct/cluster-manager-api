package poc

import (
	"github.com/juju/loggo"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/cma"
	"github.com/samsung-cnct/cluster-manager-api/pkg/apis/cma/v1alpha1"
)

type Layout struct {}

var (
	logger loggo.Logger
)

func SetLogger() {
	logger = util.GetModuleLogger("pkg.layouts.poc", loggo.INFO)
}

func NewLayout() *Layout {
	return &Layout{}
}

func (l *Layout) GenerateSDSCluster(options cma.SDSClusterOptions) v1alpha1.SDSCluster {
	cluster := cma.GenerateSDSCluster(options)
	cluster.Spec.Layout = "poc"
	return cluster
}

func (l *Layout) GenerateSDSPackageManager(options cma.SDSPackageManagerOptions, cluster *v1alpha1.SDSCluster) v1alpha1.SDSPackageManager {
	// Overriding whatever options came in
	options.Name = cluster.Name
	options.Version = "v2.9.1"
	options.Namespace = "cma-tiller"
	options.ClusterWide = true
	options.AdminNamespaces = []string{}

	packageManager := cma.GenerateSDSPackageManager(options)
	packageManager.Labels = make(map[string]string)
	packageManager.Labels["SDSCluster"] = string(cluster.ObjectMeta.UID)
	return packageManager
}

func (l *Layout) GenerateSDSApplications(cluster *v1alpha1.SDSCluster, packageManager *v1alpha1.SDSPackageManager) []v1alpha1.SDSApplication {
	var applications []v1alpha1.SDSApplication

	// Generating Prometheus Operator Application
	applications = append(applications, cma.GenerateSDSApplication(cma.SDSApplicationOptions{
		Name: "prometheus-operator",
		Namespace: "prometheus",
		PackageManager: packageManager.Name,
		Chart: cma.Chart{
			Name:       "coreos/prometheus-operator",
			Repository: cma.ChartRepository{Name: "coreos", URL: "https://s3-eu-west-1.amazonaws.com/coreos-charts/stable/"},
		},
	}))

	// Generating Kube-Prometheus Application
	applications = append(applications, cma.GenerateSDSApplication(cma.SDSApplicationOptions{
		Name: "kube-prometheus",
		Namespace: "prometheus",
		PackageManager: packageManager.Name,
		Chart: cma.Chart{
			Name:       "coreos/kube-prometheus",
			Repository: cma.ChartRepository{Name: "coreos", URL: "https://s3-eu-west-1.amazonaws.com/coreos-charts/stable/"},
		},
	}))

	// Generating Logging Client Application
	applications = append(applications, cma.GenerateSDSApplication(cma.SDSApplicationOptions{
		Name: "logging",
		Namespace: "logging",
		PackageManager: packageManager.Name,
		Chart: cma.Chart{
			Name:       "sds/logging-client",
			Repository: cma.ChartRepository{Name: "sds", URL: "https://charts.migrations.cnct.io"},
		},
	}))

	// Generating nginx-ingress Application
	applications = append(applications, cma.GenerateSDSApplication(cma.SDSApplicationOptions{
		Name:           "nginx-ingress",
		Namespace:      "ingress",
		PackageManager: packageManager.Name,
		Chart: cma.Chart{
			Name:       "stable/nginx-ingress",
			Repository: cma.ChartRepository{Name: "sds", URL: "https://charts.migrations.cnct.io"},
		},
		Values: `
## nginx configuration
## Ref: https://raw.githubusercontent.com/kubernetes/charts/master/stable/nginx-ingress/values.yaml
##
controller:
  service:
    targetPorts:
      http: 80
      https: 443
    type: NodePort
  stats:
    enabled: true
    service:
      annotations: {}
      clusterIP: ""
      ## List of IP addresses at which the stats service is available
      ## Ref: https://kubernetes.io/docs/user-guide/services/#external-ips
      ##
      servicePort: 18080
      type: ClusterIP

  ## If controller.stats.enabled = true and controller.metrics.enabled = true, Prometheus metrics will be exported
  ##
  metrics:
    enabled: true
    service:
      servicePort: 9913
      type: ClusterIP
## RBAC is now enabled by default.  disable it.
rbac:
  create: false
  createRole: false
  createClusterRole: false
  serviceAccountName: default`,
	}))

	for _, application := range applications {
		application.Labels = make(map[string]string)
		application.Labels["SDSCluster"] = string(cluster.ObjectMeta.UID)
	}

	return applications
}
