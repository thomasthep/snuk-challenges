package aggregator

import (
	"time"
)

func (a Aggregator) Start(interval time.Duration) {
	a.Ticker = time.NewTicker(interval)

	go func() {
		for {
			select {
			case <-a.Ticker.C:
				a.Storage.Set(a.Bucket.Aggregate())
			}
		}
	}()
}

func (a Aggregator) Stop() {
	a.Ticker.Stop()
}
