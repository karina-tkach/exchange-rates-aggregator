package config

type KrakenConfig struct {
	BaseURL string
}

func LoadKrakenConfig() KrakenConfig {
	return KrakenConfig{
		BaseURL: GetEnv(
			"KRAKEN_BASE_URL",
			"https://api.kraken.com",
		),
	}
}
