package responses

type StoreBalanceResponse struct {
	BaseJsonRpcResponse
	Result *StoreBalanceResult `json:"result"`
}

type StoreBalanceResult struct {
	SkuBalanceItems []SkuBalanceItem `json:"items"`
}

type SkuBalanceItem struct {
	Sku         string `json:"sku"`
	TotalQnt    int32  `json:"total_qnt"`
	ReservedQnt int32  `json:"reserved_qnt"`
}
