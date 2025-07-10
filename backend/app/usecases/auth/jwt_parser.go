package auth

import (
	"errors"
	"strings"

	"gitlab.com/a10869/api-modules/shared/cms_client"

	"gitlab.com/a10869/api-modules/backend/pkg/configs"
	"gitlab.com/a10869/api-modules/shared/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// ExtractTokenMetadata func to extract metadata from JWT.
func ExtractTokenMetadata(
	c *fiber.Ctx,
	roles []string,
) (*cms_client.SSOTokenPublicData, error) {
	token, err := verifyToken(extractToken(c))
	if err != nil {
		return nil, utils.FiberValidationException{Status: fiber.StatusUnauthorized, Exception: err}
	}
	if !token.Valid {
		return nil, utils.FiberValidationException{Status: fiber.StatusUnauthorized, Exception: errors.New("invalid token")}
	}
	tokenData, err := cms_client.SSODecodeToken(token)
	if err != nil {
		return nil, utils.FiberValidationException{
			Status:    fiber.StatusUnauthorized,
			Exception: err,
		}
	}
	if len(roles) != 0 && !cms_client.SSOHasIntersection(roles, tokenData.Roles) {
		return nil, utils.FiberValidationException{
			Status:    fiber.StatusForbidden,
			Exception: errors.New("User with current role is not allow action"),
		}
	}
	return tokenData, nil
}

func extractToken(c *fiber.Ctx) string {
	bearToken := c.Get("Authorization")

	// Normally Authorization HTTP header.
	onlyToken := strings.Split(bearToken, " ")
	if len(onlyToken) == 2 {
		return onlyToken[1]
	}

	return ""
}

func verifyToken(tokenString string) (*jwt.Token, error) {
	publicKey := configs.AppConfig.JWT.AccessKey.PublicKey
	if publicKey == nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "JWT public key not configured")
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Проверяем алгоритм подписи
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "Unexpected signing method")
		}
		return publicKey, nil
	})

	if err != nil {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid token: "+err.Error())
	}

	if !token.Valid {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
	}

	return token, nil
}
