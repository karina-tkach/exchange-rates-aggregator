package repositories

import (
	"context"
	"market-data-collector/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresPairRepository struct {
	db *pgxpool.Pool
}

func NewPostgresPairRepository(db *pgxpool.Pool) *PostgresPairRepository {
	return &PostgresPairRepository{
		db: db,
	}
}

func (r *PostgresPairRepository) GetAll(ctx context.Context) ([]models.Pair, error) {
	rows, err := r.db.Query(
		ctx,
		`
		SELECT id, base, quote
		FROM pairs
		ORDER BY id
		`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pairs []models.Pair

	for rows.Next() {
		var pair models.Pair

		if err := rows.Scan(
			&pair.ID,
			&pair.Base,
			&pair.Quote,
		); err != nil {
			return nil, err
		}

		pairs = append(pairs, pair)
	}

	return pairs, nil
}
