package sum_test

import (
	"github.com/SarthakMakhija/data-structures-and-algorithms/leetcode/sum"
	"testing"
)

func TestAdd1(t *testing.T) {
	result := sum.Add(45, 97)
	if result != 142 {
		t.Fatalf("Expected %v, found %v", 45+97, result)
	}
}

func TestAdd2(t *testing.T) {
	result := sum.Add(10, 13)
	if result != 23 {
		t.Fatalf("Expected %v, found %v", 10+13, result)
	}
}

func TestAdd3(t *testing.T) {
	result := sum.Add(45, 99)
	if result != 144 {
		t.Fatalf("Expected %v, found %v", 45+99, result)
	}
}

func TestAdd4(t *testing.T) {
	result := sum.Add(1054, 2056)
	if result != 3110 {
		t.Fatalf("Expected %v, found %v", 1054+2056, result)
	}
}
