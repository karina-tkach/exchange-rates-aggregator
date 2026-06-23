package app

import (
	"context"
	"log"
	"market-data-collector/internal/collector"
	"market-data-collector/internal/config"
	"market-data-collector/internal/factory"
	"market-data-collector/internal/repositories"
	"market-data-collector/internal/scheduler"
	"market-data-collector/internal/storage"
	"os"
	"os/signal"
	"syscall"
)

func Run() {
	config.LoadEnv()

	db := storage.ConnectPostgres()

	redis := storage.CreateRedisClient()

	repos := repositories.NewRepoManager(db, redis)

	exchangeConfig := config.LoadExchangeConfig()
	exFactory := factory.NewExchangeFactory(exchangeConfig)

	col := collector.NewCollector(
		repos.PairRepo,
		repos.ExchangeRepo,
		repos.QuoteRepo,
		repos.RateCache,
		exFactory,
	)

	sc := scheduler.NewScheduler(col)

	ctx, cancel := context.WithCancel(context.Background())

	log.Println("🚀 collector started")

	sc.Start(ctx, 30)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	<-sig

	log.Println("🛑 shutting down collector...")

	cancel()
	sc.Wait()
	db.Close()
	if err := redis.Close(); err != nil {
		log.Printf("redis close error: %v", err)
	}

	log.Println("✅ exit complete")
}
