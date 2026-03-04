package usecases

import (
	"context"
	"errors"
	"fmt"

	"github.com/rs/zerolog"
	"gitlab.com/a10869/api-modules/backend/app/addons/vault"
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries"
)

// TargetUserRotateAddonsUC – перевыпуск и обновление всех прав для пользователя
type TargetUserRotateAddonsUC struct {
	TargetUserQueries *queries.TargetUserQueries
	TargetQueries     *queries.TargetQueries
	UserQueries       *queries.UserQueries
	VaultClient       vault.ClientInterface
	Logger            *zerolog.Logger
}

// Execute – удаляет пользователя из цели
func (uc *TargetUserRotateAddonsUC) Execute(userID uint) error {
	uc.Logger.Error().Msg(fmt.Sprintf(
		"TargetUserRotateAddonsUC rotate permissions on by users %d",
		userID,
	))
	user, err := uc.UserQueries.Get(userID)
	if err != nil && !errors.Is(err, queries.UserNotActive) {
		uc.Logger.Info().Msg(fmt.Sprintf(
			"TargetUserRotateAddonsUC get by userId %d %+v", userID, err,
		))
		return err
	}
	targetsIds := make([]string, 0)
	if user.IsActive() {
		targetsIds, err = uc.TargetUserQueries.GetByUserId(userID)
	} else {
		uc.Logger.Info().Msg(fmt.Sprintf(
			"TargetUserRotateAddonsUC user %d email %s is not active",
			userID,
			user.Email,
		))
	}
	if err != nil {
		uc.Logger.Info().Msg(fmt.Sprintf(
			"TargetUserRotateAddonsUC error get targets by userId %d %+v", userID, err,
		))
		return err
	}
	uc.Logger.Info().Msg(fmt.Sprintf(
		"TargetUserRotateAddonsUC get targets %d count by userId %d",
		len(targetsIds),
		userID,
	))
	vaultBind := vault.UsersAndServices{
		UserEmail: user.Email,
		Services:  make([]vault.ServicesBindUser, 0),
	}
	for _, targetID := range targetsIds {
		uc.Logger.Info().Msg(fmt.Sprintf(
			"TargetUserRotateAddonsUC get target %s and userId %d",
			targetID,
			userID,
		))
		target, err := uc.TargetQueries.Get(targetID)
		if err != nil {
			uc.Logger.Error().Msg(fmt.Sprintf(
				"TargetUserRotateAddonsUC error target %s and userId %d %+v", targetID, userID, err,
			))
			continue
		}
		vaultBindServices := vault.ServicesBindUser{
			ServiceName: target.Name,
			Policies:    make([]string, 0),
		}
		if uc.TargetUserQueries.CheckTargetAndUserByRole(
			targetID,
			userID,
			models.UserRoleVaultViewer,
		) {
			vaultBindServices.Policies = append(vaultBindServices.Policies, vault.Read)
		}
		if uc.TargetUserQueries.CheckTargetAndUserByRole(
			targetID,
			userID,
			models.UserRoleVaultWriter,
		) {
			vaultBindServices.Policies = append(vaultBindServices.Policies, vault.Write)
		}
		vaultBind.Services = append(vaultBind.Services, vaultBindServices)
	}
	ctx := context.Background()
	err = uc.VaultClient.UserBindAccessServices(ctx, []vault.UsersAndServices{vaultBind})
	if err != nil {
		uc.Logger.Error().Msg(
			fmt.Sprintf(
				"TargetUserRotateAddonsUC error with vault by userId %d %+v",
				userID,
				err,
			),
		)
		return err
	}
	return nil
}
