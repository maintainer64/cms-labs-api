package routes

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/controllers"
)

func V1UserRoutes(a *fiber.App) {
	group := a.Group("/api/v1/user")
	group.Post("/upsert", controllers.UserCreate)
	group.Post("/list", controllers.UserList)
	group.Post("/get", controllers.UserGet)
}
