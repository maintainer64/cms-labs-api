package routes

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/controllers"
)

func V1PNETServerQueueRoutes(a *fiber.App) {
	group := a.Group("/api/v1/pnet-server-queue")
	group.Post("/upsert", controllers.PNETServerQueueCreate)
	group.Post("/list", controllers.PNETServerQueueList)
}
