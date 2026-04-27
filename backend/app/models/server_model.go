package models

import (
	"strings"
	"time"

	"gorm.io/gorm"
)

const (
	ServerTypePnet       = "pnet"
	ServerTypeOpenID     = "openid"
	ServerTypeKubernetes = "k8s"
)

type ServerBase struct {
	Name string `gorm:"type:varchar(255)" json:"name"`
	Url  string `gorm:"type:varchar(255)" json:"url"`
	// enumeration: pnet, openid, k8s
	Type                 string     `gorm:"type:varchar(255)" json:"type"`
	IsActive             bool       `gorm:"type:bool" json:"is_active"`
	MinutesForDisconnect int        `gorm:"type:int" json:"minutes_for_disconnect"`
	MaxCountUsersLimit   int        `gorm:"type:int" json:"max_count_users_limit"`
	LastOnlineStatus     *time.Time `gorm:"type:datetime(3)" json:"last_online_status"`
	LastCountUsers       int        `gorm:"type:int" json:"last_count_users"`
	UnitRate             int        `gorm:"type:int" json:"unit_rate"`
}

func (t *ServerBase) HasPrefixUrl(redirectUri string) bool {
	redirectUri = strings.TrimPrefix(redirectUri, "http://")
	redirectUri = strings.TrimPrefix(redirectUri, "https://")

	urls := strings.FieldsFunc(t.Url, func(r rune) bool {
		return r == ',' || r == ';'
	})

	for _, u := range urls {
		url := strings.TrimSpace(u)
		url = strings.TrimPrefix(url, "http://")
		url = strings.TrimPrefix(url, "https://")

		if strings.HasPrefix(redirectUri, url) {
			return true
		}
	}

	return false
}

func ServerIsRealActive(db *gorm.DB) *gorm.DB {
	return db.Where("servers.is_active = ?", true).
		Where("servers.type = ?", ServerTypePnet).
		Where("(UTC_TIMESTAMP() < DATE_ADD(servers.last_online_status, INTERVAL servers.minutes_for_disconnect MINUTE) OR servers.minutes_for_disconnect = 0)").
		Where("(servers.last_count_users < servers.max_count_users_limit OR servers.max_count_users_limit = 0)")
}

type ServerSecret struct {
	Token    string `gorm:"type:string" json:"token" valid:"required"`
	ClientID string `gorm:"type:string" json:"client_id" valid:"required"`
}

type ServerListItem struct {
	Base
	ServerBase
}

func (ServerListItem) TableName() string {
	return "servers"
}

type Server struct {
	Base
	ServerBase
	ServerSecret
}
