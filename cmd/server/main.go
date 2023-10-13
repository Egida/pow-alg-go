package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/ivasilkov/pow-alg-go/internal/config"
	"github.com/ivasilkov/pow-alg-go/internal/handler/get_quote"
	"github.com/ivasilkov/pow-alg-go/internal/storage/inmem"
	"github.com/ivasilkov/pow-alg-go/internal/transport/tcp/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	log := slog.New(slog.NewTextHandler(os.Stdout, nil))

	defer func() {
		if panicErr := recover(); panicErr != nil {
			log.Error("panic occurred", "err", panicErr)
			os.Exit(1)
		}
		os.Exit(0)
	}()

	if err := run(ctx, log); err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	log.Info("server stopped")
	os.Exit(0)
}

func run(ctx context.Context, log *slog.Logger) error {
	cfg := config.MustNewServerCfg()

	h := get_quote.NewHandler(log, inmem.NewStorage())

	srv, err := server.NewServer(cfg.Addr, log, h)
	if err != nil {
		return fmt.Errorf("server creation error: %w", err)
	}
	defer srv.Close()

	if err := srv.Run(ctx); err != nil {
		return fmt.Errorf("server running error: %w", err)
	}

	return nil
}
