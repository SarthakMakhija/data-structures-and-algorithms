package backtracking_test

import (
	"github.com/SarthakMakhija/data-structures-and-algorithms/backtracking"
	"reflect"
	"sort"
	"testing"
)

func TestAllPossibleWaysToArrangeChildrenInChairs1(t *testing.T) {
	children := []string{"B1", "B2", "B3"}
	arrangements := backtracking.AllPossibleWaysToArrangeChildrenInChairs(children)
	sort.Strings(arrangements)

	expected := []string{"B1B2B3", "B1B3B2", "B2B1B3", "B2B3B1", "B3B1B2", "B3B2B1"}
	sort.Strings(expected)

	if !reflect.DeepEqual(arrangements, expected) {
		t.Fatalf("Expecpted %v, received %v", expected, arrangements)
	}
}

func TestAllPossibleWaysToArrangeChildrenInChairs2(t *testing.T) {
	children := []string{"B1", "B2", "B3", "B4"}
	arrangements := backtracking.AllPossibleWaysToArrangeChildrenInChairs(children)
	sort.Strings(arrangements)

	var expected = []string{"B1B2B3B4", "B1B2B4B3", "B1B3B2B4", "B1B3B4B2", "B1B4B3B2", "B1B4B2B3",
		"B2B1B3B4", "B2B1B4B3", "B2B3B1B4", "B2B3B4B1", "B2B4B3B1", "B2B4B1B3",
		"B3B2B1B4", "B3B2B4B1", "B3B1B2B4", "B3B1B4B2", "B3B4B1B2", "B3B4B2B1",
		"B4B2B3B1", "B4B2B1B3", "B4B3B2B1", "B4B3B1B2", "B4B1B3B2", "B4B1B2B3"}
	sort.Strings(expected)

	if !reflect.DeepEqual(arrangements, expected) {
		t.Fatalf("Expecpted %v, received %v", expected, arrangements)
	}
}
