package routes

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/controllers"
)

func V1LTIFormRoutes(a *fiber.App) {
	group := a.Group("/api/v1/lti-form")
	group.Post("/upsert", controllers.CreateLTIForm)
	group.Post("/list", controllers.ListLTIForm)
	group.Post("/delete", controllers.DeleteLTIForm)
	group.Post("/get", controllers.GetLTIForm)
}
