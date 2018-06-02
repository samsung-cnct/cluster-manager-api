package main

import (
	"flag"
	"fmt"
	"net"
	"strings"
	"sync"

	"github.com/samsung-cnct/cma-operator/pkg/util"
	"github.com/samsung-cnct/cma-operator/pkg/util/k8sutil"
	"github.com/soheilhy/cmux"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/juju/loggo"
	"github.com/samsung-cnct/cluster-manager-api/pkg/apiserver"
	"k8s.io/client-go/rest"
	"github.com/spf13/cobra"
	"encoding/json"
	"github.com/samsung-cnct/cluster-manager-api/pkg/version"
)

var (
	logger loggo.Logger
	config *rest.Config
)


func main() {
	var rootCmd = &cobra.Command{
		Use:   "cluster-manager-api",
		Short: "Hugo is a very fast static site generator",
		Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
		Run: func(cmd *cobra.Command, args []string) {
			doSomething()
		},
	}
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Hugo is a very fast static site generator",
		Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
		Run: func(cmd *cobra.Command, args []string) {
			versionString, _ := json.Marshal(version.Get())
			fmt.Printf("%s\n", versionString)
		},
	}

	cobra.OnInitialize()
	viperInit()
	rootCmd.AddCommand(versionCmd)
	rootCmd.Execute()
}


func doSomething() {
	var err error
	logger := util.GetModuleLogger("cmd.cma-operator", loggo.INFO)

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
