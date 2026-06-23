package storage

import (
	"market-data-collector/internal/config"

	"github.com/redis/go-redis/v9"
)

func CreateRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: config.GetRedisAddress(),
	})
}
