package models

// LTIFormBase struct to describe LTIForm object.
type LTIFormBase struct {
	Name            string  `gorm:"type:varchar(255)" json:"name" validate:"required"`
	LTIClientID     string  `gorm:"type:varchar(255);column:lti_client_id" json:"lti_client_id" validate:"required"`
	LTIDeploymentID string  `gorm:"type:varchar(255);column:lti_deployment_id" json:"lti_deployment_id" validate:"required"`
	SSOURL          *string `gorm:"type:varchar(255);column:sso_url" json:"sso_url"`
}

type LTIFormSecret struct {
	BaseURI         string `gorm:"type:varchar(255)" json:"base_uri" validate:"required"`
	LTIAuthTokenURI string `gorm:"type:varchar(255)" json:"lti_auth_token_uri" validate:"required"`
	LTIAuthLoginURI string `gorm:"type:varchar(255)" json:"lti_auth_login_uri" validate:"required"`
	KeySetURI       string `gorm:"type:varchar(255)" json:"key_set_uri" validate:"required"`
	TargetLinkURI   string `gorm:"type:varchar(255)" json:"target_link_uri" validate:"required"`
	PublicKey       string `gorm:"type:text" json:"public_key" validate:"required"`
	PrivateKey      string `gorm:"type:text" json:"private_key" validate:"required"`
}

type LTIFormListItem struct {
	Base
	LTIFormBase
}

// TableName переопределяет название таблицы для LTIFormListItem на `lti_forms`
func (LTIFormListItem) TableName() string {
	return "lti_forms"
}

type LTIForm struct {
	Base
	LTIFormBase
	LTIFormSecret
}
