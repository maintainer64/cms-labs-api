package routes

import (
	fiber "github.com/gofiber/fiber/v2"

	swagger "github.com/gofiber/swagger"
)

// SwaggerRoute func for describe group of API Docs routes.
func SwaggerRoute(a *fiber.App) {
	route := a.Group("/api/docs")
	route.Get("*", swagger.HandlerDefault)
}
