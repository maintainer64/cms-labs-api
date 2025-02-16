package routes

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/pnetlabaddon/app/controllers"
)

// V1SSORoute func for describe group of openid protocol auth.
func V1SSORoute(a *fiber.App) {
	group := a.Group("/pnet-lab-addon/api/v1/sso")
	group.Get("/login", controllers.SSOFirstFactor)
	group.Get("/openid", controllers.SSOSecondFactor)
}
