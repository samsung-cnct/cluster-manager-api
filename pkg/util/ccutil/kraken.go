package ccutil

import (
	ccapi "github.com/samsung-cnct/cluster-controller/pkg/apis/clustercontroller/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/k8sutil"
	"k8s.io/client-go/rest"
)

type KrakenClusterOptions struct {
	Name string
}

func GenerateKrakenCluster(options KrakenClusterOptions) ccapi.KrakenCluster {
	return ccapi.KrakenCluster{
		ObjectMeta: metav1.ObjectMeta{
			Name: options.Name,
		},
		Spec: ccapi.KrakenClusterSpec{
			CustomerID: "myCustomerID",
			CloudProvider: ccapi.CloudProviderInfo{
				Name: "aws",
				Credentials: ccapi.CloudProviderCredentials{
					Username: "myuser",
					Password: "fakepassword1",
				},
			},
			Provisioner: ccapi.ProvisionerInfo{
				Name: "juju",
			},
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
						MachineType: "m3.medium",
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