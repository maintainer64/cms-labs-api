package auth

import (
	"errors"
	"fmt"
	"strings"

	"github.com/rs/zerolog"

	"github.com/google/uuid"
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries"
	"gitlab.com/a10869/api-modules/backend/app/usecases/response"
)

type SSOAuthorizeInputDTO struct {
	UserID       uint   `json:"user_id"`
	ClientID     string `json:"client_id"`
	RedirectUri  string `json:"redirect_uri"`
	ResponseType string `json:"response_type"`
	Scope        string `json:"scope"`
	Path         string `json:"path"`
	State        string `json:"state"`
	Extra        string `json:"extra"`
}

type SSOAuthorizeOutputDTO struct {
	RedirectUri string `json:"redirect_uri"`
	Code        string `json:"code"`
	Application string `json:"application"`
	Path        string `json:"path"`
	State       string `json:"state"`
	Extra       string `json:"extra"`
}

type SSOAuthorizeResponse = response.Response[SSOAuthorizeOutputDTO]

const (
	SSOAuthorizeResponseType = "code"
)

type SSOAuthorizeUC struct {
	TokenAttemptQueries *queries.TokenAttemptQueries
	PNETServerQueries   *queries.PNETServerQueries
	*zerolog.Logger
}

func (u *SSOAuthorizeUC) Execute(inputDTO SSOAuthorizeInputDTO) (*SSOAuthorizeOutputDTO, error) {
	if err := u.validate(inputDTO); err != nil {
		return nil, err
	}
	u.Logger.Info().Msg(fmt.Sprintf("authorize sso authorize with clientID %s and userID %v", inputDTO.ClientID, inputDTO.UserID))
	server, err := u.PNETServerQueries.GetByClientId(inputDTO.ClientID)
	if err != nil {
		return nil, err
	}
	if !server.IsActive {
		u.Logger.Info().Msg(fmt.Sprintf("server is not active sso authorize with clientID %s and userID %v", inputDTO.ClientID, inputDTO.UserID))
		return nil, errors.New("server is not active")
	}
	if !strings.HasPrefix(inputDTO.RedirectUri, server.Url) {
		u.Logger.Info().Msg(fmt.Sprintf("invalid redirect_uri sso authorize with clientID %s and userID %v", inputDTO.ClientID, inputDTO.UserID))
		return nil, errors.New("invalid redirect_uri")
	}
	entityCreate := &models.TokenAttempt{}
	entityCreate.UserID = inputDTO.UserID
	entityCreate.ServerID = server.ID
	entityCreate.Token = ""
	entityCreate.State = inputDTO.State
	entityCreate.AuthorizationCode = uuid.New().String()
	err = u.TokenAttemptQueries.Upsert(entityCreate)
	if err != nil {
		return nil, err
	}
	u.Logger.Info().Msg(fmt.Sprintf("response second factor sso authorize with clientID %s and userID %v", inputDTO.ClientID, inputDTO.UserID))
	return &SSOAuthorizeOutputDTO{
		RedirectUri: inputDTO.RedirectUri,
		Code:        entityCreate.AuthorizationCode,
		Application: inputDTO.ClientID,
		Path:        inputDTO.Path,
		State:       inputDTO.State,
		Extra:       inputDTO.Extra,
	}, nil

}

func (u *SSOAuthorizeUC) validate(inputDTO SSOAuthorizeInputDTO) error {
	if inputDTO.ResponseType == SSOAuthorizeResponseType && inputDTO.ClientID != "" && inputDTO.RedirectUri != "" && inputDTO.Scope != "" {
		return nil
	}
	return errors.New("Invalid params response_type")
}
