package models

import (
	"strings"
	"time"

	"gorm.io/gorm"
)

const (
	ServerTypePnet   = "pnet"
	ServerTypeOpenID = "openid"
)

// PNETServerBase struct to describe PNETServer object.
type PNETServerBase struct {
	Name                 string     `gorm:"type:varchar(255)" json:"name"`
	Url                  string     `gorm:"type:varchar(255)" json:"url"`
	Type                 string     `gorm:"type:varchar(255)" json:"type"`
	IsActive             bool       `gorm:"type:bool" json:"is_active"`
	MinutesForDisconnect int        `gorm:"type:int" json:"minutes_for_disconnect"`
	MaxCountUsersLimit   int        `gorm:"type:int" json:"max_count_users_limit"`
	LastOnlineStatus     *time.Time `gorm:"type:datetime(3)" json:"last_online_status"`
	LastCountUsers       int        `gorm:"type:int" json:"last_count_users"`
	UnitRate             int        `gorm:"type:int" json:"unit_rate"`
}

func (t *PNETServerBase) HasPrefixUrl(redirectUri string) bool {
	redirectUri = strings.TrimPrefix(redirectUri, "http://")
	redirectUri = strings.TrimPrefix(redirectUri, "https://")
	url := strings.TrimPrefix(t.Url, "http://")
	url = strings.TrimPrefix(url, "https://")
	return strings.HasPrefix(redirectUri, url)
}

func PNETServeIsRealActive(db *gorm.DB) *gorm.DB {
	return db.Where("pnet_servers.is_active = ?", true).
		Where("pnet_servers.type = ?", ServerTypePnet).
		Where("(UTC_TIMESTAMP() < DATE_ADD(pnet_servers.last_online_status, INTERVAL pnet_servers.minutes_for_disconnect MINUTE) OR pnet_servers.minutes_for_disconnect = 0)").
		Where("(pnet_servers.last_count_users < pnet_servers.max_count_users_limit OR pnet_servers.max_count_users_limit = 0)")
}

type PNETServerSecret struct {
	Token    string `gorm:"type:string" json:"token" valid:"required"`
	ClientID string `gorm:"type:string" json:"client_id" valid:"required"`
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
