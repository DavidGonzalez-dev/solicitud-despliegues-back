package service

import (
	"context"

	"go-solicitud-despliegues-back/internal/domain"
	pkgHttp "go-solicitud-despliegues-back/pkg/http"
	"net/http"
)

type AzureDevOpsService interface {
	GetUserAzureDVProfile(ctx context.Context, accessToken string) (*domain.UserAzureDVProfile, error)
}

type azureDevOpsService struct {
	Client      *http.Client
	oboService  OboService
}

func NewAzureDevopsService() AzureDevOpsService {
	return &azureDevOpsService {
		Client:      &http.Client{Timeout: 10 * http.DefaultClient.Timeout},
		oboService: NewOboService("499b84ac-1321-427f-aa17-267ca6975798/.default"),
	}
}

func (s *azureDevOpsService) GetUserAzureDVProfile(ctx context.Context, accessToken string) (*domain.UserAzureDVProfile, error) {
	
	// Get obo Token
	oboToken, err := s.oboService.GetOboToken(ctx, accessToken)
	if err != nil {
		return nil, err
	}


	// Make request
	req, err := http.NewRequest("GET", "https://app.vssps.visualstudio.com/_apis/profile/profiles/me?api-version=7.1-preview.3", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+oboToken)

	var result domain.UserAzureDVProfile
	if err := pkgHttp.DoHttpRequest(req, &result); err != nil {
		return nil, err
	}

	// Return info
	return &result, nil
}
