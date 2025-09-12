package domain

var UserRoles = struct {
	CLOUD     string
	DEVELOPER string
}{
	CLOUD:     "cloud",
	DEVELOPER: "developer",
}

type UserAzureDVProfile struct {
	ObjectID    string `json:"oid" gorm:"primaryKey"`
	AzureID     string `json:"azureId" gorm:"azure_id;not null"`
	DisplayName string `json:"displayName" gorm:"display_name;not null"`
	Role        string `json:"role" gorm:"role;not null"`
}
