package auth

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

type SSOTokenIntrospect struct {
	Active    bool     `json:"active"`
	ClientID  string   `json:"client_id"`
	Username  string   `json:"username"`
	Scope     []string `json:"scope"`
	TokenType string   `json:"token_type"`
	Exp       int64    `json:"exp"`
	Iat       int64    `json:"iat"`
	Sub       string   `json:"sub"`
}

type SSOTokenIntrospectResponse struct {
	Error  bool               `json:"error" validate:"required"`
	Msg    string             `json:"msg" validate:"required"`
	Result SSOTokenIntrospect `json:"result"`
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

type SsoTokenPublicDataResponse struct {
	Error  bool               `json:"error" validate:"required"`
	Msg    string             `json:"msg" validate:"required"`
	Result SSOTokenPublicData `json:"result"`
}
