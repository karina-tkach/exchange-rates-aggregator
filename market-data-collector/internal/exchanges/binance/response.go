package binance

type BookTickerResponse struct {
	BidPrice string `json:"bidPrice"`
	AskPrice string `json:"askPrice"`

	Code    int    `json:"code"`
	Message string `json:"msg"`
}

type PriceResponse struct {
	Price string `json:"price"`

	Code    int    `json:"code"`
	Message string `json:"message"`
}
