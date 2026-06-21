package kraken

type TickerResponse struct {
	Error  []string          `json:"error"`
	Result map[string]Ticker `json:"result"`
}

type Ticker struct {
	Ask  [3]string `json:"a"`
	Bid  [3]string `json:"b"`
	Last [2]string `json:"c"`
}
