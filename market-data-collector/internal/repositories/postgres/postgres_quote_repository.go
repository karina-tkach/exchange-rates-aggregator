package postgres

import (
	"context"
	"market-data-collector/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresQuoteRepository struct {
	db *pgxpool.Pool
}

func NewPostgresQuoteRepository(db *pgxpool.Pool) *PostgresQuoteRepository {
	return &PostgresQuoteRepository{
		db: db,
	}
}

func (r *PostgresQuoteRepository) SaveBatch(ctx context.Context, quotes []models.Quote) error {
	if len(quotes) == 0 {
		return nil
	}

	batch := &pgx.Batch{}

	for _, q := range quotes {
		batch.Queue(`
			INSERT INTO quotes (
				pair_id,
				price,
				bid,
				ask,
				source,
				time
			)
			VALUES ($1,$2,$3,$4,$5,$6)
		`,
			q.PairID,
			q.Price,
			q.Bid,
			q.Ask,
			q.Source,
			q.Timestamp,
		)
	}

	br := r.db.SendBatch(ctx, batch)
	defer br.Close()

	for range quotes {
		if _, err := br.Exec(); err != nil {
			return err
		}
	}

	return nil
}
