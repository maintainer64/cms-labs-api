package models

import "time"

// PNETServerBase struct to describe PNETServer object.
type PNETServerBase struct {
	Name             string    `gorm:"type:varchar(255)" json:"name"`
	Url              string    `gorm:"type:varchar(255)" json:"url"`
	IsActive         bool      `gorm:"type:bool" json:"is_active"`
	LastOnlineStatus time.Time `gorm:"type:datetime(3)" json:"last_online_status"`
	UnitRate         uint      `gorm:"type:int" json:"unit_rate"`
}

type PNETServerSecret struct {
	Token string `gorm:"type:bool" json:"token" valid:"required"`
}

type PNETServerListItem struct {
	Base
	PNETServerBase
}

// TableName переопределяет название таблицы для PNETServerListItem на `pnet_servers`
func (PNETServerListItem) TableName() string {
	return "pnet_servers"
}

type PNETServer struct {
	Base
	PNETServerBase
	PNETServerSecret
}
