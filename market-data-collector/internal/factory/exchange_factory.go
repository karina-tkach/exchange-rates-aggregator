package factory

import (
	"market-data-collector/internal/config"
	"market-data-collector/internal/exchanges"
	"market-data-collector/internal/exchanges/binance"
	"market-data-collector/internal/exchanges/coinbase"
	"market-data-collector/internal/exchanges/kraken"
)

type ExchangeFactory struct {
	config config.ExchangeConfig
}

func NewExchangeFactory(cfg config.ExchangeConfig) *ExchangeFactory {
	return &ExchangeFactory{
		config: cfg,
	}
}

func (f *ExchangeFactory) Build(name string) exchanges.Exchange {
	switch name {

	case "Binance":
		return binance.New(f.config.Binance)

	case "Kraken":
		return kraken.New(f.config.Kraken)

	case "Coinbase":
		return coinbase.New(f.config.Coinbase)

	default:
		return nil
	}
}
