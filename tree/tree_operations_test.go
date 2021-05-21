package tree_test

import (
	"testing"

	"example.com/data-structures-and-algorithms/tree"
)

func TestTreeTraversal_1(t *testing.T) {
	tree := binary_tree_1()
	output := tree.Traverse()
	if output != "DBEAFCG" {
		t.Fatalf("Expected DBEAFCG, received %v", output)
	}
}

func TestTreeTraversal_2(t *testing.T) {
	leaf_D := tree.Node{
		Value: "D",
	}
	tree := tree.BinaryTree{
		Root: &leaf_D,
	}

	output := tree.Traverse()
	if output != "D" {
		t.Fatalf("Expected D, received %v", output)
	}
}

func TestTreeTraversal_3(t *testing.T) {
	tree := binary_tree_2()
	output := tree.Traverse()
	if output != "QDBEAFCG" {
		t.Fatalf("Expected QDBEAFCG, received %v", output)
	}
}

func binary_tree_1() *tree.BinaryTree {

	leaf_D := tree.Node{
		Value: "D",
	}
	leaf_E := tree.Node{
		Value: "E",
	}
	leaf_B := tree.Node{
		Value: "B",
		Left:  &leaf_D,
		Right: &leaf_E,
	}

	leaf_F := tree.Node{
		Value: "F",
	}
	leaf_G := tree.Node{
		Value: "G",
	}
	leaf_C := tree.Node{
		Value: "C",
		Left:  &leaf_F,
		Right: &leaf_G,
	}

	leaf_A := tree.Node{
		Value: "A",
		Left:  &leaf_B,
		Right: &leaf_C,
	}
	return &tree.BinaryTree{
		Root: &leaf_A,
	}
}

func binary_tree_2() *tree.BinaryTree {

	leaf_Q := tree.Node{
		Value: "Q",
	}
	leaf_D := tree.Node{
		Value: "D",
		Left:  &leaf_Q,
	}
	leaf_E := tree.Node{
		Value: "E",
	}
	leaf_B := tree.Node{
		Value: "B",
		Left:  &leaf_D,
		Right: &leaf_E,
	}

	leaf_F := tree.Node{
		Value: "F",
	}
	leaf_G := tree.Node{
		Value: "G",
	}
	leaf_C := tree.Node{
		Value: "C",
		Left:  &leaf_F,
		Right: &leaf_G,
	}

	leaf_A := tree.Node{
		Value: "A",
		Left:  &leaf_B,
		Right: &leaf_C,
	}
	return &tree.BinaryTree{
		Root: &leaf_A,
	}
}
