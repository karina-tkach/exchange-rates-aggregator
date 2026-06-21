package repositories

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type RepoManager struct {
	PairRepo     PairRepository
	ExchangeRepo ExchangeRepository
	QuoteRepo    QuoteRepository
}

func NewRepoManager(db *pgxpool.Pool) *RepoManager {
	return &RepoManager{
		PairRepo:     NewPostgresPairRepository(db),
		ExchangeRepo: NewPostgresExchangeRepository(db),
		QuoteRepo:    NewPostgresQuoteRepository(db),
	}
}
