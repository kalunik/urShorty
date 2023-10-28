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
	GetRepo() repository.Repository
}

type UrlPairUsecase struct {
	repo repository.Repository
	log  logger.Logger
}

func NewUrlPairUsecase(repository repository.Repository, logger logger.Logger) Usecase {
	return &UrlPairUsecase{
		repo: repository,
		log:  logger,
	}
}

func (u *UrlPairUsecase) AddUrlPair(ctx context.Context, pair *entity.UrlPair) error {
	ok, err := u.repo.AddUrlPair(ctx, pair)
	if ok {
		u.log.Infof("New urlPair '%s' added to redis", pair.Short)
		return nil
	}
	if err != nil {
		return err
	}
	return nil
}

func (u *UrlPairUsecase) FindFullUrl(ctx context.Context, shortUrl string) (string, error) {
	fullUrl, err := u.repo.GetFullUrl(ctx, shortUrl)
	if err != nil {
		return "", err
	}
	return fullUrl, nil
}

func (u *UrlPairUsecase) GetRepo() repository.Repository {
	return u.repo
}
