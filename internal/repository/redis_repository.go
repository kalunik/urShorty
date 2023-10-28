package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/kalunik/urShorty/internal/entity"
	"github.com/redis/go-redis/v9"
)

type Repository interface {
	AddUrlPair(ctx context.Context, pair *entity.UrlPair) (bool, error)
	GetFullUrl(ctx context.Context, shortUrl string) (string, error)
}

type repo struct {
	client *redis.Client
}

func NewRedisRepository(client *redis.Client) Repository {
	return &repo{
		client: client,
	}
}

func (r *repo) AddUrlPair(ctx context.Context, pair *entity.UrlPair) (bool, error) {
	exists, err := r.client.Exists(ctx, pair.Short).Result()
	if err != nil {
		return false, fmt.Errorf("redis client: EXISTS failure: %w", err)
	}

	if exists == 0 {
		if err := r.client.Set(ctx, pair.Short, pair.Full, 0).Err(); err != nil {
			return false, fmt.Errorf("redis client: SET failure: %w", err)
		}
		return true, nil
	}
	return false, nil
}

func (r *repo) GetFullUrl(ctx context.Context, shortUrl string) (string, error) {
	fullUrl, err := r.client.Get(ctx, shortUrl).Result()
	if err == redis.Nil {
		return "", errors.New("redis client: url not exist")
	} else if err != nil {
		return "", fmt.Errorf("redis client: GET failure: %w", err)
	}
	return fullUrl, nil
}
