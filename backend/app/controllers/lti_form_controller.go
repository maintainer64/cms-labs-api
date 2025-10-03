package controllers

import (
	"gitlab.com/a10869/api-modules/backend/app/di"
	"gitlab.com/a10869/api-modules/backend/app/usecases"
	"gitlab.com/a10869/api-modules/backend/app/usecases/auth"
	"gitlab.com/a10869/api-modules/shared/cms_client"
	"gitlab.com/a10869/api-modules/shared/jsonrpc"
	"gitlab.com/a10869/api-modules/shared/logs"
)

// LTIFormUpsert func for creates a new LTIForm.
// @Description Create lti_form. Roles: [admin]
// @Summary create lti_form
// @Tags lti_form
// @Accept json
// @Produce json
// @Param object body usecases.LTIFormEditRequest true "lti_form form info"
// @Success 200 {object} usecases.LTIFormEditResponse
// @Security ApiKeyAuth
// @Router /api/v1/rpc/lti_form.upsert [post]
func LTIFormUpsert(c *jsonrpc.Ctx) (interface{}, error) {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	if _, err := auth.ExtractTokenMetadata(
		c,
		[]string{cms_client.SSOUsersRoleAdmin},
	); err != nil {
		return nil, err
	}
	dto := usecases.LTIFormEditInputDTO{}
	err := jsonrpc.ValidatorBase(c, &dto)
	if err != nil {
		return nil, err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return nil, err
	}
	uc := container.LTIFormEditUC()
	output, err := uc.Execute(dto)
	return output, err
}

// LTIFormList func for view of list LTIForm.
// @Description List lti_form. Roles: [admin]
// @Summary list lti_form
// @Tags lti_form
// @Accept json
// @Produce json
// @Param object body usecases.LTIFormListRequest true "lti_form list info"
// @Success 200 {object} usecases.LTIFormListResponse
// @Security ApiKeyAuth
// @Router /api/v1/rpc/lti_form.list [post]
func LTIFormList(c *jsonrpc.Ctx) (interface{}, error) {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	if _, err := auth.ExtractTokenMetadata(
		c,
		[]string{cms_client.SSOUsersRoleAdmin},
	); err != nil {
		return nil, err
	}
	dto := usecases.LTIFormListInputDTO{}
	err := jsonrpc.ValidatorBase(c, &dto)
	if err != nil {
		return nil, err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return nil, err
	}
	uc := container.LTIFormListUC()
	output, err := uc.Execute(dto)
	return output, err
}

// LTIFormDelete func for delete LTIForm.
// @Description Delete lti_form. Roles: [admin]
// @Summary delete lti_form
// @Tags lti_form
// @Accept json
// @Produce json
// @Param object body usecases.LTIFormDeleteRequest true "lti_form id"
// @Success 200 {object} usecases.LTIFormDeleteResponse
// @Security ApiKeyAuth
// @Router /api/v1/rpc/lti_form.delete [post]
func LTIFormDelete(c *jsonrpc.Ctx) (interface{}, error) {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	if _, err := auth.ExtractTokenMetadata(
		c,
		[]string{cms_client.SSOUsersRoleAdmin},
	); err != nil {
		return nil, err
	}
	dto := usecases.LTIFormDeleteInputDTO{}
	err := jsonrpc.ValidatorBase(c, &dto)
	if err != nil {
		return nil, err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return nil, err
	}
	defer container.Close()
	uc := container.LTIFormDeleteUC()
	output, err := uc.Execute(dto)
	return output, err
}

// LTIFormGet func for full model LTIForm.
// @Description get lti_form. Roles [admin]
// @Summary get lti_form
// @Tags lti_form
// @Accept json
// @Produce json
// @Param object body usecases.LTIFormGetRequest true "lti_form id"
// @Success 200 {object} usecases.LTIFormGetResponse
// @Security ApiKeyAuth
// @Router /api/v1/rpc/lti_form.get [post]
func LTIFormGet(c *jsonrpc.Ctx) (interface{}, error) {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	if _, err := auth.ExtractTokenMetadata(
		c,
		[]string{cms_client.SSOUsersRoleAdmin},
	); err != nil {
		return nil, err
	}
	dto := usecases.LTIFormGetInputDTO{}
	err := jsonrpc.ValidatorBase(c, &dto)
	if err != nil {
		return nil, err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return nil, err
	}
	defer container.Close()
	uc := container.LTIFormGetUC()
	output, err := uc.Execute(dto)
	return output, err
}

// LTIFormSSOListGet func for list of SSO URLs.
// @Description List sso url. Roles: [none]
// @Summary list sso url
// @Tags lti_form
// @Accept json
// @Produce json
// @Param object body usecases.LTIFormListSSORequest true "filter"
// @Success 200 {object} usecases.LTIFormListSSOResponse
// @Security ApiKeyAuth
// @Router /api/v1/rpc/lti_form.sso_list_get [post]
func LTIFormSSOListGet(c *jsonrpc.Ctx) (interface{}, error) {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return nil, err
	}
	uc := container.LTIFormListSSOUC()
	output, err := uc.Execute()
	return output, err
}
