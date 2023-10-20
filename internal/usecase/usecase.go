package usecase

import (
	"context"
	"github.com/kalunik/urShorty/internal/entity"
	"github.com/kalunik/urShorty/internal/repository"
	"github.com/kalunik/urShorty/pkg/logger"
)

type Usecase interface {
	AddUrlPair(ctx context.Context, pair *entity.UrlPair) error
	FindFullUrl(ctx context.Context, shortUrl string) (string, error)
}

type UrlPairUsecase struct {
	repo repository.Repository
	log  logger.Loger
}

func NewUrlPairUsecase(repository repository.Repository, loger logger.Loger) Usecase {
	return &UrlPairUsecase{
		repo: repository,
		log:  loger,
	}
}

func (u *UrlPairUsecase) AddUrlPair(ctx context.Context, pair *entity.UrlPair) error {
	if err := u.repo.AddUrlPair(ctx, pair); err != nil {
		return err
	}
	return nil
}

func (u *UrlPairUsecase) FindFullUrl(ctx context.Context, shortUrl string) (string, error) {
	fullUrl, err := u.repo.FindFullUrl(ctx, shortUrl)
	if err != nil {
		return "", err
	}
	return fullUrl, nil
}
