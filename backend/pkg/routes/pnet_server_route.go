package routes

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/controllers"
)

func V1PNETServerRoutes(a *fiber.App) {
	group := a.Group("/api/v1/pnet-server")
	group.Post("/upsert", controllers.CreatePNETServer)
	group.Post("/list", controllers.ListPNETServer)
	group.Post("/delete", controllers.DeletePNETServer)
	group.Post("/get", controllers.GetPNETServer)
}
