package awsk8sutil

import (
	"fmt"
	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	CredentialsSecretKeyIDDataMap     = "awsSecretKeyID"
	CredentialsSecretAccessKeyDataMap = "awsSecretAccessKey"
	CredentialsRegionDataMap          = "awsRegion"

	SecretAWSCredentialsSuffix = "-awscreds"
	SecretAWSCredentialsLabel  = "awscredentials"
)

type Credentials struct {
	SecretKeyID     string
	SecretAccessKey string
	Region          string
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

func (c *Client) getSecret(name string) (secret *corev1.Secret, err error) {
	if c.secretClient == nil {
		return &corev1.Secret{}, fmt.Errorf("secret client is not initialized")
	}

	secretResult, err := c.secretClient.Get(name, v1.GetOptions{})
	if err != nil {
		return
	}
	secret = secretResult
	return
}

func (c *Client) updateSecret(secret *corev1.Secret) (err error) {
	if c.secretClient == nil {
		return fmt.Errorf("secret client is not initialized")
	}

	_, err = c.secretClient.Update(secret)
	return
}

func (c *Client) ListCredentials() (result []string, err error) {
	listOption := v1.ListOptions{
		FieldSelector: "type=" + string(corev1.SecretTypeOpaque),
		LabelSelector: SecretAWSCredentialsLabel + "=true",
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
	if result.Labels[SecretAWSCredentialsLabel] != "true" {
		err = fmt.Errorf("secret %s does not have label awscredentials=true", name)
		return
	}
	var regionBytes []byte
	var secretKeyIDBytes []byte
	var secretAccessKeyBytes []byte
	var ok bool
	if regionBytes, ok = result.Data[CredentialsRegionDataMap]; !ok {
		return Credentials{}, fmt.Errorf("could not find region in secret")
	}
	if secretKeyIDBytes, ok = result.Data[CredentialsSecretKeyIDDataMap]; !ok {
		return Credentials{}, fmt.Errorf("could not find secretKeyID in secret")
	}
	if secretAccessKeyBytes, ok = result.Data[CredentialsSecretAccessKeyDataMap]; !ok {
		return Credentials{}, fmt.Errorf("could not find secretAccessKey in secret")
	}

	credentials = Credentials{
		Region:          string(regionBytes),
		SecretKeyID:     string(secretKeyIDBytes),
		SecretAccessKey: string(secretAccessKeyBytes),
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
	labelMap[SecretAWSCredentialsLabel] = "true"

	dataMap := make(map[string][]byte)
	dataMap[CredentialsRegionDataMap] = []byte(credentials.Region)
	dataMap[CredentialsSecretKeyIDDataMap] = []byte(credentials.SecretKeyID)
	dataMap[CredentialsSecretAccessKeyDataMap] = []byte(credentials.SecretAccessKey)

	secret := &corev1.Secret{
		ObjectMeta: v1.ObjectMeta{Name: adjustedName, Labels: labelMap},
		Type:       corev1.SecretTypeOpaque,
		Data:       dataMap,
	}

	return c.createSecret(secret)
}

func (c *Client) UpdateOrCreateCredentials(name string, credentials Credentials) (err error) {
	// Let's try getting the credentials
	secret, err := c.getSecret(c.getAdjustedName(name))

	// If we had an error, then we need to create,
	if err != nil {
		logrus.Errorf("Chose to create")
		return c.CreateCredentials(name, credentials)
	} else {
		// If we had no error, then we need to update
		logrus.Errorf("Chose to update")
		return c.updateCredentials(name, credentials, secret)
	}
}

func (c *Client) updateCredentials(name string, credentials Credentials, secret *corev1.Secret) error {
	// Ensuring Labels are set
	secret.Labels[SecretAWSCredentialsLabel] = "true"
	secret.Data[CredentialsRegionDataMap] = []byte(credentials.Region)
	secret.Data[CredentialsSecretKeyIDDataMap] = []byte(credentials.SecretKeyID)
	secret.Data[CredentialsSecretAccessKeyDataMap] = []byte(credentials.SecretAccessKey)

	return c.updateSecret(secret)
}

func (c *Client) getAdjustedName(name string) string {
	return name + SecretAWSCredentialsSuffix
}

func (c *Client) removeAdjustedName(name string) (string, error) {
	inputLength := len(name)
	suffixLength := len(SecretAWSCredentialsSuffix)
	if suffixLength >= inputLength {
		return "", fmt.Errorf("input string is longer than suffix")
	}
	if name[(inputLength-suffixLength):] != SecretAWSCredentialsSuffix {
		return "", fmt.Errorf("input string does not have the suffix")
	}
	return name[0:(inputLength - suffixLength - 1)], nil
}
