package models

import "time"

// LTIAttemptBase struct to describe LTIAttempt object.
type LTIAttemptBase struct {
	UserID       uint      `gorm:"type:int" json:"user_id"`
	PNETServerID uint      `gorm:"type:int" json:"pnet_server_id"`
	LTIRoutingID uint      `gorm:"type:int" json:"lti_routing_id"`
	ExpiredAt    time.Time `gorm:"type:datetime(3)" json:"expired_at" validate:"required"`
}

type LTIAttemptSecret struct {
	RoomNumber *int64 `gorm:"type:int" json:"room_number"`
}

type LTIAttemptListItem struct {
	Base
	LTIAttemptBase
	UserEmail      string `json:"user_email"`
	UserName       string `json:"user_name"`
	PNETServerName string `json:"pnet_server_name"`
	LTIRoutingName string `json:"lti_routing_name"`
}

// TableName переопределяет название таблицы для LTIAttemptListItem на `lti_attempts`
func (LTIAttemptListItem) TableName() string {
	return "lti_attempts"
}

type LTIAttempt struct {
	Base
	LTIAttemptBase
	LTIAttemptSecret
}

func (l *LTIAttempt) ExtendExpiredAt(extensionHours int) {
	// Проверяем, осталось ли до окончания сессии меньше часа
	if time.Until(l.ExpiredAt) < time.Hour {
		// Продлеваем сессию на указанное количество часов
		l.ExpiredAt = l.ExpiredAt.Add(time.Duration(extensionHours) * time.Hour)
	}
}
