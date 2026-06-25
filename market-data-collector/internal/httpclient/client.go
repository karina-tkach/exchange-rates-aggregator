package httpclient

import (
	"market-data-collector/internal/config"
	"net"
	"net/http"
)

var Client *http.Client

func Init(cfg config.HttpConfig) {
	Client = NewClient(cfg)
}

func NewClient(cfg config.HttpConfig) *http.Client {
	return &http.Client{
		Timeout: cfg.ClientTimeout,
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 50,
			IdleConnTimeout:     cfg.IdleConTimeout,
			DialContext: (&net.Dialer{
				Timeout:   cfg.DialTimeout,
				KeepAlive: cfg.KeepAlive,
			}).DialContext,
		},
	}
}
