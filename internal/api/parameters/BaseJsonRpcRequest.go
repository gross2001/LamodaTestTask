package parameters

type BaseJsonRpcRequest struct {
	Jsonrpc string `json:"jsonrpc" validate:"required"`
	Id      string `json:"id" validate:"required"`
	Method  string `json:"method" validate:"required"`
}
