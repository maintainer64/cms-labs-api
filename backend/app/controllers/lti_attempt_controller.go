package controllers

import (
	"gitlab.com/a10869/api-modules/backend/app/di"
	"gitlab.com/a10869/api-modules/backend/app/usecases"
	"gitlab.com/a10869/api-modules/backend/app/usecases/auth"
	"gitlab.com/a10869/api-modules/shared/cms_client"
	"gitlab.com/a10869/api-modules/shared/jsonrpc"
	"gitlab.com/a10869/api-modules/shared/logs"
)

// LTIAttemptCreate func for create or get exists LTIAttempt.
// @Description Create or get exists lti_attempt. Roles: [admin, instructor, student]
// @Summary create lti_attempt
// @Tags lti_attempt
// @Accept json
// @Produce json
// @Param form body usecases.LTIAttemptCreateRequest true "lti_attempt form info"
// @Success 200 {object} usecases.LTIAttemptCreateResponse
// @Security ApiKeyAuth
// @Router /api/v1/rpc/lti_attempt.create [post]
func LTIAttemptCreate(c *jsonrpc.Ctx) (interface{}, error) {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	token, err := auth.ExtractTokenMetadata(c, []string{})
	if err != nil {
		return nil, err
	}
	dto := usecases.LTIAttemptCreateInputDTO{}
	err = jsonrpc.ValidatorBase(c, &dto)
	if err != nil {
		return nil, err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return nil, err
	}
	defer container.Close()
	uc := container.LTIAttemptCreateUC()
	output, err := uc.SetContext(token).Execute(dto)
	return output, err
}

// LTIAttemptUpdate func for edit LTIAttempt.
// @Description Edit lti_attempt. Roles: [admin, instructor]
// @Summary edit lti_attempt
// @Tags lti_attempt
// @Accept json
// @Produce json
// @Param form body usecases.LTIAttemptEditRequest true "lti_attempt form info"
// @Success 200 {object} usecases.LTIAttemptEditResponse
// @Security ApiKeyAuth
// @Router /api/v1/rpc/lti_attempt.update [post]
func LTIAttemptUpdate(c *jsonrpc.Ctx) (interface{}, error) {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	if _, err := auth.ExtractTokenMetadata(
		c,
		[]string{cms_client.SSOUsersRoleAdmin, cms_client.SSOUsersRoleInstructor},
	); err != nil {
		return nil, err
	}
	dto := usecases.LTIAttemptEditInputDTO{}
	err := jsonrpc.ValidatorBase(c, &dto)
	if err != nil {
		return nil, err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return nil, err
	}
	defer container.Close()
	uc := container.LTIAttemptEditUC()
	output, err := uc.Execute(dto)
	return output, err
}

// LTIAttemptList func for view of list LTIAttempt.
// @Description List lti_attempt. Roles: [admin, instructor]
// @Summary list lti_attempt
// @Tags lti_attempt
// @Accept json
// @Produce json
// @Param form body usecases.LTIAttemptListRequest true "lti_attempt list info"
// @Success 200 {object} usecases.LTIAttemptListResponse
// @Security ApiKeyAuth
// @Router /api/v1/rpc/lti_attempt.list [post]
func LTIAttemptList(c *jsonrpc.Ctx) (interface{}, error) {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	if _, err := auth.ExtractTokenMetadata(
		c,
		[]string{cms_client.SSOUsersRoleAdmin, cms_client.SSOUsersRoleInstructor},
	); err != nil {
		return nil, err
	}
	dto := usecases.LTIAttemptListInputDTO{}
	err := jsonrpc.ValidatorBase(c, &dto)
	if err != nil {
		return nil, err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return nil, err
	}
	defer container.Close()
	uc := container.LTIAttemptListUC()
	output, err := uc.Execute(dto)
	return output, err
}

// LTIAttemptDelete func for delete LTIAttempt.
// @Description Delete lti_attempt.
// @Summary delete lti_attempt
// @Tags lti_attempt
// @Accept json
// @Produce json
// @Param form body usecases.LTIAttemptDeleteRequest true "lti_attempt id"
// @Success 200 {object} usecases.LTIAttemptDeleteResponse
// @Security ApiKeyAuth
// @Router /api/v1/rpc/lti_attempt.delete [post]
func LTIAttemptDelete(c *jsonrpc.Ctx) (interface{}, error) {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	if _, err := auth.ExtractTokenMetadata(
		c,
		[]string{cms_client.SSOUsersRoleAdmin, cms_client.SSOUsersRoleInstructor},
	); err != nil {
		return nil, err
	}
	dto := usecases.LTIAttemptDeleteInputDTO{}
	err := jsonrpc.ValidatorBase(c, &dto)
	if err != nil {
		return nil, err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return nil, err
	}
	defer container.Close()
	uc := container.LTIAttemptDeleteUC()
	output, err := uc.Execute(dto)
	return output, err
}

// LTIAttemptGet func for full model LTIAttempt.
// @Description get lti_attempt.
// @Summary get lti_attempt
// @Tags lti_attempt
// @Accept json
// @Produce json
// @Param form body usecases.LTIAttemptGetRequest true "lti_attempt id"
// @Success 200 {object} usecases.LTIAttemptGetResponse
// @Security ApiKeyAuth
// @Router /api/v1/rpc/lti_attempt.get [post]
func LTIAttemptGet(c *jsonrpc.Ctx) (interface{}, error) {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	dto := usecases.LTIAttemptGetInputDTO{}
	err := jsonrpc.ValidatorBase(c, &dto)
	if err != nil {
		return nil, err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return nil, err
	}
	defer container.Close()
	uc := container.LTIAttemptGetUC()
	output, err := uc.Execute(dto)
	return output, err
}
