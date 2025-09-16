package domain

var UserRoles = struct {
	CLOUD     string
	DEVELOPER string
}{
	CLOUD:     "cloud",
	DEVELOPER: "developer",
}

type UserAzureDVProfile struct {
	AccountID     string                    `json:"accountId" gorm:"primaryKey;not null"`
	ObjectID      string                    `json:"objectId" gorm:"object_id;not null;unique"`
	DisplayName   string                    `json:"displayName" gorm:"not null"`
	Role          string                    `json:"role" gorm:"not null"`
	Organizations []AzureDevopsOrganization `json:"organizations" gorm:"many2many:user_organizations"`
}
