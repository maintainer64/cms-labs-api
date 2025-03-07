package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/pnetlabaddon/app/di"
	"gitlab.com/a10869/api-modules/pnetlabaddon/app/usecases"
	"gitlab.com/a10869/api-modules/shared/cms_client"
	"gitlab.com/a10869/api-modules/shared/logs"
	"gitlab.com/a10869/api-modules/shared/utils"
)

// SSOFirstFactor Переадресация пользователя на сервер аутентификации.
// @Description Отправка пользователя с параметрами на сервер аутентификации.
// @Summary Отправка пользователя с параметрами на сервер аутентификации.
// @Tags SSO
// @Accept json
// @Produce json
// @Param extra query string false "Дополнительные параметры в base64"
// @Param path query string false "Путь с которого выполнился редирект"
// @Success 307
// @Router /v1/sso/login [get]
func SSOFirstFactor(c *fiber.Ctx) error {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	extra := c.Query("extra", "")
	path := c.Query("path", "/")
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return utils.FiberValidationException{
			Status:    fiber.StatusInternalServerError,
			Exception: err,
		}
	}
	defer container.Close()

	client := container.CMSClient()

	newURL := client.SSOAuthorizeURI(
		c.BaseURL()+"/pnet-lab-addon/api/v1/sso/openid",
		"default",
		path,
		extra,
	)
	return c.Redirect(newURL, fiber.StatusTemporaryRedirect)
}

// SSOSecondFactor Получение токена пользователя второй фактор OpenID.
// @Description Дополнительные параметры:
// Если не пришёл access, refresh токен - ошибка при выдаче доступа.
// Проставление cookies происходит автоматически по запросу.
// @Summary Второй фактор openid.
// @Tags SSO
// @Accept json
// @Produce json
// @Param state query string false "ID запроса по всему пути аутентификации"
// @Param path query string false "Оригинальный путь по всему пути аутентификации"
// @Param code query string false "Код OTP обмена авторизации"
// @Param application query string false "ClientID запрашиваемого из аутентификации"
// @Param extra query string false "Дополнительные данные с backend base64 по всему пути аутентификации"
// @Success 307
// @Router /v1/sso/openid [get]
func SSOSecondFactor(c *fiber.Ctx) error {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return utils.FiberValidationException{
			Status:    fiber.StatusInternalServerError,
			Exception: err,
		}
	}
	defer container.Close()
	uc := container.SSOSecondFactorUC()
	dto := usecases.SSOSecondFactorInputDTO{
		Code:        c.Query("code", ""),
		Application: c.Query("application", ""),
		Path:        c.Query("path", ""),
		State:       c.Query("state", ""),
		Extra:       c.Query("extra", ""),
		RedirectURI: c.BaseURL() + "/pnet-lab-addon/api/v1/sso/openid",
	}
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	// Выполняем редирект
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    output.CookieToken,
		Path:     "/",
		MaxAge:   output.CookieMaxAge,
		Expires:  output.CookieAge,
		HTTPOnly: true,
	})
	c.Cookie(&fiber.Cookie{
		Name:     "_session",
		Value:    "",
		Path:     "/",
		HTTPOnly: true,
	})
	c.Cookie(&fiber.Cookie{
		Name:     cms_client.SSORefreshTokenName,
		Value:    output.RefreshToken,
		MaxAge:   output.RefreshTokenMaxAge,
		Path:     "/",
		Expires:  output.RefreshTokenAge,
		HTTPOnly: true,
	})
	return c.Redirect("/", fiber.StatusTemporaryRedirect)
}
