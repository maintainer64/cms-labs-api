package controllers

import (
	"io"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/di"
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/usecases"
	"gitlab.com/a10869/api-modules/backend/app/usecases/auth"
	"gitlab.com/a10869/api-modules/backend/app/usecases/external"
	"gitlab.com/a10869/api-modules/shared/logs"
	"gitlab.com/a10869/api-modules/shared/utils"
)

// CurlRequestCreate func for creates a new CurlRequest.
// @Description Create curl_request. Roles [admin]
// @Summary create curl_request
// @Tags CurlRequest
// @Accept json
// @Produce json
// @Param form body usecases.CurlRequestEditInputDTO true "curl_request form info"
// @Success 200 {object} usecases.CurlRequestEditResponse
// @Security ApiKeyAuth
// @Router /v1/curl-request/upsert [post]
func CurlRequestCreate(c *fiber.Ctx) error {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	if _, err := auth.ExtractTokenMetadata(
		c,
		[]string{models.UsersRoleAdmin},
	); err != nil {
		return err
	}
	dto := usecases.CurlRequestEditInputDTO{}
	err := utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return err
	}
	defer container.Close()
	uc := container.CurlRequestEditUC()
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}

// CurlRequestList func for view of list CurlRequest.
// @Description List pnet_server. Roles: [admin, instructor]
// @Summary list pnet_server
// @Tags CurlRequest
// @Accept json
// @Produce json
// @Param form body usecases.CurlRequestListInputDTO true "pnet_server list info"
// @Success 200 {object} usecases.CurlRequestListResponse
// @Security ApiKeyAuth
// @Router /v1/curl-request/list [post]
func CurlRequestList(c *fiber.Ctx) error {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	if _, err := auth.ExtractTokenMetadata(
		c,
		[]string{models.UsersRoleAdmin, models.UsersRoleInstructor},
	); err != nil {
		return err
	}
	dto := usecases.CurlRequestListInputDTO{}
	err := utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return err
	}
	defer container.Close()
	uc := container.CurlRequestListUC()
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}

// CurlRequestDelete func for delete CurlRequest.
// @Description Delete pnet_server. Roles: [admin]
// @Summary delete curl_request
// @Tags CurlRequest
// @Accept json
// @Produce json
// @Param form body usecases.CurlRequestDeleteInputDTO true "curl_request_id"
// @Success 200 {object} usecases.CurlRequestDeleteResponse
// @Security ApiKeyAuth
// @Router /v1/curl-request/delete [post]
func CurlRequestDelete(c *fiber.Ctx) error {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	if _, err := auth.ExtractTokenMetadata(
		c,
		[]string{models.UsersRoleAdmin},
	); err != nil {
		return err
	}
	dto := usecases.CurlRequestDeleteInputDTO{}
	err := utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return err
	}
	defer container.Close()
	uc := container.CurlRequestDeleteUC()
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}

// CurlRequestGet func for full model CurlRequest.
// @Description get curl_request. Roles: [admin]
// @Summary get curl_request
// @Tags CurlRequest
// @Accept json
// @Produce json
// @Param form body usecases.CurlRequestGetInputDTO true "pnet_server id"
// @Success 200 {object} usecases.CurlRequestGetResponse
// @Security ApiKeyAuth
// @Router /v1/curl-request/get [post]
func CurlRequestGet(c *fiber.Ctx) error {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	if _, err := auth.ExtractTokenMetadata(
		c,
		[]string{models.UsersRoleAdmin},
	); err != nil {
		return err
	}
	dto := usecases.CurlRequestGetInputDTO{}
	err := utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return err
	}
	defer container.Close()
	uc := container.CurlRequestGetUC()
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}

// CurlRequestExecute execute prepare request.
// @Description execute request.
// @Summary execute request
// @Tags CurlRequest, EXTERNAL
// @Accept json
// @Produce json
// @Param form body external.CurlRequestExecuteInputDTO true "pnet_server id"
// @Success 200
// @Param Authorization header string true "Basic-токен, созданный клиентом"
// @Router /v1/curl-request/execute [post]
func CurlRequestExecute(c *fiber.Ctx) error {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	dto := external.CurlRequestExecuteInputDTO{}
	err := utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return err
	}
	defer container.Close()
	uc := container.CurlRequestExecuteUC()
	output, err := uc.SetContext(c.Locals("x-service-id").(string)).Execute(dto)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	if output == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "The method was not executed",
		})
	}
	defer output.Body.Close()
	c.Status(output.StatusCode)
	// Копируем заголовки из ответа
	for k, v := range output.Header {
		if len(v) > 0 {
			c.Set(k, v[0])
		}
	}
	// Читаем тело ответа
	body, err := io.ReadAll(output.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	// Возвращаем ответ "как есть" с оригинальным статус-кодом и заголовками
	return c.Send(body)
}
