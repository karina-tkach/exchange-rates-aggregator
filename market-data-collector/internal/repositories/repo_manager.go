package repositories

import (
	"market-data-collector/internal/repositories/postgres"
	redis2 "market-data-collector/internal/repositories/redis"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type RepoManager struct {
	PairRepo     PairRepository
	ExchangeRepo ExchangeRepository
	QuoteRepo    QuoteRepository
	RateCache    CacheRateRepository
}

func NewRepoManager(db *pgxpool.Pool, rdb *redis.Client) *RepoManager {
	return &RepoManager{
		PairRepo:     postgres.NewPostgresPairRepository(db),
		ExchangeRepo: postgres.NewPostgresExchangeRepository(db),
		QuoteRepo:    postgres.NewPostgresQuoteRepository(db),
		RateCache:    redis2.NewRedisRateRepository(rdb),
	}
}
