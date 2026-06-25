package postgres

import (
	"context"
	"fmt"
	"market-data-collector/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PairRepository struct {
	db *pgxpool.Pool
}

func NewPairRepository(db *pgxpool.Pool) *PairRepository {
	return &PairRepository{
		db: db,
	}
}

func (r *PairRepository) GetAll(ctx context.Context) ([]models.Pair, error) {
	rows, err := r.db.Query(
		ctx,
		`
		SELECT id, base, quote
		FROM pairs
		ORDER BY id
		`,
	)
	if err != nil {
		return nil, fmt.Errorf("get all pairs: %w", err)
	}

	pairs, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.Pair])
	if err != nil {
		return nil, fmt.Errorf("map all pairs: %w", err)
	}

	return pairs, nil
}
