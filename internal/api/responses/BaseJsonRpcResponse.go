package responses

type BaseJsonRpcResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      string `json:"id"`
	Error   *Error `json:"error,omitempty"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}
