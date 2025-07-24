package routes

import (
	fiber "github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/controllers"
)

// V1StoreRoute func for describe group of store global.
func V1StoreRoute(a *fiber.App) {
	group := a.Group("/api/v1")
	group.Get("/global-store", controllers.StoreGet)
	group.Post("/global-store", controllers.StoreSet)
}
