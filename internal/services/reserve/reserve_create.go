package reserve

import (
	"context"
	"testTaskLamoda/internal/api/responses"
)

func (r *Reserve) Create(ctx context.Context, skus []string) (responses.ReserveCreateResult, error) {
	skuStores, err := r.storage.CreateReserve(ctx, skus)
	if err != nil {
		return responses.ReserveCreateResult{}, nil
	}

	var result responses.ReserveCreateResult
	for sku, skuStore := range skuStores {
		item := responses.ReserveCreateItem{Sku: sku, Status: skuStore.Status}
		result.ReserveCreateItems = append(result.ReserveCreateItems, item)
	}
	return result, nil
}
