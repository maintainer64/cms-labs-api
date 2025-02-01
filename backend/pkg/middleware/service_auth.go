package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"gitlab.com/a10869/api-modules/backend/app/di"
)

func NewServiceAuthMiddleware() fiber.Handler {
	paths := []string{
		"/api/v1/sso/token",
		"/api/v1/sso/introspect",
	}
	return basicauth.New(
		basicauth.Config{
			Authorizer: func(username string, password string) bool {
				log.Info().Msg(fmt.Sprintf("ServiceAuthMiddleware: auth with service: %+v", username))
				repos, err := di.NewDIContainer().Queries()
				if err != nil {
					return false
				}
				serverModel, err := repos.PNETServerQueries.GetByClientId(username)
				if err != nil {
					return false
				}
				valid := serverModel.IsActive == true && serverModel.Token != "" && serverModel.ClientID != "" && serverModel.Token == password
				log.Info().Msg(
					fmt.Sprintf(
						"ServiceAuthMiddleware: auth with service: %+v is valid = %+v",
						username,
						valid,
					),
				)
				return valid
			},
			Next: func(c *fiber.Ctx) bool {
				for _, path := range paths {
					if c.Path() == path {
						// NEED CHECK
						return false
					}
				}
				// SKIP CHECK
				return true
			},
		},
	)
}
