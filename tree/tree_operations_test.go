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

func TestTreeSum(t *testing.T) {
	tree := binary_tree_3()
	output := tree.Sum()
	if output != 36 {
		t.Fatalf("Expected 36, received %v", output)
	}
}

func TestTreeNodeCount_1(t *testing.T) {
	tree := binary_tree_3()
	output := tree.Count_Nodes()
	if output != 8 {
		t.Fatalf("Expected 8, received %v", output)
	}
}

func TestTreeNodeCount_2(t *testing.T) {
	leaf_D := tree.IntNode{
		Value: 1,
	}
	tree := tree.IntBinaryTree{
		Root: &leaf_D,
	}
	output := tree.Count_Nodes()
	if output != 1 {
		t.Fatalf("Expected 1, received %v", output)
	}
}

func TestTreeLeafNodeCount_1(t *testing.T) {
	tree := binary_tree_3()
	output := tree.Count_Leaf_Nodes()
	if output != 4 {
		t.Fatalf("Expected 4, received %v", output)
	}
}

func TestTreeLeafNodeCount_2(t *testing.T) {
	leaf_D := tree.IntNode{
		Value: 1,
	}
	tree := tree.IntBinaryTree{
		Root: &leaf_D,
	}
	output := tree.Count_Leaf_Nodes()
	if output != 1 {
		t.Fatalf("Expected 1, received %v", output)
	}
}

func TestHeight_1(t *testing.T) {
	tree := binary_tree_3()
	output := tree.Height_1()

	if output != 4 {
		t.Fatalf("Expected 4, received %v", output)
	}
}

func TestHeight_2(t *testing.T) {
	tree := binary_tree_4()
	output := tree.Height_2()

	if output != 5 {
		t.Fatalf("Expected 5, received %v", output)
	}
}

func TestHeight_3(t *testing.T) {
	leaf_2 := tree.IntNode{
		Value: 2,
	}
	leaf_3 := tree.IntNode{
		Value: 3,
	}
	leaf_1 := tree.IntNode{
		Value: 1,
		Left:  &leaf_2,
		Right: &leaf_3,
	}
	tree := tree.IntBinaryTree{
		Root: &leaf_1,
	}
	output := tree.Height_2()

	if output != 2 {
		t.Fatalf("Expected 2, received %v", output)
	}
}

func TestMax_1(t *testing.T) {
	tree := binary_tree_3()
	output := tree.Max()

	if output != 8 {
		t.Fatalf("Expected 8, received %v", output)
	}
}

func TestMax_2(t *testing.T) {
	tree := binary_tree_4()
	output := tree.Max()

	if output != 9 {
		t.Fatalf("Expected 9, received %v", output)
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

func binary_tree_3() *tree.IntBinaryTree {

	leaf_8 := tree.IntNode{
		Value: 8,
	}
	leaf_4 := tree.IntNode{
		Value: 4,
		Left:  &leaf_8,
	}
	leaf_5 := tree.IntNode{
		Value: 5,
	}
	leaf_2 := tree.IntNode{
		Value: 2,
		Left:  &leaf_4,
		Right: &leaf_5,
	}

	leaf_6 := tree.IntNode{
		Value: 6,
	}
	leaf_7 := tree.IntNode{
		Value: 7,
	}
	leaf_3 := tree.IntNode{
		Value: 3,
		Left:  &leaf_6,
		Right: &leaf_7,
	}

	leaf_1 := tree.IntNode{
		Value: 1,
		Left:  &leaf_2,
		Right: &leaf_3,
	}
	return &tree.IntBinaryTree{
		Root: &leaf_1,
	}
}

func binary_tree_4() *tree.IntBinaryTree {

	leaf_9 := tree.IntNode{
		Value: 9,
	}
	leaf_8 := tree.IntNode{
		Value: 8,
		Right: &leaf_9,
	}
	leaf_4 := tree.IntNode{
		Value: 4,
		Left:  &leaf_8,
	}
	leaf_5 := tree.IntNode{
		Value: 5,
	}
	leaf_2 := tree.IntNode{
		Value: 2,
		Left:  &leaf_4,
		Right: &leaf_5,
	}

	leaf_6 := tree.IntNode{
		Value: 6,
	}
	leaf_7 := tree.IntNode{
		Value: 7,
	}
	leaf_3 := tree.IntNode{
		Value: 3,
		Left:  &leaf_6,
		Right: &leaf_7,
	}

	leaf_1 := tree.IntNode{
		Value: 1,
		Left:  &leaf_2,
		Right: &leaf_3,
	}
	return &tree.IntBinaryTree{
		Root: &leaf_1,
	}
}
