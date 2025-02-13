package cms_client

import (
	"github.com/golang-jwt/jwt/v5"
)

const SSORefreshTokenName = "cms-labs-refresh-token" // #nosec G101

type SSOToken struct {
	AccessToken  string `json:"access_token" required:"true"`
	RefreshToken string `json:"refresh_token" required:"true"`
	TokenType    string `json:"token_type" required:"true"`
	ExpiresIn    int64  `json:"expires_in" required:"true"`
	State        string `json:"state"`
	UserId       uint   `json:"user_id"`
}

type SSOTokenResponse struct {
	Error  bool     `json:"error" validate:"required"`
	Msg    string   `json:"msg" validate:"required"`
	Result SSOToken `json:"result"`
}

// SSOTokenPublicData struct to describe public payload object.
type SSOTokenPublicData struct {
	Id           uint   `json:"id"`
	ServerID     uint   `json:"server_id"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	Role         string `json:"role"`
	LastLaunchId string `json:"last_launch_id"`
	Expires      int64  `json:"exp"`
}

type SSOTokenPublicDataResponse struct {
	Error  bool               `json:"error" validate:"required"`
	Msg    string             `json:"msg" validate:"required"`
	Result SSOTokenPublicData `json:"result"`
}

func (t *SSOTokenPublicData) JWTClaims() jwt.MapClaims {
	return jwt.MapClaims{
		"id":             t.Id,
		"email":          t.Email,
		"server_id":      t.ServerID,
		"name":           t.Name,
		"role":           t.Role,
		"last_launch_id": t.LastLaunchId,
		"expires":        t.Expires,
	}
}

type SSOTokenPublicExtraParams struct {
}
