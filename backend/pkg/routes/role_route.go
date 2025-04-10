package routes

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/controllers"
)

func V1RoleRoutes(a *fiber.App) {
	group := a.Group("/api/v1/role")
	group.Post("/upsert", controllers.RoleUpsert)
	group.Post("/list", controllers.RoleList)
	group.Post("/delete", controllers.RoleDelete)
}
