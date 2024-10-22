package reserve

import (
	"context"
	"fmt"
	"testTaskLamoda/internal/api/responses"
)

func (r *Reserve) Delete(ctx context.Context, skus []string) (responses.ReserveDeleteResult, error) {
	skuStores, err := r.storage.DeleteReserve(ctx, skus)

	fmt.Printf("skuStores %v\n", skuStores)
	if err != nil {
		return responses.ReserveDeleteResult{}, nil
	}
	var result responses.ReserveDeleteResult
	for sku, skuStore := range skuStores {
		item := responses.ReserveDeleteItem{Sku: sku, Status: skuStore.Status}
		result.ReserveDeleteItem = append(result.ReserveDeleteItem, item)
	}
	return result, nil
}
