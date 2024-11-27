package models

import (
	"time"

	"gorm.io/gorm"
)

// PNETServerBase struct to describe PNETServer object.
type PNETServerBase struct {
	Name                 string     `gorm:"type:varchar(255)" json:"name"`
	Url                  string     `gorm:"type:varchar(255)" json:"url"`
	IsActive             bool       `gorm:"type:bool" json:"is_active"`
	MinutesForDisconnect uint       `gorm:"type:int" json:"minutes_for_disconnect"`
	MaxCountUsersLimit   uint       `gorm:"type:int" json:"max_count_users_limit"`
	LastOnlineStatus     *time.Time `gorm:"type:datetime(3)" json:"last_online_status"`
	LastCountUsers       uint       `gorm:"type:int" json:"last_count_users"`
	UnitRate             uint       `gorm:"type:int" json:"unit_rate"`
}

func PNETServeIsRealActive(db *gorm.DB) *gorm.DB {
	return db.Where("pnet_servers.is_active = ?", true).
		Where("UTC_TIMESTAMP() < DATE_ADD(pnet_servers.last_online_status, INTERVAL pnet_servers.minutes_for_disconnect MINUTE) OR pnet_servers.minutes_for_disconnect = 0").
		Where("pnet_servers.last_count_users < pnet_servers.max_count_users_limit OR pnet_servers.max_count_users_limit = 0")
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
