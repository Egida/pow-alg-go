package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	cl *redis.Client
}

func New(ctx context.Context, addr string) (*Redis, error) {
	cl := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	err := cl.Ping(ctx).Err()
	if err != nil {
		return nil, err
	}

	return &Redis{cl: cl}, nil
}

func (r *Redis) AddUnique(ctx context.Context, key, value string) error {
	cmd := r.cl.SetNX(ctx, key, value, 2*24*time.Hour)
	if cmd.Err() != nil {
		return fmt.Errorf("failed to SetNX: %w", cmd.Err())
	}
	if !cmd.Val() {
		return ErrAlreadyExists
	}
	return nil
}
