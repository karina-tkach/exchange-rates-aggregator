package config

type CoinbaseConfig struct {
	BaseURL string
}

func LoadCoinbaseConfig() CoinbaseConfig {
	return CoinbaseConfig{
		BaseURL: GetEnv(
			"COINBASE_BASE_URL",
			"https://api.exchange.coinbase.com",
		),
	}
}
