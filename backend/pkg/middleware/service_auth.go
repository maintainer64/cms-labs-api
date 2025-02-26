package middleware

import (
	"fmt"
	"strings"

	"gitlab.com/a10869/api-modules/shared/logs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"gitlab.com/a10869/api-modules/backend/app/di"
)

func NewServiceAuthMiddleware() fiber.Handler {
	paths := []string{
		"/api/v1/sso/token",
		"/api/v1/sso/introspect",
		"/api/v1/unl-file/sync",
	}
	return basicauth.New(
		basicauth.Config{
			ContextUsername: "x-service-id",
			Authorizer: func(username string, password string) bool {
				diLoggerConf := &logs.ZeroLoggerConf{Name: "middleware.service_auth"}
				log := logs.NewZeroLogger(diLoggerConf)
				log.Info().Msg(fmt.Sprintf("ServiceAuthMiddleware: auth with service: %+v", username))
				repos, err := di.NewDIContainer(diLoggerConf)
				if err != nil {
					return false
				}
				defer repos.Close()
				serverModel, err := repos.Queries.PNETServerQueries.GetByClientId(username)
				if err != nil {
					return false
				}
				valid := serverModel.IsActive && serverModel.Token != "" && serverModel.ClientID != "" && serverModel.Token == password
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
					if strings.HasPrefix(c.Path(), path) {
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
