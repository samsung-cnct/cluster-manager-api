package cmaaks

import (
	"context"
	"crypto/tls"
	pb "github.com/samsung-cnct/cma-aks/pkg/generated/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type AKSClient struct {
	conn   *grpc.ClientConn
	client pb.ClusterClient
}

func CreateNewClient(hostname string, insecure bool) (AKSClientInterface, error) {
	output := AKSClient{}
	err := output.CreateNewClient(hostname, insecure)
	if err != nil {
		return nil, err
	}
	return &output, err
}

func (a *AKSClient) CreateNewClient(hostname string, insecure bool) error {
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

func (a *AKSClient) Close() error {
	return a.conn.Close()
}

func (a *AKSClient) SetClient(client pb.ClusterClient) {
	a.client = client
}

func (a *AKSClient) CreateCluster(input CreateClusterInput) (CreateClusterOutput, error) {
	var instanceGroups []*pb.CreateClusterAKSSpec_AKSInstanceGroup
	for _, j := range input.Azure.InstanceGroups {
		instanceGroups = append(instanceGroups, &pb.CreateClusterAKSSpec_AKSInstanceGroup{
			Name:        j.Name,
			Type:        j.Type,
			MinQuantity: int32(j.MinQuantity),
			MaxQuantity: int32(j.MaxQuantity),
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
			Status: result.Cluster.Status,
		},
	}
	return output, nil
}

func (a *AKSClient) GetCluster(input GetClusterInput) (GetClusterOutput, error) {
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
			Status:     result.Cluster.Status,
			Kubeconfig: result.Cluster.Kubeconfig,
		},
	}
	return output, nil
}

func (a *AKSClient) DeleteCluster(input DeleteClusterInput) (DeleteClusterOutput, error) {
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
		Status: result.Status,
	}
	return output, nil
}

func (a *AKSClient) ListClusters(input ListClusterInput) (ListClusterOutput, error) {
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
			Status: j.Status,
		})
	}
	output := ListClusterOutput{
		Clusters: clusters,
	}
	return output, nil
}
