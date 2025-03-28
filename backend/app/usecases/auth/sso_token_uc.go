package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/rs/zerolog"

	"gitlab.com/a10869/api-modules/shared/cms_client"

	"gitlab.com/a10869/api-modules/backend/app/queries"
)

type SSOTokenInputDTO struct {
	GrantType    string `json:"grant_type"`
	RedirectUri  string `json:"redirect_uri"`
	Code         string `json:"code"`
	RefreshToken string `json:"refresh_token"`
}

const (
	TokenGrantTypeAuthorizationCode = "authorization_code"
	TokenGrantTypeRefreshToken      = "refresh_token"
)

type SSOTokenUC struct {
	TokenAttemptQueries *queries.TokenAttemptQueries
	TokenManager        *TokenManager
	PNETServerQueries   *queries.PNETServerQueries
	Logger              *zerolog.Logger
	IssId               string
}

// SwaggerSSOToken copy of cms_client.SSOToken
type SwaggerSSOToken struct {
	AccessToken  string `json:"access_token" required:"true"`
	RefreshToken string `json:"refresh_token" required:"true"`
	TokenType    string `json:"token_type" required:"true"`
	ExpiresIn    int64  `json:"expires_in" required:"true"`
	State        string `json:"state"`
	UserId       uint   `json:"user_id"`
}

// SwaggerSSOTokenResponse copy of cms_client.SSOTokenResponse
type SwaggerSSOTokenResponse struct {
	Error  bool            `json:"error" validate:"required"`
	Msg    string          `json:"msg" validate:"required"`
	Result SwaggerSSOToken `json:"result"`
}

// SwaggerSSOTokenPublicData copy of cms_client.SSOTokenPublicData
type SwaggerSSOTokenPublicData struct {
	// Iss. Идентификатор эмитента токена
	Iss string `json:"iss"`
	// Sub. Уникальный идентификатор пользователя в системе OpenID Provider (OP)
	Sub uint `json:"sub"`
	// Aud. Получатель токена (обычно client_id приложения, запрашивающего токен)
	Aud string `json:"aud"`
	// Exp. Время истечения срока действия токена (в Unix timestamp)
	Exp int64 `json:"exp"`
	// Iat. Время выдачи токена (в Unix timestamp)
	Iat int64 `json:"iat"`
	// Nonce.(Если запрос авторизации включал nonce) Случайное значение для предотвращения атак подмены
	Nonce string `json:"nonce"`
	// Email. Почта уникальная пользователя
	Email string `json:"email"`
	// Name. Полное ФИО пользователя
	Name string `json:"name"`
	// ServerID ID сервера аутентификации (как с Iss)
	ServerID uint `json:"server_id"`
	// Role. Роль пользователя
	Role string `json:"role"`
	// LastLaunchId. ID пользователя SSO через LMS систему
	LastLaunchId string `json:"last_launch_id"`
}

func (u *SSOTokenUC) SetContext(issId string) *SSOTokenUC {
	u.IssId = issId
	return u
}

func (u *SSOTokenUC) Execute(inputDTO SSOTokenInputDTO) (*cms_client.SSOToken, error) {
	if err := u.validate(inputDTO); err != nil {
		return &cms_client.SSOToken{}, err
	}
	if inputDTO.GrantType == TokenGrantTypeAuthorizationCode {
		return u.ByAuthCode(inputDTO)
	}
	if inputDTO.GrantType == TokenGrantTypeRefreshToken {
		return u.ByRefresh(inputDTO)
	}
	return nil, errors.New("invalid grant type")
}

func (u *SSOTokenUC) validate(inputDTO SSOTokenInputDTO) error {
	if inputDTO.GrantType == TokenGrantTypeAuthorizationCode && inputDTO.RedirectUri != "" && inputDTO.Code != "" {
		return nil
	}
	if inputDTO.GrantType == TokenGrantTypeRefreshToken && inputDTO.RefreshToken != "" {
		return nil
	}
	return errors.New("Invalid params grant_type")
}

func (u *SSOTokenUC) ByAuthCode(inputDTO SSOTokenInputDTO) (*cms_client.SSOToken, error) {
	attempt, err := u.TokenAttemptQueries.GetByAuthCode(inputDTO.Code)
	if err != nil {
		return nil, err
	}
	server, err := u.PNETServerQueries.Get(attempt.ServerID)
	if err != nil {
		return nil, err
	}
	u.Logger.Info().Msg(fmt.Sprintf("Token get by auth code by server_id: %+v", server.ID))
	if !server.IsActive {
		return nil, errors.New("server is not active")
	}
	if !server.HasPrefixUrl(inputDTO.RedirectUri) {
		return nil, errors.New("invalid redirect_uri")
	}
	return u.TokenManager.NewJWTByUserId(u.IssId, attempt.UserID, attempt.ServerID, attempt.State)
}

func (u *SSOTokenUC) ByRefresh(inputDTO SSOTokenInputDTO) (*cms_client.SSOToken, error) {
	expiresRefreshToken, err := ParseRefreshToken(inputDTO.RefreshToken)
	if err != nil {
		return nil, err
	}
	now := time.Now().Unix()
	if now >= expiresRefreshToken {
		return nil, errors.New("refresh token is expired")
	}
	attempt, err := u.TokenAttemptQueries.GetByToken(inputDTO.RefreshToken)
	if err != nil {
		return nil, err
	}
	u.Logger.Info().Msg(fmt.Sprintf("Token get by refresh token by server_id: %+v", attempt.ServerID))
	if attempt.ServerID == 0 {
		return u.TokenManager.NewJWTByUserId(u.IssId, attempt.UserID, attempt.ServerID, "")
	}
	server, err := u.PNETServerQueries.Get(attempt.ServerID)
	if err != nil {
		return nil, err
	}
	if !server.IsActive {
		return nil, errors.New("server is not active")
	}
	return u.TokenManager.NewJWTByUserId(u.IssId, attempt.UserID, attempt.ServerID, "")
}
