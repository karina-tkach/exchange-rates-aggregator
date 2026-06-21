package config

type BinanceConfig struct {
	BaseURL string
}

func LoadBinanceConfig() BinanceConfig {
	return BinanceConfig{
		BaseURL: GetEnv(
			"BINANCE_BASE_URL",
			"https://api.binance.com",
		),
	}
}
