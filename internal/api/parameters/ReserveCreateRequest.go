package parameters

type ReserveCreateRequest struct {
	BaseJsonRpcRequest
	Params ReserveCreateParams `json:"params"`
}

type ReserveCreateParams struct {
	Skus []string `json:"skus"`
}
