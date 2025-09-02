package config

import (
	"fmt"
	"go-solicitud-despliegues-back/pkg/jwks"
	"os"

	"github.com/MicahParks/keyfunc"
)

// Stores min values that must be set in env variables for auth with Azure AD and the auth middleware to work
type AuthConfig struct {
	TenantID     string `json:"tenant_id"`
	ApiAudience  string `json:"api_audience"` // ej. "api://<client-id>"
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func NewAuthConfig() *AuthConfig {
	return &AuthConfig{
		TenantID:     os.Getenv("AZURE_TENANT_ID"),
		ClientID:     os.Getenv("AZURE_CLIENT_ID"),
		ClientSecret: os.Getenv("AZURE_CLIENT_SECRET"),
	}
}

func (c *AuthConfig) Issuer() string {
	return "https://login.microsoftonline.com/" + c.TenantID + "/v2.0"
}

func (c *AuthConfig) JWKSURL() string {
	return "https://login.microsoftonline.com/" + c.TenantID + "/discovery/v2.0/keys"
}


// Authenticator encapsulates the logic to validate a JWT token and extract user info
type Authenticator struct {
	TenantID string
	Audience string
	Issuer   string
	JWKS     *keyfunc.JWKS
}

// Init jwks
func NewAuthenticator(authConfig *AuthConfig) (*Authenticator, error) {

	jwks, err := jwks.InitJWKS(authConfig.JWKSURL())
	if err != nil {
		return nil, fmt.Errorf("failed to initialize JWKS: %v", err)
	}

	return &Authenticator{
		TenantID: authConfig.TenantID,
		Audience: authConfig.ClientID,
		Issuer:   authConfig.Issuer(),
		JWKS:     jwks,
	}, nil
}

