package parameters

type ReserveDeleteRequest struct {
	BaseJsonRpcRequest
	Params ReserveDeleteParams `json:"params"`
}

type ReserveDeleteParams struct {
	Skus []string `json:"skus"`
}
