package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type KrakenConfig struct {
	BaseURL string `envconfig:"KRAKEN_BASE_URL" default:"https://api.kraken.com"`
}

func loadKrakenConfig() (KrakenConfig, error) {
	var cfg KrakenConfig
	if err := envconfig.Process("", &cfg); err != nil {
		return KrakenConfig{}, fmt.Errorf("kraken config: %w", err)
	}
	return cfg, nil
}
