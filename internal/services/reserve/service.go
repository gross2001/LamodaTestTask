package reserve

import (
	"context"
	"testTaskLamoda/internal/api/responses"
	"testTaskLamoda/internal/storage"
)

type Service interface {
	Create(ctx context.Context, skus []string) (responses.ReserveCreateResult, error)
	Delete(ctx context.Context, skus []string) (responses.ReserveDeleteResult, error)
}

type Config struct {
	Storage storage.Storage
}

type Reserve struct {
	storage storage.Storage
}

func New(c Config) Service {
	return &Reserve{
		storage: c.Storage,
	}
}
