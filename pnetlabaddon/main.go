// Дополнение в PNETLab. API для аутентификации с помощью общего компонента SSO
package main

import (
	fiber "github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload" // load .env file automatically
	"gitlab.com/a10869/api-modules/pnetlabaddon/app/di"
	_ "gitlab.com/a10869/api-modules/pnetlabaddon/docs" // load API Docs files (Swagger)
	"gitlab.com/a10869/api-modules/pnetlabaddon/pkg/configs"
	"gitlab.com/a10869/api-modules/pnetlabaddon/pkg/middleware"
	"gitlab.com/a10869/api-modules/pnetlabaddon/pkg/routes"
	"gitlab.com/a10869/api-modules/pnetlabaddon/pkg/utils"
	"gitlab.com/a10869/api-modules/shared/logs"
	"gitlab.com/a10869/api-modules/shared/scheduler"
)

// @title API
// @version 1.0
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your@mail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// Define Fiber config.
	config := configs.FiberConfig()
	logs.ZeroLogInit(configs.AppConfig.Debug)

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Define scheduler app with config
	schedulerInterval := scheduler.NewScheduler(
		configs.AppConfig.SchedulerInterval,
	).AddTask(&di.PnetServerPingTask{})

	// Middlewares.
	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.
	routes.FiberRoutes(app)         // Register Fiber's routes for app.

	// Start server (with or without graceful shutdown).
	if configs.AppConfig.Server.Layer == "dev" {
		utils.StartServer(app, schedulerInterval)
	} else {
		utils.StartServerWithGracefulShutdown(app, schedulerInterval)
	}
}
