package coinbase

import (
	"context"
	"fmt"
	"market-data-collector/internal/config"
	"market-data-collector/internal/httpclient"
	"market-data-collector/internal/models"
	"time"

	"github.com/shopspring/decimal"
)

type Coinbase struct {
	cfg config.CoinbaseConfig
}

func New(cfg config.CoinbaseConfig) *Coinbase {
	return &Coinbase{cfg: cfg}
}

func (c *Coinbase) Name() string {
	return "Coinbase"
}

func (c *Coinbase) Fetch(ctx context.Context, pair models.Pair) (models.Quote, error) {
	symbol := pair.Base + "-" + pair.Quote
	url := fmt.Sprintf(
		"%s/products/%s/ticker",
		c.cfg.BaseURL,
		symbol,
	)

	var product ProductTickerResponse
	err := httpclient.GetJSON(ctx, httpclient.Client, url, &product)
	if err != nil {
		return models.Quote{}, err
	}

	if product.Message != "" {
		return models.Quote{}, fmt.Errorf(
			"coinbase error %s: %s",
			symbol,
			product.Message,
		)
	}

	bid, err := decimal.NewFromString(product.BidPrice)
	if err != nil {
		return models.Quote{}, fmt.Errorf("parsing bid: %w", err)
	}

	ask, err := decimal.NewFromString(product.AskPrice)
	if err != nil {
		return models.Quote{}, fmt.Errorf("parsing ask: %w", err)
	}

	price, err := decimal.NewFromString(product.Price)
	if err != nil {
		return models.Quote{}, fmt.Errorf("parsing price: %w", err)
	}

	return models.Quote{
		PairID:    pair.ID,
		Price:     price,
		Bid:       bid,
		Ask:       ask,
		Source:    c.Name(),
		Timestamp: time.Now(),
	}, nil
}
