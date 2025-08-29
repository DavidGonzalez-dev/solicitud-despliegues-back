package migrations

import (
	"go-solicitud-despliegues-back/internal/domain"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&domain.UserAzureDVProfile{},
	)
}