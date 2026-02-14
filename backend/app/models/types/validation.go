package types

type UserStoreSetRequest struct {
	JSONRPC string    `json:"jsonrpc" default:"2.0" required:"true"`
	Method  string    `json:"method" default:"user.global_store_set" required:"true"`
	Params  JsonStore `json:"params,omitempty"`
	ID      string    `json:"id,omitempty" default:"1" required:"true"`
}

type UserStoreSetResponse struct {
	JSONRPC string      `json:"jsonrpc" default:"2.0" required:"true"`
	Result  bool        `json:"result,omitempty"`
	Error   interface{} `json:"error,omitempty"`
	ID      string      `json:"id,omitempty" default:"1" required:"true"`
}

type UserStoreGetRequest struct {
	JSONRPC string       `json:"jsonrpc" default:"2.0" required:"true"`
	Method  string       `json:"method" default:"user.global_store_get" required:"true"`
	Params  *interface{} `json:"params,omitempty"`
	ID      string       `json:"id,omitempty" default:"1" required:"true"`
}

type UserStoreGetResponse struct {
	JSONRPC string      `json:"jsonrpc" default:"2.0" required:"true"`
	Result  JsonStore   `json:"result,omitempty"`
	Error   interface{} `json:"error,omitempty"`
	ID      string      `json:"id,omitempty" default:"1" required:"true"`
}
