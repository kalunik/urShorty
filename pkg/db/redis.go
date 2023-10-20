package db

import (
	"context"
	"fmt"
	"github.com/kalunik/urShorty/config"
	"github.com/redis/go-redis/v9"
	"time"
)

type repo struct {
	client *redis.Client
}

func NewRedisConnection(cfg *config.AppConfig) (*redis.Client, error) {
	redisAddr := cfg.Redis.RedisAddr
	if redisAddr == "" {
		redisAddr = ":6379"
	}

	client := redis.NewClient(&redis.Options{
		Addr:         redisAddr,
		MinIdleConns: cfg.Redis.MinIdleConns,
		PoolSize:     cfg.Redis.PoolSize,
		PoolTimeout:  time.Duration(cfg.Redis.PoolTimeout) * time.Second,
		Password:     cfg.Redis.Password,
		DB:           cfg.Redis.DB,
	})

	status := client.Ping(context.Background())
	if err := status.Err(); err != nil {
		return nil, fmt.Errorf("redis: connection couldn't be established: %w", err)
	}
	return client, nil
}
