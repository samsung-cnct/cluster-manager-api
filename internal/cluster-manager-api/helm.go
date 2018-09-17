package cluster_manager_api

import (
	pb "github.com/samsung-cnct/cluster-manager-api/pkg/generated/api"
	"golang.org/x/net/context"
)

// Will install (or reinstall) tiller
func (s *Server) ProvisionTiller(ctx context.Context, in *pb.ProvisionTillerMsg) (*pb.ProvisionTillerReply, error) {
	return &pb.ProvisionTillerReply{}, nil
}

// Will install (or reinstall) helm chart
// This will be destructive if a chart has already been deployed with the same name
func (s *Server) InstallHelmChart(ctx context.Context, in *pb.InstallHelmChartMsg) (*pb.InstallHelmChartReply, error) {
	return &pb.InstallHelmChartReply{}, nil
}

// Will delete deployed helm chart
func (s *Server) DeleteHelmChart(ctx context.Context, in *pb.DeleteHelmChartMsg) (*pb.DeleteHelmChartReply, error) {
	return &pb.DeleteHelmChartReply{}, nil
}
