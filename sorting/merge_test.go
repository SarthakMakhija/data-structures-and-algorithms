package sorting_test

import (
	"github.com/SarthakMakhija/data-structures-and-algorithms/sorting"
	"reflect"
	"testing"
)

func TestMergeSortedSlices1(t *testing.T) {
	elements := []int{8, 9, 10, 11, 2, 3, 6, 7}
	merge := sorting.Merge{Elements: elements}

	merge.MergeSorted(0, 3, 7)
	expected := []int{2, 3, 6, 7, 8, 9, 10, 11}

	if !reflect.DeepEqual(expected, merge.Elements) {
		t.Fatalf("Expected %v, received %v", expected, merge.Elements)
	}
}

func TestMergeSortedSlices2(t *testing.T) {
	elements := []int{8, 9, 10, 11, 2, 3, 6, 7}
	merge := sorting.Merge{Elements: elements}

	merge.MergeSorted(0, 2, 4)
	expected := []int{8, 9, 10, 11, 2, 3, 6, 7}

	if !reflect.DeepEqual(expected, merge.Elements) {
		t.Fatalf("Expected %v, received %v", expected, merge.Elements)
	}
}

func TestSortsUsingMergeSort1(t *testing.T) {
	elements := []int{8, 3, 7, 4, 9, 2, 6, 5}
	merge := sorting.Merge{Elements: elements}
	merge.Sort()

	expected := []int{2, 3, 4, 5, 6, 7, 8, 9}
	if !reflect.DeepEqual(expected, merge.Elements) {
		t.Fatalf("Expected %v, received %v", expected, merge.Elements)
	}
}

func TestSortsUsingMergeSort2(t *testing.T) {
	elements := []int{11, 13, 7, 2, 6, 9, 4, 5, 10, 3}
	merge := sorting.Merge{Elements: elements}
	merge.Sort()

	expected := []int{2, 3, 4, 5, 6, 7, 9, 10, 11, 13}
	if !reflect.DeepEqual(expected, merge.Elements) {
		t.Fatalf("Expected %v, received %v", expected, merge.Elements)
	}
}

func TestSortsUsingMergeSort3(t *testing.T) {
	elements := []int{11, 13, 7, 2, 6, 9, 4}
	merge := sorting.Merge{Elements: elements}
	merge.Sort()

	expected := []int{2, 4, 6, 7, 9, 11, 13}
	if !reflect.DeepEqual(expected, merge.Elements) {
		t.Fatalf("Expected %v, received %v", expected, merge.Elements)
	}
}
