package app

import (
	"context"
	"log"
	"market-data-collector/internal/collector"
	"market-data-collector/internal/config"
	"market-data-collector/internal/database"
	"market-data-collector/internal/factory"
	"market-data-collector/internal/repositories"
	"market-data-collector/internal/scheduler"
	"os"
	"os/signal"
	"syscall"
)

func Run() {
	config.LoadEnv()

	db := database.ConnectPostgres()

	repos := repositories.NewRepoManager(db)

	exchangeConfig := config.LoadExchangeConfig()
	exFactory := factory.NewExchangeFactory(exchangeConfig)

	col := collector.NewCollector(
		repos.PairRepo,
		repos.ExchangeRepo,
		repos.QuoteRepo,
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

	log.Println("✅ db closed, exit complete")
}
