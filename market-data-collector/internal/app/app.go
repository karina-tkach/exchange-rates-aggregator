package app

import (
	"context"
	"log"
	"market-data-collector/internal/collector"
	"market-data-collector/internal/config"
	"market-data-collector/internal/factory"
	"market-data-collector/internal/httpclient"
	"market-data-collector/internal/repositories"
	"market-data-collector/internal/scheduler"
	"market-data-collector/internal/storage"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

func Run() {
	_ = godotenv.Load()
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db := storage.ConnectPostgres(cfg.Database)

	redis := storage.CreateRedisClient(cfg.Redis)

	httpclient.Init(cfg.Http)

	repos := repositories.NewRepoManager(db, redis)

	exFactory := factory.NewExchangeFactory(cfg.Exchange)

	col := collector.NewCollector(
		repos.PairRepo,
		repos.ExchangeRepo,
		repos.QuoteRepo,
		repos.RateCache,
		exFactory,
		cfg.Collector,
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
