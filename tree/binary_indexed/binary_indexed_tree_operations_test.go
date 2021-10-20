package binary_indexed_test

import (
	"github.com/SarthakMakhija/data-structures-and-algorithms/tree/binary_indexed"
	"testing"
)

func TestBinaryIndexedTreeFindSumOfFirst1(t *testing.T) {
	elements := []int{1, 2, 3, 4, 5, 6}
	tree := binary_indexed.New(elements)

	sum := tree.SumUptoIndex(5)
	expected := 21

	if sum != expected {
		t.Fatalf("Expected %v, received %v", expected, sum)
	}
}

func TestBinaryIndexedTreeSumUptoIndex3(t *testing.T) {
	elements := []int{1, 2, -3, -4, 5, 6, 9}
	tree := binary_indexed.New(elements)

	sum := tree.SumUptoIndex(3)
	expected := -4

	if sum != expected {
		t.Fatalf("Expected %v, received %v", expected, sum)
	}
}

func TestBinaryIndexedTreeSumUptoIndex4(t *testing.T) {
	elements := []int{1, 2, -3, -4, 5, 6, 9}
	tree := binary_indexed.New(elements)

	sum := tree.SumUptoIndex(4)
	expected := 1

	if sum != expected {
		t.Fatalf("Expected %v, received %v", expected, sum)
	}
}

func TestBinaryIndexedTreeSumUptoIndex5(t *testing.T) {
	elements := []int{1, 2, -3, -4, 5, 6, 9}
	tree := binary_indexed.New(elements)

	sum := tree.SumUptoIndex(5)
	expected := 7

	if sum != expected {
		t.Fatalf("Expected %v, received %v", expected, sum)
	}
}
