package azure

import (
	"fmt"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/cmaaks"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/k8sutil/azure"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/k8sutil/cma"
	"github.com/spf13/viper"
)

const (
	CMAAKSEndpointViperVariableName = "cmaaks-endpoint"
	CMAAKSInsecureViperVariableName = "cmaaks-insecure"
	NotEnabledErrorMessage          = "azure support is not enabled"
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

	// Adding the CMAAKS Client
	c.cmaAKSClient, err = getCMAAKSClient()
	if err != nil {
		return err
	}

	c.secretClient, err = azurek8sutil.CreateFromDefaults()
	if err != nil {
		// Closing because we created the client but then errored
		c.Close()
		return err
	}

	c.cmaK8sClient, err = cmak8sutil.CreateFromDefaults()
	if err != nil {
		// Closing because we created the client but then errored
		c.Close()
		return err
	}

	return nil
}

func getCMAAKSClient() (cmaaks.ClientInterface, error) {
	if IsEnabled() == false {
		return nil, fmt.Errorf(NotEnabledErrorMessage)
	}

	hostname := viper.GetString(CMAAKSEndpointViperVariableName)
	insecure := viper.GetBool(CMAAKSInsecureViperVariableName)
	return cmaaks.CreateNewClient(hostname, insecure)
}

func (c *Client) SetCMAAKSClient(client cmaaks.ClientInterface) {
	c.cmaAKSClient = client
}

func (c *Client) SetSecretClient(client azurek8sutil.ClientInterface) {
	c.secretClient = client
}

func (c *Client) SetCMAK8sClient(client cmak8sutil.ClientInterface) {
	c.cmaK8sClient = client
}

func (c *Client) Close() error {
	return c.cmaAKSClient.Close()
}

func IsEnabled() bool {
	if viper.GetString(CMAAKSEndpointViperVariableName) == "" {
		return false
	}
	return true
}
