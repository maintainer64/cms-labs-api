package models

import "time"

// UserToken struct to describe Save utils.TokenPublicData object.
type UserToken struct {
	UserID       uint      `gorm:"primaryKey" json:"user_id"`
	RefreshToken string    `gorm:"type:varchar(255)" json:"refresh_token"`
	HashPassword string    `gorm:"type:varchar(255)" json:"hash_password"`
	CreatedAt    time.Time `gorm:"type:datetime(3)" json:"created_at" validate:"required"`
	UpdatedAt    time.Time `gorm:"type:datetime(3)" json:"updated_at" validate:"required"`
}
