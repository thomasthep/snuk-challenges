package math

import (
	"errors"
	goMath "math"
)

var errInputEmpty = errors.New("input must not be empty")

func Mean(input []float64) (float64, error) {
	if len(input) == 0 {
		return goMath.NaN(), errInputEmpty
	}

	sum, _ := Sum(input)

	return sum / float64(len(input)), nil
}

func Sum(input []float64) (sum float64, err error) {
	if len(input) == 0 {
		return goMath.NaN(), errInputEmpty
	}

	for _, n := range input {
		sum += n
	}

	return sum, nil
}
