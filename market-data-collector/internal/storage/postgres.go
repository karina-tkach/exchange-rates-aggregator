package storage

import (
	"context"
	"fmt"
	"log"
	"market-data-collector/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectPostgres(cfg config.DatabaseConfig) *pgxpool.Pool {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)

	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatalf("parse postgres config: %v", err)
	}

	poolConfig.MaxConns = int32(cfg.MaxConns)
	poolConfig.MinConns = int32(cfg.MinConns)

	poolConfig.MaxConnLifetime = cfg.MaxConnLifetime

	poolConfig.MaxConnIdleTime = cfg.MaxConnIdleTime

	pool, err := pgxpool.NewWithConfig(
		context.Background(),
		poolConfig,
	)

	if err != nil {
		log.Fatalf("failed to create postgres pool: %v", err)
	}

	timeout := cfg.Timeout

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("failed to ping postgres: %v", err)
	}

	log.Println("postgres connected!")

	return pool
}
