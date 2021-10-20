package sorting_test

import (
	"github.com/SarthakMakhija/data-structures-and-algorithms/sorting"
	"reflect"
	"testing"
)

func TestSortsUsingBucketSort(t *testing.T) {
	elements := []int{11, 13, 7, 2, 6, 9, 4, 5, 10, 3}
	bucket := sorting.Bucket{Elements: elements}
	bucket.Sort()

	expected := []int{2, 3, 4, 5, 6, 7, 9, 10, 11, 13}
	if !reflect.DeepEqual(expected, bucket.Elements) {
		t.Fatalf("Expected %v, received %v", expected, bucket.Elements)
	}
}

func TestSortsUsingBucketSortWithDuplicates(t *testing.T) {
	elements := []int{6, 3, 9, 10, 15, 6, 8, 12, 3, 6}
	bucket := sorting.Bucket{Elements: elements}
	bucket.Sort()

	expected := []int{3, 3, 6, 6, 6, 8, 9, 10, 12, 15}
	if !reflect.DeepEqual(expected, bucket.Elements) {
		t.Fatalf("Expected %v, received %v", expected, bucket.Elements)
	}
}
