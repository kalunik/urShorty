package usecase

import (
	"context"
	"errors"
	"github.com/kalunik/urShorty/internal/entity"
	"github.com/kalunik/urShorty/internal/repository"
	"github.com/kalunik/urShorty/internal/usecase/webapi"
	"github.com/kalunik/urShorty/pkg/logger"
	"github.com/kalunik/urShorty/pkg/utils"
	"net/url"
	"time"
)

type Usecase interface {
	AddUrlPath(ctx context.Context, meta *entity.PathMeta) error
	GetFullUrl(ctx context.Context, shortPath string, visitorIP string) (string, error)
	PathVisits(ctx context.Context, shortPath string) (*entity.PathVisitsList, error)
	IsExists(ctx context.Context, shortPath string) (bool, error)
	GetRepo() repository.RedisRepository
}

type PathMetaUsecase struct {
	redisRepo      repository.RedisRepository
	clickhouseRepo repository.ClickhouseRepository
	geoIP          webapi.IPGeolocation
	log            logger.Logger
}

func NewPathMetaUsecase(redis repository.RedisRepository, clickhouse repository.ClickhouseRepository, logger logger.Logger) Usecase {
	return &PathMetaUsecase{
		redisRepo:      redis,
		clickhouseRepo: clickhouse,
		geoIP:          webapi.NewIPGeoWebAPI(),
		log:            logger,
	}
}

func (u *PathMetaUsecase) AddUrlPath(ctx context.Context, meta *entity.PathMeta) error {
	var err error
	meta.ShortPath, err = utils.GenerateHash(meta.FullUrl)
	if err != nil {
		return errors.New("addPair: generate hash fail")
	}
	parsedUrl, _ := url.Parse(meta.FullUrl)
	meta.Domain = parsedUrl.Host
	meta.CreatedAt = time.Now()

	if err := u.redisRepo.AddPathUrl(ctx, *meta); err != nil {
		return err
	}

	if err := u.clickhouseRepo.AddNewShortPath(ctx, *meta); err != nil {
		return err
	}

	return nil
}

func (u *PathMetaUsecase) GetFullUrl(ctx context.Context, shortUrl string, visitorIP string) (string, error) {
	fullUrl, err := u.redisRepo.GetFullUrl(ctx, shortUrl)
	if err != nil {
		return "", err
	}

	location, err := u.geoIP.GetIPLocation(visitorIP)
	if err != nil {
		return "", err
	}

	meta := entity.PathMeta{
		ShortPath: shortUrl,
		VisitedAt: time.Now(),
		Country:   location.Country,
		City:      location.City,
		Proxy:     location.Proxy,
	}

	err = u.clickhouseRepo.AddPathVisit(ctx, meta)
	if err != nil {
		return fullUrl, err
	}

	return fullUrl, nil
}

func (u *PathMetaUsecase) PathVisits(ctx context.Context, shortPath string) (*entity.PathVisitsList, error) {
	visits, err := u.clickhouseRepo.PathVisits(ctx, shortPath)
	if err != nil {
		return nil, err
	}
	return visits, nil
}

func (u *PathMetaUsecase) IsExists(ctx context.Context, shortPath string) (bool, error) {
	exist, err := u.redisRepo.IsExist(ctx, shortPath)
	if err != nil {
		return false, err
	}
	if exist == 1 {
		return true, nil
	}
	return false, nil
}

func (u *PathMetaUsecase) GetRepo() repository.RedisRepository {
	return u.redisRepo
}
