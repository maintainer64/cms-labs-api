package models

import "time"

var (
	UsersRoleAdmin      string = "admin"
	UsersRoleInstructor string = "instructor"
	UsersRoleStudent    string = "student"
)

func UsersRoleValidate(role string) string {
	switch role {
	case UsersRoleAdmin:
		return UsersRoleAdmin
	case UsersRoleInstructor:
		return UsersRoleInstructor
	case UsersRoleStudent:
		return UsersRoleStudent
	default:
		return UsersRoleStudent
	}
}

// UserBase struct to describe User object.
type UserBase struct {
	Email     string     `gorm:"type:varchar(255)" json:"email" valid:"required,email"`
	Name      string     `gorm:"type:varchar(255)" json:"name" valid:"required"`
	UserRole  string     `gorm:"type:varchar(255)" json:"user_role" validate:"required"`
	GroupName string     `gorm:"type:varchar(255)" json:"group_name" validate:"required"`
	LTIUserID string     `gorm:"type:varchar(255)" json:"lti_user_id"`
	DeletedAt *time.Time `gorm:"type:datetime(3)" json:"deleted_at"`
}

type UserSecret struct {
	LastLaunchID string `gorm:"type:varchar(255)" json:"last_launch_id"`
}

type UserListItem struct {
	Base
	UserBase
}

// TableName переопределяет название таблицы для UserListItem на `users`
func (UserListItem) TableName() string {
	return "users"
}

type UserListCount struct {
	Count int64 `json:"count"`
}

// TableName переопределяет название таблицы для UserListItem на `users`
func (UserListCount) TableName() string {
	return "users"
}

type User struct {
	Base
	UserBase
	UserSecret
}

func (c *UserBase) IsActive() bool {
	zeroTime := time.Time{}
	return c.DeletedAt == nil || *c.DeletedAt == zeroTime
}
