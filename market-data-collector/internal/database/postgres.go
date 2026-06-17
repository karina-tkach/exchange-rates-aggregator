package database

import (
	"context"
	"log"
	"market-data-collector/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectPostgres() *pgxpool.Pool {
	dsn := config.GetPostgresDBConnectionString()

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("failed to create postgres pool: %v", err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		log.Fatalf("failed to ping postgres: %v", err)
	}

	log.Println("postgres connected!")

	return pool
}
