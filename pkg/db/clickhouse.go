package db

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/kalunik/urShorty/config"
)

func NewClickhouseConnection(cfg *config.AppConfig) (driver.Conn, error) {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{cfg.Clickhouse.ClickhouseAddr},
		Auth: clickhouse.Auth{
			Database: cfg.Clickhouse.ClickhouseDb,
			Username: cfg.Clickhouse.Username,
			Password: cfg.Clickhouse.Password,
		},
		Settings: clickhouse.Settings{
			"max_execution_time": cfg.Clickhouse.MaxExecutionTime,
		},
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
		MaxOpenConns: cfg.Clickhouse.MaxOpenConns,
		MaxIdleConns: cfg.Clickhouse.MaxIdleConns,
	})
	if err != nil {
		return nil, err
	}
	if err := conn.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("clickhouse: connection couldn't be established: %w", err)
	}
	return conn, nil
}
