package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/di"
	"gitlab.com/a10869/api-modules/backend/app/usecases"
	"gitlab.com/a10869/api-modules/backend/pkg/utils"
)

// ServiceCardCreate func for creates a new ServiceCard.
// @Description Create service_card.
// @Summary create service_card
// @Tags ServiceCard
// @Accept json
// @Produce json
// @Param form body usecases.ServiceCardEditInputDTO true "service_card form info"
// @Success 200 {object} usecases.ServiceCardEditResponse
// @Security ApiKeyAuth
// @Router /v1/service-card/upsert [post]
func ServiceCardCreate(c *fiber.Ctx) error {
	dto := usecases.ServiceCardEditInputDTO{}
	err := utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	uc, err := di.NewDIContainer().ServiceCardEditUC()
	if err != nil {
		return err
	}
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}

// ServiceCardList func for view of list ServiceCard.
// @Description List service_card.
// @Summary list service_card
// @Tags ServiceCard
// @Accept json
// @Produce json
// @Param form body usecases.ServiceCardListInputDTO true "service_card list info"
// @Success 200 {object} usecases.ServiceCardListResponse
// @Security ApiKeyAuth
// @Router /v1/service-card/list [post]
func ServiceCardList(c *fiber.Ctx) error {
	dto := usecases.ServiceCardListInputDTO{}
	err := utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	uc, err := di.NewDIContainer().ServiceCardListUC()
	if err != nil {
		return err
	}
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}

// ServiceCardDelete func for delete ServiceCard.
// @Description Delete service_card.
// @Summary delete service_card
// @Tags ServiceCard
// @Accept json
// @Produce json
// @Param form body usecases.ServiceCardDeleteInputDTO true "service_card id"
// @Success 200 {object} usecases.ServiceCardDeleteResponse
// @Security ApiKeyAuth
// @Router /v1/service-card/delete [post]
func ServiceCardDelete(c *fiber.Ctx) error {
	dto := usecases.ServiceCardDeleteInputDTO{}
	err := utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	uc, err := di.NewDIContainer().ServiceCardDeleteUC()
	if err != nil {
		return err
	}
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}

// ServiceCardGet func for full model ServiceCard.
// @Description get service_card.
// @Summary get service_card
// @Tags ServiceCard
// @Accept json
// @Produce json
// @Param form body usecases.ServiceCardGetInputDTO true "service_card id"
// @Success 200 {object} usecases.ServiceCardGetResponse
// @Security ApiKeyAuth
// @Router /v1/service-card/get [post]
func ServiceCardGet(c *fiber.Ctx) error {
	dto := usecases.ServiceCardGetInputDTO{}
	err := utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	uc, err := di.NewDIContainer().ServiceCardGetUC()
	if err != nil {
		return err
	}
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}
