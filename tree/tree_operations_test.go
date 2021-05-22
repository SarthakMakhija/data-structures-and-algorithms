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
	leaf_D := tree.StringNode{
		Value: "D",
	}
	tree := tree.StringBinaryTree{
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

func binary_tree_1() *tree.StringBinaryTree {

	leaf_D := tree.StringNode{
		Value: "D",
	}
	leaf_E := tree.StringNode{
		Value: "E",
	}
	leaf_B := tree.StringNode{
		Value: "B",
		Left:  &leaf_D,
		Right: &leaf_E,
	}

	leaf_F := tree.StringNode{
		Value: "F",
	}
	leaf_G := tree.StringNode{
		Value: "G",
	}
	leaf_C := tree.StringNode{
		Value: "C",
		Left:  &leaf_F,
		Right: &leaf_G,
	}

	leaf_A := tree.StringNode{
		Value: "A",
		Left:  &leaf_B,
		Right: &leaf_C,
	}
	return &tree.StringBinaryTree{
		Root: &leaf_A,
	}
}

func binary_tree_2() *tree.StringBinaryTree {

	leaf_Q := tree.StringNode{
		Value: "Q",
	}
	leaf_D := tree.StringNode{
		Value: "D",
		Left:  &leaf_Q,
	}
	leaf_E := tree.StringNode{
		Value: "E",
	}
	leaf_B := tree.StringNode{
		Value: "B",
		Left:  &leaf_D,
		Right: &leaf_E,
	}

	leaf_F := tree.StringNode{
		Value: "F",
	}
	leaf_G := tree.StringNode{
		Value: "G",
	}
	leaf_C := tree.StringNode{
		Value: "C",
		Left:  &leaf_F,
		Right: &leaf_G,
	}

	leaf_A := tree.StringNode{
		Value: "A",
		Left:  &leaf_B,
		Right: &leaf_C,
	}
	return &tree.StringBinaryTree{
		Root: &leaf_A,
	}
}
