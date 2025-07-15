package routes

import (
	"github.com/gofiber/fiber/v2"
)

func FiberRoutes(app *fiber.App) {
	// Routes.
	SwaggerRoute(app)
	V1TasksRoute(app)
	V1TopologyRoute(app)
	V1TokenRoute(app)
	NotFoundRoute(app)
}
