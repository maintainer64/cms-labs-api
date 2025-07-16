package middleware

import (
	fiber "github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/shared/utils"
)

func InternalFormatterNew(debug bool) fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()
		switch e := err.(type) {
		case utils.FiberValidationException:
			return c.Status(e.Status).JSON(fiber.Map{
				"error": true,
				"msg":   e.Exception.Error(),
			})
		case utils.FiberSuccessResponse:
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"error":  false,
				"msg":    nil,
				"result": e.Result,
			})
		case nil:
			return e
		default:
			if debug {
				return e
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   "Internal Server Error",
			})
		}
	}
}
