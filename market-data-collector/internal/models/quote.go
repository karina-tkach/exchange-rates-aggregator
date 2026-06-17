package models

import "time"

type Quote struct {
	PairID    int
	Price     float64
	Bid       float64
	Ask       float64
	Source    string
	Timestamp time.Time
}
