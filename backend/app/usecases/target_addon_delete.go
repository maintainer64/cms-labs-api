package usecases

import (
	"context"

	"github.com/goccy/go-json"
	"gitlab.com/a10869/api-modules/backend/app/addons"
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries"
	"gitlab.com/a10869/api-modules/shared/cms_client"
)

// TargetAddonDeleteUC – отключение дополнения от цели
type TargetAddonDeleteUC struct {
	TargetAddonQueries  *queries.TargetAddonQueries
	TargetQueries       *queries.TargetQueries
	TargetUserQueries   *queries.TargetUserQueries
	FactoryAddonService *addons.FactoryAddonService
	User                *cms_client.SSOTokenPublicData
}

// TargetAddonDeleteInputDTO – параметры удаления
type TargetAddonDeleteInputDTO struct {
	TargetID string `json:"target_id" validate:"required"`
	AddonID  string `json:"addon_id" validate:"required"`
	IssId    string `json:"iss_id"`
}

type TargetAddonDeleteOutputDTO struct {
}

type TargetAddonDeleteRequest struct {
	JSONRPC string                    `json:"jsonrpc" default:"2.0" required:"true"`
	Method  string                    `json:"method" default:"service_card.upsert" required:"true"`
	Params  TargetAddonDeleteInputDTO `json:"params,omitempty"`
	ID      string                    `json:"id,omitempty" default:"1" required:"true"`
}

type TargetAddonDeleteResponse struct {
	JSONRPC string                     `json:"jsonrpc" default:"2.0" required:"true"`
	Result  TargetAddonDeleteOutputDTO `json:"result,omitempty"`
	Error   interface{}                `json:"error,omitempty"`
	ID      string                     `json:"id,omitempty" default:"1" required:"true"`
}

func (uc *TargetAddonDeleteUC) SetContext(user *cms_client.SSOTokenPublicData) *TargetAddonDeleteUC {
	uc.User = user
	return uc
}

// Execute – отключает дополнение
func (uc *TargetAddonDeleteUC) Execute(dto TargetAddonDeleteInputDTO) (*TargetAddonDeleteOutputDTO, error) {
	uc.FactoryAddonService.SetContext(dto.IssId)
	if !uc.TargetUserQueries.CheckTargetAndUserByRole(dto.TargetID, uc.User.UserID(), models.UserRoleEditor) {
		return nil, ErrPermissionDenied
	}
	target, err := uc.TargetQueries.Get(dto.TargetID)
	if err != nil {
		return nil, err
	}
	addon, err := uc.TargetAddonQueries.GetByTargetAndAddon(dto.TargetID, dto.AddonID)
	if err != nil {
		return nil, err
	}
	addonService := uc.FactoryAddonService.GetAddonServiceByID(dto.AddonID)
	if addonService == nil {
		return nil, ErrAddonConfigNotFound
	}
	data, _ := json.Marshal(addon.Config)
	var configAddon addons.AddonOperationConfig
	_ = json.Unmarshal(data, &configAddon)
	ctx := context.Background()
	_, err = addonService.Delete(ctx, target.Name, &configAddon)
	if err != nil {
		return nil, err
	}
	return &TargetAddonDeleteOutputDTO{}, uc.TargetAddonQueries.Delete(addon.ID)
}
