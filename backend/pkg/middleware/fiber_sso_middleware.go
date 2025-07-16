package middleware

import (
	"strings"

	fiber "github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/shared/utils"
)

func InternalSSOFormatterNew(debug bool) fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()
		if !strings.HasPrefix(c.Path(), "/api/v1/sso") {
			// Форматтер ошибок SSO общего формата
			return err
		}
		switch e := err.(type) {
		case utils.FiberValidationException:
			return c.Status(e.Status).JSON(fiber.Map{
				"error":             "Ошибка авторизации",
				"error_description": e.Exception.Error(),
			})
		case nil:
			return e
		default:
			if debug {
				return e
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":             true,
				"error_description": "Internal Server Error",
			})
		}
	}
}
