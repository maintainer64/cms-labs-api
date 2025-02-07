package models

import (
	"time"
)

type Base struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `gorm:"type:datetime(3)" json:"created_at" validate:"required"`
	UpdatedAt time.Time `gorm:"type:datetime(3)" json:"updated_at" validate:"required"`
}
