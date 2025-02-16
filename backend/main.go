// Package main CMS Labs Core проект LTI <-> SSO компонентов
package main

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload" // load .env file automatically
	"gitlab.com/a10869/api-modules/backend/app/di"
	_ "gitlab.com/a10869/api-modules/backend/docs" // load API Docs files (Swagger)
	"gitlab.com/a10869/api-modules/backend/pkg/configs"
	"gitlab.com/a10869/api-modules/backend/pkg/middleware"
	"gitlab.com/a10869/api-modules/backend/pkg/routes"
	"gitlab.com/a10869/api-modules/backend/pkg/utils"
	"gitlab.com/a10869/api-modules/shared/logs"
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
	logs.ZeroLogInit(configs.AppConfig.Debug)

	container, err := di.NewDIContainer(&logs.ZeroLoggerConf{Name: "main"})
	if err != nil {
		panic(err)
	}

	startup := container.TaskStartup()
	if err = startup.Startup(); err != nil {
		panic(err)
	}

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
