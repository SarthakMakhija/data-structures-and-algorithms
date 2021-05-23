package binarysearchtree_test

import (
	"testing"

	binarysearchtree "example.com/data-structures-and-algorithms/tree/binary_search_tree"
)

func TestSearch_1(t *testing.T) {
	tree := binary_search_tree_1()
	contains := tree.Search(21)

	if contains != true {
		t.Fatalf("Expected true bit received %v", contains)
	}
}

func TestSearch_2(t *testing.T) {
	tree := binary_search_tree_1()
	contains := tree.Search(40)

	if contains != true {
		t.Fatalf("Expected true bit received %v", contains)
	}
}

func TestSearch_3(t *testing.T) {
	tree := binary_search_tree_1()
	contains := tree.Search(41)

	if contains != false {
		t.Fatalf("Expected false bit received %v", contains)
	}
}

func binary_search_tree_1() *binarysearchtree.IntBinarySearchTree {
	node_10 := binarysearchtree.IntNode{
		Value: 10,
	}
	node_21 := binarysearchtree.IntNode{
		Value: 21,
	}
	node_20 := binarysearchtree.IntNode{
		Value: 20,
		Left:  &node_10,
		Right: &node_21,
	}
	node_40 := binarysearchtree.IntNode{
		Value: 40,
	}

	node_30 := binarysearchtree.IntNode{
		Value: 30,
		Left:  &node_20,
		Right: &node_40,
	}
	return &binarysearchtree.IntBinarySearchTree{
		Root: &node_30,
	}
}
