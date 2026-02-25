package middleware

import (
	fiber "github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/usecases/auth"
	"gitlab.com/a10869/api-modules/shared/jsonrpc"
)

// NewJWTMiddleware создаёт middleware для извлечения и проверки JWT токена без ошибок
func NewJWTMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		c := &jsonrpc.Ctx{FiberCtx: ctx}
		// Извлекаем и парсим JWT токен
		tokenData, _ := auth.ExtractTokenMetadata(c, []string{})
		// Сохраняем распарсенные данные токена в locals
		if tokenData != nil {
			c.FiberCtx.Locals("jwtUser", tokenData)
			c.FiberCtx.Locals("x-user-email", tokenData.Email)
			c.FiberCtx.Locals("x-user-role", tokenData.UserRoleMain())
		}
		// Переходим к следующему middleware или обработчику
		return c.FiberCtx.Next()
	}
}
