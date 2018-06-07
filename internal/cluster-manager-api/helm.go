package cluster_manager_api

import (
	"fmt"

	pb "github.com/samsung-cnct/cluster-manager-api/pkg/generated/api"
	"github.com/samsung-cnct/cma-operator/pkg/util/cma"
	"golang.org/x/net/context"
)

func (s *Server) InstallHelmChart(ctx context.Context, in *pb.InstallHelmChartMsg) (*pb.InstallHelmChartReply, error) {
	SetLogger()
	res, err := cma.CreateSDSApplication(cma.GenerateSDSApplication(cma.SDSApplicationOptions{
		Name:           in.Chart.Name,
		Namespace:      in.Chart.Namespace,
		Values:         in.Chart.Values,
		PackageManager: in.PackageManagerName,
		Chart: cma.Chart{
			Name: in.Chart.Chart,
			Repository: cma.ChartRepository{
				Name: in.Repo.Name,
				URL:  in.Repo.Url,
			},
			Version: in.Chart.Version,
		},
	}), "default", nil)

	message := "Successfully installed " + in.Chart.Chart + " as " + in.Chart.Name + " in " + in.Chart.Namespace
	if err != nil {
		message = fmt.Sprintf("%v", err)
	}

	return &pb.InstallHelmChartReply{Ok: res, Message: message}, nil
}

func (s *Server) DeleteHelmChart(ctx context.Context, in *pb.DeleteHelmChartMsg) (*pb.DeleteHelmChartReply, error) {
	SetLogger()
	res, err := cma.DeleteSDSApplication(in.Name, in.Namespace, nil)

	message := "Successfully deleted " + in.Name
	if err != nil {
		message = fmt.Sprintf("%v", err)
	}

	return &pb.DeleteHelmChartReply{Ok: res, Message: message}, nil
}
