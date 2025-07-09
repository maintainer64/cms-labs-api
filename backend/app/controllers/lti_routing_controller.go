package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/di"
	"gitlab.com/a10869/api-modules/backend/app/usecases"
	"gitlab.com/a10869/api-modules/backend/app/usecases/auth"
	"gitlab.com/a10869/api-modules/shared/cms_client"
	"gitlab.com/a10869/api-modules/shared/logs"
	"gitlab.com/a10869/api-modules/shared/utils"
)

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
		[]string{cms_client.SSOUsersRoleAdmin, cms_client.SSOUsersRoleAdmin},
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
