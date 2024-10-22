package responses

type ReserveCreateResponse struct {
	BaseJsonRpcResponse
	Result *ReserveCreateResult `json:"result"`
}

type ReserveCreateResult struct {
	ReserveCreateItems []ReserveCreateItem `json:"items"`
}

type ReserveCreateItem struct {
	Sku    string `json:"sku"`
	Status string `json:"status"`
}
