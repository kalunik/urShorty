package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/kalunik/urShorty/internal/entity"
	"github.com/redis/go-redis/v9"
)

type RedisRepository interface {
	AddPathUrl(ctx context.Context, pair entity.PathMeta) error
	GetFullUrl(ctx context.Context, shortUrl string) (string, error)
	IsExist(ctx context.Context, shortUrl string) (int64, error)
}

type redisRepo struct {
	client *redis.Client
}

func NewRedisRepository(client *redis.Client) RedisRepository {
	return &redisRepo{
		client: client,
	}
}

func (r *redisRepo) AddPathUrl(ctx context.Context, pair entity.PathMeta) error {
	if err := r.client.Set(ctx, pair.ShortPath, pair.FullUrl, 0).Err(); err != nil {
		return fmt.Errorf("redis client: SET failure: %w", err)
	}
	return nil
}

func (r *redisRepo) GetFullUrl(ctx context.Context, shortUrl string) (string, error) {
	fullUrl, err := r.client.Get(ctx, shortUrl).Result()
	if err == redis.Nil {
		return "", errors.New("redis client: url not exist")
	} else if err != nil {
		return "", fmt.Errorf("redis client: GET failure: %w", err)
	}
	return fullUrl, nil
}

func (r *redisRepo) IsExist(ctx context.Context, shortUrl string) (int64, error) {
	exist, err := r.client.Exists(ctx, shortUrl).Result()
	if err != nil {
		return exist, fmt.Errorf("redis client: EXISTS failure: %w", err)
	}
	return exist, nil
}
