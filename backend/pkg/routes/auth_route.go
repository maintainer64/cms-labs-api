package routes

import (
	fiber "github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/controllers"
)

// V1AuthRoute func for describe group of JWT routes.
func V1AuthRoute(a *fiber.App) {
	group := a.Group("/api/v1/token")
	group.Post("/renew", controllers.TokensRenew)
	group.Post("/password_change", controllers.TokensPasswordRecover)
	group.Post("/login", controllers.TokensByCredentials)
	group.Post("/logout", controllers.TokensRemove)
}
