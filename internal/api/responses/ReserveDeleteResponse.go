package responses

type ReserveDeleteResponse struct {
	BaseJsonRpcResponse
	Result *ReserveDeleteResult `json:"result"`
}

type ReserveDeleteResult struct {
	ReserveDeleteItem []ReserveDeleteItem `json:"items"`
}

type ReserveDeleteItem struct {
	Sku    string `json:"sku"`
	Status string `json:"status"`
}
