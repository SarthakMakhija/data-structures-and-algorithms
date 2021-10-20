package sorting_test

import (
	"github.com/SarthakMakhija/data-structures-and-algorithms/sorting"
	"reflect"
	"testing"
)

func TestSortsUsingCountSort(t *testing.T) {
	elements := []int{11, 13, 7, 2, 6, 9, 4, 5, 10, 3}
	count := sorting.Count{Elements: elements}
	count.Sort()

	expected := []int{2, 3, 4, 5, 6, 7, 9, 10, 11, 13}
	if !reflect.DeepEqual(expected, count.Elements) {
		t.Fatalf("Expected %v, received %v", expected, count.Elements)
	}
}
