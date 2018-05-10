package cluster_manager_api

import (
	pb "github.com/samsung-cnct/cluster-manager-api/pkg/api"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/k8sutil"
	"golang.org/x/net/context"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/helmutil"
	"k8s.io/client-go/rest"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/ccutil"
	"fmt"
	"io/ioutil"
	"os"
	"k8s.io/client-go/tools/clientcmd"
)

func (s *Server) ProvisionTiller(ctx context.Context, in *pb.ProvisionTillerMsg) (*pb.ProvisionTillerReply, error) {
	SetLogger()
	config, err := retrieveClusterRestConfig(in.Cluster, "default", nil)
	if err != nil {
		return &pb.ProvisionTillerReply{Ok: false, Message: fmt.Sprintf("%v", err)}, nil
	}

	k8sutil.CreateNamespace(k8sutil.GenerateNamespace(in.Namespace), config)
	k8sutil.CreateServiceAccount(k8sutil.GenerateServiceAccount("tiller-sa"), in.Namespace, config)
	if in.ClusterWide {
		k8sutil.CreateClusterRole(helmutil.GenerateClusterAdminRole("tiller-"+in.Namespace), config)
		k8sutil.CreateClusterRoleBinding(k8sutil.GenerateSingleClusterRolebinding("tiller-"+in.Namespace, "tiller-sa", in.Namespace,"tiller-"+in.Namespace ), config)
	} else {
		logger.Infof("Not cluster wide")
		namespaces := append(in.AdminNamespaces, in.Namespace)
		for _, namespace := range namespaces {
			logger.Infof("Creating namespace %s", namespace)
			k8sutil.CreateNamespace(k8sutil.GenerateNamespace(namespace), config)
			k8sutil.CreateRole(helmutil.GenerateAdminRole(in.Namespace+"-tiller"), namespace, config)
			k8sutil.CreateRoleBinding(k8sutil.GenerateSingleRolebinding(in.Namespace+"-tiller", "tiller-sa", in.Namespace, in.Namespace+"-tiller"), namespace, config)
		}
	}
	k8sutil.CreateJob(helmutil.GenerateTillerInitJob(
		helmutil.TillerInitOptions{
			BackoffLimit:   4,
			Name:           "tiller-install-job",
			Namespace:      in.Namespace,
			ServiceAccount: "tiller-sa",
			Version:        in.Version}), in.Namespace, config)

	if config == nil {
		config = nil
	}
	return &pb.ProvisionTillerReply{Ok: true, Message: "Installed Tiller"}, nil
}

func retrieveClusterRestConfig(name string, namespace string, config *rest.Config) (*rest.Config, error) {
	cluster, err := ccutil.GetKrakenCluster(name, namespace, config)
	if err != nil {
		return nil, err
	}
	// Let's create a tempfile and line it up for removal
	file, err := ioutil.TempFile(os.TempDir(), "kraken-kubeconfig")
	defer os.Remove(file.Name())
	file.WriteString(cluster.Status.Kubeconfig)

	clusterConfig, err := clientcmd.BuildConfigFromFlags("", file.Name())

	if err != nil {
		logger.Errorf("Could not load kubeconfig for cluster -->%s<-- in namespace -->%s<--", name, namespace)
		return nil, err
	}
	return clusterConfig, nil
}