package jwks

import (
	"fmt"
	"net/http"
	"time"

	"github.com/MicahParks/keyfunc"
)

// Init JWKS with cache and auto_refresh
// It must be called within the start of the application and must keep alive the *keyfunc.JWKS
func InitJWKS(jwksURL string) (*keyfunc.JWKS, error) {

	options := keyfunc.Options{
		RefreshErrorHandler: func(err error) {
			println("There was an error with the jwt.Keyfunc\nError:", err.Error())
		},
		RefreshInterval:   time.Hour,
		RefreshRateLimit:  time.Minute * 5,
		RefreshTimeout:    time.Second * 10,
		RefreshUnknownKID: true,
		Client:            &http.Client{Timeout: time.Second * 10},
	}

	jwks, err := keyfunc.Get(jwksURL, options)
	if err != nil {
		return nil, fmt.Errorf("failed to generate the JWKS from the given URL.\nError: %s", err.Error())
	}

	return jwks, nil
}