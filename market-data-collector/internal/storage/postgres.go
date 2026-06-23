package storage

import (
	"context"
	"log"
	"market-data-collector/internal/config"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectPostgres() *pgxpool.Pool {
	dsn := config.GetPostgresDBConnectionString()

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("failed to create postgres pool: %v", err)
	}

	timeout := config.GetDuration(
		"DB_TIMEOUT",
		3*time.Second,
	)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("failed to ping postgres: %v", err)
	}

	log.Println("postgres connected!")

	return pool
}
