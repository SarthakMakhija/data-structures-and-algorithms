package divide_test

import (
	"github.com/SarthakMakhija/data-structures-and-algorithms/leetcode/divide"
	"testing"
)

func TestDividePositiveDivisor1(t *testing.T) {
	result := divide.Divide(10, 3)
	if result != 3 {
		t.Fatalf("Expected 3, received %v", result)
	}
}

func TestDividePositiveDivisor2(t *testing.T) {
	result := divide.Divide(1, 1)
	if result != 1 {
		t.Fatalf("Expected 1, received %v", result)
	}
}

func TestDividePositiveDivisor3(t *testing.T) {
	result := divide.Divide(0, 3)
	if result != 0 {
		t.Fatalf("Expected 0, received %v", result)
	}
}

func TestDividePositiveDivisor4(t *testing.T) {
	result := divide.Divide(1, 3)
	if result != 0 {
		t.Fatalf("Expected 0, received %v", result)
	}
}

func TestDividePositiveDivisor5(t *testing.T) {
	result := divide.Divide(16, 4)
	if result != 4 {
		t.Fatalf("Expected 4, received %v", result)
	}
}

/////
func TestDivideNegativeDivisor1(t *testing.T) {
	result := divide.Divide(10, -3)
	if result != -3 {
		t.Fatalf("Expected -3, received %v", result)
	}
}

func TestDivideNegativeDivisor2(t *testing.T) {
	result := divide.Divide(1, -1)
	if result != -1 {
		t.Fatalf("Expected -1, received %v", result)
	}
}

func TestDivideNegativeDivisor3(t *testing.T) {
	result := divide.Divide(0, -3)
	if result != 0 {
		t.Fatalf("Expected 0, received %v", result)
	}
}

func TestDivideNegativeDivisor4(t *testing.T) {
	result := divide.Divide(1, -3)
	if result != 0 {
		t.Fatalf("Expected 0, received %v", result)
	}
}

func TestDivideNegativeDivisor5(t *testing.T) {
	result := divide.Divide(16, -4)
	if result != -4 {
		t.Fatalf("Expected -4, received %v", result)
	}
}
