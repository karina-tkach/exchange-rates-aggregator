package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"market-data-collector/internal/models"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisRateRepository struct {
	client *redis.Client
}

func NewRedisRateRepository(client *redis.Client) *RedisRateRepository {
	return &RedisRateRepository{
		client: client,
	}
}

type CachedRate struct {
	Price     string    `json:"price"`
	Bid       string    `json:"bid"`
	Ask       string    `json:"ask"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (r *RedisRateRepository) SaveCurrentRates(
	ctx context.Context,
	quotes []models.Quote,
	pairs map[uint32]string,
) error {
	grouped := make(map[string]map[string]CachedRate)

	for _, q := range quotes {
		pair := pairs[q.PairID]

		if grouped[pair] == nil {
			grouped[pair] = make(
				map[string]CachedRate,
			)
		}

		grouped[pair][q.Source] = CachedRate{
			Price:     q.Price.String(),
			Bid:       q.Bid.String(),
			Ask:       q.Ask.String(),
			UpdatedAt: q.Timestamp,
		}
	}

	for pair, value := range grouped {
		payload, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if err := r.client.Set(
			ctx,
			fmt.Sprintf("rates:%s", pair),
			payload,
			35*time.Second,
		).Err(); err != nil {
			return err
		}
	}

	return nil
}
