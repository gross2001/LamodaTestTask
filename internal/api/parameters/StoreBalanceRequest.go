package parameters

type StoreBalanceRequest struct {
	BaseJsonRpcRequest
	Params StoreBalanceParams `json:"params"`
}

type StoreBalanceParams struct {
	Store_Id int `json:"store_id"`
}
