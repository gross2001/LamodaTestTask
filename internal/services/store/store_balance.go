package store

import (
	"context"
	"fmt"
	"testTaskLamoda/internal/api/responses"
)

func (s *Store) Balance(ctx context.Context, storeId int) (responses.StoreBalanceResult, error) {
	// TODO: type assertion, change?
	fmt.Println("HELLO IN BALANCE")
	skuStores, err := s.storage.StoreBalance(ctx, uint(storeId))
	fmt.Printf("skuStores %v\n", skuStores)

	if err != nil {
		return responses.StoreBalanceResult{}, nil
	}
	var result responses.StoreBalanceResult
	for _, skuStore := range skuStores {
		item := responses.SkuBalanceItem{
			Sku:         skuStore.Sku,
			TotalQnt:    skuStore.Total_quantity,
			ReservedQnt: skuStore.Reserved,
		}
		fmt.Printf("item %v\n", item)
		result.SkuBalanceItems = append(result.SkuBalanceItems, item)
	}
	return result, nil
}
