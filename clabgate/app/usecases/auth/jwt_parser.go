package auth

import (
	"strings"

	jwt "github.com/golang-jwt/jwt/v5"

	"gitlab.com/a10869/api-modules/shared/cms_client"

	"gitlab.com/a10869/api-modules/shared/jsonrpc"
)

// ExtractTokenMetadata func to extract metadata from JWT.
func ExtractTokenMetadata(
	c *jsonrpc.Ctx,
	roles []string,
) (*cms_client.SSOTokenPublicData, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(extractToken(c), jwt.MapClaims{})
	if err != nil {
		return nil, jsonrpc.NewRpcError("unauthorized", "unauthorized")
	}
	tokenData, err := cms_client.SSODecodeToken(token)
	if err != nil {
		return nil, jsonrpc.NewRpcError("unauthorized", "unauthorized")
	}
	if len(roles) != 0 && !cms_client.SSOHasIntersection(roles, tokenData.Roles) {
		return nil, jsonrpc.NewRpcError("forbidden", "user with current role is not allow action")
	}
	return tokenData, nil
}

func extractToken(c *jsonrpc.Ctx) string {
	bearToken := c.FiberCtx.Get("Authorization")

	// Normally Authorization HTTP header.
	onlyToken := strings.Split(bearToken, " ")
	if len(onlyToken) == 2 {
		return onlyToken[1]
	}

	return ""
}
