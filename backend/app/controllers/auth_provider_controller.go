package controllers

import (
	"gitlab.com/a10869/api-modules/backend/app/di"
	"gitlab.com/a10869/api-modules/backend/app/usecases"
	"gitlab.com/a10869/api-modules/backend/app/usecases/auth"
	"gitlab.com/a10869/api-modules/shared/cms_client"
	"gitlab.com/a10869/api-modules/shared/jsonrpc"
	"gitlab.com/a10869/api-modules/shared/logs"
)

// AuthProviderUpsert func for creates a new AuthProvider.
// @Description Create lti_form. Roles: [admin]
// @Summary create lti_form
// @Tags lti_form
// @Accept json
// @Produce json
// @Param object body usecases.AuthProviderEditRequest true "lti_form form info"
// @Success 200 {object} usecases.AuthProviderEditResponse
// @Security ApiKeyAuth
// @Router /api/v1/rpc/auth_provider.upsert [post]
func AuthProviderUpsert(c *jsonrpc.Ctx) (interface{}, error) {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	if _, err := auth.ExtractTokenMetadata(
		c,
		[]string{cms_client.SSOUsersRoleAdmin},
	); err != nil {
		return nil, err
	}
	dto := usecases.AuthProviderEditInputDTO{}
	err := jsonrpc.ValidatorBase(c, &dto)
	if err != nil {
		return nil, err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return nil, err
	}
	uc := container.AuthProviderEditUC()
	output, err := uc.Execute(dto)
	return output, err
}

// AuthProviderList func for view of list AuthProvider.
// @Description List lti_form. Roles: [admin]
// @Summary list lti_form
// @Tags lti_form
// @Accept json
// @Produce json
// @Param object body usecases.AuthProviderListRequest true "lti_form list info"
// @Success 200 {object} usecases.AuthProviderListResponse
// @Security ApiKeyAuth
// @Router /api/v1/rpc/auth_provider.list [post]
func AuthProviderList(c *jsonrpc.Ctx) (interface{}, error) {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	if _, err := auth.ExtractTokenMetadata(
		c,
		[]string{cms_client.SSOUsersRoleAdmin},
	); err != nil {
		return nil, err
	}
	dto := usecases.AuthProviderListInputDTO{}
	err := jsonrpc.ValidatorBase(c, &dto)
	if err != nil {
		return nil, err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return nil, err
	}
	uc := container.AuthProviderListUC()
	output, err := uc.Execute(dto)
	return output, err
}

// AuthProviderDelete func for delete AuthProvider.
// @Description Delete lti_form. Roles: [admin]
// @Summary delete lti_form
// @Tags lti_form
// @Accept json
// @Produce json
// @Param object body usecases.AuthProviderDeleteRequest true "lti_form id"
// @Success 200 {object} usecases.AuthProviderDeleteResponse
// @Security ApiKeyAuth
// @Router /api/v1/rpc/auth_provider.delete [post]
func AuthProviderDelete(c *jsonrpc.Ctx) (interface{}, error) {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	if _, err := auth.ExtractTokenMetadata(
		c,
		[]string{cms_client.SSOUsersRoleAdmin},
	); err != nil {
		return nil, err
	}
	dto := usecases.AuthProviderDeleteInputDTO{}
	err := jsonrpc.ValidatorBase(c, &dto)
	if err != nil {
		return nil, err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return nil, err
	}
	defer container.Close()
	uc := container.AuthProviderDeleteUC()
	output, err := uc.Execute(dto)
	return output, err
}

// AuthProviderGet func for full model AuthProvider.
// @Description get lti_form. Roles [admin]
// @Summary get lti_form
// @Tags lti_form
// @Accept json
// @Produce json
// @Param object body usecases.AuthProviderGetRequest true "lti_form id"
// @Success 200 {object} usecases.AuthProviderGetResponse
// @Security ApiKeyAuth
// @Router /api/v1/rpc/auth_provider.get [post]
func AuthProviderGet(c *jsonrpc.Ctx) (interface{}, error) {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	if _, err := auth.ExtractTokenMetadata(
		c,
		[]string{cms_client.SSOUsersRoleAdmin},
	); err != nil {
		return nil, err
	}
	dto := usecases.AuthProviderGetInputDTO{}
	err := jsonrpc.ValidatorBase(c, &dto)
	if err != nil {
		return nil, err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return nil, err
	}
	defer container.Close()
	uc := container.AuthProviderGetUC()
	output, err := uc.Execute(dto)
	return output, err
}

// AuthProviderSSOListGet func for list of SSO URLs.
// @Description List sso url. Roles: [none]
// @Summary list sso url
// @Tags lti_form
// @Accept json
// @Produce json
// @Param object body usecases.AuthProviderListSSORequest true "filter"
// @Success 200 {object} usecases.AuthProviderListSSOResponse
// @Security ApiKeyAuth
// @Router /api/v1/rpc/auth_provider.sso_list_get [post]
func AuthProviderSSOListGet(c *jsonrpc.Ctx) (interface{}, error) {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return nil, err
	}
	uc := container.AuthProviderListSSOUC()
	output, err := uc.Execute()
	return output, err
}
