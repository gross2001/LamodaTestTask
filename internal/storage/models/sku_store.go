package models

type SkuStore struct {
	Id             int32
	Sku            string
	Store_id       int32
	Total_quantity int32
	Reserved       int32
}

type SkuStoreStatus struct {
	SkuStoreInfo SkuStore
	Status       string
}
