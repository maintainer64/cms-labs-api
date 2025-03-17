package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries"
	"gitlab.com/a10869/api-modules/backend/app/usecases/response"
)

type LTIRoutingEditUC struct {
	LTIRoutingQueries *queries.LTIRoutingQueries
}

type LTIRoutingEditInputDTO struct {
	ID                   uint   `json:"id"`
	Name                 string `json:"name"`
	LTITitle             string `json:"lti_title"`
	LTIDescription       string `json:"lti_description"`
	LTITaskID            string `json:"lti_task_id"`
	LTICourseID          string `json:"lti_course_id"`
	LTISubID             string `json:"lti_sub_id"`
	LTIParamsTask        string `json:"lti_params_task"`
	Collaboration        int    `json:"collaboration"`
	PinnedSessionMinutes int    `json:"pinned_session_minutes"`
	PNETLabsType         string `json:"pnet_labs_type"`
	PNETLabsPath         string `json:"pnet_labs_path"`
	PNETTestPath         string `json:"pnet_test_path"`
	PNETServerID         uint   `json:"pnet_server_id"`
	IsDefault            bool   `json:"is_default"`
}

type LTIRoutingEditOutputDTO struct {
	ID uint `json:"id" required:"true"`
}

type LTIRoutingEditResponse = response.Response[LTIRoutingEditOutputDTO]

func (u *LTIRoutingEditUC) Execute(dto LTIRoutingEditInputDTO) (LTIRoutingEditOutputDTO, error) {
	entity := &models.LTIRouting{}
	entity.ID = dto.ID
	entity.Name = dto.Name
	entity.LTITitle = dto.LTITitle
	entity.LTIDescription = dto.LTIDescription
	entity.LTITaskID = dto.LTITaskID
	entity.LTICourseID = dto.LTICourseID
	entity.LTISubID = dto.LTISubID
	entity.LTIParamsTask = dto.LTIParamsTask
	entity.Collaboration = dto.Collaboration
	entity.PinnedSessionMinutes = dto.PinnedSessionMinutes
	entity.PNETLabsType = dto.PNETLabsType
	entity.PNETLabsPath = dto.PNETLabsPath
	entity.PNETTestPath = dto.PNETTestPath
	entity.PNETServerID = dto.PNETServerID
	entity.IsDefault = dto.IsDefault
	err := u.LTIRoutingQueries.Upsert(entity)
	return LTIRoutingEditOutputDTO{ID: entity.ID}, err
}
