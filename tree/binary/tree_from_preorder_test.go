package binarytree_test

import (
	binarytree "github.com/SarthakMakhija/data-structures-and-algorithms/tree/binary"
	"testing"
)

func TestBinaryTreeFromPreorderTraversalWithMarkers1(t *testing.T) {
	traversal := []string{"H", "B", "F", "", "", "E", "A", "", "", "", "C", "", "D", "", "G", "I", "", "", ""}
	tree := binarytree.FromPreOrderTraversalWithMarkers(traversal)

	output := tree.Preorder()
	expected := "HBFEACDGI"
	if output != expected {

		t.Fatalf("Expected %v, received %v", expected, output)
	}
}

func TestBinaryTreeFromPreorderTraversalWithMarkers2(t *testing.T) {
	traversal := []string{"A", "B", "", "", "C", "", ""}
	tree := binarytree.FromPreOrderTraversalWithMarkers(traversal)

	output := tree.Preorder()
	expected := "ABC"
	if output != expected {

		t.Fatalf("Expected %v, received %v", expected, output)
	}
}

func TestBinaryTreeFromPreorderTraversalWithMarkers3(t *testing.T) {
	traversal := []string{"A", "B", "E", "", "", "H", "", "L", "C", "P", "", "", "Q", "", ""}
	tree := binarytree.FromPreOrderTraversalWithMarkers(traversal)

	output := tree.Preorder()
	expected := "ABEHLCPQ"
	if output != expected {

		t.Fatalf("Expected %v, received %v", expected, output)
	}
}
