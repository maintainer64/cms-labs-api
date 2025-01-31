package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/di"
	"gitlab.com/a10869/api-modules/backend/app/usecases/auth"
)

// LTILaunch функция инициализирует подключение к LTI.
// @Description инициализация подключения к LTI.
// @Summary инициализация подключения к LTI
// @Tags LTI
// @Accept json
// @Produce json
// @Success 302
// @Router /v2/lti/launch [post]
// @Router /v2/lti/launch [get]
func LTILaunch(c *fiber.Ctx) error {
	diContainer := di.NewDIContainer()
	uc, err := diContainer.LTIProtocolLaunch()
	if err != nil {
		return err
	}
	if err = uc.ServeHTTP(c); err != nil {
		return err
	}
	authManager, err := diContainer.AuthTokenManager()
	if err != nil {
		return err
	}
	token, err := authManager.NewJWTByLaunchID(c.Locals("LTILaunchID").(string))
	if err != nil {
		return err
	}
	c.Cookie(&fiber.Cookie{
		Name:     "refresh-token",
		Value:    token.RefreshToken,
		Path:     "/",
		SameSite: fiber.CookieSameSiteNoneMode,
		Expires:  auth.ExpiresRefreshCookie(),
		Secure:   true,
	})
	return c.Redirect("/", fiber.StatusFound)
}

// LTILogin функция аутентификация пользователя по LTI.
// @Description аутентификация LTI.
// @Summary аутентификация через LTI
// @Tags LTI
// @Accept json
// @Produce json
// @Success 200
// @Router /v2/lti/login [post]
// @Router /v2/lti/login [get]
func LTILogin(c *fiber.Ctx) error {
	uc, err := di.NewDIContainer().LTIProtocolLogin()
	if err != nil {
		return err
	}
	return uc.ServeHTTP(c)
}
