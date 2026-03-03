package controllers

import (
	"gitlab.com/a10869/api-modules/backend/app/di"
	"gitlab.com/a10869/api-modules/backend/app/usecases/auth"
	"gitlab.com/a10869/api-modules/shared/cms_client"
	"gitlab.com/a10869/api-modules/shared/jsonrpc"
	"gitlab.com/a10869/api-modules/shared/logs"
)

// PNETServerQueueUpsert func for update on new queue RoundRobin.
// @Description Create pnet_server_queue. Roles: [admin, instructor]
// @Summary create pnet_server_queue
// @Tags server_queue
// @Accept json
// @Produce json
// @Param object body round_queue_pool_pnet.RoundQueuePoolPnetUpsertRequest true "request"
// @Success 200 {object} round_queue_pool_pnet.RoundQueuePoolPnetUpsertResponse
// @Security ApiKeyAuth
// @Router /api/v1/rpc/server_queue.upsert [post]
func PNETServerQueueUpsert(c *jsonrpc.Ctx) (interface{}, error) {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	if _, err := auth.ExtractTokenMetadata(
		c,
		[]string{cms_client.SSOUsersRoleAdmin, cms_client.SSOUsersRoleInstructor},
	); err != nil {
		return nil, err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return nil, err
	}
	defer container.Close()
	uc := container.PnetServerChangeDistributionUC()
	err = uc.Execute()
	if err != nil {
		return nil, err
	}
	return true, nil
}

// PNETServerQueueList func for list distribution queue
// @Description List pnet_server_queue. Roles: [admin, instructor]
// @Summary list pnet_server_queue
// @Tags server_queue
// @Accept json
// @Produce json
// @Param object body round_queue_pool_pnet.RoundQueuePoolPnetListRequest true "request"
// @Success 200 {object} round_queue_pool_pnet.RoundQueuePoolPnetListResponse
// @Security ApiKeyAuth
// @Router /api/v1/rpc/server_queue.list [post]
func PNETServerQueueList(c *jsonrpc.Ctx) (interface{}, error) {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	if _, err := auth.ExtractTokenMetadata(
		c,
		[]string{cms_client.SSOUsersRoleAdmin, cms_client.SSOUsersRoleInstructor},
	); err != nil {
		return nil, err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return nil, err
	}
	defer container.Close()
	uc := container.PnetServerListDistributionUC()
	output, err := uc.Execute()
	return output, err
}
