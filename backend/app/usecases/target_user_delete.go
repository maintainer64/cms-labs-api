package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries"
	"gitlab.com/a10869/api-modules/shared/cms_client"
)

// TargetUserDeleteUC – удаление пользователя из цели
type TargetUserDeleteUC struct {
	TargetUserQueries *queries.TargetUserQueries
	TargetQueries     *queries.TargetQueries
	User              *cms_client.SSOTokenPublicData
}

// TargetUserDeleteInputDTO – параметры удаления
type TargetUserDeleteInputDTO struct {
	TargetID string `json:"target_id" validate:"required"`
	UserID   uint   `json:"user_id" validate:"required"`
}

type TargetUserDeleteOutputDTO struct {
}

type TargetUserDeleteRequest struct {
	JSONRPC string                   `json:"jsonrpc" default:"2.0" required:"true"`
	Method  string                   `json:"method" default:"service_card.upsert" required:"true"`
	Params  TargetUserDeleteInputDTO `json:"params,omitempty"`
	ID      string                   `json:"id,omitempty" default:"1" required:"true"`
}

type TargetUserDeleteResponse struct {
	JSONRPC string                    `json:"jsonrpc" default:"2.0" required:"true"`
	Result  TargetUserDeleteOutputDTO `json:"result,omitempty"`
	Error   interface{}               `json:"error,omitempty"`
	ID      string                    `json:"id,omitempty" default:"1" required:"true"`
}

func (uc *TargetUserDeleteUC) SetContext(user *cms_client.SSOTokenPublicData) *TargetUserDeleteUC {
	uc.User = user
	return uc
}

// Execute – удаляет пользователя из цели
func (uc *TargetUserDeleteUC) Execute(dto TargetUserDeleteInputDTO) (*TargetUserDeleteOutputDTO, error) {
	// 1. Проверка прав: удаляющий должен быть редактором цели
	if !uc.TargetUserQueries.CheckTargetAndUserByRole(dto.TargetID, uc.User.UserID(), models.UserRoleEditor) {
		return nil, ErrPermissionDenied
	}
	// 4. Удаляем из БД
	err := uc.TargetUserQueries.DeleteByTargetAndUser(dto.TargetID, dto.UserID)
	if err != nil {
		return nil, err
	}

	// TODO: Синхронизация прав всех пользоваетелей сервиса

	return &TargetUserDeleteOutputDTO{}, nil
}
