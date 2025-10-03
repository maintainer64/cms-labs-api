package routes

import (
	fiber "github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/shared/jsonrpc"
)

func FiberRoutes(app *fiber.App) {
	// Routes.
	SwaggerRoute(app)
	V1SSORoute(app)
	V1RpcRoute(app)
	V2LTIRoutes(app)
	jsonrpc.EmptyRoutes(app)
}
