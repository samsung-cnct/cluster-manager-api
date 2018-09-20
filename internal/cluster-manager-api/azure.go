package cluster_manager_api

import (
	"fmt"
	"github.com/samsung-cnct/cluster-manager-api/internal/cluster-manager-api/azure"
	pb "github.com/samsung-cnct/cluster-manager-api/pkg/generated/api"
	"golang.org/x/net/context"
)

func (s *Server) UpdateAzureCredentials(ctx context.Context, in *pb.UpdateAzureCredentialsMsg) (*pb.UpdateAzureCredentialsReply, error) {
	if s.azure != nil {
		return s.azure.UpdateCredentials(in)
	} else {
		return nil, fmt.Errorf(azure.NotEnabledErrorMessage)
	}
}
