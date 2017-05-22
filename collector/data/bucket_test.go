package data

import "testing"

func TestBucketAggregateEmpty(t *testing.T) {
	// Mathematically incorrect

	var bucket = Bucket{
		Name:   "Test",
		Values: &[]float64{},
	}

	if avg := bucket.Aggregate(); avg != 0 {
		t.Error("Expected average value of 0, got ", avg)
	}
}

func TestBucketAggregate(t *testing.T) {
	var bucket = Bucket{
		Name: "Test",
		Values: &[]float64{
			1.1,
			2.2,
			3.3,
			4.4,
			5.5,
		},
	}

	if avg := bucket.Aggregate(); avg != 3.3 {
		t.Error("Expected average value of 3.3, got ", avg)
	}
}

func TestBucketAppend(t *testing.T) {
	var value = 1234.5678
	var bucket Bucket = Bucket{
		Name:   "Test",
		Values: &[]float64{},
	}

	if l := len(*bucket.Values); l != 0 {
		t.Error("Expected empty values, got ", l)
	}

	bucket.Append(value)

	if l := len(*bucket.Values); l != 1 {
		t.Error("Expected empty values, got ", l)
	}
}
