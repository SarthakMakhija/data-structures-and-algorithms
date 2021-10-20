package sorting_test

import (
	"github.com/SarthakMakhija/data-structures-and-algorithms/sorting"
	"reflect"
	"testing"
)

func TestSortsUsingRadixSort1(t *testing.T) {
	elements := []int{11, 13, 7, 2, 6, 9, 4, 5, 10, 3}
	radix := sorting.Radix{Elements: elements}
	radix.Sort()

	expected := []int{2, 3, 4, 5, 6, 7, 9, 10, 11, 13}
	if !reflect.DeepEqual(expected, radix.Elements) {
		t.Fatalf("Expected %v, received %v", expected, radix.Elements)
	}
}

func TestSortsUsingRadixSort2(t *testing.T) {
	elements := []int{237, 146, 259, 348, 152, 163, 235, 48, 36, 62}
	radix := sorting.Radix{Elements: elements}
	radix.Sort()

	expected := []int{36, 48, 62, 146, 152, 163, 235, 237, 259, 348}
	if !reflect.DeepEqual(expected, radix.Elements) {
		t.Fatalf("Expected %v, received %v", expected, radix.Elements)
	}
}

func TestSortsUsingRadixSortWithDuplicates(t *testing.T) {
	elements := []int{6, 3, 9, 10, 15, 6, 8, 12, 3, 6}
	radix := sorting.Radix{Elements: elements}
	radix.Sort()

	expected := []int{3, 3, 6, 6, 6, 8, 9, 10, 12, 15}
	if !reflect.DeepEqual(expected, radix.Elements) {
		t.Fatalf("Expected %v, received %v", expected, radix.Elements)
	}
}
