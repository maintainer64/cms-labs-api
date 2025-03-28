package auth

import (
	"errors"
	"slices"
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
	token, err := verifyToken(c)
	if err != nil {
		return nil, utils.FiberValidationException{Status: fiber.StatusUnauthorized, Exception: err}
	}

	// Setting and checking token and credentials.
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, utils.FiberValidationException{Status: fiber.StatusUnauthorized, Exception: err}
	}
	tokenData := cms_client.SSOTokenPublicData{
		Iss:          claims["iss"].(string),
		Sub:          uint(claims["sub"].(float64)),
		Aud:          claims["aud"].(string),
		Exp:          int64(claims["sub"].(float64)),
		Iat:          int64(claims["iat"].(float64)),
		Nonce:        claims["nonce"].(string),
		Email:        claims["email"].(string),
		Name:         claims["name"].(string),
		ServerID:     uint(claims["server_id"].(float64)),
		Role:         claims["role"].(string),
		LastLaunchId: claims["last_launch_id"].(string),
	}
	if len(roles) != 0 && !slices.Contains(roles, tokenData.Role) {
		return nil, utils.FiberValidationException{
			Status:    fiber.StatusUnauthorized,
			Exception: errors.New("User with current role is not allow action"),
		}
	}
	return &tokenData, nil
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

func verifyToken(c *fiber.Ctx) (*jwt.Token, error) {
	tokenString := extractToken(c)

	token, err := jwt.Parse(tokenString, jwtKeyFunc)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(configs.AppConfig.JWT.SecretKey), nil
}
