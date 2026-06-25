package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type CoinbaseConfig struct {
	BaseURL string `envconfig:"COINBASE_BASE_URL" default:"https://api.exchange.coinbase.com"`
}

func loadCoinbaseConfig() (CoinbaseConfig, error) {
	var cfg CoinbaseConfig
	if err := envconfig.Process("", &cfg); err != nil {
		return CoinbaseConfig{}, fmt.Errorf("coinbase config: %w", err)
	}
	return cfg, nil
}
