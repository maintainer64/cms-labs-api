package routes

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/clabgate/app/controllers"
)

func V1TokenRoute(a *fiber.App) {
	group := a.Group("/clabgate/api/v1/tokens")
	group.Get("/yaml", controllers.TokenAccessYaml)
	group.Get("/json", controllers.TokenAccessJson)
}
