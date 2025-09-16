package domain

type AzureDevopsOrganization struct {
	AccountId   string `json:"accountId" gorm:"account_id;primaryKey"`
	AccountName string `json:"accountName" gorm:"account_name;not null"`
}

type UserOrganizations struct {
	UserAccountId string 
	OrgAccountId  string 
}
