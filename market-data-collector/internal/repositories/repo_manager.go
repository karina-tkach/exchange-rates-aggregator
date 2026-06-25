package repositories

import (
	"market-data-collector/internal/repositories/postgres"
	redisPkg "market-data-collector/internal/repositories/redis"

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
		PairRepo:     postgres.NewPairRepository(db),
		ExchangeRepo: postgres.NewExchangeRepository(db),
		QuoteRepo:    postgres.NewQuoteRepository(db),
		RateCache:    redisPkg.NewRateRepository(rdb),
	}
}
