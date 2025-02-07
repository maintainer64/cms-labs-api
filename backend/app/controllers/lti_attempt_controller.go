package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/di"
	"gitlab.com/a10869/api-modules/backend/app/usecases"
	"gitlab.com/a10869/api-modules/backend/app/usecases/auth"
	"gitlab.com/a10869/api-modules/shared/utils"
)

// LTIAttemptCreate func for create or get exists LTIAttempt.
// @Description Create or get exists lti_attempt. Roles: [admin, instructor, student]
// @Summary create lti_attempt
// @Tags LTIAttempt
// @Accept json
// @Produce json
// @Param form body usecases.LTIAttemptCreateInputDTO true "lti_attempt form info"
// @Success 200 {object} usecases.LTIAttemptCreateResponse
// @Security ApiKeyAuth
// @Router /v1/lti-attempt/create [post]
func LTIAttemptCreate(c *fiber.Ctx) error {
	token, err := auth.ExtractTokenMetadata(c, []string{})
	if err != nil {
		return err
	}
	dto := usecases.LTIAttemptCreateInputDTO{}
	err = utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	uc, err := di.NewDIContainer().LTIAttemptCreateUC()
	if err != nil {
		return err
	}
	output, err := uc.SetContext(token).Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}

// LTIAttemptEdit func for edit LTIAttempt.
// @Description Edit lti_attempt.
// @Summary edit lti_attempt
// @Tags LTIAttempt
// @Accept json
// @Produce json
// @Param form body usecases.LTIAttemptEditInputDTO true "lti_attempt form info"
// @Success 200 {object} usecases.LTIAttemptEditResponse
// @Security ApiKeyAuth
// @Router /v1/lti-attempt/edit [post]
func LTIAttemptEdit(c *fiber.Ctx) error {
	dto := usecases.LTIAttemptEditInputDTO{}
	err := utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	uc, err := di.NewDIContainer().LTIAttemptEditUC()
	if err != nil {
		return err
	}
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}

// LTIAttemptList func for view of list LTIAttempt.
// @Description List lti_attempt.
// @Summary list lti_attempt
// @Tags LTIAttempt
// @Accept json
// @Produce json
// @Param form body usecases.LTIAttemptListInputDTO true "lti_attempt list info"
// @Success 200 {object} usecases.LTIAttemptListResponse
// @Security ApiKeyAuth
// @Router /v1/lti-attempt/list [post]
func LTIAttemptList(c *fiber.Ctx) error {
	dto := usecases.LTIAttemptListInputDTO{}
	err := utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	uc, err := di.NewDIContainer().LTIAttemptListUC()
	if err != nil {
		return err
	}
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}

// LTIAttemptDelete func for delete LTIAttempt.
// @Description Delete lti_attempt.
// @Summary delete lti_attempt
// @Tags LTIAttempt
// @Accept json
// @Produce json
// @Param form body usecases.LTIAttemptDeleteInputDTO true "lti_attempt id"
// @Success 200 {object} usecases.LTIAttemptDeleteResponse
// @Security ApiKeyAuth
// @Router /v1/lti-attempt/delete [post]
func LTIAttemptDelete(c *fiber.Ctx) error {
	dto := usecases.LTIAttemptDeleteInputDTO{}
	err := utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	uc, err := di.NewDIContainer().LTIAttemptDeleteUC()
	if err != nil {
		return err
	}
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}

// LTIAttemptGet func for full model LTIAttempt.
// @Description get lti_attempt.
// @Summary get lti_attempt
// @Tags LTIAttempt
// @Accept json
// @Produce json
// @Param form body usecases.LTIAttemptGetInputDTO true "lti_attempt id"
// @Success 200 {object} usecases.LTIAttemptGetResponse
// @Security ApiKeyAuth
// @Router /v1/lti-attempt/get [post]
func LTIAttemptGet(c *fiber.Ctx) error {
	dto := usecases.LTIAttemptGetInputDTO{}
	err := utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	uc, err := di.NewDIContainer().LTIAttemptGetUC()
	if err != nil {
		return err
	}
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}
