package routes

import (
	"github.com/gofiber/fiber/v2"
)

func FiberRoutes(app *fiber.App) {
	// Routes.
	SwaggerRoute(app)
	V1AuthRoute(app)
	V1SSORoute(app)
	V1CurlRequestRoutes(app)
	V1UserRoutes(app)
	V1RoleRoutes(app)
	V1LTIFormRoutes(app)
	V1LTIAttemptRoutes(app)
	V1PNETServerRoutes(app)
	V1PNETServerQueueRoutes(app)
	V1ServiceCardRoutes(app)
	V2LTIRoutes(app)
	V1LTIRoutingRoutes(app)
	NotFoundRoute(app)
}
