package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/di"
	"gitlab.com/a10869/api-modules/backend/app/usecases"
	"gitlab.com/a10869/api-modules/backend/app/usecases/auth"
	"gitlab.com/a10869/api-modules/backend/app/usecases/external"
	"gitlab.com/a10869/api-modules/shared/logs"
	"gitlab.com/a10869/api-modules/shared/utils"
)

// PNETServerCreate func for creates a new PNETServer.
// @Description Create pnet_server. Roles [admin]
// @Summary create pnet_server
// @Tags PNETServer
// @Accept json
// @Produce json
// @Param form body usecases.PNETServerEditInputDTO true "pnet_server form info"
// @Success 200 {object} usecases.PNETServerEditResponse
// @Security ApiKeyAuth
// @Router /v1/pnet-server/upsert [post]
func PNETServerCreate(c *fiber.Ctx) error {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	if _, err := auth.ExtractTokenMetadata(
		c,
		[]string{cms_client.SSOUsersRoleAdmin},
	); err != nil {
		return err
	}
	dto := usecases.PNETServerEditInputDTO{}
	err := utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return err
	}
	defer container.Close()
	uc := container.PNETServerEditUC()
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	ucDistribution := container.PnetServerChangeDistributionUC()
	err = ucDistribution.Execute()
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}

// PNETServerList func for view of list PNETServer.
// @Description List pnet_server. Roles: [admin, instructor]
// @Summary list pnet_server
// @Tags PNETServer
// @Accept json
// @Produce json
// @Param form body usecases.PNETServerListInputDTO true "pnet_server list info"
// @Success 200 {object} usecases.PNETServerListResponse
// @Security ApiKeyAuth
// @Router /v1/pnet-server/list [post]
func PNETServerList(c *fiber.Ctx) error {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	if _, err := auth.ExtractTokenMetadata(
		c,
		[]string{cms_client.SSOUsersRoleAdmin, cms_client.SSOUsersRoleAdmin},
	); err != nil {
		return err
	}
	dto := usecases.PNETServerListInputDTO{}
	err := utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return err
	}
	defer container.Close()
	uc := container.PNETServerListUC()
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}

// PNETServerDelete func for delete PNETServer.
// @Description Delete pnet_server. Roles: [admin]
// @Summary delete pnet_server
// @Tags PNETServer
// @Accept json
// @Produce json
// @Param form body usecases.PNETServerDeleteInputDTO true "pnet_server id"
// @Success 200 {object} usecases.PNETServerDeleteResponse
// @Security ApiKeyAuth
// @Router /v1/pnet-server/delete [post]
func PNETServerDelete(c *fiber.Ctx) error {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	if _, err := auth.ExtractTokenMetadata(
		c,
		[]string{cms_client.SSOUsersRoleAdmin},
	); err != nil {
		return err
	}
	dto := usecases.PNETServerDeleteInputDTO{}
	err := utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return err
	}
	defer container.Close()
	uc := container.PNETServerDeleteUC()
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	ucDistribution := container.PnetServerChangeDistributionUC()
	err = ucDistribution.Execute()
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}

// PNETServerGet func for full model PNETServer.
// @Description get pnet_server. Roles: [admin, instructor]
// @Summary get pnet_server
// @Tags PNETServer
// @Accept json
// @Produce json
// @Param form body usecases.PNETServerGetInputDTO true "pnet_server id"
// @Success 200 {object} usecases.PNETServerGetResponse
// @Security ApiKeyAuth
// @Router /v1/pnet-server/get [post]
func PNETServerGet(c *fiber.Ctx) error {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	if _, err := auth.ExtractTokenMetadata(
		c,
		[]string{cms_client.SSOUsersRoleAdmin, cms_client.SSOUsersRoleAdmin},
	); err != nil {
		return err
	}
	dto := usecases.PNETServerGetInputDTO{}
	err := utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return err
	}
	defer container.Close()
	uc := container.PNETServerGetUC()
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}

// PNETServerPing ping from external servers.
// @Description ing from external servers.
// @Summary ping from pnet_server
// @Tags PNETServer, EXTERNAL
// @Accept json
// @Produce json
// @Param form body external.PNETServerPingInputDTO true "pnet_server id"
// @Success 200 {object} external.PNETServerPingResponse
// @Param Authorization header string true "Basic-токен, созданный клиентом"
// @Router /v1/pnet-server/ping [post]
func PNETServerPing(c *fiber.Ctx) error {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	dto := external.PNETServerPingInputDTO{}
	err := utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return err
	}
	defer container.Close()
	uc := container.PNETServerPingUC()
	output, err := uc.SetContext(c.Locals("x-service-id").(string)).Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}
