package cmassh

import (
	"context"
	"crypto/tls"
	pb "github.com/samsung-cnct/cma-ssh/pkg/generated/api"
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
	var workerNodes []*pb.SshMachineSpec
	var controlPlaneNodes []*pb.SshMachineSpec

	for _, j := range input.ControlPlaneNodes {
		var labels []*pb.KubernetesLabel
		for _, k := range j.Labels {
			labels = append(labels, &pb.KubernetesLabel{Name: k.Name, Value: k.Value})
		}
		controlPlaneNodes = append(controlPlaneNodes, &pb.SshMachineSpec{
			Host:     j.Host,
			Port:     int32(j.Port),
			Username: j.Username,
			Password: j.Password,
			Labels:   labels,
		})
	}
	for _, j := range input.WorkerNodes {
		var labels []*pb.KubernetesLabel
		for _, k := range j.Labels {
			labels = append(labels, &pb.KubernetesLabel{Name: k.Name, Value: k.Value})
		}
		workerNodes = append(workerNodes, &pb.SshMachineSpec{
			Host:     j.Host,
			Port:     int32(j.Port),
			Username: j.Username,
			Password: j.Password,
			Labels:   labels,
		})
	}
	result, err := a.client.CreateCluster(context.Background(), &pb.CreateClusterMsg{
		Name:              input.Name,
		K8SVersion:        input.K8SVersion,
		ControlPlaneNodes: controlPlaneNodes,
		WorkerNodes:       workerNodes,
		ApiEndpoint:       input.APIEndpoint,
		HighAvailability:  input.HighAvailability,
		NetworkFabric:     input.NetworkFabric,
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
	})
	if err != nil {
		return DeleteClusterOutput{}, err
	}
	output := DeleteClusterOutput{
		Status: result.Status,
	}
	return output, nil
}

func (a *Client) ListClusters(input ListClusterInput) (ListClusterOutput, error) {
	var clusters []ClusterItem
	result, err := a.client.GetClusterList(context.Background(), &pb.GetClusterListMsg{})
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

func (a *Client) AdjustCluster(input AdjustClusterInput) (AdjustClusterOutput, error) {
	var addNodes []*pb.SshMachineSpec
	var removeNodes []*pb.AdjustClusterMsg_SshRemoveMachineSpec

	for _, j := range input.AddNodes {
		var labels []*pb.KubernetesLabel
		for _, k := range j.Labels {
			labels = append(labels, &pb.KubernetesLabel{Name: k.Name, Value: k.Value})
		}
		addNodes = append(addNodes, &pb.SshMachineSpec{
			Host:     j.Host,
			Port:     int32(j.Port),
			Username: j.Username,
			Password: j.Password,
			Labels:   labels,
		})
	}
	for _, j := range input.RemoveNodes {
		removeNodes = append(removeNodes, &pb.AdjustClusterMsg_SshRemoveMachineSpec{
			Host: j.Host,
		})
	}
	_, err := a.client.AdjustClusterNodes(context.Background(), &pb.AdjustClusterMsg{
		Name:        input.Name,
		AddNodes:    addNodes,
		RemoveNodes: removeNodes,
	})
	if err != nil {
		return AdjustClusterOutput{}, err
	}
	return AdjustClusterOutput{}, nil
}

func (a *Client) GetClusterUpgrades(input GetClusterUpgradesInput) (GetClusterUpgradesOutput, error) {
	result, err := a.client.GetUpgradeClusterInformation(context.Background(), &pb.GetUpgradeClusterInformationMsg{
		Name: input.Name,
	})
	if err != nil {
		return GetClusterUpgradesOutput{}, err
	}

	output := GetClusterUpgradesOutput{}
	for _, j := range result.Versions {
		output.Versions = append(output.Versions, j)
	}
	return output, nil
}

func (a *Client) ClusterUpgrade(input ClusterUpgradeInput) (ClusterUpgradeOutput, error) {
	_, err := a.client.UpgradeCluster(context.Background(), &pb.UpgradeClusterMsg{
		Name:    input.Name,
		Version: input.Version,
	})
	return ClusterUpgradeOutput{}, err
}
