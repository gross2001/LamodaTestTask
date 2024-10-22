package store

import (
	"context"
	"testTaskLamoda/internal/api/responses"
	"testTaskLamoda/internal/storage"
)

type Service interface {
	Balance(ctx context.Context, storeId int) (responses.StoreBalanceResult, error)
}

type Config struct {
	Storage storage.Storage
}

type Store struct {
	storage storage.Storage
}

func New(c Config) Service {
	return &Store{
		storage: c.Storage,
	}
}
