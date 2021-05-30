package bitwise_test

import (
	"github.com/SarthakMakhija/data-structures-and-algorithms/bitwise"
	"testing"
)

func TestIsKthDigitSet_1(t *testing.T) {
	number := 5
	result := bitwise.IsKthBitSet(3, number)

	if result != true {
		t.Fatalf("Expected true received %v", result)
	}
}

func TestIsKthDigitSet_2(t *testing.T) {
	number := 5
	result := bitwise.IsKthBitSet(1, number)

	if result != true {
		t.Fatalf("Expected true received %v", result)
	}
}

func TestIsKthDigitSet_3(t *testing.T) {
	number := 10
	result := bitwise.IsKthBitSet(4, number)

	if result != true {
		t.Fatalf("Expected true received %v", result)
	}
}

func TestIsKthDigitSet_4(t *testing.T) {
	number := 10
	result := bitwise.IsKthBitSet(3, number)

	if result != false {
		t.Fatalf("Expected false received %v", result)
	}
}

func TestIsKthDigitSetAlternate_1(t *testing.T) {
	number := 5
	result := bitwise.IsKthBitSetAlternate(3, number)

	if result != true {
		t.Fatalf("Expected true received %v", result)
	}
}

func TestIsKthDigitSetAlternate_2(t *testing.T) {
	number := 5
	result := bitwise.IsKthBitSetAlternate(1, number)

	if result != true {
		t.Fatalf("Expected true received %v", result)
	}
}

func TestIsKthDigitSetAlternate_3(t *testing.T) {
	number := 10
	result := bitwise.IsKthBitSetAlternate(4, number)

	if result != true {
		t.Fatalf("Expected true received %v", result)
	}
}

func TestIsKthDigitSetAlternate_4(t *testing.T) {
	number := 10
	result := bitwise.IsKthBitSetAlternate(3, number)

	if result != false {
		t.Fatalf("Expected false received %v", result)
	}
}

func TestPowerOf2By_1(t *testing.T) {
	result := bitwise.PowerOf2By(3)

	if result != 8 {
		t.Fatalf("Expected %v received %v", 8, result)
	}
}

func TestPowerOf2By_2(t *testing.T) {
	result := bitwise.PowerOf2By(10)

	if result != 1024 {
		t.Fatalf("Expected %v received %v", 1024, result)
	}
}

func TestSetBitsCount_1(t *testing.T) {
	result := bitwise.SetBitsCountIn(10)

	if result != 2 {
		t.Fatalf("Expected %v received %v", 2, result)
	}
}

func TestSetBitsCount_2(t *testing.T) {
	result := bitwise.SetBitsCountIn(0)

	if result != 0 {
		t.Fatalf("Expected %v received %v", 0, result)
	}
}

func TestSetBitsCount_3(t *testing.T) {
	result := bitwise.SetBitsCountIn(7)

	if result != 3 {
		t.Fatalf("Expected %v received %v", 3, result)
	}
}

func TestSetBitsCount_4(t *testing.T) {
	result := bitwise.SetBitsCountIn(1)

	if result != 1 {
		t.Fatalf("Expected %v received %v", 1, result)
	}
}

func TestSetBitsCount_5(t *testing.T) {
	result := bitwise.SetBitsCountIn(100)

	if result != 3 {
		t.Fatalf("Expected %v received %v", 3, result)
	}
}
