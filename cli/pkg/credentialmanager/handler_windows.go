package credentialmanager

import (
	"net/url"

	"github.com/docker/docker-credential-helpers/wincred"
)

type CredentialManager struct {
}

func NewCredentialManager() (cm *CredentialManager) {
	return &CredentialManager{}
}

// SetCreds stores the credentials consisting of an endpoint and an api token in the keychain.
func (cm *CredentialManager) SetCreds(endPoint url.URL, apiToken string) error {
	return setCreds(wincred.Wincred{}, endPoint, apiToken)
}

// GetCreds reads the credentials and returns an endpoint, the api token, or potentially an error.
func (cm *CredentialManager) GetCreds() (url.URL, string, error) {
	return getCreds(wincred.Wincred{})
}

// SetInstallCreds sets the install credentials
func (cm *CredentialManager) SetInstallCreds(creds string) error {
	return setInstallCreds(wincred.Wincred{}, creds)
}

// GetInstallCreds gets the install credentials
func (cm *CredentialManager) GetInstallCreds() (string, error) {
	return getInstallCreds(wincred.Wincred{})
}
