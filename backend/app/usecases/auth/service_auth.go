package auth

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/rs/zerolog"
	"gitlab.com/a10869/api-modules/backend/app/queries"
	"gitlab.com/a10869/api-modules/shared/jsonrpc"
)

type ServiceAuthorizeUC struct {
	PNETServerQueries *queries.PNETServerQueries
	*zerolog.Logger
}

var (
	ServiceNotPermissions = jsonrpc.NewRpcError("unauthorized", "service not permissions")
)

func (u *ServiceAuthorizeUC) Execute(c *jsonrpc.Ctx) error {
	authHeader := c.FiberCtx.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Basic ") {
		c.FiberCtx.Set("WWW-Authenticate", "Basic realm=\"Restricted\"")
		return ServiceNotPermissions
	}
	encodedCredentials := strings.TrimPrefix(authHeader, "Basic ")
	decoded, err := base64.StdEncoding.DecodeString(encodedCredentials)
	if err != nil {
		u.Logger.Info().Msg(fmt.Sprintf("Error decoding basic auth credentials: %v", err))
		return ServiceNotPermissions
	}
	credentials := strings.SplitN(string(decoded), ":", 2)
	if len(credentials) != 2 {
		return ServiceNotPermissions
	}
	username := credentials[0]
	password := credentials[1]
	serverModel, err := u.PNETServerQueries.GetByClientId(username)
	if err != nil {
		u.Logger.Info().Msg(fmt.Sprintf("Error getting client basic auth model by username: %v", username))
		return ServiceNotPermissions
	}
	valid := serverModel.IsActive && serverModel.Token != "" && serverModel.ClientID != "" && serverModel.Token == password
	u.Logger.Info().Msg(
		fmt.Sprintf(
			"auth with service: %+v is valid = %+v",
			username,
			valid,
		),
	)
	if !valid {
		return ServiceNotPermissions
	}
	c.FiberCtx.Locals("x-service-id", serverModel.ClientID)
	return nil
}
