package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/di"
	"gitlab.com/a10869/api-modules/backend/app/usecases/auth"
	"gitlab.com/a10869/api-modules/backend/pkg/utils"
)

// TokensRenew method for renew access and refresh tokens.
// @Description Renew access and refresh tokens.
// @Summary renew access and refresh tokens
// @Tags Token
// @Accept json
// @Produce json
// @Param form body auth.RenewManagerInputDTO true "renew token form info"
// @Success 200 {object} auth.SSOTokenResponse
// @Router /v1/token/renew [post]
func TokensRenew(c *fiber.Ctx) error {
	refreshToken := c.Cookies("refresh-token", "")
	if refreshToken == "" {
		dto := auth.RenewManagerInputDTO{}
		err := utils.FiberValidatorBase(c, &dto)
		if err != nil {
			return utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
		}
		refreshToken = dto.RefreshToken
	}
	uc, err := di.NewDIContainer().SSOTokenUC()
	if err != nil {
		return utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	token, err := uc.Execute(
		auth.SSOTokenInputDTO{
			GrantType:    auth.TokenGrantTypeRefreshToken,
			RefreshToken: refreshToken,
		},
	)
	if err != nil {
		return utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	c.Cookie(&fiber.Cookie{
		Name:     "refresh-token",
		Value:    token.RefreshToken,
		Path:     "/",
		Expires:  auth.ExpiresRefreshCookie(),
		SameSite: fiber.CookieSameSiteNoneMode,
		Secure:   true,
	})
	return utils.FiberSuccessResponse{Result: token}
}

// TokensByCredentials method for grant access by email and password
// @Description Login by email and password.
// @Summary login by email and password
// @Tags Token
// @Accept json
// @Produce json
// @Param form body auth.RenewManagerCredentialsInputDTO true "credentials form info"
// @Success 200 {object} auth.SSOTokenResponse
// @Router /v1/token/login [post]
func TokensByCredentials(c *fiber.Ctx) error {
	dto := auth.RenewManagerCredentialsInputDTO{}
	err := utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	uc, err := di.NewDIContainer().AuthTokenManager()
	if err != nil {
		return utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	token, err := uc.NewJWTByCredentials(dto.Email, dto.Password)
	if err != nil {
		return utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	c.Cookie(&fiber.Cookie{
		Name:     "refresh-token",
		Value:    token.RefreshToken,
		Path:     "/",
		Expires:  auth.ExpiresRefreshCookie(),
		SameSite: fiber.CookieSameSiteNoneMode,
		Secure:   true,
	})
	return utils.FiberSuccessResponse{Result: token}
}

// TokensRemove method for remove access and refresh token
// @Description Logout.
// @Summary logout
// @Tags Token
// @Accept json
// @Produce json
// @Success 200 {object} auth.SSOTokenResponse
// @Router /v1/token/logout [post]
func TokensRemove(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "refresh-token",
		Value:    "",
		Path:     "/",
		Expires:  auth.ExpiresRefreshCookie(),
		SameSite: fiber.CookieSameSiteNoneMode,
		Secure:   true,
	})
	return utils.FiberSuccessResponse{Result: nil}
}

// TokensPasswordRecover method for change password
// @Description Change password.
// @Summary change password
// @Tags Token
// @Accept json
// @Produce json
// @Param form body auth.UserPasswordChangeInputDTO true "credentials form info"
// @Success 200 {object} auth.UserPasswordRecoverResponse
// @Router /v1/token/password_change [post]
func TokensPasswordRecover(c *fiber.Ctx) error {
	claims, err := auth.ExtractTokenMetadata(c, []string{})
	if err != nil {
		return err
	}
	dto := auth.UserPasswordChangeInputDTO{}
	if err := utils.FiberValidatorBase(c, &dto); err != nil {
		return err
	}
	uc, err := di.NewDIContainer().UserPasswordRecoverUC()
	if err != nil {
		return utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	response, err := uc.SetContext(claims).Execute(dto)
	if err != nil {
		return utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	return utils.FiberSuccessResponse{Result: response}
}
