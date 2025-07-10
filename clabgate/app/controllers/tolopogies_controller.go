package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/clabgate/app/di"
	"gitlab.com/a10869/api-modules/clabgate/app/usecases"
	"gitlab.com/a10869/api-modules/clabgate/app/usecases/auth"
	"gitlab.com/a10869/api-modules/shared/logs"
	"gitlab.com/a10869/api-modules/shared/utils"
)

// TopologiesGet func for view of list tasks.
// @Description List tasks. Roles: [student, admin, instructor]
// @Summary list tasks
// @Tags Tasks
// @Accept json
// @Produce json
// @Param form body usecases.TopologiesGetInputDTO true "topology namespace"
// @Success 200 {object} usecases.TopologiesGetResponse
// @Security ApiKeyAuth
// @Router /v1/topologies/get [post]
func TopologiesGet(c *fiber.Ctx) error {
	user, err := auth.ExtractTokenMetadata(c, []string{})
	if err != nil {
		return err
	}
	diLoggerConf := logs.NewZeroLoggerConf(c)
	dto := usecases.TopologiesGetInputDTO{}
	err = utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return err
	}
	defer container.Close()
	uc, err := container.TopologiesGetUC()
	if err != nil {
		return err
	}
	output, err := uc.SetContext(user).Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}

// TopologiesCreate func for create personal topologies.
// @Description List tasks. Roles: [student, admin, instructor]
// @Summary list tasks
// @Tags Tasks
// @Accept json
// @Produce json
// @Param form body usecases.TopologiesGetInputDTO true "topology namespace"
// @Success 200 {object} usecases.TopologiesGetResponse
// @Security ApiKeyAuth
// @Router /v1/topologies/create [post]
func TopologiesCreate(c *fiber.Ctx) error {
	user, err := auth.ExtractTokenMetadata(c, []string{})
	if err != nil {
		return err
	}
	diLoggerConf := logs.NewZeroLoggerConf(c)
	dto := usecases.TopologiesGetInputDTO{}
	err = utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return err
	}
	defer container.Close()
	uc, err := container.TopologiesGetUC()
	if err != nil {
		return err
	}
	output, err := uc.SetContext(user).Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}
