package cluster_manager_api

import (
	"github.com/samsung-cnct/cluster-controller/pkg/client/clientset/versioned"
	pb "github.com/samsung-cnct/cluster-manager-api/pkg/api"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/k8sutil"
	"golang.org/x/net/context"
	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/kubernetes"

	"github.com/juju/loggo"
	clusterController "github.com/samsung-cnct/cluster-manager-api/pkg/cluster-controller"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

)

var (
	logger loggo.Logger
)

type Server struct{}

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



	//tempCluster, err := clusterControllerClient.KrakenCRCli.SamsungV1alpha1().KrakenClusters("default").Get("my-test-cluster", metav1.GetOptions{})
	//tempCluster.ObjectMeta.Labels


	logger.Infof("Was asked to get pods on -->%s<-- namespace, answer was -->%d<--", in.Namespace, int32(len(pods.Items)))
	return &pb.GetPodCountReply{Pods: int32(len(pods.Items))}, nil

}


func SetLogger() {
	logger = util.GetModuleLogger("internal.cluster-manager-api", loggo.INFO)
}
