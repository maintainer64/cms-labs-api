package models

import "time"

// LTILaunchDataBase struct to describe LTILaunchData object.
type LTILaunchDataBase struct {
	ID        string    `gorm:"type:varchar(255)" json:"id"`
	LTIFormID uint      `json:"lti_form_id"`
	CreatedAt time.Time `gorm:"type:datetime(3)" json:"created_at" validate:"required"`
	UpdatedAt time.Time `gorm:"type:datetime(3)" json:"updated_at" validate:"required"`
}

type LTILaunchDataSecret struct {
	LaunchData string `gorm:"type:text" json:"launch_data"`
}

type LTILaunchData struct {
	LTILaunchDataBase
	LTILaunchDataSecret
}
