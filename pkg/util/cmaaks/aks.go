package cmaaks

import (
	"context"
	"crypto/tls"

	pb "github.com/samsung-cnct/cma-aks/pkg/generated/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func CreateNewClient(hostname string, insecure bool) (ClientInterface, error) {
	output := Client{}
	err := output.CreateNewClient(hostname, insecure)
	if err != nil {
		return nil, err
	}
	return &output, err
}

func (a *Client) CreateNewClient(hostname string, insecure bool) error {
	var err error
	if insecure {
		// This is for non TLS traffic
		a.conn, err = grpc.Dial(hostname, grpc.WithInsecure())
		if err != nil {
			return err
		}
	} else {
		// If TLS is enabled, we're going to create credentials, also using built in certificates
		var tlsConf tls.Config
		creds := credentials.NewTLS(&tlsConf)

		a.conn, err = grpc.Dial(hostname, grpc.WithTransportCredentials(creds))
		if err != nil {
			return err
		}
	}
	a.client = pb.NewClusterClient(a.conn)
	return nil
}

func (a *Client) Close() error {
	return a.conn.Close()
}

func (a *Client) SetClient(client pb.ClusterClient) {
	a.client = client
}

func (a *Client) CreateCluster(input CreateClusterInput) (CreateClusterOutput, error) {
	var instanceGroups []*pb.CreateClusterAKSSpec_AKSInstanceGroup
	for _, j := range input.Azure.InstanceGroups {
		instanceGroups = append(instanceGroups, &pb.CreateClusterAKSSpec_AKSInstanceGroup{
			Name:        j.Name,
			Type:        j.Type,
			MinQuantity: int32(j.MinQuantity),
		})
	}
	result, err := a.client.CreateCluster(context.Background(), &pb.CreateClusterMsg{
		Name: input.Name,
		Provider: &pb.CreateClusterProviderSpec{
			Name:       AKSProvider,
			K8SVersion: input.K8SVersion,
			Azure: &pb.CreateClusterAKSSpec{
				Location: input.Azure.Location,
				Credentials: &pb.AzureCredentials{
					AppId:          input.Azure.Credentials.AppID,
					Tenant:         input.Azure.Credentials.Tenant,
					Password:       input.Azure.Credentials.Password,
					SubscriptionId: input.Azure.Credentials.SubscriptionID,
				},
				ClusterAccount: &pb.AzureClusterServiceAccount{
					ClientId:     input.Azure.ClusterServiceAccount.ClientID,
					ClientSecret: input.Azure.ClusterServiceAccount.ClientSecret,
				},
				InstanceGroups: instanceGroups,
			},
			HighAvailability: input.HighAvailability,
			NetworkFabric:    input.NetworkFabric,
		},
	})
	if err != nil {
		return CreateClusterOutput{}, err
	}
	output := CreateClusterOutput{
		Cluster: ClusterItem{
			ID:     result.Cluster.Id,
			Name:   result.Cluster.Name,
			Status: result.Cluster.Status.String(),
		},
	}
	return output, nil
}

func (a *Client) GetCluster(input GetClusterInput) (GetClusterOutput, error) {
	result, err := a.client.GetCluster(context.Background(), &pb.GetClusterMsg{
		Name: input.Name,
		Credentials: &pb.AzureCredentials{
			AppId:          input.Credentials.AppID,
			Tenant:         input.Credentials.Tenant,
			Password:       input.Credentials.Password,
			SubscriptionId: input.Credentials.SubscriptionID,
		},
	})
	if err != nil {
		return GetClusterOutput{}, err
	}
	output := GetClusterOutput{
		Cluster: ClusterDetailItem{
			ID:         result.Cluster.Id,
			Name:       result.Cluster.Name,
			Status:     result.Cluster.Status.String(),
			Kubeconfig: result.Cluster.Kubeconfig,
		},
	}
	return output, nil
}

func (a *Client) DeleteCluster(input DeleteClusterInput) (DeleteClusterOutput, error) {
	result, err := a.client.DeleteCluster(context.Background(), &pb.DeleteClusterMsg{
		Name: input.Name,
		Credentials: &pb.AzureCredentials{
			AppId:          input.Credentials.AppID,
			Tenant:         input.Credentials.Tenant,
			Password:       input.Credentials.Password,
			SubscriptionId: input.Credentials.SubscriptionID,
		},
	})
	if err != nil {
		return DeleteClusterOutput{}, err
	}
	output := DeleteClusterOutput{
		Status: result.Status.String(),
	}
	return output, nil
}

func (a *Client) ListClusters(input ListClusterInput) (ListClusterOutput, error) {
	var clusters []ClusterItem
	result, err := a.client.GetClusterList(context.Background(), &pb.GetClusterListMsg{
		Credentials: &pb.AzureCredentials{
			AppId:          input.Credentials.AppID,
			Tenant:         input.Credentials.Tenant,
			Password:       input.Credentials.Password,
			SubscriptionId: input.Credentials.SubscriptionID,
		},
	})
	if err != nil {
		return ListClusterOutput{}, err
	}
	for _, j := range result.Clusters {
		clusters = append(clusters, ClusterItem{
			ID:     j.Id,
			Name:   j.Name,
			Status: j.Status.String(),
		})
	}
	output := ListClusterOutput{
		Clusters: clusters,
	}
	return output, nil
}

func (a *Client) GetClusterUpgrades(input GetClusterUpgradesInput) (output GetClusterUpgradesOutput, err error) {
	output = GetClusterUpgradesOutput{}
	result, err := a.client.GetClusterUpgrades(context.Background(), &pb.GetClusterUpgradesMsg{
		Name: input.Name,
		Credentials: &pb.AzureCredentials{
			AppId:          input.Credentials.AppID,
			Tenant:         input.Credentials.Tenant,
			Password:       input.Credentials.Password,
			SubscriptionId: input.Credentials.SubscriptionID,
		},
	})
	if err != nil {
		return
	}

	// Processing the output
	for _, j := range result.Upgrades {
		output.Versions = append(output.Versions, j.Version)
	}
	return
}

func (a *Client) ClusterUpgrade(input ClusterUpgradeInput) (output ClusterUpgradeOutput, err error) {
	output = ClusterUpgradeOutput{}
	_, err = a.client.UpgradeCluster(context.Background(), &pb.UpgradeClusterMsg{
		Name: input.Name,
		Provider: &pb.UpgradeClusterProviderSpec{
			Name:       AKSProvider,
			K8SVersion: input.Version,
			Azure: &pb.UpgradeClusterAKSSpec{
				Credentials: &pb.AzureCredentials{
					AppId:          input.Credentials.AppID,
					Tenant:         input.Credentials.Tenant,
					Password:       input.Credentials.Password,
					SubscriptionId: input.Credentials.SubscriptionID,
				},
			},
		},
	})
	if err != nil {
		return
	}

	return
}

func (a *Client) GetClusterNodeCount(input GetClusterNodeCountInput) (output GetClusterNodeCountOutput, err error) {
	output = GetClusterNodeCountOutput{}
	result, err := a.client.GetClusterNodeCount(context.Background(), &pb.GetClusterNodeCountMsg{
		Name: input.Name,
		Credentials: &pb.AzureCredentials{
			AppId:          input.Credentials.AppID,
			Tenant:         input.Credentials.Tenant,
			Password:       input.Credentials.Password,
			SubscriptionId: input.Credentials.SubscriptionID,
		},
	})
	if err != nil {
		return
	}
	output = GetClusterNodeCountOutput{
		Name:  result.Name,
		Count: result.Count,
	}
	return output, nil
}

func (a *Client) ScaleCluster(input ScaleClusterInput) (output ScaleClusterOutput, err error) {
	output = ScaleClusterOutput{}
	_, err = a.client.ScaleCluster(context.Background(), &pb.ScaleClusterMsg{
		Name: input.Name,
		Credentials: &pb.AzureCredentials{
			AppId:          input.Credentials.AppID,
			Tenant:         input.Credentials.Tenant,
			Password:       input.Credentials.Password,
			SubscriptionId: input.Credentials.SubscriptionID,
		},
		NodePool: input.NodePool,
		Count:    input.Count,
	})
	if err != nil {
		return output, err
	}

	return output, nil
}
