package api

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"testTaskLamoda/internal/services"
	"testTaskLamoda/internal/storage"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type ServerAPI struct {
	Server   *http.Server
	Services services.Services
	log      *slog.Logger
	port     uint
}

func NewServer(ctx context.Context, log *slog.Logger, storage storage.Storage, port uint) *ServerAPI {

	r := chi.NewRouter()

	server := &http.Server{
		Addr:              fmt.Sprintf("0.0.0.0:%d", port),
		Handler:           r,
		ReadTimeout:       1 * time.Second,
		WriteTimeout:      1 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
		BaseContext: func(_ net.Listener) context.Context {
			return ctx
		},
	}

	services, _ := services.NewServices(storage)

	serverApi := ServerAPI{
		Server:   server,
		Services: *services,
		port:     port,
		log:      log,
	}

	r.Use(middleware.RequestID)

	r.Post("/reserve/create", serverApi.CreateReserve())
	r.Post("/reserve/delete", serverApi.DeleteReserve())
	r.Post("/store/balance", serverApi.StoreBalace())

	return &serverApi
}
