package exchanges

import (
	"context"
	"market-data-collector/internal/models"
)

type Exchange interface {
	Name() string
	Fetch(ctx context.Context, pair models.Pair) (models.Quote, error)
}
