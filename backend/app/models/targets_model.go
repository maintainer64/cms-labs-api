package models

import (
	"time"

	"gitlab.com/a10869/api-modules/backend/app/models/types"
	"gitlab.com/a10869/api-modules/shared/connection"
)

// роль пользователя в TargetUser

const (
	UserRoleVaultViewer = "vault_viewer"
	UserRoleVaultWriter = "vault_writer"
	UserRoleEditor      = "editor"
	UserRoleNominal     = "nominal"
)

// TargetType — тип целевого объекта (сервис, сервер, модуль, виртуальный)
const (
	TargetTypeService = "SERVICE"
	TargetTypeServer  = "SERVER"
	TargetTypeModule  = "MODULE"
	TargetTypeVirtual = "VIRTUAL"
)

// Target представляет собой сервис, сервер, модуль или виртуальный объект
type Target struct {
	ID           string          `gorm:"type:varchar(255);primaryKey" json:"id"`
	Type         string          `gorm:"type:varchar(50);index;not null" json:"type"`
	Name         string          `gorm:"type:varchar(255);index;not null" json:"name"`
	Description  *string         `gorm:"type:text" json:"description"`
	Links        types.JsonStore `gorm:"type:json" json:"links"`         // список ссылок
	Tags         types.JsonStore `gorm:"type:json" json:"tags"`          // массив тегов
	InternalTags types.JsonStore `gorm:"type:json" json:"internal_tags"` // массив внутренних тегов
	CreatedAt    time.Time       `gorm:"type:datetime(3)" json:"created_at" validate:"required"`
	UpdatedAt    time.Time       `gorm:"type:datetime(3)" json:"updated_at" validate:"required"`
}

// TargetRelation представляет связь «многие‑ко‑многим» между двумя Target
type TargetRelation struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	FromTargetID string    `gorm:"column:from_target_id;type:varchar(255);index;not null" json:"from_target_id"`
	ToTargetID   string    `gorm:"column:to_target_id;type:varchar(255);index;not null" json:"to_target_id"`
	RelationType string    `gorm:"column:relation_type;type:varchar(255);not null" json:"relation_type"`
	CreatedAt    time.Time `gorm:"type:datetime(3)" json:"created_at" validate:"required"`
	UpdatedAt    time.Time `gorm:"type:datetime(3)" json:"updated_at" validate:"required"`
}

type TargetUserBase struct {
	TargetID string          `gorm:"type:varchar(255);index;not null" json:"target_id"`
	UserID   uint            `gorm:"type:int" json:"user_id"`
	Roles    types.JsonStore `gorm:"type:json;not null" json:"roles"` // массив строк UserRole
}

type TargetAddonBase struct {
	TargetID  string               `gorm:"column:target_id;type:varchar(255);index;not null" json:"target_id"`
	AddonType connection.AddonType `gorm:"column:addon_type;type:varchar(50);not null" json:"addon_type"`
	AddonID   string               `gorm:"column:addon_id;type:varchar(36);not null" json:"addon_id"`
	Config    types.JsonStore      `gorm:"type:json;not null" json:"config"`
}

type TargetAddon struct {
	Base
	TargetAddonBase
}

type TargetUser struct {
	Base
	TargetUserBase
}
