package repositories

import "market-data-collector/internal/models"

type PairRepository interface {
	GetAll() ([]models.Pair, error)
}

type ExchangeRepository interface {
	GetEnabled() ([]string, error)
}

type QuoteRepository interface {
	SaveBatch([]models.Quote) error
}
