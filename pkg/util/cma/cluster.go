package cma

import (
	sdsapi "github.com/samsung-cnct/cluster-manager-api/pkg/apis/cma/v1alpha1"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/k8sutil"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

type SDSClusterOptions struct {
	Name     string
	Provider string
	AWS      AWSOptions
	MaaS     MaaSOptions
}

type AWSOptions struct {
	SecretKeyId     string
	SecretAccessKey string
	Region          string
}

type MaaSOptions struct {
	Endpoint string
	Username string
	OAuthKey string
}

func GenerateSDSCluster(options SDSClusterOptions) sdsapi.SDSCluster {
	return sdsapi.SDSCluster{
		ObjectMeta: metav1.ObjectMeta{
			Name: options.Name,
		},
		Spec: sdsapi.SDSClusterSpec{
			Provider: sdsapi.ProviderSpec{
				Name: options.Provider,
				AWS:  sdsapi.AWSSpec{Region: options.AWS.Region, SecretAccessKey: options.AWS.SecretAccessKey, SecretKeyId: options.AWS.SecretKeyId},
				MaaS: sdsapi.MaaSSpec{Endpoint: options.MaaS.Endpoint, Username: options.MaaS.Username, OAuthKey: options.MaaS.OAuthKey},
			},
			PackageManager: sdsapi.SDSPackageManagerSpec{
				Version:   "v2.9.0",
				Name:      options.Name,
				Namespace: "cma-tiller",
				Permissions: sdsapi.PackageManagerPermissions{
					ClusterWide: true,
				},
			},
			Applications: []sdsapi.SDSApplicationSpec{
				{
					Name:           "prometheus-operator",
					Namespace:      "prometheus",
					PackageManager: sdsapi.SDSPackageManagerRef{Name: options.Name},
					Chart: sdsapi.Chart{
						Name:       "coreos/prometheus-operator",
						Repository: sdsapi.ChartRepository{Name: "coreos", URL: "https://s3-eu-west-1.amazonaws.com/coreos-charts/stable/"},
					},
				},
				{
					Name:           "kube-prometheus",
					Namespace:      "prometheus",
					PackageManager: sdsapi.SDSPackageManagerRef{Name: options.Name},
					Chart: sdsapi.Chart{
						Name:       "coreos/kube-prometheus",
						Repository: sdsapi.ChartRepository{Name: "coreos", URL: "https://s3-eu-west-1.amazonaws.com/coreos-charts/stable/"},
					},
				},
				{
					Name:           "logging",
					Namespace:      "logging",
					PackageManager: sdsapi.SDSPackageManagerRef{Name: options.Name},
					Chart: sdsapi.Chart{
						Name:       "sds/logging-client",
						Repository: sdsapi.ChartRepository{Name: "sds", URL: "https://charts.migrations.cnct.io"},
					},
				},
				{
					Name:           "nginx-ingress",
					Namespace:      "ingress",
					PackageManager: sdsapi.SDSPackageManagerRef{Name: options.Name},
					Chart: sdsapi.Chart{
						Name:       "stable/nginx-ingress",
						Repository: sdsapi.ChartRepository{Name: "stable", URL: "https://charts.migrations.cnct.io"},
					},
					Values: `## nginx configuration
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
				},
			},
		},
	}
}

func CreateSDSCluster(cluster sdsapi.SDSCluster, namespace string, config *rest.Config) (bool, error) {
	var err error
	SetLogger()
	client := prepareRestClient(config)

	_, err = client.CmaV1alpha1().SDSClusters(namespace).Create(&cluster)
	if err != nil && !k8sutil.IsResourceAlreadyExistsError(err) {
		logger.Infof("SDSCluster -->%s<-- Cannot be created, error was %v", cluster.ObjectMeta.Name, err)
		return false, err
	} else if k8sutil.IsResourceAlreadyExistsError(err) {
		logger.Infof("SDSCluster -->%s<-- Already exists, cannot recreate", cluster.ObjectMeta.Name)
		return false, err
	}

	return true, err
}

func DeleteSDSCluster(name string, namespace string, config *rest.Config) (bool, error) {
	var err error
	SetLogger()
	client := prepareRestClient(config)

	err = client.CmaV1alpha1().SDSClusters(namespace).Delete(name, &metav1.DeleteOptions{})
	if err != nil {
		logger.Infof("SDSCluster -->%s<-- Cannot be Deleted, error was %v", name, err)
		return false, err
	}
	return true, err
}

func UpdateSDSCluster(cluster *sdsapi.SDSCluster, namespace string, config *rest.Config) (*sdsapi.SDSCluster, error) {
	var err error
	SetLogger()
	client := prepareRestClient(config)

	updatedCluster, err := client.CmaV1alpha1().SDSClusters(namespace).Update(cluster)
	if err != nil {
		logger.Infof("SDSCluster -->%s<-- Cannot be Updated, error was %v", cluster.Name, err)
		return updatedCluster, err
	}
	return nil, err
}

func GetSDSCluster(name string, namespace string, config *rest.Config) (*sdsapi.SDSCluster, error) {
	var err error
	SetLogger()
	client := prepareRestClient(config)

	cluster, err := client.CmaV1alpha1().SDSClusters(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		logger.Infof("SDSCluster -->%s<-- Cannot be Retrieved, error was %v", name, err)
		return nil, err
	}
	return cluster, err
}

func ListSDSClusters(namespace string, config *rest.Config) ([]sdsapi.SDSCluster, error) {
	var err error
	SetLogger()
	client := prepareRestClient(config)

	list, err := client.CmaV1alpha1().SDSClusters(namespace).List(metav1.ListOptions{})
	if err != nil {
		logger.Infof("SDSCluster List Cannot be Retrieved, error was %v", err)
		return nil, err
	}
	return list.Items, nil

}
