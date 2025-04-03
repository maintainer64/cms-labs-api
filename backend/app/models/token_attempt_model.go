package models

import "time"

// UserPassword struct to describe Save response.TokenPublicData object.
type UserPassword struct {
	UserID       uint      `gorm:"primaryKey" json:"user_id"`
	HashPassword string    `gorm:"type:varchar(255)" json:"hash_password"`
	CreatedAt    time.Time `gorm:"type:datetime(3)" json:"created_at" validate:"required"`
	UpdatedAt    time.Time `gorm:"type:datetime(3)" json:"updated_at" validate:"required"`
}

// TokenAttemptBase struct to describe TokenAttempt object.
type TokenAttemptBase struct {
	UserID   uint `gorm:"type:int" json:"user_id"`
	ServerID uint `gorm:"type:int" json:"server_id"`
}

type TokenAttemptSecret struct {
	Token             string `gorm:"type:varchar(255)" json:"token"`
	State             string `gorm:"type:varchar(255)" json:"state"`
	Nonce             string `gorm:"type:varchar(255)" json:"nonce"`
	AuthorizationCode string `gorm:"type:varchar(255)" json:"authorization_code"`
}

type TokenAttempt struct {
	Base
	TokenAttemptBase
	TokenAttemptSecret
}
