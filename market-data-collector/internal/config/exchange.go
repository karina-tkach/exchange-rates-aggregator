package config

type ExchangeConfig struct {
	Binance  BinanceConfig
	Kraken   KrakenConfig
	Coinbase CoinbaseConfig
}

func loadExchangeConfig() (ExchangeConfig, error) {
	binance, err := loadBinanceConfig()
	if err != nil {
		return ExchangeConfig{}, err
	}

	kraken, err := loadKrakenConfig()
	if err != nil {
		return ExchangeConfig{}, err
	}

	coinbase, err := loadCoinbaseConfig()
	if err != nil {
		return ExchangeConfig{}, err
	}

	return ExchangeConfig{
		Binance:  binance,
		Kraken:   kraken,
		Coinbase: coinbase,
	}, nil
}
