package azurek8sutil

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	CredentialsAppIDDataMap          = "azureAppID"
	CredentialsTenantDataMap         = "azureTenant"
	CredentialsPasswordDataMap       = "azurePassword"
	CredentialsSubscriptionIDDataMap = "azureSubscriptionID"

	SecretAzureCredentialsSuffix = "-azurecreds"
	SecretAzureCredentialsLabel  = "azurecredentials"
)

type Credentials struct {
	AppID          string
	Tenant         string
	Password       string
	SubscriptionID string
}

func (c *Client) getSecretList(options v1.ListOptions) (result []corev1.Secret, err error) {
	if c.secretClient == nil {
		return nil, fmt.Errorf("secret client is not initialized")
	}
	secrets, err := c.secretClient.List(options)

	result = secrets.Items
	return
}

func (c *Client) deleteSecret(name string) (err error) {
	if c.secretClient == nil {
		return fmt.Errorf("secret client is not initialized")
	}

	err = c.secretClient.Delete(name, &v1.DeleteOptions{})
	return
}

func (c *Client) createSecret(secret *corev1.Secret) (err error) {
	if c.secretClient == nil {
		return fmt.Errorf("secret client is not initialized")
	}

	_, err = c.secretClient.Create(secret)
	return
}

func (c *Client) getSecret(name string) (secret corev1.Secret, err error) {
	if c.secretClient == nil {
		return corev1.Secret{}, fmt.Errorf("secret client is not initialized")
	}

	secretResult, err := c.secretClient.Get(name, v1.GetOptions{})
	if err != nil {
		return
	}
	secret = *secretResult
	return
}

func (c *Client) ListCredentials() (result []string, err error) {
	listOption := v1.ListOptions{
		FieldSelector: "type=" + string(corev1.SecretTypeOpaque),
		LabelSelector: SecretAzureCredentialsLabel + "=true",
	}
	credentials, err := c.getSecretList(listOption)
	if err != nil {
		return nil, err
	}
	for _, j := range credentials {
		adjustedName, err := c.removeAdjustedName(j.Name)
		if err == nil {
			result = append(result, adjustedName)
		}
	}
	return result, nil
}

func (c *Client) GetCredentials(name string) (credentials Credentials, err error) {
	adjustedName := c.getAdjustedName(name)
	result, err := c.getSecret(adjustedName)
	if err != nil {
		return
	}
	if result.Type != corev1.SecretTypeOpaque {
		err = fmt.Errorf("secret %s is not of type %s, but rather is of type %s", name, corev1.SecretTypeOpaque, result.Type)
		return
	}
	if result.Labels[SecretAzureCredentialsLabel] != "true" {
		err = fmt.Errorf("secret %s does not have label "+SecretAzureCredentialsLabel+"=true", name)
		return
	}
	var appIDBytes []byte
	var tenantBytes []byte
	var passwordBytes []byte
	var subscriptionIDBytes []byte
	var ok bool
	if appIDBytes, ok = result.Data[CredentialsAppIDDataMap]; !ok {
		return Credentials{}, fmt.Errorf("could not find appID in secret")
	}
	if tenantBytes, ok = result.Data[CredentialsTenantDataMap]; !ok {
		return Credentials{}, fmt.Errorf("could not find tenant in secret")
	}
	if passwordBytes, ok = result.Data[CredentialsPasswordDataMap]; !ok {
		return Credentials{}, fmt.Errorf("could not find password in secret")
	}
	if subscriptionIDBytes, ok = result.Data[CredentialsSubscriptionIDDataMap]; !ok {
		return Credentials{}, fmt.Errorf("could not find subscriptionID in secret")
	}

	credentials = Credentials{
		AppID:          string(appIDBytes),
		Tenant:         string(tenantBytes),
		Password:       string(passwordBytes),
		SubscriptionID: string(subscriptionIDBytes),
	}
	return
}

func (c *Client) DeleteCredentials(name string) (err error) {
	adjustedName := c.getAdjustedName(name)
	return c.deleteSecret(adjustedName)
}

func (c *Client) CreateCredentials(name string, credentials Credentials) (err error) {
	adjustedName := c.getAdjustedName(name)
	labelMap := make(map[string]string)
	labelMap["cma"] = "true"
	labelMap[SecretAzureCredentialsLabel] = "true"

	dataMap := make(map[string][]byte)
	dataMap[CredentialsAppIDDataMap] = []byte(credentials.AppID)
	dataMap[CredentialsTenantDataMap] = []byte(credentials.Tenant)
	dataMap[CredentialsPasswordDataMap] = []byte(credentials.Password)
	dataMap[CredentialsSubscriptionIDDataMap] = []byte(credentials.SubscriptionID)

	secret := &corev1.Secret{
		ObjectMeta: v1.ObjectMeta{Name: adjustedName, Labels: labelMap},
		Type:       corev1.SecretTypeOpaque,
		Data:       dataMap,
	}

	return c.createSecret(secret)
}

func (c *Client) getAdjustedName(name string) string {
	return name + SecretAzureCredentialsSuffix
}

func (c *Client) removeAdjustedName(name string) (string, error) {
	inputLength := len(name)
	suffixLength := len(SecretAzureCredentialsSuffix)
	if suffixLength >= inputLength {
		return "", fmt.Errorf("input string is longer than suffix")
	}
	if name[(inputLength-suffixLength):] != SecretAzureCredentialsSuffix {
		return "", fmt.Errorf("input string does not have the suffix")
	}
	return name[0:(inputLength - suffixLength - 1)], nil
}
