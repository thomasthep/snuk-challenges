package data

import (
	"github.com/thomasthep/snuk-challenges/collector/math"
)

type Bucket struct {
	Name   string
	Values *[]float64
}

func (b *Bucket) Aggregate() float64 {
	if len(*b.Values) == 0 {
		return 0
	}

	mean, _ := math.Mean(*b.Values)

	return mean
}

func (b *Bucket) Append(value float64) {
	*b.Values = append(*b.Values, value)
}
