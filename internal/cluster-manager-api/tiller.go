package cluster_manager_api

import (
	pb "github.com/samsung-cnct/cluster-manager-api/pkg/api"
	"golang.org/x/net/context"
	"k8s.io/client-go/rest"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/ccutil"
	"io/ioutil"
	"os"
	"k8s.io/client-go/tools/clientcmd"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/cma"
)

func (s *Server) ProvisionTiller(ctx context.Context, in *pb.ProvisionTillerMsg) (*pb.ProvisionTillerReply, error) {
	SetLogger()
	cma.CreateSDSPackageManager(cma.GenerateSDSPackageManager(cma.SDSPackageManagerOptions{
		Name: in.Cluster,
		Namespace: in.Namespace,
		Version: in.Version,
		ClusterWide: in.ClusterWide,
		AdminNamespaces: in.AdminNamespaces,
	}), "default", nil)
	return &pb.ProvisionTillerReply{Ok: true, Message: "Queued Tiller Install"}, nil
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