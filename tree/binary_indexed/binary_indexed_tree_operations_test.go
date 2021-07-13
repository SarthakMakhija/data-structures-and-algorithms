package binary_indexed_test

import (
	"fmt"
	"github.com/SarthakMakhija/data-structures-and-algorithms/tree/binary_indexed"
	"testing"
)

func TestBinaryIndexedTreeFindSumOfFirst1(t *testing.T) {
	elements := []int{1, 2, 3, 4, 5, 6}
	tree := binary_indexed.New(elements)

	fmt.Println(tree.Tree)

	sum := tree.FindSumOfFirst(5)
	expected := 15

	if sum != expected {
		t.Fatalf("Expected %v, received %v", expected, sum)
	}
}

func TestBinaryIndexedTreeFindSumOfFirst2(t *testing.T) {
	elements := []int{1, 2, 3, 4, 5, 6}
	tree := binary_indexed.New(elements)

	fmt.Println(tree.Tree)

	sum := tree.FindSumOfFirst(6)
	expected := 21

	if sum != expected {
		t.Fatalf("Expected %v, received %v", expected, sum)
	}
}

func TestBinaryIndexedTreeFindSumOfFirst3(t *testing.T) {
	elements := []int{1, 2, -3, -4, 5, 6, 9}
	tree := binary_indexed.New(elements)

	fmt.Println(tree.Tree)

	sum := tree.FindSumOfFirst(3)
	expected := 0

	if sum != expected {
		t.Fatalf("Expected %v, received %v", expected, sum)
	}
}

func TestBinaryIndexedTreeFindSumOfFirst4(t *testing.T) {
	elements := []int{1, 2, -3, -4, 5, 6, 9}
	tree := binary_indexed.New(elements)

	fmt.Println(tree.Tree)

	sum := tree.FindSumOfFirst(5)
	expected := 1

	if sum != expected {
		t.Fatalf("Expected %v, received %v", expected, sum)
	}
}

func TestBinaryIndexedTreeFindSumOfFirst5(t *testing.T) {
	elements := []int{1, 2, -3, -4, 5, 6, 9}
	tree := binary_indexed.New(elements)

	fmt.Println(tree.Tree)

	sum := tree.FindSumOfFirst(7)
	expected := 16

	if sum != expected {
		t.Fatalf("Expected %v, received %v", expected, sum)
	}
}
