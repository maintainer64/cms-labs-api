package models

// LTIRoutingBase struct to describe LTIRouting object.
type LTIRoutingBase struct {
	Name string `gorm:"type:varchar(255);column:name" json:"name"`
}

type LTIRoutingSecret struct {
	// LTI Params
	LTITitle       string `gorm:"type:varchar(255);column:lti_title" json:"lti_title"`
	LTIDescription string `gorm:"type:varchar(255);column:lti_description" json:"lti_description"`
	LTITaskID      string `gorm:"type:varchar(255);column:lti_task_id" json:"lti_task_id"`
	LTIParamsTask  string `gorm:"type:varchar(255);column:lti_params_task" json:"lti_params_task"`
	// Параметры
	Collaboration        int    `gorm:"type:int;column:collaboration" json:"collaboration"`
	PinnedSessionMinutes int    `gorm:"type:int;column:pinned_session_minutes" json:"pinned_session_minutes"`
	PNETLabsPath         string `gorm:"type:varchar(255);column:pnet_labs_path" json:"pnet_labs_path"`
	PNETTestPath         string `gorm:"type:varchar(255);column:pnet_test_path" json:"pnet_test_path"`
}

type LTIRoutingListItem struct {
	Base
	LTIRoutingBase
}

// TableName переопределяет название таблицы для LTIRoutingListItem на `lti_routings`
func (LTIRoutingListItem) TableName() string {
	return "lti_routings"
}

type LTIRouting struct {
	Base
	LTIRoutingBase
	LTIRoutingSecret
}
