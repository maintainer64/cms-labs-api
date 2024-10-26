package models

// LTINonceTokenBase struct to describe LTINonceToken object.
type LTINonceTokenBase struct {
	Nonce         string `gorm:"type:varchar(255)" json:"nonce" validate:"required"`
	TargetLinkURI string `gorm:"type:varchar(255)" json:"target_link_uri" validate:"required"`
}

type LTINonceToken struct {
	Base
	LTINonceTokenBase
}
