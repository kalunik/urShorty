package repository

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/kalunik/urShorty/internal/entity"
)

type ClickhouseRepository interface {
	AddNewShortPath(ctx context.Context, meta entity.PathMeta) error
	AddPathVisit(ctx context.Context, meta entity.PathMeta) error
	PathVisits(ctx context.Context, shortPath string) (*entity.PathVisitsList, error)
}

type clickhouseRepo struct {
	client driver.Conn
}

func NewClickhouseRepository(client driver.Conn) ClickhouseRepository {
	return &clickhouseRepo{
		client: client,
	}
}

func (c clickhouseRepo) AddNewShortPath(ctx context.Context, meta entity.PathMeta) error {
	if err := c.client.Exec(ctx, addNewPath,
		meta.ShortPath, meta.Domain, meta.CreatedAt,
		meta.Latitude, meta.Longitude, meta.Country,
		meta.City, meta.Proxy); err != nil {
		return fmt.Errorf("clickhouse client: AddNewShortPath failure: %w", err)
	}
	return nil
}

func (c clickhouseRepo) AddPathVisit(ctx context.Context, meta entity.PathMeta) error {
	if err := c.client.Exec(ctx, addPathVisit,
		meta.ShortPath, meta.VisitedAt,
		meta.Latitude, meta.Longitude, meta.Country,
		meta.City, meta.Proxy); err != nil {
		return fmt.Errorf("clickhouse client: AddPathVisit failure: %w", err)
	}
	return nil
}

func (c clickhouseRepo) PathVisits(ctx context.Context, shortPath string) (*entity.PathVisitsList, error) {
	meta := entity.VisitMeta{}
	response := &entity.PathVisitsList{}
	rows, err := c.client.Query(ctx, listPathVisits, shortPath)
	if err != nil {
		return nil, fmt.Errorf("clickhouse client: PathVisits failure: %w", err)
	}

	for rows.Next() {
		if err := rows.Scan(&meta.VisitedAt, &meta.Latitude, &meta.Longitude,
			&meta.Country, &meta.City, &meta.Proxy); err != nil {
			return nil, err
		}
		response.Visits = append(response.Visits, &meta)
		response.TotalCount++
	}
	return response, nil
}
