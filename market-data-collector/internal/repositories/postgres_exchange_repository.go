package repositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresExchangeRepository struct {
	db *pgxpool.Pool
}

func NewPostgresExchangeRepository(db *pgxpool.Pool) *PostgresExchangeRepository {
	return &PostgresExchangeRepository{
		db: db,
	}
}

func (r *PostgresExchangeRepository) GetEnabled() ([]string, error) {
	rows, err := r.db.Query(
		context.Background(),
		`
		SELECT name
		FROM exchanges
		WHERE is_enabled = true
		ORDER BY id
		`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var exchanges []string

	for rows.Next() {
		var name string

		if err := rows.Scan(&name); err != nil {
			return nil, err
		}

		exchanges = append(exchanges, name)
	}

	return exchanges, nil
}
