package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/ivasilkov/pow-alg-go/internal/config"
	"github.com/ivasilkov/pow-alg-go/internal/handler/server/get_quote"
	"github.com/ivasilkov/pow-alg-go/internal/pow/hashcash"
	"github.com/ivasilkov/pow-alg-go/internal/storage/inmem"
	"github.com/ivasilkov/pow-alg-go/internal/storage/redis"
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
		return
	}

	log.Info("server stopped")
}

func run(ctx context.Context, log *slog.Logger) error {
	cfg := config.MustNewServerCfg()

	redisStorage, err := redis.New(ctx, cfg.GetRedisAddr())
	if err != nil {
		return fmt.Errorf("redis creation error: %w", err)
	}
	inMemStorage := inmem.NewStorage()
	powAlg := hashcash.New()

	h := get_quote.NewHandler(log, inMemStorage, redisStorage, powAlg)

	srv, err := server.NewServer(cfg.GetAddr(), log, h)
	if err != nil {
		return fmt.Errorf("server creation error: %w", err)
	}
	defer srv.Close() //nolint:errcheck //close fn

	if err := srv.Run(ctx); err != nil {
		return fmt.Errorf("server running error: %w", err)
	}

	return nil
}
