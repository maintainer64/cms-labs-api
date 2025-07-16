package middleware

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"gitlab.com/a10869/api-modules/clabgate/pkg/configs"
	"gitlab.com/a10869/api-modules/shared/logs"
	"gitlab.com/a10869/api-modules/shared/middleware"
)

// FiberMiddleware provide Fiber's built-in middlewares.
// See: https://docs.gofiber.io/api/middleware
func FiberMiddleware(a *fiber.App) {
	a.Use(
		// Add CORS to each route.
		cors.New(),
		// Options запросы всегда 204
		logs.NewOptionsFiberNoContent(),
		// Header X-Request-ID
		requestid.New(requestid.Config{
			ContextKey: "x-request-id",
		}),
		// Add simple healthcheck.
		healthcheck.New(),
		// Service JWT extractor
		NewJWTMiddleware(),
		// Add simple logger.
		logs.NewFiberZerologLogger(),
		// InternalFormatterException
		middleware.InternalFormatterNew(configs.AppConfig.Debug),
	)
}
