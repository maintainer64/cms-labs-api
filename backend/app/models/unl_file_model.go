package models

import (
	"time"
)

// UNLFileBase struct to describe UNLFile object.
type UNLFileBase struct {
	Path      string     `gorm:"uniqueIndex,type:text" json:"path" validate:"required"`
	Type      string     `gorm:"type:varchar(255)" json:"type"`
	DeletedAt *time.Time `gorm:"type:datetime(3)" json:"deleted_at"`
}

type UNLFileSecret struct {
	SyncedId string `gorm:"type:varchar(255)" json:"synced_id"`
	Content  []byte `gorm:"type:mediumblob"` // Содержимое файла
}

type UNLFileListItem struct {
	Base
	UNLFileBase
}

// TableName переопределяет название таблицы для UNLFileListItem на `unl_files`
func (UNLFileListItem) TableName() string {
	return "unl_files"
}

type UNLFile struct {
	Base
	UNLFileBase
	UNLFileSecret
}
