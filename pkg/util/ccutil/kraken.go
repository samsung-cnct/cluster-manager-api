package ccutil

import (
	ccapi "github.com/samsung-cnct/cluster-controller/pkg/apis/clustercontroller/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/k8sutil"
	"k8s.io/client-go/rest"
)

type KrakenClusterOptions struct {
	Name string
	Provider string
	AWS AWSOptions
	MaaS MaaSOptions
}

type AWSOptions struct {
	SecretKeyId	string
	SecretAccessKey string
	Region string
}

type MaaSOptions struct {
	Endpoint string
	Username string
	OAuthKey	string
}


func GenerateKrakenCluster(options KrakenClusterOptions) ccapi.KrakenCluster {
	return ccapi.KrakenCluster{
		ObjectMeta: metav1.ObjectMeta{
			Name: options.Name,
		},
		Spec: ccapi.KrakenClusterSpec{
			CustomerID: "myCustomerID",
			CloudProvider: generateProviderBlock(options),
			Provisioner: generateProvisionerBlock(options),
			Cluster: ccapi.ClusterInfo{
				ClusterName: "my-test-cluster",
				NodePools: []ccapi.NodeProperties{
					{
						Name:        "worker",
						PublicIPs:   false,
						Size:        1,
						MachineType: "m4.xlarge",
						Os:          "ubuntu:16:04",
					},
					{
						Name:        "master",
						PublicIPs:   false,
						Size:        1,
						MachineType: "m4.xlarge",
						Os:          "ubuntu:16:04",
					},
					{
						Name:        "etcd",
						PublicIPs:   false,
						Size:        3,
						MachineType: "m2.medium",
						Os:          "ubuntu:16:04",
					},
				},
				Fabric: ccapi.FabricInfo{
					Name: "canal",
				},
			},
		},
	}
}

func generateProviderBlock(options KrakenClusterOptions) ccapi.CloudProviderInfo {
	switch options.Provider {
	case "aws":
		return ccapi.CloudProviderInfo{
			Name: "aws",
			Credentials: ccapi.CloudProviderCredentials{
				Accesskey: options.AWS.SecretAccessKey,
				Password: options.AWS.SecretKeyId,
			},
			Region: "aws/"+options.AWS.Region,
		}
	default:
		return ccapi.CloudProviderInfo{
			Name: "maas",
			Credentials: ccapi.CloudProviderCredentials{
				Username: options.MaaS.Username,
				Password: options.MaaS.OAuthKey,
			},
		}
	}
}

func generateProvisionerBlock(options KrakenClusterOptions) ccapi.ProvisionerInfo {
	switch options.Provider {
	case "aws":
		return ccapi.ProvisionerInfo{
			Name: "juju",
			Bundle: "cs:bundle/kubernetes-core-306",
		}
	default:
		return ccapi.ProvisionerInfo{
			Name: "juju",
			Bundle: "cs:bundle/kubernetes-core-306",
			MaasEndpoint: options.MaaS.Endpoint,
		}
	}
}


func CreateKrakenCluster(cluster ccapi.KrakenCluster, namespace string, config *rest.Config) (bool, error) {
	var err error
	SetLogger()
	client := prepareRestClient(config)

	_, err = client.SamsungV1alpha1().KrakenClusters(namespace).Create(&cluster)
	if err != nil && !k8sutil.IsResourceAlreadyExistsError(err) {
		logger.Infof("KrakenCluster -->%s<-- Cannot be created, error was %v", cluster.ObjectMeta.Name, err)
		return false, err
	} else if k8sutil.IsResourceAlreadyExistsError(err) {
		logger.Infof("KrakenCluster -->%s<-- Already exists, cannot recreate", cluster.ObjectMeta.Name)
		return false, err
	}

	return true, err
}

func DeleteKrakenCluster(name string, namespace string, config *rest.Config) (bool, error) {
	var err error
	SetLogger()
	client := prepareRestClient(config)

	err = client.SamsungV1alpha1().KrakenClusters(namespace).Delete(name, &metav1.DeleteOptions{})
	if err != nil {
		logger.Infof("KrakenCluster -->%s<-- Cannot be Deleted, error was %v", name, err)
		return false, err
	}
	return true, err
}

func GetKrakenCluster(name string, namespace string, config *rest.Config) (*ccapi.KrakenCluster, error) {
	var err error
	SetLogger()
	client := prepareRestClient(config)

	cluster, err := client.SamsungV1alpha1().KrakenClusters(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		logger.Infof("KrakenCluster -->%s<-- Cannot be Retrieved, error was %v", name, err)
		return nil, err
	}
	return cluster, err
}