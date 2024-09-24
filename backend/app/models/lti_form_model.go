package models

// LTIForm struct to describe LTIForm object.

type LTIFromBase struct {
	Name    string `gorm:"type:varchar(255)" json:"name" validate:"required"`
	Version string `gorm:"type:varchar(255)" json:"version" validate:"required"`
}

type LTIFromSecret struct {
	LTIKey    string `gorm:"type:varchar(255),index" json:"lti_key" validate:"required"`
	LTISecret string `gorm:"type:varchar(255)" json:"lti_secret" validate:"required"`
}

type LTIFormListItem struct {
	Base
	LTIFromBase
}

// TableName переопределяет название таблицы для LTIFormListItem на `lti_forms`
func (LTIFormListItem) TableName() string {
	return "lti_forms"
}

type LTIForm struct {
	Base
	LTIFromBase
	LTIFromSecret
}
