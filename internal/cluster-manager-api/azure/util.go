package azure

import (
	"fmt"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/cmaaks"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/k8sutil/azure"
	"github.com/spf13/viper"
)

const (
	CMAAKEndpointViperVariableName  = "cmaaks-endpoint"
	CMAAKSInsecureViperVariableName = "cmaaks-insecure"
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

	return nil
}

func getCMAAKSClient() (cmaaks.AKSClientInterface, error) {
	hostname := viper.GetString(CMAAKEndpointViperVariableName)
	if hostname == "" {
		return nil, fmt.Errorf("azure support is not enabled")
	}
	insecure := viper.GetBool(CMAAKSInsecureViperVariableName)
	return cmaaks.CreateNewClient(hostname, insecure)
}

func (c *Client) Close() error {
	return c.cmaAKSClient.Close()
}
