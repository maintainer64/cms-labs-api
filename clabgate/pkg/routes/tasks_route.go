package routes

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/clabgate/app/controllers"
)

func V1TasksRoute(a *fiber.App) {
	group := a.Group("/clabgate/api/v1/tasks")
	group.Post("/list", controllers.TasksList)
}
