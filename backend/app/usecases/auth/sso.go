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

type SSOError struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}
