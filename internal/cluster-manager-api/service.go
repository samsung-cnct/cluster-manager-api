package cluster_manager_api

import (
	"golang.org/x/net/context"
	pb "github.com/samsung-cnct/cluster-manager-api/pkg/api"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/k8sutil"
	"k8s.io/client-go/kubernetes"
	"github.com/samsung-cnct/cluster-controller/pkg/client/clientset/versioned"
	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"

	clusterController "github.com/samsung-cnct/cluster-manager-api/pkg/cluster-controller"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/juju/loggo"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util"

	ccapi "github.com/samsung-cnct/cluster-controller/pkg/apis/clustercontroller/v1alpha1"
)

var (
	logger loggo.Logger
)

type Server struct {}

func (s *Server) HelloWorld(ctx context.Context, in *pb.HelloWorldMsg) (*pb.HelloWorldReply, error) {
	return &pb.HelloWorldReply{Message: "Hello " + in.Name}, nil
}

func (s *Server) GetPodCount(ctx context.Context, in *pb.GetPodCountMsg) (*pb.GetPodCountReply, error) {
	SetLogger()
	// create the clientSet
	clientSet, err := kubernetes.NewForConfig(k8sutil.DefaultConfig)
	if err != nil {
		logger.Errorf("Cannot establish a client connection to kubernetes: %v", err)
		return nil, err
	}

	clusterControllerClient := clusterController.New(clusterController.Config{
		KubeCli:     clientSet,
		KubeExtCli:  apiextensionsclient.NewForConfigOrDie(k8sutil.DefaultConfig),
		KrakenCRCli: versioned.NewForConfigOrDie(k8sutil.DefaultConfig),
	})

	pods, err := clusterControllerClient.KubeCli.CoreV1().Pods(in.Namespace).List(metav1.ListOptions{})
	if err != nil {
		logger.Errorf("Cannot establish a client connection to kubernetes: %v", err)
		return nil, err
	}

	dummy := &ccapi.KrakenCluster{
		ObjectMeta: metav1.ObjectMeta{
			Name: "my-test-cluster",
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
						Name: "worker",
						PublicIPs: false,
						Size: 1,
						MachineType: "m4.xlarge",
						Os: "ubuntu:16:04",
					},
					{
						Name: "master",
						PublicIPs: false,
						Size: 1,
						MachineType: "m4.xlarge",
						Os: "ubuntu:16:04",
					},
					{
						Name: "etcd",
						PublicIPs: false,
						Size: 3,
						MachineType: "m3.medium",
						Os: "ubuntu:16:04",
					},
				},
				Fabric: ccapi.FabricInfo{
					Name: "canal",
				},
			},
		},
	}

	_, err = clusterControllerClient.KrakenCRCli.SamsungV1alpha1().KrakenClusters("default").Create(dummy)
	if err != nil && !k8sutil.IsResourceAlreadyExistsError(err) {
		logger.Infof("KrakenCluster -->%s<-- Cannot be created, error was %v", dummy.ObjectMeta.Name, err)
	} else if k8sutil.IsResourceAlreadyExistsError(err) {
		logger.Infof("KrakenCluster -->%s<-- Already exists, cannot recreate", dummy.ObjectMeta.Name)
	}

	logger.Infof("Was asked to get pods on -->%s<-- namespace, answer was -->%d<--", in.Namespace, int32(len(pods.Items)))
	return &pb.GetPodCountReply{Pods: int32(len(pods.Items))}, nil

}

func SetLogger() {
	logger = util.GetModuleLogger("internal.cluster-manager-api", loggo.INFO)
}