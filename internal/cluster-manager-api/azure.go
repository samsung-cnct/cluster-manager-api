package cluster_manager_api

import (
	"fmt"
	pb "github.com/samsung-cnct/cluster-manager-api/pkg/generated/api"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/cmaaks"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/k8sutil/azure"
	"github.com/spf13/viper"
)

func azureGetClient() (cmaaks.AKSClientInterface, error) {
	hostname := viper.GetString("cmaaks-endpoint")
	if hostname == "" {
		return nil, fmt.Errorf("azure support is not enabled")
	}
	insecure := viper.GetBool("cmaaks-insecure")
	return cmaaks.CreateNewClient(hostname, insecure)
}

func azureCreateCluster(in *pb.CreateClusterMsg) (*pb.CreateClusterReply, error) {
	var instanceGroups []cmaaks.InstanceGroup
	client, err := azureGetClient()
	if err != nil {
		return &pb.CreateClusterReply{}, err
	}
	defer client.Close()
	azureSecretClient, err := azurek8sutil.CreateFromDefaults()
	if err != nil {
		return &pb.CreateClusterReply{}, err
	}
	for _, j := range in.Provider.GetAzure().InstanceGroups {
		instanceGroups = append(instanceGroups, cmaaks.InstanceGroup{
			Name:        j.Name,
			Type:        j.Type,
			MinQuantity: int(j.MinQuantity),
			MaxQuantity: int(j.MaxQuantity),
		})
	}
	result, err := client.CreateCluster(cmaaks.CreateClusterInput{
		Name:       in.Name,
		K8SVersion: in.Provider.K8SVersion,
		Azure: cmaaks.AzureSpec{
			Location: in.Provider.GetAzure().Location,
			Credentials: cmaaks.Credentials{
				AppID:          in.Provider.GetAzure().Credentials.AppId,
				Tenant:         in.Provider.GetAzure().Credentials.Tenant,
				Password:       in.Provider.GetAzure().Credentials.Password,
				SubscriptionID: in.Provider.GetAzure().Credentials.SubscriptionId,
			},
			ClusterServiceAccount: cmaaks.ClusterServiceAccount{
				ClientID:     in.Provider.GetAzure().ClusterAccount.ClientId,
				ClientSecret: in.Provider.GetAzure().ClusterAccount.ClientSecret,
			},
			InstanceGroups: instanceGroups,
		},
		HighAvailability: in.Provider.HighAvailability,
		NetworkFabric:    in.Provider.NetworkFabric,
	})
	if err != nil {
		return &pb.CreateClusterReply{}, err
	}

	// Cluster Creation was successful, going to save the credentials
	err = azureSecretClient.CreateCredentials(in.Name, azurek8sutil.Credentials{
		AppID:          in.Provider.GetAzure().Credentials.AppId,
		Tenant:         in.Provider.GetAzure().Credentials.Tenant,
		Password:       in.Provider.GetAzure().Credentials.Password,
		SubscriptionID: in.Provider.GetAzure().Credentials.SubscriptionId,
	})

	if err != nil {
		// TODO Unsure what to do if we suddenly can't persist the credentails to kubernetes
		// TODO Going to log for now
		logger.Errorf("Could not set AWS credentials into kubernetes, this is bad")
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

func azureGetCluster(in *pb.GetClusterMsg) (*pb.GetClusterReply, error) {
	client, err := azureGetClient()
	if err != nil {
		return &pb.GetClusterReply{}, err
	}
	defer client.Close()
	result, err := client.GetCluster(cmaaks.GetClusterInput{
		Name: in.Name,
		Credentials: cmaaks.Credentials{
			AppID:          in.GetAzure().AppId,
			Tenant:         in.GetAzure().Tenant,
			Password:       in.GetAzure().Password,
			SubscriptionID: in.GetAzure().SubscriptionId,
		},
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

func azureGetClusterList(in *pb.GetClusterListMsg) (*pb.GetClusterListReply, error) {
	var clusters []*pb.ClusterItem
	client, err := azureGetClient()
	if err != nil {
		return &pb.GetClusterListReply{}, err
	}
	defer client.Close()
	result, err := client.ListClusters(cmaaks.ListClusterInput{
		Credentials: cmaaks.Credentials{
			AppID:          in.GetAzure().AppId,
			Tenant:         in.GetAzure().Tenant,
			Password:       in.GetAzure().Password,
			SubscriptionID: in.GetAzure().SubscriptionId,
		},
	})
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

func azureDeleteCluster(in *pb.DeleteClusterMsg) (*pb.DeleteClusterReply, error) {
	client, err := azureGetClient()
	if err != nil {
		return &pb.DeleteClusterReply{}, err
	}
	defer client.Close()
	result, err := client.DeleteCluster(cmaaks.DeleteClusterInput{
		Name: in.Name,
		Credentials: cmaaks.Credentials{
			AppID:          in.GetAzure().AppId,
			Tenant:         in.GetAzure().Tenant,
			Password:       in.GetAzure().Password,
			SubscriptionID: in.GetAzure().SubscriptionId,
		},
	})
	if err != nil {
		return &pb.DeleteClusterReply{}, err
	}
	return &pb.DeleteClusterReply{
		Ok:     true,
		Status: result.Status,
	}, nil
}
