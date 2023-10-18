package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"time"
)

type AppConfig struct {
	Server ServerConfig
	Redis  RedisConfig
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
	Db             int
}

type ConfigDriver struct {
	v *viper.Viper
}

func LoadNewConfig() (*ConfigDriver, error) {
	v := viper.New()
	v.SetConfigFile(findConfigPath())

	if err := v.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError

		if errors.As(err, &configFileNotFoundError) {
			return nil, errors.New("Config: file not found")
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
		return nil, fmt.Errorf("Config: unable to decode into struct: %w", err)
	}
	return config, nil
}

func findConfigPath() string {
	configPaths := map[string]string{
		"docker": "config/config-docker.yml",
		"local":  "config/config-local.yml",
	}

	pathKey := os.Getenv("CONFIG")
	if pathKey == "" {
		return configPaths["local"]
	}
	return configPaths[pathKey]
}
