package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"testTaskLamoda/internal/api"
	"testTaskLamoda/internal/config"
	"testTaskLamoda/internal/storage/postgres"

	"golang.org/x/sync/errgroup"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustReadEnv()

	log := setupLogger(cfg.Env)

	storage, err := postgres.New(cfg.Service_DB_DSN)
	if err != nil {
		log.Error("Failed to initialize storage", "error:", err)
		os.Exit(1)
	}

	mainCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	server := api.NewServer(mainCtx, log, storage, cfg.APP_ADDR)

	g, gCtx := errgroup.WithContext(mainCtx)
	g.Go(func() error {
		log.Info("server started...")
		return server.Server.ListenAndServe()
	})
	g.Go(func() error {
		<-gCtx.Done()
		log.Info("server shutdown started...")
		return server.Server.Shutdown(context.Background())
	})

	if err := g.Wait(); err != nil {
		log.Error("server stoped", "reason", err)
	}
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
