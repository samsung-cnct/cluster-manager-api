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

	logger.Infof("Was asked to get pods on -->%s<-- namespace, answer was -->%d<--", in.Namespace, int32(len(pods.Items)))
	return &pb.GetPodCountReply{Pods: int32(len(pods.Items))}, nil

}

func SetLogger() {
	logger = util.GetModuleLogger("internal.cluster-manager-api", loggo.INFO)
}