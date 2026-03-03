package models

// AuthProviderBase struct to describe AuthProvider object.

const (
	AuthProviderTypeLTI  = "lti"
	AuthProviderTypeLDAP = "ldap"
)

type AuthProviderBase struct {
	Name   string  `gorm:"type:varchar(255)" json:"name" validate:"required"`
	Type   string  `gorm:"type:varchar(255)" json:"type" validate:"required"`
	SSOURL *string `gorm:"type:varchar(255);column:sso_url" json:"sso_url"`
}

type AuthProviderSecret struct {
	BaseURI         string `gorm:"type:varchar(255)" json:"base_uri" validate:"required"`
	LTIClientID     string `gorm:"type:varchar(255);column:lti_client_id" json:"lti_client_id"`
	LTIDeploymentID string `gorm:"type:varchar(255);column:lti_deployment_id" json:"lti_deployment_id"`
	LTIAuthTokenURI string `gorm:"type:varchar(255)" json:"lti_auth_token_uri"`
	LTIAuthLoginURI string `gorm:"type:varchar(255)" json:"lti_auth_login_uri"`
	KeySetURI       string `gorm:"type:varchar(255)" json:"key_set_uri"`
	TargetLinkURI   string `gorm:"type:varchar(255)" json:"target_link_uri"`
	PublicKey       string `gorm:"type:text" json:"public_key"`
	PrivateKey      string `gorm:"type:text" json:"private_key"`
}

type AuthProviderListItem struct {
	Base
	AuthProviderBase
}

// TableName переопределяет название таблицы для AuthProviderListItem на `lti_forms`
func (AuthProviderListItem) TableName() string {
	return "auth_providers"
}

type AuthProvider struct {
	Base
	AuthProviderBase
	AuthProviderSecret
}
