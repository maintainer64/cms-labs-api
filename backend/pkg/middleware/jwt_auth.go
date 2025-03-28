package middleware

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/usecases/auth"
)

// NewJWTMiddleware создаёт middleware для извлечения и проверки JWT токена без ошибок
func NewJWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Извлекаем и парсим JWT токен
		tokenData, _ := auth.ExtractTokenMetadata(c, []string{})
		// Сохраняем распарсенные данные токена в locals
		if tokenData != nil {
			c.Locals("jwtUser", tokenData)
			c.Locals("x-user-email", tokenData.Email)
			c.Locals("x-user-role", tokenData.UserRoleMain())
		}
		// Переходим к следующему middleware или обработчику
		return c.Next()
	}
}
