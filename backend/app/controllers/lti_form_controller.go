package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/di"
	"gitlab.com/a10869/api-modules/backend/app/usecases"
	"gitlab.com/a10869/api-modules/backend/pkg/utils"
)

// CreateLTIForm func for creates a new LTI integration.
// @Description Create a lti integration.
// @Summary create lti integration
// @Tags LTIForm
// @Accept json
// @Produce json
// @Param form body usecases.LtiFormEditInputDTO true "lti form info"
// @Success 200 {object} usecases.LtiFormEditResponse
// @Security ApiKeyAuth
// @Router /v1/lti-form/upsert [post]
func CreateLTIForm(c *fiber.Ctx) error {
	dto := usecases.LtiFormEditInputDTO{}
	err := utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	uc, err := di.NewDIContainer().LtiFormEditUC()
	if err != nil {
		return err
	}
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}

// ListLTIForm func for view of list LTI integration.
// @Description List lti integration.
// @Summary list lti integration
// @Tags LTIForm
// @Accept json
// @Produce json
// @Param form body usecases.LtiFormListInputDTO true "lti list info"
// @Success 200 {object} usecases.LtiFormEditResponse
// @Security ApiKeyAuth
// @Router /v1/lti-form/list [post]
func ListLTIForm(c *fiber.Ctx) error {
	dto := usecases.LtiFormListInputDTO{}
	err := utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	uc, err := di.NewDIContainer().LtiFormListUC()
	if err != nil {
		return err
	}
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}

// DeleteLTIForm func for delete LTI integration.
// @Description Delete lti integration.
// @Summary delete lti integration
// @Tags LTIForm
// @Accept json
// @Produce json
// @Param form body usecases.LtiFormDeleteInputDTO true "lti id"
// @Success 200 {object} usecases.LtiFormDeleteResponse
// @Security ApiKeyAuth
// @Router /v1/lti-form/delete [post]
func DeleteLTIForm(c *fiber.Ctx) error {
	dto := usecases.LtiFormDeleteInputDTO{}
	err := utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	uc, err := di.NewDIContainer().LtiFormDeleteUC()
	if err != nil {
		return err
	}
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}

// GetLTIForm func for full model LTI integration.
// @Description get lti integration.
// @Summary get lti integration
// @Tags LTIForm
// @Accept json
// @Produce json
// @Param form body usecases.LtiFormGetInputDTO true "lti id"
// @Success 200 {object} usecases.LtiFormGetResponse
// @Security ApiKeyAuth
// @Router /v1/lti-form/get [post]
func GetLTIForm(c *fiber.Ctx) error {
	dto := usecases.LtiFormGetInputDTO{}
	err := utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	uc, err := di.NewDIContainer().LtiFormGetUC()
	if err != nil {
		return err
	}
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}
