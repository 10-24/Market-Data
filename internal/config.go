package config

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/BurntSushi/toml"
)

const CandleDurationMs = 500

var GetConfigDirPath = sync.OnceValue(func() string {
	if path := os.Getenv("CONFIG_DIR"); path != "" {
		return path
	}
	return "./config"
})

var Get = sync.OnceValue(func() *Config {
	configPath := filepath.Join(GetConfigDirPath(), "config.toml")
	data, fileErr := os.ReadFile(configPath)
	if fileErr != nil {
		panic(fileErr)
	}

	var config Config
	if err := toml.Unmarshal(data, &config); err != nil {
		panic(err)
	}
	return &config
})

type Config struct {
	DbPath   string `toml:"db_path"`
	Settings struct {
		WatchTrades   WatchTradesConfig   `toml:"watch_trades"`
		CreateCandles CreateCandlesConfig `toml:"create_candles"`
	} `toml:"settings"`
}

type WatchTradesConfig struct {
	BatchSize       int `toml:"batch_size"`
	BatchBufferSize int `toml:"batch_buffer_size"`
}
type CreateCandlesConfig struct {
	BatchSize       int `toml:"batch_size"`
	BatchBufferSize int `toml:"batch_buffer_size"`
}
