package main

import (
	"flag"
	"fmt"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/k8sutil"
	"github.com/soheilhy/cmux"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"net"
	"strings"

	"github.com/juju/loggo"
	"github.com/samsung-cnct/cluster-manager-api/pkg/apiserver"
	"k8s.io/client-go/rest"
	ccworkqueue "github.com/samsung-cnct/cluster-manager-api/pkg/util/ccutil/workqueue"
	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/cma"
)

var (
	logger loggo.Logger
	config *rest.Config
)

func main() {
	var err error
	logger := util.GetModuleLogger("cmd.redis-operator", loggo.INFO)
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
	k8sutil.CreateCRD(apiextensionsclient.NewForConfigOrDie(k8sutil.DefaultConfig), cma.GenerateSDSClusterCRD() )
	// Install the CMA SDSPackageManager CRD
	k8sutil.CreateCRD(apiextensionsclient.NewForConfigOrDie(k8sutil.DefaultConfig), cma.GenerateSDSPackageManagerCRD() )
	// Install the CMA SDSApplication CRD
	k8sutil.CreateCRD(apiextensionsclient.NewForConfigOrDie(k8sutil.DefaultConfig), cma.GenerateSDSApplicationCRD() )

	// Start the KrakenCluster CR Watcher
	go func() {
		ccworkqueue.ListenToKrakenClusterChanges(nil)
	}()

	logger.Infof("Creating Web Server")
	tcpMux := createWebServer(&apiserver.ServerOptions{PortNumber: portNumber})

	logger.Infof("Starting to serve requests on port %d", portNumber)
	tcpMux.Serve()
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
