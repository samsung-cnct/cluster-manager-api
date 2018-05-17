package main

import (
	"flag"
	"fmt"
	"net"
	"strings"
	"sync"

	"github.com/samsung-cnct/cluster-manager-api/pkg/controllers/cluster"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/k8sutil"
	"github.com/soheilhy/cmux"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/juju/loggo"
	"github.com/samsung-cnct/cluster-manager-api/pkg/apiserver"
	ccworkqueue "github.com/samsung-cnct/cluster-manager-api/pkg/util/ccutil/workqueue"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/cma"
	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/rest"
	"github.com/samsung-cnct/cluster-manager-api/pkg/controllers/sds-package-manager"
	"github.com/samsung-cnct/cluster-manager-api/pkg/controllers/sds-application"
)

var (
	logger loggo.Logger
	config *rest.Config
)

func main() {
	var err error
	logger := util.GetModuleLogger("cmd.cma-operator", loggo.INFO)
	viperInit()

	// get flags
	portNumber := viper.GetInt("port")
	kubeconfigLocation := viper.GetString("kubeconfig")

	// Debug for now
	logger.Infof("Parsed Variables: \n  Port: %d \n  Kubeconfig: %s", portNumber, kubeconfigLocation)

	k8sutil.KubeConfigLocation = kubeconfigLocation
	k8sutil.DefaultConfig, err = k8sutil.GenerateKubernetesConfig()

	if err != nil {
		logger.Infof("Was unable to generate a valid kubernetes default config, some functionality may be broken.  Error was %v", err)
	}

	// Install the CMA SDSCluster CRD
	k8sutil.CreateCRD(apiextensionsclient.NewForConfigOrDie(k8sutil.DefaultConfig), cma.GenerateSDSClusterCRD())
	// Install the CMA SDSPackageManager CRD
	k8sutil.CreateCRD(apiextensionsclient.NewForConfigOrDie(k8sutil.DefaultConfig), cma.GenerateSDSPackageManagerCRD())
	// Install the CMA SDSApplication CRD
	k8sutil.CreateCRD(apiextensionsclient.NewForConfigOrDie(k8sutil.DefaultConfig), cma.GenerateSDSApplicationCRD())

	var wg sync.WaitGroup
	stop := make(chan struct{})

	logger.Infof("Starting the SDSCluster Controller")
	sdsClusterController := cluster.NewSDSClusterController(nil)
	wg.Add(1)
	go func() {
		defer wg.Done()
		sdsClusterController.Run(3, stop)
	}()

	sdsPackageManagerController := sds_package_manager.NewSDSPackageManagerController(nil)
	// Start the SDSPackageManager Controller
	wg.Add(1)
	go func() {
		defer wg.Done()
		sdsPackageManagerController.Run(3, stop)
	}()
	// TODO: Start the SDSApplication Controller

	sdsApplicationController := sds_application.NewSDSApplicationController(nil)
	// Start the SDSPackageManager Controller
	wg.Add(1)
	go func() {
		defer wg.Done()
		sdsApplicationController.Run(3, stop)
	}()

	logger.Infof("Starting KrakenCluster Watcher")
	wg.Add(1)
	go func() {
		defer wg.Done()
		ccworkqueue.ListenToKrakenClusterChanges(nil)
	}()

	logger.Infof("Creating Web Server")
	tcpMux := createWebServer(&apiserver.ServerOptions{PortNumber: portNumber})
	wg.Add(1)
	go func() {
		defer wg.Done()
		logger.Infof("Starting to serve requests on port %d", portNumber)
		tcpMux.Serve()
	}()

	<-stop
	logger.Infof("Wating for controllers to shut down gracefully")
	wg.Wait()
}

func createWebServer(options *apiserver.ServerOptions) cmux.CMux {
	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", options.PortNumber))
	if err != nil {
		panic(err)
	}
	tcpMux := cmux.New(conn)

	apiserver.AddServersToMux(tcpMux, options)

	return tcpMux
}

func viperInit() {
	viper.SetEnvPrefix("clustermanagerapi")
	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)

	// using standard library "flag" package
	flag.Int("port", 9050, "Port to listen on")
	flag.String("kubeconfig", "", "Location of kubeconfig file")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	viper.AutomaticEnv()
}
