package controllers

import (
	"gitlab.com/a10869/api-modules/backend/app/di"
	"gitlab.com/a10869/api-modules/backend/app/queries"
	"gitlab.com/a10869/api-modules/backend/app/usecases"
	"gitlab.com/a10869/api-modules/backend/app/usecases/auth"
	"gitlab.com/a10869/api-modules/backend/app/usecases/external"
	"gitlab.com/a10869/api-modules/shared/cms_client"
	"gitlab.com/a10869/api-modules/shared/jsonrpc"
	"gitlab.com/a10869/api-modules/shared/logs"
)

// PNETServerUpsert func for creates a new PNETServer.
// @Description Create pnet_server. Roles [admin]
// @Summary create pnet_server
// @Tags server
// @Accept json
// @Produce json
// @Param object body usecases.PNETServerEditRequest true "pnet_server form info"
// @Success 200 {object} usecases.PNETServerEditResponse
// @Security ApiKeyAuth
// @Router /api/v1/rpc/server.upsert [post]
func PNETServerUpsert(c *jsonrpc.Ctx) (interface{}, error) {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	if _, err := auth.ExtractTokenMetadata(
		c,
		[]string{cms_client.SSOUsersRoleAdmin},
	); err != nil {
		return nil, err
	}
	dto := usecases.PNETServerEditInputDTO{}
	err := jsonrpc.ValidatorBase(c, &dto)
	if err != nil {
		return nil, err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return nil, err
	}
	defer container.Close()
	uc := container.PNETServerEditUC()
	output, err := uc.Execute(dto)
	if err != nil {
		return nil, err
	}
	ucDistribution := container.PnetServerChangeDistributionUC()
	err = ucDistribution.Execute()
	if err != nil {
		return nil, err
	}
	return output, nil
}

// PNETServerList func for view of list PNETServer.
// @Description List pnet_server. Roles: [admin, instructor]
// @Summary list pnet_server
// @Tags server
// @Accept json
// @Produce json
// @Param object body usecases.PNETServerListRequest true "pnet_server list info"
// @Success 200 {object} usecases.PNETServerListResponse
// @Security ApiKeyAuth
// @Router /api/v1/rpc/server.list [post]
func PNETServerList(c *jsonrpc.Ctx) (interface{}, error) {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	if _, err := auth.ExtractTokenMetadata(
		c,
		[]string{cms_client.SSOUsersRoleAdmin, cms_client.SSOUsersRoleInstructor},
	); err != nil {
		return nil, err
	}
	dto := queries.PNETServerQueriesListDTO{}
	err := jsonrpc.ValidatorBase(c, &dto)
	if err != nil {
		return nil, err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return nil, err
	}
	defer container.Close()
	uc := container.PNETServerListUC()
	output, err := uc.Execute(dto)
	return output, err
}

// PNETServerDelete func for delete PNETServer.
// @Description Delete pnet_server. Roles: [admin]
// @Summary delete pnet_server
// @Tags server
// @Accept json
// @Produce json
// @Param object body usecases.PNETServerDeleteRequest true "pnet_server id"
// @Success 200 {object} usecases.PNETServerDeleteResponse
// @Security ApiKeyAuth
// @Router /api/v1/rpc/server.delete [post]
func PNETServerDelete(c *jsonrpc.Ctx) (interface{}, error) {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	if _, err := auth.ExtractTokenMetadata(
		c,
		[]string{cms_client.SSOUsersRoleAdmin},
	); err != nil {
		return nil, err
	}
	dto := usecases.PNETServerDeleteInputDTO{}
	err := jsonrpc.ValidatorBase(c, &dto)
	if err != nil {
		return nil, err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return nil, err
	}
	defer container.Close()
	uc := container.PNETServerDeleteUC()
	output, err := uc.Execute(dto)
	if err != nil {
		return nil, err
	}
	ucDistribution := container.PnetServerChangeDistributionUC()
	err = ucDistribution.Execute()
	if err != nil {
		return nil, err
	}
	return output, nil
}

// PNETServerGet func for full model PNETServer.
// @Description get pnet_server. Roles: [admin, instructor]
// @Summary get pnet_server
// @Tags server
// @Accept json
// @Produce json
// @Param object body usecases.PNETServerGetRequest true "pnet_server id"
// @Success 200 {object} usecases.PNETServerGetResponse
// @Security ApiKeyAuth
// @Router /api/v1/rpc/server.get [post]
func PNETServerGet(c *jsonrpc.Ctx) (interface{}, error) {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	if _, err := auth.ExtractTokenMetadata(
		c,
		[]string{cms_client.SSOUsersRoleAdmin, cms_client.SSOUsersRoleInstructor},
	); err != nil {
		return nil, err
	}
	dto := usecases.PNETServerGetInputDTO{}
	err := jsonrpc.ValidatorBase(c, &dto)
	if err != nil {
		return nil, err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return nil, err
	}
	defer container.Close()
	uc := container.PNETServerGetUC()
	output, err := uc.Execute(dto)
	return output, err
}

// PNETServerPing ping from external servers.
// @Description ing from external servers.
// @Summary ping from pnet_server
// @Tags server, EXTERNAL
// @Accept json
// @Produce json
// @Param object body external.PNETServerPingRequest true "pnet_server id"
// @Success 200 {object} external.PNETServerPingResponse
// @Param Authorization header string true "Basic-токен, созданный клиентом"
// @Router /api/v1/rpc/server.ping [post]
func PNETServerPing(c *jsonrpc.Ctx) (interface{}, error) {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	dto := external.PNETServerPingInputDTO{}
	err := jsonrpc.ValidatorBase(c, &dto)
	if err != nil {
		return nil, err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return nil, err
	}
	defer container.Close()
	err = container.ServiceAuthorizeUC().Execute(c)
	if err != nil {
		return nil, err
	}
	uc := container.PNETServerPingUC()
	output, err := uc.SetContext(c.FiberCtx.Locals("x-service-id").(string)).Execute(dto)
	return output, err
}
