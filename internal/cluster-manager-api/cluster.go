package cluster_manager_api

import (
	"fmt"

	pb "github.com/samsung-cnct/cluster-manager-api/pkg/api"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/ccutil"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/cma"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/k8sutil"
	"golang.org/x/net/context"
)

func (s *Server) CreateCluster(ctx context.Context, in *pb.CreateClusterMsg) (*pb.CreateClusterReply, error) {
	options := cma.SDSClusterOptions{
		Name:     in.Name,
		Provider: in.Provider.Name,
		AWS: cma.AWSOptions{
			Region:          in.Provider.Aws.Region,
			SecretAccessKey: in.Provider.Aws.SecretAccessKey,
			SecretKeyId:     in.Provider.Aws.SecretKeyId,
		},
		MaaS: cma.MaaSOptions{
			Endpoint: in.Provider.Maas.Endpoint,
			Username: in.Provider.Maas.Username,
			OAuthKey: in.Provider.Maas.OauthKey,
		},
	}
	_, err := cma.CreateSDSCluster(cma.GenerateSDSCluster(options), "default", nil)
	if err == nil {
		return &pb.CreateClusterReply{Ok: true, Status: "Creating"}, nil
	} else if k8sutil.IsResourceAlreadyExistsError(err) {
		return &pb.CreateClusterReply{Ok: true, Status: "Cluster already exists"}, nil
	} else {
		return &pb.CreateClusterReply{Ok: false, Status: "Could not create cluster, reason is " + fmt.Sprintf("%s", err)}, nil
	}
}

func (s *Server) GetCluster(ctx context.Context, in *pb.GetClusterMsg) (*pb.GetClusterReply, error) {
	cluster, err := ccutil.GetKrakenCluster(in.Name, "default", nil)
	if err != nil {
		return &pb.GetClusterReply{Ok: false, Status: fmt.Sprintf("%v", err)}, nil
	}
	return &pb.GetClusterReply{Ok: true, Status: string(cluster.Status.State), Kubeconfig: cluster.Status.Kubeconfig}, nil

}

func (s *Server) DeleteCluster(ctx context.Context, in *pb.DeleteClusterMsg) (*pb.DeleteClusterReply, error) {
	ok, err := cma.DeleteSDSCluster(in.Name, "default", nil)

	// Shouldn't be needed, but just doing it for now
	ok, err = ccutil.DeleteKrakenCluster(in.Name, "default", nil)
	if err != nil {
		return &pb.DeleteClusterReply{Ok: ok, Status: fmt.Sprintf("%v", err)}, nil
	}
	return &pb.DeleteClusterReply{Ok: ok, Status: "Deleting"}, nil
}

func (s *Server) GetClusterList(ctx context.Context, in *pb.GetClusterListMsg) (reply *pb.GetClusterListReply, err error) {
	reply = &pb.GetClusterListReply{}
	list, err := cma.ListSDSClusters("default", nil)
	if err != nil {
		reply.Ok = false
		return
	}
	reply.Ok = true
	for _, cluster := range list {
		reply.Clusters = append(reply.Clusters, &pb.ClusterListItem{Name: cluster.Name, Status: string(cluster.Status.Phase)})
	}
	return
}
