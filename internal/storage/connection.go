package storage

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func NewConnectionPool(ctx context.Context, connString string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		return pool, fmt.Errorf("failed to create pgxpool: %w", err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		return pool, fmt.Errorf("failed to ping postgres: %w", err)
	}

	slog.Info("Connected postgres to database")

	return pool, nil
}

func NewRedisClient(ctx context.Context, connString string) (*redis.Client, error) {
	opt, err := redis.ParseURL(connString)
	if err != nil {
		return nil, fmt.Errorf("failed to parse redis url: %w", err)
	}

	rdb := redis.NewClient(opt)

	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to ping redis: %w", err)
	}

	slog.Info("Connected redis to database")

	return rdb, nil
}
