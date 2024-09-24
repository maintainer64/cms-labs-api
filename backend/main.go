package main

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/pkg/configs"
	"gitlab.com/a10869/api-modules/backend/pkg/middleware"
	"gitlab.com/a10869/api-modules/backend/pkg/routes"
	"gitlab.com/a10869/api-modules/backend/pkg/utils"

	_ "gitlab.com/a10869/api-modules/backend/docs" // load API Docs files (Swagger)

	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)

// @title API
// @version 1.0
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your@mail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Middlewares.
	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	// Routes.
	routes.SwaggerRoute(app)
	routes.V1LTIFormRoutes(app)
	routes.V1PNETServerRoutes(app)
	routes.V2LTIRoutes(app)
	routes.NotFoundRoute(app)

	// Start server (with or without graceful shutdown).
	if configs.AppConfig.Server.Layer == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
