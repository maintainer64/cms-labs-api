package routes

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/controllers"
)

func V1UNLFileRoutes(a *fiber.App) {
	group := a.Group("/api/v1/unl-file")
	group.Post("/list", controllers.UNLFileList)
	group.Post("/get", controllers.UNLFileGet)
	group.Post("/sync", controllers.UNLFileSync)
}
