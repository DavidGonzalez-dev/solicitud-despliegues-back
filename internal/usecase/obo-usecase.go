package usecase

import (
	"context"
	"go-solicitud-despliegues-back/internal/domain"
	"go-solicitud-despliegues-back/internal/service"
)


type OboUseCase interface {
	LoginOnBehalfOf (req domain.OboRequest) (domain.OboResponse, error)
}

type oboUseCase struct {
	OboService service.OboService
}

func NewOboUsecase(s service.OboService) OboUseCase {
	return &oboUseCase {
		OboService: s,
	}
}

func (u *oboUseCase) LoginOnBehalfOf (req domain.OboRequest) (domain.OboResponse, error) {
	
	// Create context and get new token
	ctx := context.Background()
	token, err := u.OboService.GetOboToken(ctx, req.AccessToken)
	if err != nil {
		return domain.OboResponse{}, err
	}

	return domain.OboResponse{OboToken: token}, nil
	
}
