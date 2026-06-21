package httpclient

import (
	"market-data-collector/internal/config"
	"net"
	"net/http"
	"time"
)

func NewClient() *http.Client {
	return &http.Client{
		Timeout: config.GetDuration(
			"HTTP_CLIENT_TIMEOUT",
			3*time.Second,
		),
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 50,
			IdleConnTimeout: config.GetDuration(
				"HTTP_IDLE_CONN_TIMEOUT",
				90*time.Second,
			),
			DialContext: (&net.Dialer{
				Timeout: config.GetDuration(
					"HTTP_DIAL_TIMEOUT",
					2*time.Second,
				),
				KeepAlive: config.GetDuration(
					"HTTP_KEEP_ALIVE",
					30*time.Second,
				),
			}).DialContext,
		},
	}
}

var Client = NewClient()
