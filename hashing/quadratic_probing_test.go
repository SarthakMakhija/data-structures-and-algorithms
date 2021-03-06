package hashing_test

import (
	"github.com/SarthakMakhija/data-structures-and-algorithms/hashing"
	"testing"
)

func TestQuadraticProbingContains_1(t *testing.T) {
	bucket := hashing.QuadraticProbingFixedSizedBucket{}

	_, _ = bucket.Add(47)
	_, _ = bucket.Add(57)
	_, _ = bucket.Add(67)
	_, _ = bucket.Add(77)

	result := bucket.Contains(77)
	if result != true {
		t.Fatalf("Expected %v, received %v", true, result)
	}
}

func TestQuadraticProbingContains_2(t *testing.T) {
	bucket := hashing.QuadraticProbingFixedSizedBucket{}

	_, _ = bucket.Add(47)
	_, _ = bucket.Add(57)
	_, _ = bucket.Add(67)
	_, _ = bucket.Add(77)
	_, _ = bucket.Add(87)
	_, _ = bucket.Add(97)

	result := bucket.Contains(97)
	if result != true {
		t.Fatalf("Expected %v, received %v", true, result)
	}
}

func TestQuadraticProbingContains_3(t *testing.T) {
	bucket := hashing.QuadraticProbingFixedSizedBucket{}

	_, _ = bucket.Add(47)
	_, _ = bucket.Add(57)
	_, _ = bucket.Add(67)
	_, _ = bucket.Add(77)
	_, _ = bucket.Add(87)
	_, _ = bucket.Add(97)

	result := bucket.Contains(37)
	if result != false {
		t.Fatalf("Expected %v, received %v", false, result)
	}
}

func TestQuadraticProbingAdd(t *testing.T) {
	bucket := hashing.QuadraticProbingFixedSizedBucket{}

	_, _ = bucket.Add(1)
	_, _ = bucket.Add(2)
	_, _ = bucket.Add(3)
	_, _ = bucket.Add(4)
	_, _ = bucket.Add(5)
	_, _ = bucket.Add(6)
	_, _ = bucket.Add(7)
	_, _ = bucket.Add(8)
	_, _ = bucket.Add(9)
	_, _ = bucket.Add(10)

	_, err := bucket.Add(11)

	if err == nil {
		t.Fatalf("Expected an error received no error")
	}
}
