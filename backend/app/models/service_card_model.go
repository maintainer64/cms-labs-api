package models

// ServiceCardBase struct to describe ServiceCard object.
type ServiceCardBase struct {
	ImageUrl    string `gorm:"type:varchar(255)" json:"image_url" validate:"required"`
	Url         string `gorm:"type:varchar(255)" json:"url" validate:"required"`
	Name        string `gorm:"type:varchar(255)" json:"name" validate:"required"`
	Description string `gorm:"type:varchar(255)" json:"description" validate:"required"`
	Order       uint   `gorm:"type:int" json:"order"`
	IsActive    bool   `gorm:"type:bool" json:"is_active"`
}

type ServiceCardListItem struct {
	Base
	ServiceCardBase
}

// TableName переопределяет название таблицы для ServiceCardListItem на `service_cards`
func (ServiceCardListItem) TableName() string {
	return "service_cards"
}

type ServiceCard struct {
	Base
	ServiceCardBase
}
