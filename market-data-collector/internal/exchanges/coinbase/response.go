package coinbase

type ProductTickerResponse struct {
	BidPrice string `json:"bid"`
	AskPrice string `json:"ask"`
	Price    string `json:"price"`

	Message string `json:"message"`
}
