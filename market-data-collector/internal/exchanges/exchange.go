package exchanges

import "market-data-collector/internal/models"

type Exchange interface {
	Name() string
	Fetch(pair models.Pair) (models.Quote, error)
}
