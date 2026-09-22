package routes

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/maintainer64/cms-labs-api/backend/app/controllers"
)

// V2LTIRoutes func for describe group of LTI routes.
func V2LTIRoutes(a *fiber.App) {
	group := a.Group("/api/v2/lti")
	group.Post("/login", controllers.LTILogin)
	group.Get("/login", controllers.LTILogin)
	group.Post("/launch", controllers.LTILaunch)
	group.Get("/launch", controllers.LTILaunch)
}
