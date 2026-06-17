package collector

import (
	"log"
	"market-data-collector/internal/exchanges"
	"market-data-collector/internal/factory"
	"market-data-collector/internal/models"
	"market-data-collector/internal/repositories"
	"sync"
)

type Collector struct {
	pairRepo        repositories.PairRepository
	exchangeRepo    repositories.ExchangeRepository
	quoteRepo       repositories.QuoteRepository
	exchangeFactory *factory.ExchangeFactory
}

func NewCollector(pairRepo repositories.PairRepository, exchangeRepo repositories.ExchangeRepository,
	quoteRepo repositories.QuoteRepository, exchangeFactory *factory.ExchangeFactory) *Collector {
	return &Collector{
		pairRepo:        pairRepo,
		exchangeRepo:    exchangeRepo,
		quoteRepo:       quoteRepo,
		exchangeFactory: exchangeFactory,
	}
}

func (c *Collector) RunCycle() {
	pairs, err := c.pairRepo.GetAll()
	if err != nil {
		log.Println("pairs error:", err)
		return
	}

	exNames, err := c.exchangeRepo.GetEnabled()
	if err != nil {
		log.Println("exchanges error:", err)
		return
	}

	var exs []exchanges.Exchange
	for _, name := range exNames {
		ex := c.exchangeFactory.Build(name)
		if ex != nil {
			exs = append(exs, ex)
		}
	}

	var wg sync.WaitGroup
	quotesChan := make(chan models.Quote, 500)

	for _, ex := range exs {
		for _, pair := range pairs {
			wg.Add(1)

			go func(e exchanges.Exchange, p models.Pair) {
				defer wg.Done()

				quote, err := e.Fetch(p)
				if err != nil {
					log.Printf("[%s] error: %v\n", e.Name(), err)
					return
				}

				quotesChan <- quote

			}(ex, pair)
		}
	}

	wg.Wait()
	close(quotesChan)

	var quotes []models.Quote

	for q := range quotesChan {
		quotes = append(quotes, q)
	}

	if len(quotes) == 0 {
		return
	}

	if err := c.quoteRepo.SaveBatch(quotes); err != nil {
		log.Println("save batch error:", err)
	}
}
