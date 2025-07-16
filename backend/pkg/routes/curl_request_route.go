package routes

import (
	fiber "github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/controllers"
)

func V1CurlRequestRoutes(a *fiber.App) {
	group := a.Group("/api/v1/curl-request")
	group.Post("/upsert", controllers.CurlRequestCreate)
	group.Post("/list", controllers.CurlRequestList)
	group.Post("/delete", controllers.CurlRequestDelete)
	group.Post("/get", controllers.CurlRequestGet)
	group.Post("/execute", controllers.CurlRequestExecute)
}
