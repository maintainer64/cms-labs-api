package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"gitlab.com/a10869/api-modules/shared/cms_client"

	"gitlab.com/a10869/api-modules/backend/pkg/configs"

	"github.com/golang-jwt/jwt/v5"
)

type TokenDataWithExp struct {
	Token string `json:"token"`
	Exp   int64  `json:"exp"`
}

// GenerateNewTokens func for generate a new Access & Refresh tokens.
func GenerateNewTokens(entity *cms_client.SSOTokenPublicData, state string) (*cms_client.SSOToken, error) {
	// Generate JWT Access token.
	accessToken, err := generateNewAccessToken(entity)
	if err != nil {
		// Return token generation error.
		return nil, err
	}

	// Generate JWT Refresh token.
	refreshToken, err := generateNewRefreshToken()
	if err != nil {
		// Return token generation error.
		return nil, err
	}

	return &cms_client.SSOToken{
		AccessToken:  accessToken.Token,
		RefreshToken: refreshToken.Token,
		TokenType:    "bearer",
		ExpiresIn:    accessToken.Exp,
		State:        state,
		UserId:       entity.Sub,
	}, nil
}

func generateNewAccessToken(entity *cms_client.SSOTokenPublicData) (*TokenDataWithExp, error) {
	// Set secret key from .env file.
	secret := configs.AppConfig.JWT.SecretKey

	// Set expires minutes count for secret key from .env file.
	minutesCount := configs.AppConfig.JWT.SecretKeyExpireMinutes

	// Set public claims:
	entity.Iat = time.Now().Unix()
	entity.Exp = time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()

	// Create a new claims.
	claims := entity.JWTClaims()

	// Create a new JWT access token with claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate token.
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		// Return error, it JWT token generation failed.
		return nil, err
	}

	return &TokenDataWithExp{
		Token: t,
		Exp:   entity.Exp,
	}, nil
}

func generateNewRefreshToken() (*TokenDataWithExp, error) {
	// Create a new SHA256 hash.
	hash := sha256.New()

	// Create a new now date and time string with salt.
	refresh := configs.AppConfig.JWT.SecretRefresh + time.Now().String()

	// See: https://pkg.go.dev/io#Writer.Write
	_, err := hash.Write([]byte(refresh))
	if err != nil {
		// Return error, it refresh token generation failed.
		return nil, err
	}

	// Set expires hours count for refresh key from .env file.
	hoursCount := configs.AppConfig.JWT.SecretRefreshExpireHours

	// Set expiration time.
	expireTime := time.Now().Add(time.Hour * time.Duration(hoursCount)).Unix()

	// Create a new refresh token (sha256 string with salt + expire time).
	t := hex.EncodeToString(hash.Sum(nil)) + "." + fmt.Sprint(expireTime)

	return &TokenDataWithExp{
		Token: t,
		Exp:   expireTime,
	}, nil
}

// ParseRefreshToken func for parse second argument from refresh token.
func ParseRefreshToken(refreshToken string) (int64, error) {
	splitToken := strings.Split(refreshToken, ".")
	if len(splitToken) < 2 {
		return 0, errors.New("invalid refresh token")
	}
	return strconv.ParseInt(splitToken[1], 0, 64)
}
