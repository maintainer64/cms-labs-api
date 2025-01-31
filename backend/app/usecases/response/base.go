package response

type Response[T any] struct {
	Error  bool   `json:"error" validate:"required"`
	Msg    string `json:"msg" validate:"required"`
	Result T      `json:"result"`
}
