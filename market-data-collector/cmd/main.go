package main

import (
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

func main() {
	config.LoadEnv()

	db := database.ConnectPostgres()

	pairRepo := repositories.NewPostgresPairRepository(db)
	exchangeRepo := repositories.NewPostgresExchangeRepository(db)
	quoteRepo := repositories.NewPostgresQuoteRepository(db)

	exFactory := factory.NewExchangeFactory()

	col := collector.NewCollector(
		pairRepo,
		exchangeRepo,
		quoteRepo,
		exFactory,
	)

	sc := scheduler.NewScheduler(col)

	log.Println("🚀 collector started")

	go sc.Start(30)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	<-sig

	log.Println("🛑 shutting down collector...")

	db.Close()

	log.Println("✅ db closed, exit complete")
}
