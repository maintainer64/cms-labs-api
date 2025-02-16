package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/di"
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/usecases"
	"gitlab.com/a10869/api-modules/backend/app/usecases/auth"
	"gitlab.com/a10869/api-modules/shared/logs"
	"gitlab.com/a10869/api-modules/shared/utils"
)

// LTIRoutingCreate func for creates a new LTIRouting.
// @Description Create lti_routing. Roles [admin, instructor]
// @Summary create lti_routing
// @Tags LTIRouting
// @Accept json
// @Produce json
// @Param form body usecases.LTIRoutingEditInputDTO true "lti_routing form info"
// @Success 200 {object} usecases.LTIRoutingEditResponse
// @Security ApiKeyAuth
// @Router /v1/lti-routing/upsert [post]
func LTIRoutingCreate(c *fiber.Ctx) error {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	if _, err := auth.ExtractTokenMetadata(
		c,
		[]string{models.UsersRoleAdmin, models.UsersRoleInstructor},
	); err != nil {
		return err
	}
	dto := usecases.LTIRoutingEditInputDTO{}
	err := utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return err
	}
	defer container.Close()
	uc := container.LTIRoutingEditUC()
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}

// LTIRoutingList func for view of list LTIRouting.
// @Description List lti_routing. Roles [admin, instructor]
// @Summary list lti_routing
// @Tags LTIRouting
// @Accept json
// @Produce json
// @Param form body usecases.LTIRoutingListInputDTO true "lti_routing list info"
// @Success 200 {object} usecases.LTIRoutingListResponse
// @Security ApiKeyAuth
// @Router /v1/lti-routing/list [post]
func LTIRoutingList(c *fiber.Ctx) error {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	if _, err := auth.ExtractTokenMetadata(
		c,
		[]string{models.UsersRoleAdmin, models.UsersRoleInstructor},
	); err != nil {
		return err
	}
	dto := usecases.LTIRoutingListInputDTO{}
	err := utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return err
	}
	defer container.Close()
	uc := container.LTIRoutingListUC()
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}

// LTIRoutingDelete func for delete LTIRouting.
// @Description Delete lti_routing. Roles [admin, instructor]
// @Summary delete lti_routing
// @Tags LTIRouting
// @Accept json
// @Produce json
// @Param form body usecases.LTIRoutingDeleteInputDTO true "lti_routing id"
// @Success 200 {object} usecases.LTIRoutingDeleteResponse
// @Security ApiKeyAuth
// @Router /v1/lti-routing/delete [post]
func LTIRoutingDelete(c *fiber.Ctx) error {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	if _, err := auth.ExtractTokenMetadata(
		c,
		[]string{models.UsersRoleAdmin, models.UsersRoleInstructor},
	); err != nil {
		return err
	}
	dto := usecases.LTIRoutingDeleteInputDTO{}
	err := utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return err
	}
	defer container.Close()
	uc := container.LTIRoutingDeleteUC()
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}

// LTIRoutingGet func for full model LTIRouting.
// @Description get lti_routing. Roles [admin, instructor]
// @Summary get lti_routing
// @Tags LTIRouting
// @Accept json
// @Produce json
// @Param form body usecases.LTIRoutingGetInputDTO true "lti_routing id"
// @Success 200 {object} usecases.LTIRoutingGetResponse
// @Security ApiKeyAuth
// @Router /v1/lti-routing/get [post]
func LTIRoutingGet(c *fiber.Ctx) error {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	if _, err := auth.ExtractTokenMetadata(
		c,
		[]string{models.UsersRoleAdmin, models.UsersRoleInstructor},
	); err != nil {
		return err
	}
	dto := usecases.LTIRoutingGetInputDTO{}
	err := utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return err
	}
	defer container.Close()
	uc := container.LTIRoutingGetUC()
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}
