package repositories

import (
	"context"
	"market-data-collector/internal/models"
)

type PairRepository interface {
	GetAll(ctx context.Context) ([]models.Pair, error)
}

type ExchangeRepository interface {
	GetEnabled(ctx context.Context) ([]string, error)
}

type QuoteRepository interface {
	SaveBatch(ctx context.Context, quotes []models.Quote) error
}

type CacheRateRepository interface {
	SaveCurrentRates(ctx context.Context, quotes []models.Quote, pairs map[uint32]string) error
}
