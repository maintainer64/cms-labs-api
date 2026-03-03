package controllers

import (
	fiber "github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/di"
	"gitlab.com/a10869/api-modules/backend/app/usecases/auth"
	"gitlab.com/a10869/api-modules/shared/cms_client"
	"gitlab.com/a10869/api-modules/shared/jsonrpc"
	"gitlab.com/a10869/api-modules/shared/logs"
)

// LTILaunch функция инициализирует подключение к LTI.
// @Description инициализация подключения к LTI.
// @Summary инициализация подключения к LTI
// @Tags LTI
// @Accept json
// @Produce json
// @Success 302
// @Router /api/v2/lti/launch [post]
// @Router /api/v2/lti/launch [get]
func LTILaunch(ctx *fiber.Ctx) error {
	c := &jsonrpc.Ctx{FiberCtx: ctx}
	issuer, err := auth.IssuerURLByBaseUrl(c.FiberCtx.BaseURL())
	if err != nil {
		return err
	}
	diLoggerConf := logs.NewZeroLoggerConf(c)
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return err
	}
	defer container.Close()
	uc := container.LTIProtocolLaunch()
	if err = uc.ServeHTTP(c.FiberCtx); err != nil {
		return err
	}
	authManager := container.AuthTokenManager()
	token, err := authManager.NewJWTByLaunchID(issuer, c.FiberCtx.Locals("LTILaunchID").(string))
	if err != nil {
		return err
	}
	c.FiberCtx.Cookie(
		&fiber.Cookie{
			Name:     cms_client.SSORefreshTokenName,
			Value:    token.RefreshToken,
			Path:     "/",
			SameSite: fiber.CookieSameSiteNoneMode,
			Expires:  auth.ExpiresRefreshCookie(),
			Secure:   true,
		},
	)
	return c.FiberCtx.Redirect("/lti-redirect", fiber.StatusFound)
}

// LTILogin функция аутентификация пользователя по LTI.
// @Description аутентификация LTI.
// @Summary аутентификация через LTI
// @Tags LTI
// @Accept json
// @Produce json
// @Success 200
// @Router /api/v2/lti/login [post]
// @Router /api/v2/lti/login [get]
func LTILogin(ctx *fiber.Ctx) error {
	c := &jsonrpc.Ctx{FiberCtx: ctx}
	diLoggerConf := logs.NewZeroLoggerConf(c)
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return err
	}
	defer container.Close()
	uc := container.LTIProtocolLogin()
	return uc.ServeHTTP(c.FiberCtx)
}
