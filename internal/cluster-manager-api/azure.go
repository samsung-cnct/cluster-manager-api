package cluster_manager_api

import (
	pb "github.com/samsung-cnct/cluster-manager-api/pkg/generated/api"
	"golang.org/x/net/context"
)

func (s *Server) UpdateAzureCredentials(ctx context.Context, in *pb.UpdateAzureCredentialsMsg) (*pb.UpdateAzureCredentialsReply, error) {
	return s.azure.UpdateCredentials(in)
}
