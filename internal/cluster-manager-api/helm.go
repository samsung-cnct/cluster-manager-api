package cluster_manager_api

import (
	"fmt"

	pb "github.com/samsung-cnct/cluster-manager-api/pkg/api"
	"golang.org/x/net/context"
)

func (s *Server) InstallHelmChart(ctx context.Context, in *pb.InstallHelmChartMsg) (*pb.InstallHelmChartReply, error) {
	SetLogger()
	config, err := retrieveClusterRestConfig(in.Cluster, "default", nil)
	if err != nil {
		return &pb.InstallHelmChartReply{Ok: false, Message: fmt.Sprintf("%v", err)}, nil
	}

	// TODO Have not implemented this yet - just a stub
	if config == nil {
		return &pb.InstallHelmChartReply{Ok: true, Message: "I am a stub message"}, nil
	}
	return &pb.InstallHelmChartReply{Ok: false, Message: "I am a stub message"}, nil
}

func (s *Server) DeleteHelmChart(ctx context.Context, in *pb.DeleteHelmChartMsg) (*pb.DeleteHelmChartReply, error) {
	SetLogger()
	config, err := retrieveClusterRestConfig(in.Cluster, "default", nil)
	if err != nil {
		return &pb.DeleteHelmChartReply{Ok: false, Message: fmt.Sprintf("%v", err)}, nil
	}

	// TODO Have not implemented this yet - just a stub
	if config == nil {
		return &pb.DeleteHelmChartReply{Ok: true, Message: "I am a stub message"}, nil
	}
	return &pb.DeleteHelmChartReply{Ok: false, Message: "I am a stub message"}, nil
}
