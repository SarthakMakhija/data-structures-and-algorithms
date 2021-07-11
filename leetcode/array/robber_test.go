package array_test

import (
	"github.com/SarthakMakhija/data-structures-and-algorithms/leetcode/array"
	"testing"
)

func TestMaxAmountToBeRobbed1(t *testing.T) {
	maxAmountThatCanBeRobbed := array.MaxAmountToBeRobbed([]int{1, 2, 3, 1})
	if maxAmountThatCanBeRobbed != 4 {
		t.Fatalf("Expected 4, recevied %v", maxAmountThatCanBeRobbed)
	}
}

func TestMaxAmountToBeRobbed2(t *testing.T) {
	maxAmountThatCanBeRobbed := array.MaxAmountToBeRobbed([]int{2, 7, 9, 3, 1})
	if maxAmountThatCanBeRobbed != 12 {
		t.Fatalf("Expected 12, recevied %v", maxAmountThatCanBeRobbed)
	}
}
