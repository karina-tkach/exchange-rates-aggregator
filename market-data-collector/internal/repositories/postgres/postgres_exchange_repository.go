package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ExchangeRepository struct {
	db *pgxpool.Pool
}

func NewExchangeRepository(db *pgxpool.Pool) *ExchangeRepository {
	return &ExchangeRepository{
		db: db,
	}
}

func (r *ExchangeRepository) GetEnabled(ctx context.Context) ([]string, error) {
	rows, err := r.db.Query(
		ctx,
		`
		SELECT name
		FROM exchanges
		WHERE is_enabled = true
		ORDER BY id
		`,
	)
	if err != nil {
		return nil, fmt.Errorf("get enabled exchanges: %w", err)
	}
	defer rows.Close()

	var exchanges []string

	for rows.Next() {
		var name string

		if err := rows.Scan(&name); err != nil {
			return nil, fmt.Errorf("scan exchange name: %w", err)
		}

		exchanges = append(exchanges, name)
	}

	return exchanges, nil
}
