package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/di"
	"gitlab.com/a10869/api-modules/backend/app/usecases"
	"gitlab.com/a10869/api-modules/backend/pkg/utils"
)

// CreatePNETServer func for creates a new PNETServer connection.
// @Description Create a pnet server connection.
// @Summary create PNETServer integration
// @Tags PNETServer
// @Accept json
// @Produce json
// @Param form body usecases.PNETServerEditInputDTO true "pnet server info"
// @Success 200 {object} usecases.PNETServerEditResponse
// @Security ApiKeyAuth
// @Router /v1/pnet-server/upsert [post]
func CreatePNETServer(c *fiber.Ctx) error {
	dto := usecases.PNETServerEditInputDTO{}
	err := utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	uc, err := di.NewDIContainer().PNETServerEditUC()
	if err != nil {
		return err
	}
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}

// ListPNETServer func for view of list PNET Servers integration.
// @Description List pnet server integration.
// @Summary list pnet server integration
// @Tags PNETServer
// @Accept json
// @Produce json
// @Param form body usecases.PNETServerListInputDTO true "pnet-server list info"
// @Success 200 {object} usecases.PNETServerEditResponse
// @Security ApiKeyAuth
// @Router /v1/pnet-server/list [post]
func ListPNETServer(c *fiber.Ctx) error {
	dto := usecases.PNETServerListInputDTO{}
	err := utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	uc, err := di.NewDIContainer().PNETServerListUC()
	if err != nil {
		return err
	}
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}

// DeletePNETServer func for delete PNET server integration.
// @Description Delete pnet-server integration.
// @Summary delete pnet-server integration
// @Tags PNETServer
// @Accept json
// @Produce json
// @Param form body usecases.PNETServerDeleteInputDTO true "pnet-server id"
// @Success 200 {object} usecases.PNETServerDeleteResponse
// @Security ApiKeyAuth
// @Router /v1/pnet-server/delete [post]
func DeletePNETServer(c *fiber.Ctx) error {
	dto := usecases.PNETServerDeleteInputDTO{}
	err := utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	uc, err := di.NewDIContainer().PNETServerDeleteUC()
	if err != nil {
		return err
	}
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}

// GetPNETServer func for full model PNETServer integration.
// @Description get lti integration.
// @Summary get lti integration
// @Tags PNETServer
// @Accept json
// @Produce json
// @Param form body usecases.PNETServerGetInputDTO true "pnet-server id"
// @Success 200 {object} usecases.PNETServerGetResponse
// @Security ApiKeyAuth
// @Router /v1/pnet-server/get [post]
func GetPNETServer(c *fiber.Ctx) error {
	dto := usecases.PNETServerGetInputDTO{}
	err := utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	uc, err := di.NewDIContainer().PNETServerGetUC()
	if err != nil {
		return err
	}
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}
