package config

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Database  DatabaseConfig
	Redis     RedisConfig
	Http      HttpConfig
	Collector CollectorConfig
	Exchange  ExchangeConfig
}

type DatabaseConfig struct {
	Host     string `envconfig:"DB_HOST" default:"localhost"`
	Port     string `envconfig:"DB_PORT" default:"5432"`
	User     string `envconfig:"DB_USER" default:"postgres"`
	Password string `envconfig:"DB_PASSWORD" default:"postgres"`
	Name     string `envconfig:"DB_NAME" default:"postgres"`

	Timeout         time.Duration `envconfig:"DB_TIMEOUT" default:"5m"`
	MaxConns        int           `envconfig:"DB_MAX_CONNS" default:"10"`
	MinConns        int           `envconfig:"DB_MIN_CONNS" default:"10"`
	MaxConnLifetime time.Duration `envconfig:"DB_MAX_CONN_LIFETIME" default:"1h"`
	MaxConnIdleTime time.Duration `envconfig:"DB_MAX_CONN_IDLE_TIME" default:"30m"`
}

type RedisConfig struct {
	RedisAddress string `envconfig:"REDIS_ADDRESS" default:"localhost:6379"`
}

type HttpConfig struct {
	ClientTimeout  time.Duration `envconfig:"HTTP_CLIENT_TIMEOUT" default:"10s"`
	DialTimeout    time.Duration `envconfig:"HTTP_DIAL_TIMEOUT" default:"10s"`
	KeepAlive      time.Duration `envconfig:"HTTP_KEEP_ALIVE" default:"10s"`
	IdleConTimeout time.Duration `envconfig:"HTTP_IDLE_CON_TIMEOUT" default:"10s"`
}

type CollectorConfig struct {
	CollectorCycleTimeout time.Duration `envconfig:"COLLECTOR_CYCLE_TIMEOUT" default:"35s"`
	CollectorWorkers      int           `envconfig:"COLLECTOR_WORKERS" default:"50"`
}

func Load() (*Config, error) {
	db, err := loadDatabaseConfig()
	if err != nil {
		return nil, err
	}

	redis, err := loadRedisConfig()
	if err != nil {
		return nil, err
	}

	httpCfg, err := loadHttpConfig()
	if err != nil {
		return nil, err
	}

	collector, err := loadCollectorConfig()
	if err != nil {
		return nil, err
	}

	exchange, err := loadExchangeConfig()
	if err != nil {
		return nil, err
	}

	return &Config{
		Database:  db,
		Redis:     redis,
		Http:      httpCfg,
		Collector: collector,
		Exchange:  exchange,
	}, nil
}

func loadDatabaseConfig() (DatabaseConfig, error) {
	var cfg DatabaseConfig
	if err := envconfig.Process("", &cfg); err != nil {
		return DatabaseConfig{}, fmt.Errorf("database config: %w", err)
	}
	return cfg, nil
}

func loadRedisConfig() (RedisConfig, error) {
	var cfg RedisConfig
	if err := envconfig.Process("", &cfg); err != nil {
		return RedisConfig{}, fmt.Errorf("redis config: %w", err)
	}
	return cfg, nil
}

func loadHttpConfig() (HttpConfig, error) {
	var cfg HttpConfig
	if err := envconfig.Process("", &cfg); err != nil {
		return HttpConfig{}, fmt.Errorf("http config: %w", err)
	}
	return cfg, nil
}

func loadCollectorConfig() (CollectorConfig, error) {
	var cfg CollectorConfig
	if err := envconfig.Process("", &cfg); err != nil {
		return CollectorConfig{}, fmt.Errorf("collector config: %w", err)
	}
	return cfg, nil
}
