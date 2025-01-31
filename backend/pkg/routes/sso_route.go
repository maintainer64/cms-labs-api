package routes

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/controllers"
)

// V1SSORoute func for describe group of openid protocol auth.
func V1SSORoute(a *fiber.App) {
	group := a.Group("/api/v1/sso")
	group.Get("/authorize", controllers.SSOAuthorize)
	group.Post("/authorize", controllers.SSOAuthorizePost)
	group.Post("/token", controllers.SSOToken)
	group.Post("/introspect", controllers.SSOIntrospect)
	group.Post("/userinfo", controllers.SSOUserInfo)
}
