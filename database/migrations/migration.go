package migrations

import (
	"go-solicitud-despliegues-back/internal/domain"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	
	err := db.Migrator().DropTable(&domain.UserAzureDVProfile{})
	if err != nil {
		return err
	}

	err = db.Migrator().DropTable(&domain.AzureDevopsOrganization{})
	if err != nil {
		return err
	}

	err = db.Migrator().DropTable(&domain.UserOrganizations{})
	if err != nil {
		return err
	}

	return db.AutoMigrate(
		&domain.UserAzureDVProfile{},
		&domain.AzureDevopsOrganization{},
	)
}