package models

// LTIRoomBase struct to describe LTIRoom object.
type LTIRoomBase struct {
	RoomNumber int64 `gorm:"type:int" json:"room_number" validate:"required"`
}

type LTIRoom struct {
	Base
	LTIRoomBase
}
