package controllers

import (
	"net/url"

	"gitlab.com/a10869/api-modules/shared/logs"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/di"
	"gitlab.com/a10869/api-modules/backend/app/usecases/auth"
	"gitlab.com/a10869/api-modules/shared/utils"
)

// SSOAuthorize Получить код авторизации.
// @Description Получение кода авторизации по OpenID.
// @Summary Получение кода авторизации по OpenID.
// @Tags SSO
// @Accept json
// @Produce json
// @Param client_id query string true "Код приложения клиента"
// @Param redirect_uri query string true "Адрес переадресации клиента"
// @Param response_type query string true "Тип ответа. Возможные значения: [code]"  Enums(code)
// @Param scope query string true "Область доступа (значения, разделённые запятой). Возможные значения: [default, openid]"  Enums(default, openid)
// @Param path query string false "Путь, возвращается в параметрах редиректа"
// @Param state query string false "Состояние, возвращается в параметрах редиректа, участвует в генерации кода авторизации"
// @Param extra query string false "Дополнительные параметры вида base64"
// @Success 200 {object} auth.SSOAuthorizeResponse
// @Security ApiKeyAuth
// @Router /v1/sso/authorize [get]
func SSOAuthorize(c *fiber.Ctx) error {
	// Получаем QueryParams в виде map
	queryParams := c.Queries()

	// Формируем строку QueryParams
	queryString := "?"
	for key, value := range queryParams {
		// Кодируем ключ и значение, чтобы избежать проблем с символами
		queryString += url.QueryEscape(key) + "=" + url.QueryEscape(value) + "&"
	}

	// Убираем последний символ '&', если он есть
	if len(queryString) > 1 {
		queryString = queryString[:len(queryString)-1]
	}

	// Формируем новый URL
	newURL := "/login" + queryString

	// Выполняем редирект
	return c.Redirect(newURL)
}

// SSOAuthorizePost Получить код авторизации.
// @Description Получение кода авторизации по OpenID.
// @Summary Получение кода авторизации по OpenID.
// @Tags SSO
// @Accept json
// @Produce json
// @Param form body auth.SSOAuthorizeInputDTO true "renew token form info"
// @Success 200 {object} auth.SSOAuthorizeResponse
// @Security ApiKeyAuth
// @Router /v1/sso/authorize [post]
func SSOAuthorizePost(c *fiber.Ctx) error {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	claims, err := auth.ExtractTokenMetadata(c, []string{})
	if err != nil {
		return utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	dto := auth.SSOAuthorizeInputDTO{}
	err = utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	dto.UserID = claims.Id
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return err
	}
	defer container.Close()
	uc := container.SSOAuthorizeUC()
	output, err := uc.Execute(dto)
	if err != nil {
		return utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	return utils.FiberSuccessResponse{Result: output}
}

// SSOToken Получить токен доступа и токен обновления.
// @Description Получение токена доступа и токена обновления.
// @Summary Получение токена доступа и токена обновления.
// @Tags SSO, EXTERNAL
// @Accept json
// @Produce json
// @Param grant_type formData string true "Тип авторизации [authorization_code, refresh_token]" Enums(authorization_code, refresh_token)
// @Param redirect_uri formData string false "Адрес переадресации клиента"
// @Param code formData string false "Код авторизации"
// @Param refresh_token formData string false "Токен обновления"
// @Param Authorization header string true "Basic-токен, созданный клиентом"
// @Success 200 {object} auth.SwaggerSSOTokenResponse
// @Security ApiKeyAuth
// @Router /v1/sso/token [post]
func SSOToken(c *fiber.Ctx) error {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	dto := auth.SSOTokenInputDTO{}
	dto.GrantType = c.FormValue("grant_type", "")
	dto.RedirectUri = c.FormValue("redirect_uri", "")
	dto.Code = c.FormValue("code", "")
	dto.RefreshToken = c.FormValue("refresh_token", "")
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return err
	}
	defer container.Close()
	uc := container.SSOTokenUC()
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}

// SSOIntrospect Проверка состояния токена.
// @Description Проверка состояния токена.
// @Summary Проверка состояния токена.
// @Tags SSO, EXTERNAL
// @Accept json
// @Produce json
// @Param token formData string true "Токен доступа или токен обновления"
// @Param Authorization header string true "Basic-токен, созданный клиентом"
// @Success 200 {object} auth.SSOTokenIntrospectResponse
// @Security ApiKeyAuth
// @Router /v1/sso/introspect [post]
func SSOIntrospect(c *fiber.Ctx) error {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	dto := auth.SSOIntrospectInputDTO{}
	dto.Token = c.FormValue("token", "")
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return err
	}
	defer container.Close()
	uc := container.SSOIntrospectUC()
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}

// SSOUserInfo Получить информацию о пользователе.
// @Description Получить информацию о пользователе.
// @Summary Получить информацию о пользователе.
// @Tags SSO
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer токен"
// @Success 200 {object} auth.SwaggerSSOTokenPublicDataResponse
// @Security ApiKeyAuth
// @Router /v1/sso/userinfo [post]
func SSOUserInfo(c *fiber.Ctx) error {
	claims, err := auth.ExtractTokenMetadata(c, []string{})
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: claims}
}
