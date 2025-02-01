package auth

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gitlab.com/a10869/api-modules/backend/app/queries"
	"gitlab.com/a10869/api-modules/backend/pkg/utils"
)

type SSOIntrospectInputDTO struct {
	Token string `json:"token"`
}

type SSOIntrospectUC struct {
	TokenAttemptQueries *queries.TokenAttemptQueries
	PNETServerQueries   *queries.PNETServerQueries
}

func (u *SSOIntrospectUC) Execute(inputDTO SSOIntrospectInputDTO) (*SSOTokenIntrospect, error) {
	entity, err := u.ByRefresh(inputDTO)
	if err == nil {
		return entity, err
	}
	entity, err = u.ByAccess(inputDTO)
	return entity, err
}

func (u *SSOIntrospectUC) ByRefresh(inputDTO SSOIntrospectInputDTO) (*SSOTokenIntrospect, error) {
	introspect := &SSOTokenIntrospect{}
	introspect.Active = true
	introspect.Scope = []string{"default"}
	introspect.TokenType = "refresh"
	expiresRefreshToken, err := ParseRefreshToken(inputDTO.Token)
	if err != nil {
		return nil, err
	}
	now := time.Now().Unix()
	introspect.Exp = now
	if now >= expiresRefreshToken {
		introspect.Active = false
	}
	attempt, err := u.TokenAttemptQueries.GetByToken(inputDTO.Token)
	if err != nil {
		return introspect, err
	}
	introspect.Sub = fmt.Sprintf("%v", attempt.UserID)
	introspect.Iat = attempt.CreatedAt.Unix()
	if attempt.ServerID == 0 {
		return introspect, nil
	}
	server, err := u.PNETServerQueries.Get(attempt.ServerID)
	if err != nil {
		return nil, err
	}
	introspect.Username = server.Name
	introspect.ClientID = server.ClientID
	if !server.IsActive {
		introspect.Active = false
	}
	return introspect, nil
}

func (u *SSOIntrospectUC) ByAccess(inputDTO SSOIntrospectInputDTO) (*SSOTokenIntrospect, error) {
	introspect := &SSOTokenIntrospect{}
	introspect.TokenType = "access"
	introspect.Active = true
	introspect.Scope = []string{"default"}
	token, err := jwt.Parse(inputDTO.Token, jwtKeyFunc)
	if err != nil {
		return nil, utils.FiberValidationException{Status: fiber.StatusUnauthorized, Exception: err}
	}
	// Setting and checking token and credentials.
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, utils.FiberValidationException{Status: fiber.StatusUnauthorized, Exception: err}
	}
	userID := uint(claims["id"].(float64))
	var serverID uint = 0
	if claims["serverID"] != nil {
		serverID = uint(claims["serverID"].(float64))
	}
	introspect.ClientID = fmt.Sprintf("%v", serverID)
	introspect.Sub = fmt.Sprintf("%v", userID)
	introspect.Exp = int64(claims["expires"].(float64))
	attempt, err := u.TokenAttemptQueries.GetByParams(userID, serverID)
	if err != nil {
		introspect.Active = false
		return introspect, nil
	}
	introspect.Iat = attempt.CreatedAt.Unix()
	if attempt.ServerID == 0 {
		return introspect, nil
	}
	server, err := u.PNETServerQueries.Get(attempt.ServerID)
	if err != nil {
		return introspect, nil
	}
	introspect.Username = server.Name
	introspect.ClientID = server.ClientID
	if !server.IsActive {
		introspect.Active = false
	}
	return introspect, nil
}
