package recursion_test

import (
	"github.com/SarthakMakhija/data-structures-and-algorithms/recursion"
	"testing"
)

func TestPossibleWaysToTravelEndOfGridSized_1(t *testing.T) {
	possibleWays := recursion.PossibleWaysToTravelEndOfGridSized(2, 3)
	if possibleWays != 3 {
		t.Fatalf("Expecting possible ways to be 3, got %v", possibleWays)
	}
}

func TestPossibleWaysToTravelEndOfGridSized_2(t *testing.T) {
	possibleWays := recursion.PossibleWaysToTravelEndOfGridSized(2, 2)
	if possibleWays != 2 {
		t.Fatalf("Expecting possible ways to be 2, got %v", possibleWays)
	}
}

func TestPossibleWaysToTravelEndOfGridSized_3(t *testing.T) {
	possibleWays := recursion.PossibleWaysToTravelEndOfGridSized(3, 3)
	if possibleWays != 6 {
		t.Fatalf("Expecting possible ways to be 6, got %v", possibleWays)
	}
}
