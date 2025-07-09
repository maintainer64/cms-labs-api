package models

// RoleBase struct to describe Role object.
type RoleBase struct {
	Code string `gorm:"type:varchar(255)" json:"code" valid:"required"`
	Name string `gorm:"type:varchar(255)" json:"name" valid:"required"`
}

type Role struct {
	Base
	RoleBase
}

// RoleRelationBase struct to describe RoleRelation object.
type RoleRelationBase struct {
	RoleID   uint  `gorm:"type:int" json:"role_id"`
	UserID   *uint `gorm:"type:int" json:"user_id"`
	ServerID *uint `gorm:"type:int" json:"server_id"`
}

type RoleRelation struct {
	Base
	RoleRelationBase
}
