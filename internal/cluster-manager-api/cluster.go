package cluster_manager_api

import (
	"fmt"
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
		return awsCreateCluster(in)
	}
	vmware := in.Provider.GetVmware()
	if vmware != nil {
		return vmwareCreateCluster(in)
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
	if in.GetAzure() != nil && in.GetAzure().AppId != "" {
		return azureGetCluster(in)
	}
	if in.GetAws() != nil && in.GetAws().Region != "" {
		return awsGetCluster(in)
	} else {
		return vmwareGetCluster(in)
	}
	return &pb.GetClusterReply{
		Ok: true,
		Cluster: &pb.ClusterDetailItem{
			Id:         "abc123",
			Name:       "dummyName",
			Status:     "Placeholder",
			Kubeconfig: "undefined still",
		},
	}, nil
}

func (s *Server) DeleteCluster(ctx context.Context, in *pb.DeleteClusterMsg) (*pb.DeleteClusterReply, error) {
	if in.GetAzure() != nil && in.GetAzure().AppId != "" {
		return azureDeleteCluster(in)
	}
	if in.GetAws() != nil && in.GetAws().Region != "" {
		return awsDeleteCluster(in)
	} else {
		return vmwareDeleteCluster(in)
	}
	return &pb.DeleteClusterReply{Ok: true, Status: "Deleting, but not really"}, nil
}

func (s *Server) GetClusterList(ctx context.Context, in *pb.GetClusterListMsg) (reply *pb.GetClusterListReply, err error) {
	if in.GetAzure() != nil && in.GetAzure().AppId != "" {
		return azureGetClusterList(in)
	}
	if in.GetAws() != nil && in.GetAws().Region != "" {
		return awsGetClusterList(in)
	} else {
		return vmwareGetClusterList(in)
	}
	reply = &pb.GetClusterListReply{}
	return
}

// Will return upgrade options for a given cluster
func (s *Server) GetUpgradeClusterInformation(ctx context.Context, in *pb.GetUpgradeClusterInformationMsg) (*pb.GetUpgradeClusterInformationReply, error) {
	//TODO this is just temporary stuff until cma-vmware catches up
	if in.Provider == "vmware" {
		return &pb.GetUpgradeClusterInformationReply{Ok: true, Versions: []string{"1.10.8", "1.11.3"}}, nil
	}

	return &pb.GetUpgradeClusterInformationReply{}, fmt.Errorf("upgrades not supported yet")
}

// Will attempt to upgrade a cluster
func (s *Server) UpgradeCluster(ctx context.Context, in *pb.UpgradeClusterMsg) (*pb.UpgradeClusterReply, error) {
	//TODO this is just temporary stuff until cma-vmware catches up
	if in.Provider == "vmware" {
		return &pb.UpgradeClusterReply{Ok: false}, nil
	}
	return &pb.UpgradeClusterReply{}, fmt.Errorf("upgrades not supported yet")
}

// Will adjust a provision a cluster
func (s *Server) AdjustClusterNodes(ctx context.Context, in *pb.AdjustClusterMsg) (*pb.AdjustClusterReply, error) {
	vmware := in.GetVmware()
	if vmware != nil {
		return vmwareAdjustCluster(in)
	}
	return &pb.AdjustClusterReply{
		Ok: false,
	}, nil

}
