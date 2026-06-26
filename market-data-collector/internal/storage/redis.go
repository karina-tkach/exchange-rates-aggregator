package storage

import (
	"market-data-collector/internal/config"

	"github.com/redis/go-redis/v9"
)

func CreateRedisClient(cfg config.RedisConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: cfg.RedisAddress,
	})
}
