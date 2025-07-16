package routes

import (
	fiber "github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/controllers"
)

func V1PNETServerRoutes(a *fiber.App) {
	group := a.Group("/api/v1/pnet-server")
	group.Post("/upsert", controllers.PNETServerCreate)
	group.Post("/list", controllers.PNETServerList)
	group.Post("/delete", controllers.PNETServerDelete)
	group.Post("/get", controllers.PNETServerGet)
	group.Post("/ping", controllers.PNETServerPing)
}
