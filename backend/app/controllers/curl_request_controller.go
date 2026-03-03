package controllers

import (
	"io"

	"gitlab.com/a10869/api-modules/backend/app/queries"

	"gitlab.com/a10869/api-modules/shared/cms_client"

	fiber "github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/di"
	"gitlab.com/a10869/api-modules/backend/app/usecases"
	"gitlab.com/a10869/api-modules/backend/app/usecases/auth"
	"gitlab.com/a10869/api-modules/backend/app/usecases/external"
	"gitlab.com/a10869/api-modules/shared/jsonrpc"
	"gitlab.com/a10869/api-modules/shared/logs"
)

// CurlRequestUpsert func for creates a new CurlRequest.
// @Description Create curl_request. Roles [admin]
// @Summary create curl_request
// @Tags curl_request
// @Accept json
// @Produce json
// @Param object body usecases.CurlRequestEditRequest true "curl_request form info"
// @Success 200 {object} usecases.CurlRequestEditResponse
// @Security ApiKeyAuth
// @Router /api/v1/rpc/curl_request.upsert [post]
func CurlRequestUpsert(c *jsonrpc.Ctx) (interface{}, error) {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	if _, err := auth.ExtractTokenMetadata(
		c,
		[]string{cms_client.SSOUsersRoleAdmin},
	); err != nil {
		return nil, err
	}
	dto := usecases.CurlRequestEditInputDTO{}
	err := jsonrpc.ValidatorBase(c, &dto)
	if err != nil {
		return nil, err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return nil, err
	}
	defer container.Close()
	uc := container.CurlRequestEditUC()
	return uc.Execute(dto)
}

// CurlRequestList func for view of list CurlRequest.
// @Description List pnet_server. Roles: [admin, instructor]
// @Summary list pnet_server
// @Tags curl_request
// @Accept json
// @Produce json
// @Param form body usecases.CurlRequestListRequest true "pnet_server list info"
// @Success 200 {object} usecases.CurlRequestListResponse
// @Security ApiKeyAuth
// @Router /api/v1/rpc/curl_request.list [post]
func CurlRequestList(c *jsonrpc.Ctx) (interface{}, error) {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	if _, err := auth.ExtractTokenMetadata(
		c,
		[]string{cms_client.SSOUsersRoleAdmin, cms_client.SSOUsersRoleInstructor},
	); err != nil {
		return nil, err
	}
	dto := queries.CurlRequestQueriesListDTO{}
	err := jsonrpc.ValidatorBase(c, &dto)
	if err != nil {
		return nil, err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return nil, err
	}
	defer container.Close()
	uc := container.CurlRequestListUC()
	return uc.Execute(dto)
}

// CurlRequestDelete func for delete CurlRequest.
// @Description Delete pnet_server. Roles: [admin]
// @Summary delete curl_request
// @Tags curl_request
// @Accept json
// @Produce json
// @Param form body usecases.CurlRequestDeleteRequest true "curl_request_id"
// @Success 200 {object} usecases.CurlRequestDeleteResponse
// @Security ApiKeyAuth
// @Router /api/v1/rpc/curl_request.delete [post]
func CurlRequestDelete(c *jsonrpc.Ctx) (interface{}, error) {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	if _, err := auth.ExtractTokenMetadata(
		c,
		[]string{cms_client.SSOUsersRoleAdmin},
	); err != nil {
		return nil, err
	}
	dto := usecases.CurlRequestDeleteInputDTO{}
	err := jsonrpc.ValidatorBase(c, &dto)
	if err != nil {
		return nil, err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return nil, err
	}
	defer container.Close()
	uc := container.CurlRequestDeleteUC()
	return uc.Execute(dto)
}

// CurlRequestGet func for full model CurlRequest.
// @Description get curl_request. Roles: [admin]
// @Summary get curl_request
// @Tags curl_request
// @Accept json
// @Produce json
// @Param form body usecases.CurlRequestGetRequest true "pnet_server id"
// @Success 200 {object} usecases.CurlRequestGetResponse
// @Security ApiKeyAuth
// @Router /api/v1/rpc/curl_request.get [post]
func CurlRequestGet(c *jsonrpc.Ctx) (interface{}, error) {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	if _, err := auth.ExtractTokenMetadata(
		c,
		[]string{cms_client.SSOUsersRoleAdmin},
	); err != nil {
		return nil, err
	}
	dto := usecases.CurlRequestGetInputDTO{}
	err := jsonrpc.ValidatorBase(c, &dto)
	if err != nil {
		return nil, err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return nil, err
	}
	defer container.Close()
	uc := container.CurlRequestGetUC()
	return uc.Execute(dto)
}

// CurlRequestExecute execute prepare request.
// @Description execute request.
// @Summary execute request
// @Tags curl_request, EXTERNAL
// @Accept json
// @Produce json
// @Param form body external.CurlRequestExecuteInputDTO true "pnet_server id"
// @Success 200
// @Param Authorization header string true "Basic-токен, созданный клиентом"
// @Router /api/v1/curl-request/execute [post]
func CurlRequestExecute(ctx *fiber.Ctx) error {
	c := &jsonrpc.Ctx{
		FiberCtx: ctx,
		Params:   ctx.Body(),
	}
	diLoggerConf := logs.NewZeroLoggerConf(c)
	dto := external.CurlRequestExecuteInputDTO{}
	err := jsonrpc.ValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return err
	}
	defer container.Close()
	err = container.ServiceAuthorizeUC().Execute(c)
	if err != nil {
		return err
	}
	uc := container.CurlRequestExecuteUC()
	output, err := uc.SetContext(
		c.FiberCtx.Locals("x-service-id").(string),
	).Execute(dto)
	if err != nil {
		return err
	}
	if output == nil {
		return jsonrpc.NewRpcError("curl_request_not_execute", "the method was not executed")
	}
	defer output.Body.Close()
	c.FiberCtx.Status(output.StatusCode)
	// Копируем заголовки из ответа
	for k, v := range output.Header {
		if len(v) > 0 {
			c.FiberCtx.Set(k, v[0])
		}
	}
	// Читаем тело ответа
	body, err := io.ReadAll(output.Body)
	if err != nil {
		return err
	}
	// Возвращаем ответ "как есть" с оригинальным статус-кодом и заголовками
	return c.FiberCtx.Send(body)
}
