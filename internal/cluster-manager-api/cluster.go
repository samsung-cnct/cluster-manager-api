package cluster_manager_api


import (
	pb "github.com/samsung-cnct/cluster-manager-api/pkg/api"
	"golang.org/x/net/context"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/ccutil"
	"fmt"
)

func (s *Server) CreateCluster(ctx context.Context, in *pb.CreateClusterMsg) (*pb.CreateClusterReply, error) {
	ccutil.CreateKrakenCluster(
		ccutil.GenerateKrakenCluster(
			ccutil.KrakenClusterOptions{Name: in.Name},
		), "default", nil)
	return &pb.CreateClusterReply{Ok: true, Status: "Creating"}, nil
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
