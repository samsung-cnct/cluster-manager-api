package cluster_manager_api

import (
	"golang.org/x/net/context"
	pb "github.com/samsung-cnct/cluster-manager-api/pkg/api"
)

type Server struct {}

func (s *server) HellWorld(ctx context.Context, in *pb.HelloWorldMsg) (*pb.HelloWorldReply, error) {
	return &pb.HelloWorldReply{Message: "Hello " + in.Name}, nil
}