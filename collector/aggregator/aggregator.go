package aggregator

import (
	"time"

	"github.com/thomasthep/snuk-challenges/collector/data"
)

type Tickable interface {
	Start()
	Stop()
}

type Aggregator struct {
	Ticker  *time.Ticker
	Storage *data.Storage
	Bucket  *data.Bucket
}
