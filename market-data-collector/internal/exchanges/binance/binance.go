package binance

import (
	"context"
	"encoding/json"
	"fmt"
	"market-data-collector/internal/config"
	"market-data-collector/internal/httpclient"
	"market-data-collector/internal/models"
	"net/http"
	"time"

	"github.com/shopspring/decimal"
)

type Binance struct {
	cfg config.BinanceConfig
}

func New(cfg config.BinanceConfig) *Binance {
	return &Binance{cfg: cfg}
}

func (b *Binance) Name() string {
	return "Binance"
}

func (b *Binance) Fetch(ctx context.Context, pair models.Pair) (models.Quote, error) {
	symbol := pair.Base + pair.Quote
	url := fmt.Sprintf(
		"%s/api/v3/ticker/bookTicker?symbol=%s",
		b.cfg.BaseURL,
		symbol,
	)
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		url,
		nil,
	)

	if err != nil {
		return models.Quote{}, err
	}

	resp, err := httpclient.Client.Do(req)
	if err != nil {
		return models.Quote{}, err
	}
	defer resp.Body.Close()

	var book BookTickerResponse
	if err := json.NewDecoder(resp.Body).Decode(&book); err != nil {
		return models.Quote{}, err
	}

	if book.Code != 0 {
		return models.Quote{}, fmt.Errorf(
			"binance bookTicker error %s: %s",
			symbol,
			book.Message,
		)
	}

	bid, err := decimal.NewFromString(book.BidPrice)
	if err != nil {
		return models.Quote{}, err
	}

	ask, err := decimal.NewFromString(book.AskPrice)
	if err != nil {
		return models.Quote{}, err
	}

	priceURL := fmt.Sprintf(
		"%s/api/v3/ticker/price?symbol=%s",
		b.cfg.BaseURL,
		symbol,
	)
	priceResp, err := httpclient.Client.Get(priceURL)
	if err != nil {
		return models.Quote{}, err
	}
	defer priceResp.Body.Close()

	var priceData PriceResponse
	if err := json.NewDecoder(priceResp.Body).Decode(&priceData); err != nil {
		return models.Quote{}, err
	}

	if priceData.Code != 0 {
		return models.Quote{}, fmt.Errorf(
			"binance price error %s: %s",
			symbol,
			priceData.Message,
		)
	}

	price, err := decimal.NewFromString(priceData.Price)
	if err != nil {
		return models.Quote{}, err
	}

	return models.Quote{
		PairID:    pair.ID,
		Price:     price,
		Bid:       bid,
		Ask:       ask,
		Source:    b.Name(),
		Timestamp: time.Now(),
	}, nil
}
