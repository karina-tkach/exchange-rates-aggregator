package scheduler

import (
	"log"
	"market-data-collector/internal/collector"
	"time"
)

type Scheduler struct {
	collector *collector.Collector
}

func NewScheduler(collector *collector.Collector) *Scheduler {
	return &Scheduler{collector: collector}
}

func (s *Scheduler) Start(tickSeconds int) {
	ticker := time.NewTicker(time.Duration(tickSeconds) * time.Second)
	defer ticker.Stop()

	s.run()

	for range ticker.C {
		s.run()
	}
}

func (s *Scheduler) run() {
	start := time.Now()

	log.Println("📡 running collector cycle...")

	s.collector.RunCycle()

	log.Printf("✅ cycle finished in %s\n", time.Since(start))
}
