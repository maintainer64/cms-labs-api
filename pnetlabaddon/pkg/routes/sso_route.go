package routes

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/maintainer64/cms-labs-api/pnetlabaddon/app/controllers"
)

// V1SSORoute func for describe group of openid protocol auth.
func V1SSORoute(a *fiber.App) {
	group := a.Group("/pnet-lab-addon/api/v1/sso")
	group.Get("/login", controllers.SSOFirstFactor)
	group.Get("/openid", controllers.SSOSecondFactor)
}
