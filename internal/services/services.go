package services

import (
	"testTaskLamoda/internal/services/reserve"
	"testTaskLamoda/internal/services/store"
	"testTaskLamoda/internal/storage"
)

type Services struct {
	Reserve reserve.Service
	Store   store.Service
}

func NewServices(storage storage.Storage) (*Services, error) {
	services := &Services{}
	services.Reserve = reserve.New(reserve.Config{Storage: storage})
	services.Store = store.New(store.Config{Storage: storage})
	return services, nil
}
