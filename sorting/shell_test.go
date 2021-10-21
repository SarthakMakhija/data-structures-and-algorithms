package sorting_test

import (
	"github.com/SarthakMakhija/data-structures-and-algorithms/sorting"
	"reflect"
	"testing"
)

func TestSortsUsingShellSort1(t *testing.T) {
	elements := []int{11, 13, 7, 2, 6, 9, 4, 5, 10, 3}
	shell := sorting.Shell{Elements: elements}
	shell.Sort()

	expected := []int{2, 3, 4, 5, 6, 7, 9, 10, 11, 13}
	if !reflect.DeepEqual(expected, shell.Elements) {
		t.Fatalf("Expected %v, received %v", expected, shell.Elements)
	}
}

func TestSortsUsingShellSort2(t *testing.T) {
	elements := []int{9, 5, 16, 8, 13, 6, 12, 10, 4, 2, 3}
	shell := sorting.Shell{Elements: elements}
	shell.Sort()

	expected := []int{2, 3, 4, 5, 6, 8, 9, 10, 12, 13, 16}
	if !reflect.DeepEqual(expected, shell.Elements) {
		t.Fatalf("Expected %v, received %v", expected, shell.Elements)
	}
}

func TestSortsUsingShellSortWithDuplicates(t *testing.T) {
	elements := []int{6, 3, 9, 10, 15, 6, 8, 12, 3, 6}
	shell := sorting.Shell{Elements: elements}
	shell.Sort()

	expected := []int{3, 3, 6, 6, 6, 8, 9, 10, 12, 15}
	if !reflect.DeepEqual(expected, shell.Elements) {
		t.Fatalf("Expected %v, received %v", expected, shell.Elements)
	}
}
