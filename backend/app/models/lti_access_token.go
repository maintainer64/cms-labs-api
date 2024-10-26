package models

// LTIAccessToken struct to describe LTINonceToken object.
type LTIAccessToken struct {
	Base
	Index   string `gorm:"type:text" json:"index" validate:"required"`
	Payload string `gorm:"type:text" json:"target_link_uri" validate:"required"`
}
