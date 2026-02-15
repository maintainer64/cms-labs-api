package usecases

import (
	"errors"
	"time"

	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/models/types"
	"gitlab.com/a10869/api-modules/backend/app/queries"
	"gitlab.com/a10869/api-modules/shared/cms_client"
)

// TargetUserUpsertUC – добавление или обновление пользователя цели
type TargetUserUpsertUC struct {
	TargetUserQueries *queries.TargetUserQueries
	TargetQueries     *queries.TargetQueries
	User              *cms_client.SSOTokenPublicData
}

// TargetUserUpsertInputDTO – входные данные
type TargetUserUpsertInputDTO struct {
	TargetID string   `json:"target_id" validate:"required"`
	UserID   uint     `json:"user_id" validate:"required"`
	Roles    []string `json:"roles" validate:"required"`
}

// TargetUserUpsertOutputDTO – результат
type TargetUserUpsertOutputDTO struct {
	ID uint `json:"id"`
}

type TargetUserUpsertRequest struct {
	JSONRPC string                   `json:"jsonrpc" default:"2.0" required:"true"`
	Method  string                   `json:"method" default:"service_card.upsert" required:"true"`
	Params  TargetUserUpsertInputDTO `json:"params,omitempty"`
	ID      string                   `json:"id,omitempty" default:"1" required:"true"`
}

type TargetUserUpsertResponse struct {
	JSONRPC string                    `json:"jsonrpc" default:"2.0" required:"true"`
	Result  TargetUserUpsertOutputDTO `json:"result,omitempty"`
	Error   interface{}               `json:"error,omitempty"`
	ID      string                    `json:"id,omitempty" default:"1" required:"true"`
}

func (uc *TargetUserUpsertUC) SetContext(user *cms_client.SSOTokenPublicData) *TargetUserUpsertUC {
	uc.User = user
	return uc
}

// Execute – основной метод
func (uc *TargetUserUpsertUC) Execute(dto TargetUserUpsertInputDTO) (*TargetUserUpsertOutputDTO, error) {
	// 1. Проверка прав: текущий пользователь должен быть редактором цели
	hasEditor, err := uc.TargetUserQueries.HasEditor(dto.TargetID)
	if err != nil {
		return nil, err
	}
	if hasEditor && !uc.TargetUserQueries.CheckTargetAndUserByRole(
		dto.TargetID,
		uc.User.UserID(),
		models.UserRoleEditor,
	) {
		return nil, ErrPermissionDenied
	}

	// 3. Сериализуем роли в JSON
	rolesJSON := types.JsonStore{
		"roles": dto.Roles,
	}

	// 4. Ищем существующую запись
	existing, err := uc.TargetUserQueries.GetByTargetAndUser(dto.TargetID, dto.UserID)
	if err != nil && !errors.Is(err, queries.TargetUserNotFoundError) {
		return nil, err
	}

	// 6. Создаём или обновляем сущность
	entity := &models.TargetUser{
		TargetUserBase: models.TargetUserBase{
			TargetID: dto.TargetID,
			UserID:   dto.UserID,
			Roles:    rolesJSON,
		},
	}
	if existing.ID != 0 {
		entity.ID = existing.ID
		entity.CreatedAt = existing.CreatedAt
	}
	entity.UpdatedAt = time.Now().UTC()

	if err := uc.TargetUserQueries.Upsert(entity); err != nil {
		return nil, err
	}

	// TODO: Здесь синхронизировать права всех пользователей сервиса

	return &TargetUserUpsertOutputDTO{ID: entity.ID}, nil
}
