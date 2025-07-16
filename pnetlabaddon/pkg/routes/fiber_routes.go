package routes

import (
	fiber "github.com/gofiber/fiber/v2"
)

func FiberRoutes(app *fiber.App) {
	// Routes.
	SwaggerRoute(app)
	V1SSORoute(app)
	V1TaskRoute(app)
}
