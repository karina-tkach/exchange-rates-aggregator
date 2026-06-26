package kraken

import (
	"context"
	"fmt"
	"market-data-collector/internal/config"
	"market-data-collector/internal/httpclient"
	"market-data-collector/internal/models"
	"time"

	"github.com/shopspring/decimal"
)

type Kraken struct {
	cfg config.KrakenConfig
}

func New(cfg config.KrakenConfig) *Kraken {
	return &Kraken{cfg: cfg}
}

func (k *Kraken) Name() string {
	return "Kraken"
}

func (k *Kraken) Fetch(ctx context.Context, pair models.Pair) (models.Quote, error) {
	symbol := pair.Base + pair.Quote
	url := fmt.Sprintf(
		"%s/0/public/Ticker?pair=%s",
		k.cfg.BaseURL,
		symbol,
	)

	var data TickerResponse
	err := httpclient.GetJSON(ctx, httpclient.Client, url, &data)
	if err != nil {
		return models.Quote{}, err
	}

	if len(data.Error) > 0 {
		return models.Quote{}, fmt.Errorf(
			"kraken error (%s): %v",
			symbol,
			data.Error,
		)
	}

	if len(data.Result) == 0 {
		return models.Quote{}, fmt.Errorf(
			"kraken empty result for %s",
			symbol,
		)
	}

	var ticker Ticker

	for _, value := range data.Result {
		ticker = value
		break
	}

	ask, err := decimal.NewFromString(ticker.Ask[0])
	if err != nil {
		return models.Quote{}, fmt.Errorf("parsing ask: %w", err)
	}

	bid, err := decimal.NewFromString(ticker.Bid[0])
	if err != nil {
		return models.Quote{}, fmt.Errorf("parsing bid: %w", err)
	}

	price, err := decimal.NewFromString(ticker.Last[0])
	if err != nil {
		return models.Quote{}, fmt.Errorf("parsing price: %w", err)
	}

	return models.Quote{
		PairID:    pair.ID,
		Price:     price,
		Bid:       bid,
		Ask:       ask,
		Source:    k.Name(),
		Timestamp: time.Now(),
	}, nil
}
