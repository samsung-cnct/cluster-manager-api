package cluster_manager_api

import (
	"fmt"
	"github.com/samsung-cnct/cluster-manager-api/internal/cluster-manager-api/aws"
	"github.com/samsung-cnct/cluster-manager-api/internal/cluster-manager-api/azure"
	"github.com/samsung-cnct/cluster-manager-api/internal/cluster-manager-api/vmware"
	pb "github.com/samsung-cnct/cluster-manager-api/pkg/generated/api"
	"golang.org/x/net/context"
)

func (s *Server) CreateCluster(ctx context.Context, in *pb.CreateClusterMsg) (*pb.CreateClusterReply, error) {
	if in.Provider.GetAzure() != nil {
		return s.azure.CreateCluster(in)
	}
	if in.Provider.GetAws() != nil {
		return s.aws.CreateCluster(in)
	}
	if in.Provider.GetVmware() != nil {
		return s.vmware.CreateCluster(in)
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
	switch in.Provider {
	case pb.Provider_aws:
		if s.aws != nil {
			return s.aws.GetCluster(in)
		} else {
			return nil, fmt.Errorf(aws.NotEnabledErrorMessage)
		}
	case pb.Provider_azure:
		if s.azure != nil {
			return s.azure.GetCluster(in)
		} else {
			return nil, fmt.Errorf(azure.NotEnabledErrorMessage)
		}
	case pb.Provider_vmware:
		if s.vmware != nil {
			return s.vmware.GetCluster(in)
		} else {
			return nil, fmt.Errorf(vmware.NotEnabledErrorMessage)
		}
	}
	return &pb.GetClusterReply{Ok: false}, fmt.Errorf("no provider selected")
}

func (s *Server) DeleteCluster(ctx context.Context, in *pb.DeleteClusterMsg) (*pb.DeleteClusterReply, error) {
	switch in.Provider {
	case pb.Provider_aws:
		if s.aws != nil {
			return s.aws.DeleteCluster(in)
		} else {
			return nil, fmt.Errorf(aws.NotEnabledErrorMessage)
		}
	case pb.Provider_azure:
		if s.azure != nil {
			return s.azure.DeleteCluster(in)
		} else {
			return nil, fmt.Errorf(azure.NotEnabledErrorMessage)
		}
	case pb.Provider_vmware:
		if s.vmware != nil {
			return s.vmware.DeleteCluster(in)
		} else {
			return nil, fmt.Errorf(vmware.NotEnabledErrorMessage)
		}
	}
	return &pb.DeleteClusterReply{Ok: false}, fmt.Errorf("no provider selected")
}

func (s *Server) GetClusterList(ctx context.Context, in *pb.GetClusterListMsg) (reply *pb.GetClusterListReply, err error) {
	switch in.Provider {
	case pb.Provider_aws:
		if s.aws != nil {
			return s.aws.GetClusterList(in)
		} else {
			return nil, fmt.Errorf(aws.NotEnabledErrorMessage)
		}
	case pb.Provider_azure:
		if s.azure != nil {
			return s.azure.GetClusterList(in)
		} else {
			return nil, fmt.Errorf(azure.NotEnabledErrorMessage)
		}
	case pb.Provider_vmware:
		if s.vmware != nil {
			return s.vmware.GetClusterList(in)
		} else {
			return nil, fmt.Errorf(vmware.NotEnabledErrorMessage)
		}
	}
	reply = &pb.GetClusterListReply{}
	return
}

// Will return upgrade options for a given cluster
func (s *Server) GetUpgradeClusterInformation(ctx context.Context, in *pb.GetUpgradeClusterInformationMsg) (*pb.GetUpgradeClusterInformationReply, error) {
	switch in.Provider {
	case pb.Provider_azure:
		if s.azure != nil {
			return s.azure.GetClusterUpgrades(in)
		} else {
			return nil, fmt.Errorf(azure.NotEnabledErrorMessage)
		}
	case pb.Provider_vmware:
		if s.vmware != nil {
			return s.vmware.GetClusterUpgrades(in)
		} else {
			return nil, fmt.Errorf(vmware.NotEnabledErrorMessage)
		}
	}

	return &pb.GetUpgradeClusterInformationReply{}, fmt.Errorf("upgrades not supported yet")
}

// Will attempt to upgrade a cluster
func (s *Server) UpgradeCluster(ctx context.Context, in *pb.UpgradeClusterMsg) (*pb.UpgradeClusterReply, error) {
	switch in.Provider {
	case pb.Provider_azure:
		if s.azure != nil {
			return s.azure.ClusterUpgrade(in)
		} else {
			return nil, fmt.Errorf(azure.NotEnabledErrorMessage)
		}
	case pb.Provider_vmware:
		if s.vmware != nil {
			return s.vmware.ClusterUpgrade(in)
		} else {
			return nil, fmt.Errorf(vmware.NotEnabledErrorMessage)
		}
	}
	return &pb.UpgradeClusterReply{}, fmt.Errorf("upgrades not supported yet")
}

// Will adjust a provision a cluster
func (s *Server) AdjustClusterNodes(ctx context.Context, in *pb.AdjustClusterMsg) (*pb.AdjustClusterReply, error) {
	switch in.Provider {
	case pb.Provider_vmware:
		if s.vmware != nil {
			return s.vmware.AdjustCluster(in)
		} else {
			return nil, fmt.Errorf(vmware.NotEnabledErrorMessage)
		}
	}
	return &pb.AdjustClusterReply{
		Ok: false,
	}, fmt.Errorf("upgrades not supported yet")

}
