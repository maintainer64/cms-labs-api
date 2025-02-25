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

// ServiceCardCreate func for creates a new ServiceCard.
// @Description Create service_card. Roles: [admin, instructor]
// @Summary create service_card
// @Tags ServiceCard
// @Accept json
// @Produce json
// @Param form body usecases.ServiceCardEditInputDTO true "service_card form info"
// @Success 200 {object} usecases.ServiceCardEditResponse
// @Security ApiKeyAuth
// @Router /v1/service-card/upsert [post]
func ServiceCardCreate(c *fiber.Ctx) error {
	if _, err := auth.ExtractTokenMetadata(
		c,
		[]string{models.UsersRoleAdmin, models.UsersRoleInstructor},
	); err != nil {
		return err
	}
	diLoggerConf := logs.NewZeroLoggerConf(c)
	dto := usecases.ServiceCardEditInputDTO{}
	err := utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return err
	}
	defer container.Close()
	uc := container.ServiceCardEditUC()
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}

// ServiceCardList func for view of list ServiceCard.
// @Description List service_card. Roles: [admin, instructor]
// @Summary list service_card
// @Tags ServiceCard
// @Accept json
// @Produce json
// @Param form body usecases.ServiceCardListInputDTO true "service_card list info"
// @Success 200 {object} usecases.ServiceCardListResponse
// @Security ApiKeyAuth
// @Router /v1/service-card/list [post]
func ServiceCardList(c *fiber.Ctx) error {
	if _, err := auth.ExtractTokenMetadata(
		c,
		[]string{models.UsersRoleAdmin, models.UsersRoleInstructor},
	); err != nil {
		return err
	}
	diLoggerConf := logs.NewZeroLoggerConf(c)
	dto := usecases.ServiceCardListInputDTO{}
	err := utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return err
	}
	defer container.Close()
	uc := container.ServiceCardListUC()
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}

// ServiceCardDelete func for delete ServiceCard.
// @Description Delete service_card. Roles: [admin, instructor]
// @Summary delete service_card
// @Tags ServiceCard
// @Accept json
// @Produce json
// @Param form body usecases.ServiceCardDeleteInputDTO true "service_card id"
// @Success 200 {object} usecases.ServiceCardDeleteResponse
// @Security ApiKeyAuth
// @Router /v1/service-card/delete [post]
func ServiceCardDelete(c *fiber.Ctx) error {
	if _, err := auth.ExtractTokenMetadata(
		c,
		[]string{models.UsersRoleAdmin, models.UsersRoleInstructor},
	); err != nil {
		return err
	}
	diLoggerConf := logs.NewZeroLoggerConf(c)
	dto := usecases.ServiceCardDeleteInputDTO{}
	err := utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return err
	}
	defer container.Close()
	uc := container.ServiceCardDeleteUC()
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}

// ServiceCardGet func for full model ServiceCard.
// @Description get service_card. Roles: [admin, instructor]
// @Summary get service_card
// @Tags ServiceCard
// @Accept json
// @Produce json
// @Param form body usecases.ServiceCardGetInputDTO true "service_card id"
// @Success 200 {object} usecases.ServiceCardGetResponse
// @Security ApiKeyAuth
// @Router /v1/service-card/get [post]
func ServiceCardGet(c *fiber.Ctx) error {
	if _, err := auth.ExtractTokenMetadata(
		c,
		[]string{models.UsersRoleAdmin, models.UsersRoleInstructor},
	); err != nil {
		return err
	}
	diLoggerConf := logs.NewZeroLoggerConf(c)
	dto := usecases.ServiceCardGetInputDTO{}
	err := utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return err
	}
	defer container.Close()
	uc := container.ServiceCardGetUC()
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}
