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
	tokenData, err := decodeToken(token)
	if err != nil {
		return nil, utils.FiberValidationException{
			Status:    fiber.StatusUnauthorized,
			Exception: err,
		}
	}
	if len(roles) != 0 && !hasIntersection(roles, tokenData.Roles) {
		return nil, utils.FiberValidationException{
			Status:    fiber.StatusUnauthorized,
			Exception: errors.New("User with current role is not allow action"),
		}
	}
	return tokenData, nil
}

// Вспомогательная функция для проверки пересечения ролей
func hasIntersection(allowedRoles, userRoles []string) bool {
	for _, userRole := range userRoles {
		for _, allowedRole := range allowedRoles {
			if userRole == allowedRole {
				return true
			}
		}
	}
	return false
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

func getStringSlice(claims map[string]interface{}, key string) []string {
	var result []string
	if val, ok := claims[key]; ok {
		if slice, ok := val.([]interface{}); ok {
			for _, item := range slice {
				if str, ok := item.(string); ok {
					result = append(result, str)
				}
			}
		}
	}
	return result
}

func decodeToken(jwtToken *jwt.Token) (*cms_client.SSOTokenPublicData, error) {
	// Setting and checking token and credentials.
	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok || !jwtToken.Valid {
		return nil, utils.FiberValidationException{
			Status:    fiber.StatusUnauthorized,
			Exception: errors.New("token invalid"),
		}
	}
	tokenData := cms_client.SSOTokenPublicData{
		Iss:          claims["iss"].(string),
		Sub:          claims["sub"].(string),
		Aud:          claims["aud"].(string),
		Exp:          int64(claims["exp"].(float64)),
		Iat:          int64(claims["iat"].(float64)),
		Nonce:        claims["nonce"].(string),
		Email:        claims["email"].(string),
		Name:         claims["name"].(string),
		ServerID:     uint(claims["server_id"].(float64)),
		Roles:        getStringSlice(claims, "roles"),
		LastLaunchId: claims["last_launch_id"].(string),
	}
	return &tokenData, nil
}
