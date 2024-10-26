package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/di"
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/usecases"
	"gitlab.com/a10869/api-modules/backend/app/usecases/auth"
	"gitlab.com/a10869/api-modules/backend/pkg/utils"
)

// LTIFormCreate func for creates a new LTIForm.
// @Description Create lti_form. Roles: [admin]
// @Summary create lti_form
// @Tags LTIForm
// @Accept json
// @Produce json
// @Param form body usecases.LTIFormEditInputDTO true "lti_form form info"
// @Success 200 {object} usecases.LTIFormEditResponse
// @Security ApiKeyAuth
// @Router /v1/lti-form/upsert [post]
func LTIFormCreate(c *fiber.Ctx) error {
	if _, err := auth.ExtractTokenMetadata(
		c,
		[]string{models.UsersRoleAdmin},
	); err != nil {
		return err
	}
	dto := usecases.LTIFormEditInputDTO{}
	err := utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	uc, err := di.NewDIContainer().LTIFormEditUC()
	if err != nil {
		return err
	}
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}

// LTIFormList func for view of list LTIForm.
// @Description List lti_form. Roles: [admin]
// @Summary list lti_form
// @Tags LTIForm
// @Accept json
// @Produce json
// @Param form body usecases.LTIFormListInputDTO true "lti_form list info"
// @Success 200 {object} usecases.LTIFormEditResponse
// @Security ApiKeyAuth
// @Router /v1/lti-form/list [post]
func LTIFormList(c *fiber.Ctx) error {
	if _, err := auth.ExtractTokenMetadata(
		c,
		[]string{models.UsersRoleAdmin},
	); err != nil {
		return err
	}
	dto := usecases.LTIFormListInputDTO{}
	err := utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	uc, err := di.NewDIContainer().LTIFormListUC()
	if err != nil {
		return err
	}
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}

// LTIFormDelete func for delete LTIForm.
// @Description Delete lti_form. Roles: [admin]
// @Summary delete lti_form
// @Tags LTIForm
// @Accept json
// @Produce json
// @Param form body usecases.LTIFormDeleteInputDTO true "lti_form id"
// @Success 200 {object} usecases.LTIFormDeleteResponse
// @Security ApiKeyAuth
// @Router /v1/lti-form/delete [post]
func LTIFormDelete(c *fiber.Ctx) error {
	if _, err := auth.ExtractTokenMetadata(
		c,
		[]string{models.UsersRoleAdmin},
	); err != nil {
		return err
	}
	dto := usecases.LTIFormDeleteInputDTO{}
	err := utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	uc, err := di.NewDIContainer().LTIFormDeleteUC()
	if err != nil {
		return err
	}
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}

// LTIFormGet func for full model LTIForm.
// @Description get lti_form. Roles [admin]
// @Summary get lti_form
// @Tags LTIForm
// @Accept json
// @Produce json
// @Param form body usecases.LTIFormGetInputDTO true "lti_form id"
// @Success 200 {object} usecases.LTIFormGetResponse
// @Security ApiKeyAuth
// @Router /v1/lti-form/get [post]
func LTIFormGet(c *fiber.Ctx) error {
	if _, err := auth.ExtractTokenMetadata(
		c,
		[]string{models.UsersRoleAdmin},
	); err != nil {
		return err
	}
	dto := usecases.LTIFormGetInputDTO{}
	err := utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	uc, err := di.NewDIContainer().LTIFormGetUC()
	if err != nil {
		return err
	}
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}
