package routes

import (
	fiber "github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/controllers"
)

func V1LTIAttemptRoutes(a *fiber.App) {
	group := a.Group("/api/v1/lti-attempt")
	group.Post("/create", controllers.LTIAttemptCreate)
	group.Post("/edit", controllers.LTIAttemptEdit)
	group.Post("/list", controllers.LTIAttemptList)
	group.Post("/delete", controllers.LTIAttemptDelete)
	group.Post("/get", controllers.LTIAttemptGet)
}
