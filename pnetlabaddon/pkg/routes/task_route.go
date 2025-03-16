package routes

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/pnetlabaddon/app/controllers"
)

// V1TaskRoute func for describe group of task routes.
func V1TaskRoute(a *fiber.App) {
	group := a.Group("/pnet-lab-addon/api/v1/task")
	group.Post("/pnet-server-ping", controllers.PNETServerPing)
}
