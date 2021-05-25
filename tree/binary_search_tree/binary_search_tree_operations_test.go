package binarysearchtree_test

import (
	"testing"

	binarysearchtree "github.com/SarthakMakhija/data-structures-and-algorithms/tree/binary_search_tree"
)

func TestSearch_1(t *testing.T) {
	tree := binarySearchTree1()
	contains := tree.Search(21)

	if contains != true {
		t.Fatalf("Expected true bit received %v", contains)
	}
}

func TestSearch_2(t *testing.T) {
	tree := binarySearchTree1()
	contains := tree.Search(40)

	if contains != true {
		t.Fatalf("Expected true bit received %v", contains)
	}
}

func TestSearch_3(t *testing.T) {
	tree := binarySearchTree1()
	contains := tree.Search(41)

	if contains != false {
		t.Fatalf("Expected false bit received %v", contains)
	}
}

func TestInorderTraversal_1(t *testing.T) {
	tree := binarySearchTree1()
	output := tree.InOrder_Traversal()

	if output != "1020213040" {
		t.Fatalf("Expected 1020213040 bit received %v", output)
	}
}

func TestInsert_1(t *testing.T) {
	tree := binarySearchTree1()
	tree.Insert1(41)

	output := tree.InOrder_Traversal()
	if output != "102021304041" {
		t.Fatalf("Expected 102021304041 bit received %v", output)
	}
}

func TestInsert_2(t *testing.T) {
	tree := binarySearchTree1()
	tree.Insert1(31)

	output := tree.InOrder_Traversal()
	if output != "102021303140" {
		t.Fatalf("Expected 102021303140 bit received %v", output)
	}
}

func TestInsert_3(t *testing.T) {
	tree := binarySearchTree1()
	tree.Insert1(21)

	output := tree.InOrder_Traversal()
	if output != "1020213040" {
		t.Fatalf("Expected 1020213040 bit received %v", output)
	}
}

func TestInsert_4(t *testing.T) {
	tree := binarySearchTree1()
	tree.Insert2(41)

	output := tree.InOrder_Traversal()
	if output != "102021304041" {
		t.Fatalf("Expected 102021304041 bit received %v", output)
	}
}

func TestInsert_5(t *testing.T) {
	tree := binarySearchTree1()
	tree.Insert2(31)

	output := tree.InOrder_Traversal()
	if output != "102021303140" {
		t.Fatalf("Expected 102021303140 bit received %v", output)
	}
}

func TestInsert_6(t *testing.T) {
	tree := binarySearchTree1()
	tree.Insert2(21)

	output := tree.InOrder_Traversal()
	if output != "1020213040" {
		t.Fatalf("Expected 1020213040 bit received %v", output)
	}
}

func binarySearchTree1() *binarysearchtree.IntBinarySearchTree {
	node10 := binarysearchtree.IntNode{
		Value: 10,
	}
	node21 := binarysearchtree.IntNode{
		Value: 21,
	}
	node20 := binarysearchtree.IntNode{
		Value: 20,
		Left:  &node10,
		Right: &node21,
	}
	node40 := binarysearchtree.IntNode{
		Value: 40,
	}

	node30 := binarysearchtree.IntNode{
		Value: 30,
		Left:  &node20,
		Right: &node40,
	}
	return &binarysearchtree.IntBinarySearchTree{
		Root: &node30,
	}
}
