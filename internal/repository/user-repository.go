package repository

import (
	"go-solicitud-despliegues-back/internal/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByID(object_id string) (*domain.UserAzureDVProfile, error)
	StoreUserProfile(*domain.UserAzureDVProfile) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetUserByID(object_id string) (*domain.UserAzureDVProfile, error) {

	var user domain.UserAzureDVProfile
	if err := r.db.First(&user, "object_id = ?", object_id).Error; err != nil {
		return nil, err
	}
	return &user, nil

}

func (r *userRepository) StoreUserProfile(user *domain.UserAzureDVProfile) error {
	if err := r.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}