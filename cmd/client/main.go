package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/ivasilkov/pow-alg-go/internal/config"
	"github.com/ivasilkov/pow-alg-go/internal/handler/client/get_quote"
	"github.com/ivasilkov/pow-alg-go/internal/pow/hashcash"
	"github.com/ivasilkov/pow-alg-go/internal/transport/tcp/client"
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

	log.Info("client stopped")
	os.Exit(0)
}

func run(ctx context.Context, log *slog.Logger) error {
	cfg := config.MustNewClientCfg()

	h := get_quote.NewHandler(log, hashcash.New())

	cl, err := client.NewClient(cfg.GetServerAddr(), log, h)
	if err != nil {
		return fmt.Errorf("server creation error: %w", err)
	}
	defer cl.Close()

	if err := cl.Run(ctx); err != nil {
		return fmt.Errorf("server running error: %w", err)
	}

	return nil
}
