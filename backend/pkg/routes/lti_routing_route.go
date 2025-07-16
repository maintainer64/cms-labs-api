package routes

import (
	fiber "github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/controllers"
)

func V1LTIRoutingRoutes(a *fiber.App) {
	group := a.Group("/api/v1/lti-routing")
	group.Post("/upsert", controllers.LTIRoutingCreate)
	group.Post("/list", controllers.LTIRoutingList)
	group.Post("/delete", controllers.LTIRoutingDelete)
	group.Post("/get", controllers.LTIRoutingGet)
}
