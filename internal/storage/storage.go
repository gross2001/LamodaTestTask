package storage

import (
	"context"
	"errors"
	"testTaskLamoda/internal/storage/models"
)

var ErrSkuNotFound = errors.New("skus not found")

//go:generate go run github.com/vektra/mockery/v2@v2.46.0 --name Storage
type Storage interface {
	CreateReserve(ctx context.Context, skus []string) (map[string]models.SkuStoreStatus, error)
	DeleteReserve(ctx context.Context, skus []string) (map[string]models.SkuStoreStatus, error)
	StoreBalance(ctx context.Context, storeId uint) ([]models.SkuStore, error)
}
