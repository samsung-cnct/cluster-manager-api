package cluster_manager_api

import (
	"fmt"
	"github.com/samsung-cnct/cluster-manager-api/internal/cluster-manager-api/azure"
	pb "github.com/samsung-cnct/cluster-manager-api/pkg/generated/api"
	"golang.org/x/net/context"
)

func (s *Server) UpdateAWSCredentials(ctx context.Context, in *pb.UpdateAWSCredentialsMsg) (*pb.UpdateAWSCredentialsReply, error) {
	if s.azure != nil {
		return s.aws.UpdateCredentials(in)
	} else {
		return nil, fmt.Errorf(azure.NotEnabledErrorMessage)
	}

}
