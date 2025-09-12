package usecase

import (
	"context"
	"errors"
	"go-solicitud-despliegues-back/internal/domain"
	"go-solicitud-despliegues-back/internal/repository"
	"go-solicitud-despliegues-back/internal/service"
	customContext "go-solicitud-despliegues-back/pkg/context"

	"gorm.io/gorm"
)

type UserUseCase interface {
	GetUserInfo(ctx context.Context) (*domain.UserAzureDVProfile, error)
}

type userUsecase struct {
	AzureDevOpsService service.AzureDevOpsService
	repo               repository.UserRepository
}

func NewUserUseCase(s service.AzureDevOpsService, r repository.UserRepository) UserUseCase {
	return &userUsecase{
		AzureDevOpsService: s,
		repo:               r,
	}
}

func (uc *userUsecase) GetUserInfo(ctx context.Context) (*domain.UserAzureDVProfile, error) {

	// Get user from context
	userContext := ctx.Value(customContext.UserCtxKey).(customContext.ContextUser)
	
	// Check users id
	if userContext.OID == "" {
		return nil, errors.New("user id (oid) is missing in token")
	}
	
	// Verify if the user exist in the database
	user, err := uc.repo.GetUserByID(userContext.OID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Call the azure devops service to get the user profile
			userAzureDVProfile, err := uc.AzureDevOpsService.GetUserAzureDVProfile(ctx, userContext.AccessToken)
			if err != nil {
				return nil, err
			}

			userAzureDVProfile.ObjectID = userContext.OID
			userAzureDVProfile.Role = userContext.Role

			// Store the profile in the database
			if err := uc.repo.StoreUserProfile(userAzureDVProfile); err != nil {
				return nil, err
			}

			user = userAzureDVProfile
		} else {
			return nil, err
		}
	}

	return user, nil
}
