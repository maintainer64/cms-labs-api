// Package main clabgate проект для работы с c9s внутри k8s кластера
package main

import (
	fiber "github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"           // load .env file automatically
	_ "gitlab.com/a10869/api-modules/clabgate/docs" // load API Docs files (Swagger)
	"gitlab.com/a10869/api-modules/clabgate/pkg/configs"
	"gitlab.com/a10869/api-modules/clabgate/pkg/middleware"
	"gitlab.com/a10869/api-modules/clabgate/pkg/routes"
	"gitlab.com/a10869/api-modules/clabgate/pkg/utils"
	"gitlab.com/a10869/api-modules/shared/logs"
)

// @title Clabernetes Gate API
// @version 1.0
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your@mail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /clabgate/api
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// Define Fiber config.
	config := configs.FiberConfig()
	logs.ZeroLogInit(configs.AppConfig.Debug)

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Middlewares.
	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.
	routes.FiberRoutes(app)         // Register Fiber's routes for app.

	// Start server (with or without graceful shutdown).
	if configs.AppConfig.Server.Layer == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
