package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/di"
	"gitlab.com/a10869/api-modules/backend/app/usecases/auth"
	"gitlab.com/a10869/api-modules/shared/cms_client"
	"gitlab.com/a10869/api-modules/shared/logs"
	"gitlab.com/a10869/api-modules/shared/utils"
)

// PNETServerQueueCreate func for update on new queue RoundRobin.
// @Description Create pnet_server_queue. Roles: [admin, instructor]
// @Summary create pnet_server_queue
// @Tags PNETServerQueue
// @Accept json
// @Produce json
// @Success 200 {object} round_queue_pool_pnet.RoundQueuePoolPnetUpsertResponse
// @Security ApiKeyAuth
// @Router /v1/pnet-server-queue/upsert [post]
func PNETServerQueueCreate(c *fiber.Ctx) error {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	if _, err := auth.ExtractTokenMetadata(
		c,
		[]string{cms_client.SSOUsersRoleAdmin, cms_client.SSOUsersRoleAdmin},
	); err != nil {
		return err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return err
	}
	defer container.Close()
	uc := container.PnetServerChangeDistributionUC()
	err = uc.Execute()
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: true}
}

// PNETServerQueueList func for list distribution queue
// @Description List pnet_server_queue. Roles: [admin, instructor]
// @Summary list pnet_server_queue
// @Tags PNETServerQueue
// @Accept json
// @Produce json
// @Success 200 {object} round_queue_pool_pnet.RoundQueuePoolPnetListResponse
// @Security ApiKeyAuth
// @Router /v1/pnet-server-queue/list [post]
func PNETServerQueueList(c *fiber.Ctx) error {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	if _, err := auth.ExtractTokenMetadata(
		c,
		[]string{cms_client.SSOUsersRoleAdmin, cms_client.SSOUsersRoleAdmin},
	); err != nil {
		return err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return err
	}
	defer container.Close()
	uc := container.PnetServerListDistributionUC()
	output, err := uc.Execute()
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}
