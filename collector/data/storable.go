package data

type Storable interface {
	Set(float64)
	Get() float64
}
