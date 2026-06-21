package scheduler

import (
	"context"
	"log"
	"market-data-collector/internal/collector"
	"sync"
	"time"
)

type Scheduler struct {
	collector *collector.Collector
	wg        sync.WaitGroup
}

func NewScheduler(collector *collector.Collector) *Scheduler {
	return &Scheduler{collector: collector}
}

func (s *Scheduler) Start(ctx context.Context, tickSeconds int) {
	s.wg.Add(1)

	go func() {
		defer s.wg.Done()

		ticker := time.NewTicker(time.Duration(tickSeconds) * time.Second)
		defer ticker.Stop()

		s.run(ctx)

		for {
			select {
			case <-ctx.Done():
				log.Println("scheduler stopped")
				return

			case <-ticker.C:
				s.run(ctx)
			}
		}
	}()
}

func (s *Scheduler) Wait() {
	s.wg.Wait()
}

func (s *Scheduler) run(ctx context.Context) {
	start := time.Now()

	log.Println("📡 running collector cycle...")

	s.collector.RunCycle(ctx)

	log.Printf("✅ cycle finished in %s\n", time.Since(start))
}
