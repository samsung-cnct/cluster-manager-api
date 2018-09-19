package cluster_manager_api

import (
	pb "github.com/samsung-cnct/cluster-manager-api/pkg/generated/api"
	"golang.org/x/net/context"
)

func (s *Server) UpdateAWSCredentials(ctx context.Context, in *pb.UpdateAWSCredentialsMsg) (*pb.UpdateAWSCredentialsReply, error) {
	return s.aws.UpdateCredentials(in)
}
