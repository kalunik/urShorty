package repository

import (
	"context"
	"github.com/kalunik/urShorty/internal/entity"
	"github.com/redis/go-redis/v9"
)

type Repository interface {
	AddUrlPair(ctx context.Context, pair *entity.UrlPair) error
	FindFullUrl(ctx context.Context, shortUrl string) (string, error)
}

type repo struct {
	client *redis.Client
}

func NewRedisRepository(client *redis.Client) Repository {
	return &repo{
		client: client,
	}
}

func (r *repo) AddUrlPair(ctx context.Context, pair *entity.UrlPair) error {

	return nil
}

func (r *repo) FindFullUrl(ctx context.Context, shortUrl string) (string, error) {
	return "", nil
}
