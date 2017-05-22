package data

type Aggregatable interface {
	Aggregate() float64
	Append(float64)
}
