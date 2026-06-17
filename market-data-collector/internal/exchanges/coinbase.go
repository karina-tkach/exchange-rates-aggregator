package exchanges

import (
	"encoding/json"
	"fmt"
	"market-data-collector/internal/httpclient"
	"market-data-collector/internal/models"
	"strconv"
	"time"
)

type Coinbase struct{}

func (c Coinbase) Name() string {
	return "Coinbase"
}

func (c Coinbase) Fetch(pair models.Pair) (models.Quote, error) {
	symbol := pair.Base + "-" + pair.Quote
	resp, err := httpclient.Client.Get("https://api.exchange.coinbase.com/products/" + symbol + "/ticker")

	if err != nil {
		return models.Quote{}, err
	}
	defer resp.Body.Close()

	var raw map[string]any

	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return models.Quote{}, err
	}

	if msg, ok := raw["message"]; ok {
		return models.Quote{}, fmt.Errorf(
			"coinbase error (%s): %v",
			symbol,
			msg,
		)
	}

	priceStr, ok := raw["price"].(string)
	if !ok {
		return models.Quote{}, fmt.Errorf("coinbase: missing price for %s", symbol)
	}

	bidStr, ok := raw["bid"].(string)
	if !ok {
		return models.Quote{}, fmt.Errorf("coinbase: missing bid for %s", symbol)
	}

	askStr, ok := raw["ask"].(string)
	if !ok {
		return models.Quote{}, fmt.Errorf("coinbase: missing ask for %s", symbol)
	}

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return models.Quote{}, err
	}

	bid, err := strconv.ParseFloat(bidStr, 64)
	if err != nil {
		return models.Quote{}, err
	}

	ask, err := strconv.ParseFloat(askStr, 64)
	if err != nil {
		return models.Quote{}, err
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
