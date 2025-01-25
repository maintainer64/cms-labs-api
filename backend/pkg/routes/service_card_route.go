package routes

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/controllers"
)

func V1ServiceCardRoutes(a *fiber.App) {
	group := a.Group("/api/v1/service-card")
	group.Post("/upsert", controllers.ServiceCardCreate)
	group.Post("/list", controllers.ServiceCardList)
	group.Post("/delete", controllers.ServiceCardDelete)
	group.Post("/get", controllers.ServiceCardGet)
}
