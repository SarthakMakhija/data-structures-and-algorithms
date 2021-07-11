package general_test

import (
	"github.com/SarthakMakhija/data-structures-and-algorithms/leetcode/general"
	"testing"
)

func TestLargestNumberFrom1(t *testing.T) {
	largestNumberAsString := general.LargestNumberFrom([]int{10, 2})
	expected := "210"

	if largestNumberAsString != expected {
		t.Fatalf("Expected %v, recevied %v", expected, largestNumberAsString)
	}
}

func TestLargestNumberFrom2(t *testing.T) {
	largestNumberAsString := general.LargestNumberFrom([]int{3, 30, 34, 5, 9})
	expected := "9534330"

	if largestNumberAsString != expected {
		t.Fatalf("Expected %v, recevied %v", expected, largestNumberAsString)
	}
}

func TestLargestNumberFrom3(t *testing.T) {
	largestNumberAsString := general.LargestNumberFrom([]int{30, 34, 3, 5, 9})
	expected := "9534330"

	if largestNumberAsString != expected {
		t.Fatalf("Expected %v, recevied %v", expected, largestNumberAsString)
	}
}

func TestLargestNumberFrom4(t *testing.T) {
	largestNumberAsString := general.LargestNumberFrom([]int{3, 2, 1, 5, 6, 8, 90, 80})
	expected := "9088065321"

	if largestNumberAsString != expected {
		t.Fatalf("Expected %v, recevied %v", expected, largestNumberAsString)
	}
}
