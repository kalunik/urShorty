package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLoadNewConfig(t *testing.T) {
}

func TestConfigDriver_ParseConfig(t *testing.T) {
	configKeyEnv := "custom"
	configPathEnv := "config-sample.yml"
	expectedConfig := &AppConfig{
		Server: ServerConfig{":4000", false, 12},
		Redis: RedisConfig{"localhost:6379", "", "0", "0",
			200, 12000, 240, "", 0},
		Clickhouse: ClickhouseConfig{"localhost:9000", "default", "default",
			"", 60, 5, 5},
	}

	if err := os.Setenv("CONFIG", configKeyEnv); err != nil {
		assert.NoError(t, err, "err not expected")
	}
	if err := os.Setenv("CONFIG_PATH", configPathEnv); err != nil {
		assert.NoError(t, err, "err not expected")
	}

	configDriver, err := LoadNewConfig()
	if err != nil {
		assert.NoError(t, err, "load config err not expected")
	}
	config, _ := configDriver.ParseConfig()
	if err != nil {
		assert.NoError(t, err, "parse config err not expected")
	}

	assert.Equal(t, expectedConfig, config,
		"configs should be the same: expected %s, got %s\n", expectedConfig, config)
}

func TestFindConfigPath(t *testing.T) {
	type setenv struct {
		key   string
		value string
	}
	testCollection := []struct {
		setenv
		expectedPath string
	}{
		{
			setenv: setenv{
				key:   "CONFIG",
				value: "docker",
			},
			expectedPath: "config/config-docker.yml",
		},
		{
			setenv: setenv{
				key:   "CONFIG",
				value: "",
			},
			expectedPath: "config/config-local.yml",
		},
		{
			setenv: setenv{
				key:   "CONFIG",
				value: "unknown",
			},
			expectedPath: "config/config-local.yml",
		},
	}

	for _, test := range testCollection {
		err := os.Setenv(test.setenv.key, test.setenv.value)
		if err != nil {
			return
		}
		path := findConfigPath()
		assert.Equalf(t, test.expectedPath, path,
			"paths should be the same: expected %s, got %s\n", test.expectedPath, path)
	}
}
