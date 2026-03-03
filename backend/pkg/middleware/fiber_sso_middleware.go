package middleware

import (
	"strings"

	fiber "github.com/gofiber/fiber/v2"
)

func InternalSSOFormatterNew() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()
		if err == nil {
			return nil
		}
		if !strings.HasPrefix(c.Path(), "/api/v1/sso") {
			// Форматтер ошибок SSO общего формата
			return err
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":             "Ошибка авторизации",
			"error_description": err.Error(),
		})
	}
}
