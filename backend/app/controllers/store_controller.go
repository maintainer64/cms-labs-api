package controllers

import (
	"github.com/goccy/go-json"
	fiber "github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/di"
	"gitlab.com/a10869/api-modules/backend/app/models/types"
	"gitlab.com/a10869/api-modules/backend/app/usecases/auth"
	"gitlab.com/a10869/api-modules/shared/logs"
	"gitlab.com/a10869/api-modules/shared/utils"
)

// StoreGet func for get global store.
// @Description Global store of user get
// @Summary get store by user
// @Tags Store
// @Accept json
// @Produce json
// @Success 200 {object} types.UserStore
// @Security ApiKeyAuth
// @Router /v1/global-store [get]
func StoreGet(c *fiber.Ctx) error {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	user, err := auth.ExtractTokenMetadata(c, []string{})
	if err != nil {
		return err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return err
	}
	defer container.Close()
	store, err := container.Queries.UserQueries.StoreGetByUserId(user.UserID())
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(store)
}

// StoreSet func for get global store.
// @Description Global store of user get
// @Summary set store by user
// @Tags Store
// @Accept json
// @Produce json
// @Param form body types.UserStore true "store of create"
// @Success 200 {object} types.UserStore
// @Security ApiKeyAuth
// @Router /v1/global-store [post]
func StoreSet(c *fiber.Ctx) error {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	user, err := auth.ExtractTokenMetadata(c, []string{})
	if err != nil {
		return err
	}
	store := types.UserStore{}
	err = json.Unmarshal(c.Body(), &store)
	if err != nil {
		return utils.FiberValidationException{
			Status:    fiber.StatusBadRequest,
			Exception: err,
		}
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return err
	}
	defer container.Close()
	err = container.Queries.UserQueries.StoreSetByUserId(user.UserID(), store)
	return err
}
