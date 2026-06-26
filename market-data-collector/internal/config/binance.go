package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type BinanceConfig struct {
	BaseURL string `envconfig:"BINANCE_BASE_URL" default:"https://api.binance.com"`
}

func loadBinanceConfig() (BinanceConfig, error) {
	var cfg BinanceConfig
	if err := envconfig.Process("", &cfg); err != nil {
		return BinanceConfig{}, fmt.Errorf("binance config: %w", err)
	}
	return cfg, nil
}
