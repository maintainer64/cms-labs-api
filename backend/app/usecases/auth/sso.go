package auth

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
