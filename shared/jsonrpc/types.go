package jsonrpc

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type Map map[string]interface{}

type Ctx struct {
	FiberCtx *fiber.Ctx
	Params   json.RawMessage
}

type Handler = func(*Ctx) (interface{}, error)

type RequestJSONRPC struct {
	JSONRPC string          `json:"jsonrpc"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params"`
	ID      interface{}     `json:"id"`
}

type RpcError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (c RpcError) Error() string {
	return c.Message
}

func NewRpcError(code, message string) RpcError {
	return RpcError{Code: code, Message: message}
}

const (
	ParseError     = -32700
	InvalidRequest = -32600
	MethodNotFound = -32601
	InvalidParams  = -32602
	InternalError  = -32603
	RPCLogicError  = -32010 // -32000..32099
)
