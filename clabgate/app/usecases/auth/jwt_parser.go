package auth

import (
	"errors"
	"strings"

	"github.com/golang-jwt/jwt/v5"

	"gitlab.com/a10869/api-modules/shared/cms_client"

	"gitlab.com/a10869/api-modules/shared/utils"

	"github.com/gofiber/fiber/v2"
)

// ExtractTokenMetadata func to extract metadata from JWT.
func ExtractTokenMetadata(
	c *fiber.Ctx,
	roles []string,
) (*cms_client.SSOTokenPublicData, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(extractToken(c), jwt.MapClaims{})
	if err != nil {
		return nil, utils.FiberValidationException{Status: fiber.StatusUnauthorized, Exception: err}
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
