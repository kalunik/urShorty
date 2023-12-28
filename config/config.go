package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"time"
)

//go:generate mockgen -source=config.go -destination=mocks/mock.go

type AppConfig struct {
	Server     ServerConfig
	Redis      RedisConfig
	Clickhouse ClickhouseConfig
}

type ServerConfig struct {
	Port              string
	Ssl               bool
	CtxDefaultTimeout time.Duration
}

type RedisConfig struct {
	RedisAddr      string
	RedisPassword  string
	RedisDB        string
	RedisDefaultdb string
	MinIdleConns   int
	PoolSize       int
	PoolTimeout    int
	Password       string
	DB             int
}

type ClickhouseConfig struct {
	ClickhouseAddr   string
	ClickhouseDb     string
	Username         string
	Password         string
	MaxExecutionTime int
	MaxOpenConns     int
	MaxIdleConns     int
}

type Config interface {
	ParseConfig() (*AppConfig, error)
}

type ConfigDriver struct {
	v *viper.Viper
}

func LoadNewConfig() (Config, error) {
	v := viper.New()
	v.SetConfigFile(findConfigPath())

	if err := v.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError

		if errors.As(err, &configFileNotFoundError) {
			return nil, errors.New("config: file not found")
		}
		return nil, err
	}

	return &ConfigDriver{
		v: v,
	}, nil
}

func (c *ConfigDriver) ParseConfig() (*AppConfig, error) {
	config := &AppConfig{}

	err := c.v.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("config: unable to decode into struct: %w", err)
	}
	return config, nil
}

func findConfigPath() string {
	customPath := os.Getenv("CONFIG_PATH")

	configPaths := map[string]string{
		"docker": "config/config-docker.yml",
		"local":  "config/config-local.yml",
		"custom": customPath,
	}
	defaultPath := configPaths["local"]

	pathKey := os.Getenv("CONFIG")
	if pathKey == "" {
		return defaultPath
	}
	if path, ok := configPaths[pathKey]; ok {
		return path
	}
	return defaultPath
}
