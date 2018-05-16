package cluster_manager_api


import (
	pb "github.com/samsung-cnct/cluster-manager-api/pkg/api"
	"golang.org/x/net/context"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/ccutil"
	"fmt"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/cma"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/k8sutil"
)

func (s *Server) CreateCluster(ctx context.Context, in *pb.CreateClusterMsg) (*pb.CreateClusterReply, error) {
	options := cma.SDSClusterOptions{Name: in.Name, Provider: "aws"}
	_, err := cma.CreateSDSCluster(cma.GenerateSDSCluster(options), "default", nil)
	if err == nil {
		return &pb.CreateClusterReply{Ok: true, Status: "Creating"}, nil
	} else if k8sutil.IsResourceAlreadyExistsError(err) {
		return &pb.CreateClusterReply{Ok: true, Status: "Cluster already exists"}, nil
	} else {
		return &pb.CreateClusterReply{Ok: false, Status: "Could not create cluster, reason is " + fmt.Sprintf("%s", err)}, nil
	}
}

func (s *Server) GetCluster(ctx context.Context, in *pb.GetClusterMsg) (*pb.GetClusterReply, error) {
	cluster, err := ccutil.GetKrakenCluster(in.Name, "default", nil)
	if err != nil {
		return &pb.GetClusterReply{Ok: false, Status: fmt.Sprintf("%v", err)}, nil
	}
	return &pb.GetClusterReply{Ok: true, Status: string(cluster.Status.State), Kubeconfig: cluster.Status.Kubeconfig}, nil

}

func (s *Server) DeleteCluster(ctx context.Context, in *pb.DeleteClusterMsg) (*pb.DeleteClusterReply, error) {
	ok, err := ccutil.DeleteKrakenCluster(in.Name, "default", nil)
	if err != nil {
		return &pb.DeleteClusterReply{Ok: ok, Status: fmt.Sprintf("%v", err)}, nil
	}
	return &pb.DeleteClusterReply{Ok: ok, Status: "Deleting"}, nil

}
