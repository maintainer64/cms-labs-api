package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/pnetlabaddon/app/di"
	"gitlab.com/a10869/api-modules/shared/logs"
	"gitlab.com/a10869/api-modules/shared/utils"
)

// PNETServerPing Пинг в core-backend для синхронизации попыток.
// @Description Пинг в core-backend для синхронизации попыток.
// @Summary Пинг в core-backend для синхронизации попыток.
// @Tags Task
// @Accept json
// @Produce json
// @Success 200
// @Router /v1/task/pnet-server-ping [post]
func PNETServerPing(c *fiber.Ctx) error {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return utils.FiberValidationException{
			Status:    fiber.StatusInternalServerError,
			Exception: err,
		}
	}
	defer container.Close()

	client := container.PnetServerPingUC()
	err = client.Execute()

	if err != nil {
		return utils.FiberValidationException{
			Status:    fiber.StatusInternalServerError,
			Exception: err,
		}
	}
	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{})
}
