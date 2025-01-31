package middleware

import (
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"gitlab.com/a10869/api-modules/backend/pkg/configs"
	"gitlab.com/a10869/api-modules/shared/logs"
)

// FiberMiddleware provide Fiber's built-in middlewares.
// See: https://docs.gofiber.io/api/middleware
func FiberMiddleware(a *fiber.App) {
	a.Use(
		// Add CORS to each route.
		cors.New(),
		// Header X-Request-ID
		requestid.New(),
		// Add simple logger.
		fiberzerolog.New(fiberzerolog.Config{
			Logger: &logs.ZeroLog,
		}),
		// Add simple healthcheck.
		healthcheck.New(),
		// Service middleware
		NewServiceAuthMiddleware(),
		// InternalFormatterException
		InternalFormatterNew(configs.AppConfig.Debug),
	)
}
