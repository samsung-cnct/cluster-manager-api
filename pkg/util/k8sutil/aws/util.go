package awsk8sutil

import (
	"github.com/spf13/viper"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/typed/core/v1"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)

const (
	kubeconfigDir  = ".kube"
	kubeconfigFile = "config"
)

func CreateFromDefaults() (ClientInterface, error) {
	output := Client{}
	err := output.CreateNewClients()
	if err != nil {
		return nil, err
	}
	return &output, nil
}

func (c *Client) CreateNewClients() error {
	var err error
	if c.config == nil {
		err = c.createNewConfig()
		if err != nil {
			return err
		}
	}
	coreClient, err := kubernetes.NewForConfig(c.config)
	if err != nil {
		return err
	}
	c.secretClient = coreClient.CoreV1().Secrets(viper.GetString("kubernetes-namespace"))
	return err
}

func (c *Client) SetSecretClient(client v1.SecretInterface) {
	c.secretClient = client
}

func (c *Client) createNewConfig() error {
	var err error
	if c.kubeConfigLocation != "" {
		c.config, err = clientcmd.BuildConfigFromFlags("", c.kubeConfigLocation)
	} else {
		configPath := filepath.Join(homeDir(), kubeconfigDir, kubeconfigFile)
		_, err := os.Stat(configPath)
		if err == nil {
			c.config, err = clientcmd.BuildConfigFromFlags("", configPath)
		} else {
			c.config, err = rest.InClusterConfig()
		}
	}
	if err != nil {
		return err
	}
	return nil
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

func (c *Client) SetConfig(config *rest.Config) {
	c.config = config
}
