package cluster_manager_api

import (
	"fmt"
	pb "github.com/samsung-cnct/cluster-manager-api/pkg/generated/api"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/cmavmware"
	"github.com/spf13/viper"
)

func vmwareGetClient() (cmavmware.VMWareClientInterface, error) {
	hostname := viper.GetString("cmavmware-endpoint")
	if hostname == "" {
		return nil, fmt.Errorf("vmware support is not enabled")
	}
	insecure := viper.GetBool("cmavmware-insecure")
	return cmavmware.CreateNewClient(hostname, insecure)
}

func vmwareCreateCluster(in *pb.CreateClusterMsg) (*pb.CreateClusterReply, error) {
	var machines []cmavmware.MachineSpec
	client, err := vmwareGetClient()
	if err != nil {
		return &pb.CreateClusterReply{}, err
	}
	defer client.Close()
	for _, j := range in.Provider.GetVmware().Machines {
		machines = append(machines, cmavmware.MachineSpec{
			Host:                j.Host,
			Username:            j.Username,
			Port:                int(j.Port),
			ControlPlaneVersion: j.ControlPlaneVersion,
		})
	}
	result, err := client.CreateCluster(cmavmware.CreateClusterInput{
		Name:       in.Name,
		K8SVersion: in.Provider.K8SVersion,
		VMWare: cmavmware.VMWareSpec{
			Namespace:  in.Provider.GetVmware().Namespace,
			PrivateKey: in.Provider.GetVmware().PrivateKey,
			Machines:   machines,
		},
		HighAvailability: in.Provider.HighAvailability,
		NetworkFabric:    in.Provider.NetworkFabric,
	})
	if err != nil {
		return &pb.CreateClusterReply{}, err
	}
	return &pb.CreateClusterReply{
		Ok: true,
		Cluster: &pb.ClusterItem{
			Id:     result.Cluster.ID,
			Name:   result.Cluster.Name,
			Status: result.Cluster.Status,
		},
	}, nil
}

func vmwareGetCluster(in *pb.GetClusterMsg) (*pb.GetClusterReply, error) {
	client, err := vmwareGetClient()
	if err != nil {
		return &pb.GetClusterReply{}, err
	}
	defer client.Close()
	result, err := client.GetCluster(cmavmware.GetClusterInput{
		Name: in.Name,
	})
	if err != nil {
		return &pb.GetClusterReply{}, err
	}
	return &pb.GetClusterReply{
		Ok: true,
		Cluster: &pb.ClusterDetailItem{
			Id:         result.Cluster.ID,
			Name:       result.Cluster.Name,
			Status:     result.Cluster.Status,
			Kubeconfig: result.Cluster.Kubeconfig,
		},
	}, nil
}

func vmwareGetClusterList(in *pb.GetClusterListMsg) (*pb.GetClusterListReply, error) {
	var clusters []*pb.ClusterItem
	client, err := vmwareGetClient()
	if err != nil {
		return &pb.GetClusterListReply{}, err
	}
	defer client.Close()
	result, err := client.ListClusters(cmavmware.ListClusterInput{})
	if err != nil {
		return &pb.GetClusterListReply{}, err
	}
	for _, j := range result.Clusters {
		clusters = append(clusters, &pb.ClusterItem{
			Id:     j.ID,
			Name:   j.Name,
			Status: j.Status,
		})
	}
	return &pb.GetClusterListReply{
		Ok:       true,
		Clusters: clusters,
	}, nil
}

func vmwareDeleteCluster(in *pb.DeleteClusterMsg) (*pb.DeleteClusterReply, error) {
	client, err := vmwareGetClient()
	if err != nil {
		return &pb.DeleteClusterReply{}, err
	}
	defer client.Close()
	result, err := client.DeleteCluster(cmavmware.DeleteClusterInput{
		Name: in.Name,
	})
	if err != nil {
		return &pb.DeleteClusterReply{}, err
	}
	return &pb.DeleteClusterReply{
		Ok:     true,
		Status: result.Status,
	}, nil
}
