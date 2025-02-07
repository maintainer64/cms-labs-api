package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/di"
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/usecases/auth"
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
	if _, err := auth.ExtractTokenMetadata(
		c,
		[]string{models.UsersRoleAdmin, models.UsersRoleInstructor},
	); err != nil {
		return err
	}
	uc, err := di.NewDIContainer().PnetServerChangeDistributionUC()
	if err != nil {
		return err
	}
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
	if _, err := auth.ExtractTokenMetadata(
		c,
		[]string{models.UsersRoleAdmin, models.UsersRoleInstructor},
	); err != nil {
		return err
	}
	uc, err := di.NewDIContainer().PnetServerListDistributionUC()
	if err != nil {
		return err
	}
	output, err := uc.Execute()
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}
