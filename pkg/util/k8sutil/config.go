package k8sutil

import (
	"os"
	"path/filepath"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	KubeConfigLocation string
	DefaultConfig      *rest.Config
)

const (
	kubeconfigDir  = ".kube"
	kubeconfigFile = "config"
)

type promptedCredentials struct {
	username string
	password string
}

func GenerateKubernetesConfig() (*rest.Config, error) {
	if KubeConfigLocation != "" {
		return clientcmd.BuildConfigFromFlags("", KubeConfigLocation)
	} else {
		configPath := filepath.Join(homeDir(), kubeconfigDir, kubeconfigFile)
		_, err := os.Stat(configPath)
		if err == nil {
			return clientcmd.BuildConfigFromFlags("", configPath)
		} else {
			return rest.InClusterConfig()
		}
	}
}
