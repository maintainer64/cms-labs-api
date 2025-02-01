package auth

import (
	"errors"
	"fmt"
	"strings"
	"time"

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
}

func (u *SSOTokenUC) Execute(inputDTO SSOTokenInputDTO) (*SSOToken, error) {
	if err := u.validate(inputDTO); err != nil {
		return &SSOToken{}, err
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

func (u *SSOTokenUC) ByAuthCode(inputDTO SSOTokenInputDTO) (*SSOToken, error) {
	attempt, err := u.TokenAttemptQueries.GetByAuthCode(inputDTO.Code)
	if err != nil {
		return nil, err
	}
	server, err := u.PNETServerQueries.Get(attempt.ServerID)
	if err != nil {
		return nil, err
	}
	log.Info().Msg(fmt.Sprintf("Token get by auth code by server_id: %+v", server.ID))
	if !server.IsActive {
		return nil, errors.New("server is not active")
	}
	if !strings.HasPrefix(inputDTO.RedirectUri, server.Url) {
		return nil, errors.New("invalid redirect_uri")
	}
	return u.TokenManager.NewJWTByUserId(attempt.UserID, attempt.ServerID, attempt.State)
}

func (u *SSOTokenUC) ByRefresh(inputDTO SSOTokenInputDTO) (*SSOToken, error) {
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
	log.Info().Msg(fmt.Sprintf("Token get by refresh token by server_id: %+v", attempt.ServerID))
	if attempt.ServerID == 0 {
		return u.TokenManager.NewJWTByUserId(attempt.UserID, attempt.ServerID, "")
	}
	server, err := u.PNETServerQueries.Get(attempt.ServerID)
	if err != nil {
		return nil, err
	}
	if !server.IsActive {
		return nil, errors.New("server is not active")
	}
	return u.TokenManager.NewJWTByUserId(attempt.UserID, attempt.ServerID, "")
}
