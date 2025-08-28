package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

type OboService interface {
	GetOboToken(ctx context.Context, accessToken string) (string, error)
}

type oboService struct {
	Client       *http.Client
	ClientID     string
	ClientSecret string
	TentantID    string
	Scope        string
}

func NewOboService() OboService {	
	return &oboService{
		Client:       &http.Client{Timeout: 10 * http.DefaultClient.Timeout},
		ClientID:     os.Getenv("AZURE_CLIENT_ID"),
		ClientSecret: os.Getenv("AZURE_CLIENT_SECRET"),
		TentantID:    os.Getenv("AZURE_TENANT_ID"),
		Scope:        os.Getenv("AZURE_API_SCOPE"),
	}
}

func (s *oboService) GetOboToken(ctx context.Context, accessToken string) (string, error) {
	// Build the endpoint url and request body
	url := "https://login.microsoftonline.com/" + s.TentantID + "/oauth2/v2.0/token"
	data := map[string]string{
		"client_id":           s.ClientID,
		"client_secret":       s.ClientSecret,
		"grant_type":          "urn:ietf:params:oauth:grant-type:jwt-bearer",
		"requested_token_use": "on_behalf_of",
		"scope":               s.Scope,
		"assertion":           accessToken,
	}

	// Build the request body
	reqBody := ""
	for k, v := range data {
		if reqBody != "" {
			reqBody += "&"
		}
		reqBody += k + "=" + v
	}

	// Create request
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBufferString(reqBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Do the request
	resp, err := s.Client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Check for response status
	if resp.StatusCode != http.StatusOK {

		// Print the response body
		var errResp struct {
			Error            string `json:"error"`
			ErrorDescription string `json:"error_description"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
			fmt.Printf("Error: %+v\n", resp.Body)
			return "", errors.New("failed to get OBO token")
		}

		fmt.Println("Error while getting the OBO token")
		fmt.Printf("Error: %+v\n", errResp)
		return "", errors.New("failed to get OBO token")
	}

	// Read the response body
	var result struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result.AccessToken, nil
}
