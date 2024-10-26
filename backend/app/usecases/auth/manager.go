package auth

import (
	"errors"
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries"
	"gitlab.com/a10869/api-modules/backend/pkg/configs"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type TokenManager struct {
	UserQueries      *queries.UserQueries
	UserTokenQueries *queries.UserTokenQueries
}

type RenewManagerInputDTO struct {
	RefreshToken string `json:"refresh_token"`
}

type RenewManagerCredentialsInputDTO struct {
	Email    string `json:"email" required:"true"`
	Password string `json:"password" required:"true"`
}

type RenewManagerResponse struct {
	Error  bool   `json:"error" validate:"required"`
	Msg    string `json:"msg" validate:"required"`
	Result Tokens `json:"result"`
}

type RenewManagerUserResponse struct {
	Error  bool            `json:"error" validate:"required"`
	Msg    string          `json:"msg" validate:"required"`
	Result TokenPublicData `json:"result"`
}

// NewJWTByUserId генерирует новый JWT по пользователю
func (m *TokenManager) NewJWTByUserId(
	userId uint,
) (*Tokens, error) {
	return m.newJWTByUserId(userId)
}

// NewJWTByRefreshToken генерирует новый JWT по Refresh token'у
func (m *TokenManager) NewJWTByRefreshToken(
	refreshToken string,
) (*Tokens, error) {
	expiresRefreshToken, err := ParseRefreshToken(refreshToken)
	if err != nil {
		return nil, err
	}
	now := time.Now().Unix()
	if now >= expiresRefreshToken {
		return nil, errors.New("refresh token is expired")
	}
	refreshTokenEntity, err := m.UserTokenQueries.GetByRefreshToken(refreshToken)
	if err != nil {
		return nil, err
	}
	return m.newJWTByUserId(refreshTokenEntity.UserID)
}

// NewJWTByCredentials генерирует новый JWT по логину/паролю
func (m *TokenManager) NewJWTByCredentials(email string, password string) (*Tokens, error) {
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
	creds, err := m.UserTokenQueries.Get(entity.ID)
	if err != nil {
		return nil, invalidCreds
	}
	err = bcrypt.CompareHashAndPassword([]byte(creds.HashPassword), []byte(password))
	if err != nil {
		return nil, invalidCreds
	}
	return m.newJWTByUserId(entity.ID)
}

// NewJWTByLaunchID генерирует новый JWT по входу LTI
func (m *TokenManager) NewJWTByLaunchID(launchID string) (*Tokens, error) {
	entity, err := m.UserQueries.GetByLaunchID(launchID)
	if err != nil {
		return nil, err
	}
	return m.newJWTByUserId(entity.ID)
}

func (m *TokenManager) newJWTByUserId(userId uint) (*Tokens, error) {
	userModel, err := m.UserQueries.Get(userId)
	if err != nil {
		return nil, err
	}
	tokens, err := GenerateNewTokens(&TokenPublicData{
		Id:    userModel.ID,
		Email: userModel.Email,
		Name:  userModel.Name,
		Role:  userModel.UserRole,
	})
	if err != nil {
		return nil, err
	}
	if err = m.UserTokenQueries.Upsert(
		&models.UserToken{
			UserID:       userModel.ID,
			RefreshToken: tokens.Refresh.Token,
		},
	); err != nil {
		return nil, err
	}
	return tokens, nil
}

func ExpiresRefreshCookie() time.Time {
	return time.Now().Add(time.Hour * time.Duration(configs.AppConfig.JWT.SecretRefreshExpireHours))
}
