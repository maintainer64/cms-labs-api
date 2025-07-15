package routes

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/clabgate/app/controllers"
)

func V1TokenRoute(a *fiber.App) {
	group := a.Group("/clabgate/api/v1/tokens")
	group.Post("/yaml", controllers.TokenAccessYaml)
	group.Post("/json", controllers.TokenAccessJson)
}
