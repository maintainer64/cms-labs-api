package auth

import (
	"errors"
	"time"

	"gitlab.com/a10869/api-modules/shared/cms_client"

	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries"
	"gitlab.com/a10869/api-modules/backend/pkg/configs"
	"golang.org/x/crypto/bcrypt"
)

type TokenManager struct {
	UserQueries         *queries.UserQueries
	UserPasswordQueries *queries.UserPasswordQueries
	TokenAttemptQueries *queries.TokenAttemptQueries
}

type RenewManagerInputDTO struct {
	RefreshToken string `json:"refresh_token"`
}

type RenewManagerCredentialsInputDTO struct {
	Email    string `json:"email" required:"true"`
	Password string `json:"password" required:"true"`
}

// NewJWTByCredentials генерирует новый JWT по логину/паролю
func (m *TokenManager) NewJWTByCredentials(email string, password string) (*cms_client.SSOToken, error) {
	invalidCreds := errors.New("username or password is incorrect")
	if email == "" {
		return nil, invalidCreds
	}
	if password == "" {
		return nil, invalidCreds
	}

	entity, err := m.UserQueries.GetByEmail(email)
	if err != nil {
		return nil, invalidCreds
	}
	creds, err := m.UserPasswordQueries.Get(entity.ID)
	if err != nil {
		return nil, invalidCreds
	}
	err = bcrypt.CompareHashAndPassword([]byte(creds.HashPassword), []byte(password))
	if err != nil {
		return nil, invalidCreds
	}
	return m.NewJWTByUserId(entity.ID, 0, "")
}

func (m *TokenManager) NewJWTByLaunchID(launchID string) (*cms_client.SSOToken, error) {
	invalidCreds := errors.New("LTI process is incorrect")
	if launchID == "" {
		return nil, invalidCreds
	}
	entity, err := m.UserQueries.GetByLaunchID(launchID)
	if err != nil {
		return nil, invalidCreds
	}
	return m.NewJWTByUserId(entity.ID, 0, "")
}

func (m *TokenManager) NewJWTByUserId(userId uint, serverID uint, state string) (*cms_client.SSOToken, error) {
	userModel, err := m.UserQueries.Get(userId)
	if err != nil {
		return nil, err
	}
	tokens, err := GenerateNewTokens(&cms_client.SSOTokenPublicData{
		Id:           userModel.ID,
		Email:        userModel.Email,
		ServerID:     serverID,
		Name:         userModel.Name,
		LastLaunchId: userModel.LastLaunchID,
		Role:         userModel.UserRole,
	}, state)
	if err != nil {
		return nil, err
	}
	refreshModel := &models.TokenAttempt{}
	refreshModel.UserID = userId
	refreshModel.ServerID = serverID
	refreshModel.Token = tokens.RefreshToken
	if err = m.TokenAttemptQueries.Upsert(refreshModel); err != nil {
		return nil, err
	}
	return tokens, nil
}

func ExpiresRefreshCookie() time.Time {
	return time.Now().Add(time.Hour * time.Duration(configs.AppConfig.JWT.SecretRefreshExpireHours))
}
