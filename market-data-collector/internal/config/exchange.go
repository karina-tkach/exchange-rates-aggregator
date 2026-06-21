package config

type ExchangeConfig struct {
	Binance  BinanceConfig
	Kraken   KrakenConfig
	Coinbase CoinbaseConfig
}

func LoadExchangeConfig() ExchangeConfig {
	return ExchangeConfig{
		Binance:  LoadBinanceConfig(),
		Kraken:   LoadKrakenConfig(),
		Coinbase: LoadCoinbaseConfig(),
	}
}
