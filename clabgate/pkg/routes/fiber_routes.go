package routes

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/maintainer64/cms-labs-api/shared/jsonrpc"
)

func FiberRoutes(app *fiber.App) {
	// Routes.
	SwaggerRoute(app)
	WorkspaceAuthRoutes(app)
	V1RpcRoute(app)
	jsonrpc.EmptyRoutes(app)
}
