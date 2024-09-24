package routes

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/controllers"
)

// V2LTIRoutes func for describe group of LTI routes.
func V2LTIRoutes(a *fiber.App) {
	group := a.Group("/api/v2/lti")
	group.Post("/launch/authenticate/", controllers.AuthLTI)
	group.Get("/launch/public_keys/", controllers.GetLTIPublicKeys)
	group.Post("/launch", controllers.LaunchLTI)
	group.Get("/", controllers.GetLTIInfo)
}
