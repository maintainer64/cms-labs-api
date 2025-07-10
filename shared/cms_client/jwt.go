package cms_client

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gitlab.com/a10869/api-modules/shared/utils"
)

func SSOTokenGetStringSlice(claims map[string]interface{}, key string) []string {
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

func SSODecodeToken(jwtToken *jwt.Token) (*SSOTokenPublicData, error) {
	// Setting and checking token and credentials.
	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, utils.FiberValidationException{
			Status:    fiber.StatusUnauthorized,
			Exception: errors.New("token invalid"),
		}
	}
	tokenData := SSOTokenPublicData{
		Iss:          claims["iss"].(string),
		Sub:          claims["sub"].(string),
		Aud:          claims["aud"].(string),
		Azp:          claims["azp"].(string),
		Exp:          int64(claims["exp"].(float64)),
		Iat:          int64(claims["iat"].(float64)),
		Nonce:        claims["nonce"].(string),
		Email:        claims["email"].(string),
		Name:         claims["name"].(string),
		ServerID:     uint(claims["server_id"].(float64)),
		Roles:        SSOTokenGetStringSlice(claims, "roles"),
		LastLaunchId: claims["last_launch_id"].(string),
	}
	return &tokenData, nil
}

// SSOHasIntersection Вспомогательная функция для проверки пересечения ролей
func SSOHasIntersection(allowedRoles, userRoles []string) bool {
	for _, userRole := range userRoles {
		for _, allowedRole := range allowedRoles {
			if userRole == allowedRole {
				return true
			}
		}
	}
	return false
}
