package collector

import (
	"context"
	"log"
	"market-data-collector/internal/config"
	"market-data-collector/internal/exchanges"
	"market-data-collector/internal/factory"
	"market-data-collector/internal/models"
	"market-data-collector/internal/repositories"
	"sync"
)

type job struct {
	ex   exchanges.Exchange
	pair models.Pair
}

type Collector struct {
	pairRepo        repositories.PairRepository
	exchangeRepo    repositories.ExchangeRepository
	quoteRepo       repositories.QuoteRepository
	rateCache       repositories.CacheRateRepository
	exchangeFactory *factory.ExchangeFactory
	CollectorConfig config.CollectorConfig
}

func NewCollector(pairRepo repositories.PairRepository, exchangeRepo repositories.ExchangeRepository,
	quoteRepo repositories.QuoteRepository, rateCache repositories.CacheRateRepository,
	exchangeFactory *factory.ExchangeFactory, collectorConfig config.CollectorConfig) *Collector {
	return &Collector{
		pairRepo:        pairRepo,
		exchangeRepo:    exchangeRepo,
		quoteRepo:       quoteRepo,
		rateCache:       rateCache,
		exchangeFactory: exchangeFactory,
		CollectorConfig: collectorConfig,
	}
}

func (c *Collector) RunCycle(ctx context.Context) {
	pairs, err := c.pairRepo.GetAll(ctx)
	if err != nil {
		log.Printf("collector: load pairs failed: %s\n", err)
		return
	}

	exNames, err := c.exchangeRepo.GetEnabled(ctx)
	if err != nil {
		log.Printf("collector: load pairs failed: %s\n", err)
		return
	}

	var exs []exchanges.Exchange
	for _, name := range exNames {
		ex := c.exchangeFactory.Build(name)
		if ex != nil {
			exs = append(exs, ex)
		}
	}

	pairMap := make(map[uint32]string)
	for _, p := range pairs {
		pairMap[p.ID] = p.Base + "-" + p.Quote
	}

	jobs := make(chan job, len(pairs)*len(exs))
	results := make(chan models.Quote, len(pairs)*len(exs))

	var wg sync.WaitGroup

	for i := 0; i < c.CollectorConfig.CollectorWorkers; i++ {
		wg.Add(1)

		go c.worker(ctx, jobs, results, &wg)
	}

	go func() {
		defer close(jobs)

		for _, ex := range exs {
			for _, pair := range pairs {
				select {
				case jobs <- job{ex: ex, pair: pair}:
				case <-ctx.Done():
					return
				}
			}
		}
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	var quotes []models.Quote

	for q := range results {
		quotes = append(quotes, q)
	}

	if len(quotes) == 0 {
		return
	}

	if err := c.quoteRepo.SaveBatch(ctx, quotes); err != nil {
		log.Printf("save batch error: %v\n", err)
		return
	}

	if err := c.rateCache.SaveCurrentRates(ctx, quotes, pairMap); err != nil {
		log.Printf("redis save error: %v\n", err)
	}
}

func (c *Collector) worker(ctx context.Context, jobs <-chan job, results chan<- models.Quote, wg *sync.WaitGroup) {
	defer wg.Done()

	for j := range jobs {
		quote, err := j.ex.Fetch(ctx, j.pair)
		if err != nil {
			log.Printf("[%s] error: %v", j.ex.Name(), err)
			continue
		}

		select {
		case results <- quote:
		case <-ctx.Done():
			return
		}
	}
}
