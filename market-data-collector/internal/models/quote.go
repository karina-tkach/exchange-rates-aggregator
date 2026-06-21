package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Quote struct {
	PairID    uint32
	Price     decimal.Decimal
	Bid       decimal.Decimal
	Ask       decimal.Decimal
	Source    string
	Timestamp time.Time
}
