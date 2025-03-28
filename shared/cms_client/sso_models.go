package cms_client

import (
	"encoding/base64"

	"github.com/goccy/go-json"
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

// SSOTokenPublicData struct to describe public payload object.
type SSOTokenPublicData struct {
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

func (t *SSOTokenPublicData) JWTClaims() jwt.MapClaims {
	return jwt.MapClaims{
		"iss":            t.Iss,
		"sub":            t.Sub,
		"aud":            t.Aud,
		"exp":            t.Exp,
		"iat":            t.Iat,
		"nonce":          t.Nonce,
		"email":          t.Email,
		"name":           t.Name,
		"server_id":      t.ServerID,
		"role":           t.Role,
		"last_launch_id": t.LastLaunchId,
	}
}

const (
	PNETLabsTypeDefault     = "default"
	PNETLabsTypeEnumeration = "enumeration"
	PNETLabsTypeFile        = "file"
)

type SSOTokenPublicExtraParams struct {
	AttemptID string `json:"attempt_id"`
	// enumeration PNETLabsTypeDefault...
	PNETLabsType string `json:"pnet_labs_type"`
	PNETLabsPath string `json:"pnet_labs_path"`
	PNETTestPath string `json:"pnet_test_path"`
}

func (t *SSOTokenPublicExtraParams) Marshal() string {
	b, err := json.Marshal(t)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(b)
}

func UnmarshalSSOTokenPublicExtraParams(base64String string) *SSOTokenPublicExtraParams {
	var decode SSOTokenPublicExtraParams
	data, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return &decode
	}
	_ = json.Unmarshal(data, &decode)
	return &decode
}
