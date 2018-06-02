package cluster_manager_api

import (
	pb "github.com/samsung-cnct/cluster-manager-api/pkg/generated/api"
	"golang.org/x/net/context"

	"github.com/samsung-cnct/cluster-manager-api/pkg/version"
)

func (s *Server) GetVersionInformation(ctx context.Context, in *pb.GetVersionMsg) (*pb.GetVersionReply, error) {
	SetLogger()
	versionInformation := version.Get()
	reply := &pb.GetVersionReply{
		Ok: true,
		VersionInformation: &pb.GetVersionReply_VersionInformation{
			GitVersion: versionInformation.GitVersion,
			GitCommit: versionInformation.GitCommit,
			GitTreeState: versionInformation.GitTreeState,
			BuildDate: versionInformation.BuildDate,
			GoVersion: versionInformation.GoVersion,
			Compiler: versionInformation.Compiler,
			Platform: versionInformation.Platform,
		},
	}
	return reply, nil
}