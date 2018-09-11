package cluster_manager_api

import (
	pb "github.com/samsung-cnct/cluster-manager-api/pkg/generated/api"
	"golang.org/x/net/context"
)

func (s *Server) CreateCluster(ctx context.Context, in *pb.CreateClusterMsg) (*pb.CreateClusterReply, error) {
	azure := in.Provider.GetAzure()
	if azure != nil {
		return azureCreateCluster(in)
	}
	aws := in.Provider.GetAws()
	if aws != nil {

	}
	return &pb.CreateClusterReply{
		Ok: true,
		Cluster: &pb.ClusterItem{
			Id:     "abc123",
			Name:   "dummyName",
			Status: "Placeholder",
		},
	}, nil
}

func (s *Server) GetCluster(ctx context.Context, in *pb.GetClusterMsg) (*pb.GetClusterReply, error) {
	if in.GetAzure().AppId != "" {
		return azureGetCluster(in)
	}
	aws := in.GetAws()
	if aws != nil {

	}
	return &pb.GetClusterReply{
		Ok: true,
		Cluster: &pb.ClusterDetailItem{
			Id:     "abc123",
			Name:   "dummyName",
			Status: "Placeholder",
			Kubeconfig: "undefined still",
		},
	}, nil
}

func (s *Server) DeleteCluster(ctx context.Context, in *pb.DeleteClusterMsg) (*pb.DeleteClusterReply, error) {
	if in.GetAzure().AppId != "" {
		return azureDeleteCluster(in)
	}
	aws := in.GetAws()
	if aws != nil {

	}
	return &pb.DeleteClusterReply{Ok: true, Status: "Deleting, but not really"}, nil
}

func (s *Server) GetClusterList(ctx context.Context, in *pb.GetClusterListMsg) (reply *pb.GetClusterListReply, err error) {
	if in.GetAzure().AppId != "" {
		return azureGetClusterList(in)
	}
	aws := in.GetAws()
	if aws != nil {

	}
	reply = &pb.GetClusterListReply{}
	return
}

