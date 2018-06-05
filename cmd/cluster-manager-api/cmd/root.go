package cmd

import (
	"flag"
	"fmt"
	"net"
	"strings"
	"sync"

	"github.com/samsung-cnct/cma-operator/pkg/util"
	"github.com/samsung-cnct/cma-operator/pkg/util/k8sutil"
	"github.com/soheilhy/cmux"
	"github.com/spf13/viper"

	"github.com/juju/loggo"
	"github.com/samsung-cnct/cluster-manager-api/pkg/apiserver"
	"k8s.io/client-go/rest"
	"github.com/spf13/cobra"
	"os"
)

var (
	logger loggo.Logger
	config *rest.Config

	rootCmd = &cobra.Command{
		Use:   "cluster-manager-api",
		Short: "The cluster manager API",
		Long: `The Cluster Manager API

Running this by itself will invoke the webserver to run.
See subcommands for additional features`,
		Run: func(cmd *cobra.Command, args []string) {
			runWebServer()
		},
	}
)

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}


func runWebServer() {
	var err error
	logger := util.GetModuleLogger("cmd.cluster-manager-api", loggo.INFO)

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

	var wg sync.WaitGroup
	stop := make(chan struct{})

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

func init() {

	viper.SetEnvPrefix("clustermanagerapi")
	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)

	// using standard library "flag" package
	rootCmd.Flags().Int("port", 9050, "Port to listen on")
	rootCmd.Flags().String("kubeconfig", "", "Location of kubeconfig file")

	//pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	//pflag.Parse()
	viper.BindPFlag("port", rootCmd.Flags().Lookup("port"))
	viper.BindPFlag("kubeconfig", rootCmd.Flags().Lookup("kubeconfig"))

	viper.AutomaticEnv()
	rootCmd.Flags().AddGoFlagSet(flag.CommandLine)

}


