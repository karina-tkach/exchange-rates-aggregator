package factory

import (
	"market-data-collector/internal/exchanges"
)

type ExchangeFactory struct{}

func NewExchangeFactory() *ExchangeFactory {
	return &ExchangeFactory{}
}

func (f *ExchangeFactory) Build(name string) exchanges.Exchange {
	switch name {

	case "Binance":
		return exchanges.Binance{}

	case "Kraken":
		return exchanges.Kraken{}

	case "Coinbase":
		return exchanges.Coinbase{}

	default:
		return nil
	}
}
