package cluster_manager_api

import (
	"fmt"
	pb "github.com/samsung-cnct/cluster-manager-api/pkg/generated/api"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/k8sutil/cma"
	"golang.org/x/net/context"
)

// Will install (or reinstall) tiller
func (s *Server) ProvisionTiller(ctx context.Context, in *pb.ProvisionTillerMsg) (*pb.ProvisionTillerReply, error) {
	// TODO: Clean this up better!
	// Ensuring the cluster is around
	cluster, err := s.GetCluster(ctx, &pb.GetClusterMsg{
		Name:     in.Cluster,
		Azure:    in.Azure,
		Aws:      in.Aws,
		Provider: in.Provider,
	})
	// Check for errors
	if err != nil {
		return &pb.ProvisionTillerReply{}, err
	}
	if cluster.Cluster.Kubeconfig == "" {
		return &pb.ProvisionTillerReply{}, fmt.Errorf("cluster %s is not ready for a request to provision tiller", in.Cluster)
	}

	err = s.cmak8s.UpdateOrCreatePackageManager(in.Name, cmak8sutil.PackageManager{
		AdminNamespaces: in.AdminNamespaces,
		CallbackURL:     in.Callback.Url,
		Cluster:         in.Cluster,
		ClusterWide:     in.ClusterWide,
		Image:           in.Image,
		Namespace:       in.Namespace,
		RequestID:       in.Callback.RequestId,
		Version:         in.Version,
	})
	if err != nil {
		return &pb.ProvisionTillerReply{}, err
	}
	return &pb.ProvisionTillerReply{Ok: true, Message: "tiller provisioned"}, nil
}

// Will install (or reinstall) helm chart
// This will be destructive if a chart has already been deployed with the same name
func (s *Server) InstallHelmChart(ctx context.Context, in *pb.InstallHelmChartMsg) (*pb.InstallHelmChartReply, error) {
	// TODO: Clean this up better!
	// Ensuring the cluster is around
	cluster, err := s.GetCluster(ctx, &pb.GetClusterMsg{
		Name:     in.Cluster,
		Azure:    in.Azure,
		Aws:      in.Aws,
		Provider: in.Provider,
	})
	// Check for errors
	if err != nil {
		return &pb.InstallHelmChartReply{}, err
	}
	if cluster.Cluster.Kubeconfig == "" {
		return &pb.InstallHelmChartReply{}, fmt.Errorf("cluster %s is not ready for a request to provision tiller", in.Cluster)
	}

	err = s.cmak8s.UpdateOrCreateApplication(in.Chart.Name, cmak8sutil.Application{
		CallbackURL: in.Callback.Url,
		Cluster:     in.Cluster,
		Chart: cmak8sutil.Chart{
			Name: in.Chart.Chart,
			Repository: cmak8sutil.ChartRepository{
				Name: in.Chart.Repo.Name,
				URL:  in.Chart.Repo.Url,
			},
			ChartPayload: in.Chart.ChartPayload,
			Version:      in.Chart.Version,
		},
		Namespace:      in.Chart.Namespace,
		PackageManager: in.PackageManger,
		RequestID:      in.Callback.RequestId,
	})
	if err != nil {
		return &pb.InstallHelmChartReply{}, err
	}
	return &pb.InstallHelmChartReply{Ok: true, Message: "installing helm chart"}, nil
}

// Will delete deployed helm chart
func (s *Server) DeleteHelmChart(ctx context.Context, in *pb.DeleteHelmChartMsg) (*pb.DeleteHelmChartReply, error) {
	// TODO: Clean this up better!
	// Ensuring the cluster is around
	cluster, err := s.GetCluster(ctx, &pb.GetClusterMsg{
		Name:     in.Cluster,
		Azure:    in.Azure,
		Aws:      in.Aws,
		Provider: in.Provider,
	})
	// Check for errors
	if err != nil {
		return &pb.DeleteHelmChartReply{}, err
	}
	if cluster.Cluster.Kubeconfig == "" {
		return &pb.DeleteHelmChartReply{}, fmt.Errorf("cluster %s is not ready for a request to provision tiller", in.Cluster)
	}

	err = s.cmak8s.DeleteApplication(in.Chart, in.PackageManager, in.Cluster)
	if err != nil {
		return &pb.DeleteHelmChartReply{}, err
	}

	return &pb.DeleteHelmChartReply{Ok: true, Message: "deleting helm chart"}, nil
}
