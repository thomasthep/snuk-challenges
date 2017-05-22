package data

type Storage struct {
	Value float64
}

func (s *Storage) Set(value float64) {
	s.Value = value
}

func (s *Storage) Get() float64 {
	return s.Value
}
