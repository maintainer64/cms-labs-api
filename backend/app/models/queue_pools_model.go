package models

import "time"

const (
	RoundQueuePoolTypePNET = "pnet"
)

// RoundQueuePoolBase struct to describe RoundQueuePool object.
type RoundQueuePoolBase struct {
	Type        string    `gorm:"type:varchar(255)" json:"type" validate:"required"`
	ConnectedAt time.Time `gorm:"type:datetime(3)" json:"connected_at" validate:"required"`
	LastUsed    bool      `json:"last_used" validate:"required"`
	IsActive    bool      `gorm:"type:bool" json:"is_active" validate:"required"`
}

type RoundQueuePoolSecret struct {
	ServerID uint `gorm:"type:int" json:"server_id"`
}

type RoundQueuePoolListItem struct {
	Base
	ServerBase
}

// TableName overrides table name for RoundQueuePoolListItem to `round_queue_pools`
func (RoundQueuePoolListItem) TableName() string {
	return "round_queue_pools"
}

type RoundQueuePool struct {
	Base
	RoundQueuePoolBase
	RoundQueuePoolSecret
}
