package auth

import (
	"fmt"
	"strings"

	jwt "github.com/golang-jwt/jwt/v5"

	"github.com/maintainer64/cms-labs-api/clabgate/pkg/configs"
	"github.com/maintainer64/cms-labs-api/shared/cms_client"

	"github.com/maintainer64/cms-labs-api/shared/jsonrpc"
)

// ExtractTokenMetadata func to extract metadata from JWT.
func ExtractTokenMetadata(
	c *jsonrpc.Ctx,
	roles []string,
) (*cms_client.SSOTokenPublicData, error) {
	publicKey := configs.AppConfig.JWT.PublicKey
	if publicKey == nil {
		return nil, jsonrpc.NewRpcError("unauthorized", "JWT public key is not configured")
	}

	options := []jwt.ParserOption{
		jwt.WithValidMethods([]string{jwt.SigningMethodRS256.Alg()}),
		jwt.WithExpirationRequired(),
	}
	if configs.AppConfig.JWT.Issuer != "" {
		options = append(options, jwt.WithIssuer(configs.AppConfig.JWT.Issuer))
	}
	if configs.AppConfig.JWT.Audience != "" {
		options = append(options, jwt.WithAudience(configs.AppConfig.JWT.Audience))
	}

	token, err := jwt.Parse(extractToken(c), func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodRS256 {
			return nil, fmt.Errorf("unexpected signing method %q", token.Method.Alg())
		}
		return publicKey, nil
	}, options...)
	if err != nil {
		return nil, jsonrpc.NewRpcError("unauthorized", "unauthorized")
	}
	if !token.Valid {
		return nil, jsonrpc.NewRpcError("unauthorized", "invalid token")
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
