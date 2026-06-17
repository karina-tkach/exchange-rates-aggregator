package exchanges

import (
	"encoding/json"
	"fmt"
	"market-data-collector/internal/httpclient"
	"market-data-collector/internal/models"
	"strconv"
	"time"
)

type Kraken struct{}

func (k Kraken) Name() string {
	return "Kraken"
}

func (k Kraken) Fetch(pair models.Pair) (models.Quote, error) {
	symbol := pair.Base + pair.Quote
	resp, err := httpclient.Client.Get("https://api.kraken.com/0/public/Ticker?pair=" + symbol)
	if err != nil {
		return models.Quote{}, err
	}
	defer resp.Body.Close()

	var data struct {
		Error  []string       `json:"error"`
		Result map[string]any `json:"result"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
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

	for _, v := range data.Result {
		m, ok := v.(map[string]any)
		if !ok {
			return models.Quote{}, fmt.Errorf("invalid kraken response format")
		}

		ask, err := strconv.ParseFloat(m["a"].([]any)[0].(string), 64)
		if err != nil {
			return models.Quote{}, err
		}

		bid, err := strconv.ParseFloat(m["b"].([]any)[0].(string), 64)
		if err != nil {
			return models.Quote{}, err
		}

		price, err := strconv.ParseFloat(m["c"].([]any)[0].(string), 64)
		if err != nil {
			return models.Quote{}, err
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

	return models.Quote{}, nil
}
