package jsonrpc

import "github.com/gofiber/fiber/v2"

func EmptyRoutes(app *fiber.App) {
	app.Use(
		// Anonymous function.
		func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"jsonrpc": "2.0",
				"error":   fiber.Map{"code": InvalidRequest, "message": "Entrypoint is incorrect"},
				"id":      nil,
			})
		},
	)
}
