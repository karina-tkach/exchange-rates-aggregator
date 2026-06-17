package exchanges

import (
	"encoding/json"
	"fmt"
	"market-data-collector/internal/httpclient"
	"market-data-collector/internal/models"
	"strconv"
	"time"
)

type Binance struct{}

func (b Binance) Name() string {
	return "Binance"
}

func (b Binance) Fetch(pair models.Pair) (models.Quote, error) {
	symbol := pair.Base + pair.Quote
	resp, err := httpclient.Client.Get("https://api.binance.com/api/v3/ticker/bookTicker?symbol=" + symbol)
	if err != nil {
		return models.Quote{}, err
	}
	defer resp.Body.Close()

	var book map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&book); err != nil {
		return models.Quote{}, err
	}

	if code, ok := book["code"]; ok {
		return models.Quote{}, fmt.Errorf(
			"binance bookTicker error (%s): %v - %v",
			symbol,
			code,
			book["msg"],
		)
	}

	bidStr, ok := book["bidPrice"].(string)
	if !ok {
		return models.Quote{}, fmt.Errorf("binance missing bidPrice for %s", symbol)
	}

	askStr, ok := book["askPrice"].(string)
	if !ok {
		return models.Quote{}, fmt.Errorf("binance missing askPrice for %s", symbol)
	}

	bid, err := strconv.ParseFloat(bidStr, 64)
	if err != nil {
		return models.Quote{}, err
	}

	ask, err := strconv.ParseFloat(askStr, 64)
	if err != nil {
		return models.Quote{}, err
	}

	priceResp, err := httpclient.Client.Get(
		"https://api.binance.com/api/v3/ticker/price?symbol=" + symbol,
	)
	if err != nil {
		return models.Quote{}, err
	}
	defer priceResp.Body.Close()

	var priceData map[string]any
	if err := json.NewDecoder(priceResp.Body).Decode(&priceData); err != nil {
		return models.Quote{}, err
	}

	if code, ok := priceData["code"]; ok {
		return models.Quote{}, fmt.Errorf(
			"binance price error (%s): %v - %v",
			symbol,
			code,
			priceData["msg"],
		)
	}

	priceStr, ok := priceData["price"].(string)
	if !ok {
		return models.Quote{}, fmt.Errorf("binance missing price for %s", symbol)
	}

	price, err := strconv.ParseFloat(priceStr, 64)
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
