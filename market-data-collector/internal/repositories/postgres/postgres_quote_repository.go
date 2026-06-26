package postgres

import (
	"context"
	"fmt"
	"market-data-collector/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type QuoteRepository struct {
	db *pgxpool.Pool
}

func NewQuoteRepository(db *pgxpool.Pool) *QuoteRepository {
	return &QuoteRepository{
		db: db,
	}
}

func (r *QuoteRepository) SaveBatch(ctx context.Context, quotes []models.Quote) error {
	if len(quotes) == 0 {
		return nil
	}

	rows := make([][]any, 0, len(quotes))

	for _, q := range quotes {
		rows = append(rows, []any{
			q.Timestamp,
			q.PairID,
			q.Price,
			q.Bid,
			q.Ask,
			q.Source,
		})
	}

	_, err := r.db.CopyFrom(
		ctx, pgx.Identifier{"quotes"},
		[]string{
			"time",
			"pair_id",
			"price",
			"bid",
			"ask",
			"source",
		},
		pgx.CopyFromRows(rows),
	)

	if err != nil {
		return fmt.Errorf("save quotes: %w", err)
	}

	return nil
}
