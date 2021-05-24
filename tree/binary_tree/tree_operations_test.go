package binarytree_test

import (
	"testing"

	tree "example.com/data-structures-and-algorithms/tree/binary_tree"
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
	output := tree.Max_1()

	if output != 8 {
		t.Fatalf("Expected 8, received %v", output)
	}
}

func TestMax_2(t *testing.T) {
	tree := binary_tree_4()
	output := tree.Max_1()

	if output != 9 {
		t.Fatalf("Expected 9, received %v", output)
	}
}

func TestMax_3(t *testing.T) {
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
	output := tree.Max_2()

	if output != 3 {
		t.Fatalf("Expected 3, received %v", output)
	}
}

func TestMax_4(t *testing.T) {
	tree := binary_tree_3()
	output := tree.Max_2()

	if output != 8 {
		t.Fatalf("Expected 8, received %v", output)
	}
}

func TestMax_5(t *testing.T) {
	tree := binary_tree_4()
	output := tree.Max_2()

	if output != 9 {
		t.Fatalf("Expected 9, received %v", output)
	}
}

func TestContains_1(t *testing.T) {
	tree := binary_tree_4()
	output := tree.Search(9)

	if output.Contains != true {
		t.Fatalf("Expected true, received %v", output)
	}
}

func TestContains_2(t *testing.T) {
	tree := binary_tree_4()
	output := tree.Search(7)

	if output.Contains != true {
		t.Fatalf("Expected true, received %v", output)
	}
}

func TestContains_3(t *testing.T) {
	tree := binary_tree_4()
	output := tree.Search(70)

	if output != nil {
		t.Fatalf("Expected nil, received %v", output)
	}
}

func TestContains_4(t *testing.T) {
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
	output := tree.Search(3)
	if output.Contains != true {
		t.Fatalf("Expected true, received %v", output)
	}
}

func TestInsert_1(t *testing.T) {
	tree := binary_tree_3()
	tree.Insert_Left_Of(8, 9)

	output := tree.Traverse()

	if output != "984251637" {
		t.Fatalf("Expected 984251637, received %v", output)
	}
}

func TestInsert_2(t *testing.T) {
	leaf_3 := tree.IntNode{
		Value: 3,
	}
	leaf_1 := tree.IntNode{
		Value: 1,
		Right: &leaf_3,
	}
	tree := tree.IntBinaryTree{
		Root: &leaf_1,
	}
	tree.Insert_Left_Of(1, 2)

	output := tree.Traverse()

	if output != "213" {
		t.Fatalf("Expected 213, received %v", output)
	}
}

func TestDelete_1(t *testing.T) {
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
	tree.Delete(2)

	output := tree.Traverse()

	if output != "13" {
		t.Fatalf("Expected 13, received %v", output)
	}
}

func TestDelete_2(t *testing.T) {
	tree := binary_tree_3()
	tree.Delete(2)

	output := tree.Traverse()

	if output != "1637" {
		t.Fatalf("Expected 1637, received %v", output)
	}
}

func TestEvaluatePostfix_1(t *testing.T) {
	leaf_6 := tree.StringNode{
		Value: "6",
	}
	leaf_4 := tree.StringNode{
		Value: "4",
	}
	sum := tree.StringNode{
		Value: "+",
		Left:  &leaf_6,
		Right: &leaf_4,
	}
	tree := tree.StringBinaryTree{
		Root: &sum,
	}
	result, _ := tree.Evaluate_Postfix()
	if result != 10 {
		t.Fatalf("Expected 10, received %v", result)
	}
}

func TestEvaluatePostfix_2(t *testing.T) {
	leaf_2 := tree.StringNode{
		Value: "2",
	}
	leaf_6 := tree.StringNode{
		Value: "6",
	}
	leaf_4 := tree.StringNode{
		Value: "4",
	}
	sum := tree.StringNode{
		Value: "+",
		Left:  &leaf_6,
		Right: &leaf_4,
	}
	multiply := tree.StringNode{
		Value: "*",
		Left:  &leaf_2,
		Right: &sum,
	}
	tree := tree.StringBinaryTree{
		Root: &multiply,
	}
	result, _ := tree.Evaluate_Postfix()
	if result != 20 {
		t.Fatalf("Expected 20, received %v", result)
	}
}

func TestEvaluatePostfix_3(t *testing.T) {
	leaf_2 := tree.StringNode{
		Value: "2",
	}
	leaf_6 := tree.StringNode{
		Value: "6",
	}
	leaf_4 := tree.StringNode{
		Value: "4",
	}
	sum := tree.StringNode{
		Value: "+",
		Left:  &leaf_6,
		Right: &leaf_4,
	}
	minus := tree.StringNode{
		Value: "-",
		Left:  &leaf_2,
		Right: &sum,
	}
	tree := tree.StringBinaryTree{
		Root: &minus,
	}
	result, _ := tree.Evaluate_Postfix()
	if result != -8 {
		t.Fatalf("Expected -8, received %v", result)
	}
}

func TestEvaluatePostfix_4(t *testing.T) {
	leaf_40 := tree.StringNode{
		Value: "40",
	}
	leaf_6 := tree.StringNode{
		Value: "6",
	}
	leaf_4 := tree.StringNode{
		Value: "4",
	}
	sum := tree.StringNode{
		Value: "+",
		Left:  &leaf_6,
		Right: &leaf_4,
	}
	divide := tree.StringNode{
		Value: "/",
		Left:  &leaf_40,
		Right: &sum,
	}
	tree := tree.StringBinaryTree{
		Root: &divide,
	}
	result, _ := tree.Evaluate_Postfix()
	if result != 4 {
		t.Fatalf("Expected 4, received %v", result)
	}
}

func TestKthElement_1(t *testing.T) {
	//Assume there is a linked list which will be converted to a tree represented below
	tree := binary_tree_5()
	output := tree.K_element(11)
	if output != 105 {
		t.Fatalf("Expected 105, received %v", output)
	}
}

func TestKthElement_2(t *testing.T) {
	//Assume there is a linked list which will be converted to a tree represented below
	tree := binary_tree_5()
	output := tree.K_element(1)
	if output != 1 {
		t.Fatalf("Expected 1, received %v", output)
	}
}

func TestKthElement_3(t *testing.T) {
	//Assume there is a linked list which will be converted to a tree represented below
	tree := binary_tree_5()
	output := tree.K_element(9)
	if output != 90 {
		t.Fatalf("Expected 90, received %v", output)
	}
}

func TestKthElement_4(t *testing.T) {
	//Assume there is a linked list which will be converted to a tree represented below
	tree := binary_tree_5()
	output := tree.K_element(2)
	if output != 20 {
		t.Fatalf("Expected 20, received %v", output)
	}
}

func TestKthElement_5(t *testing.T) {
	//Assume there is a linked list which will be converted to a tree represented below
	tree := binary_tree_5()
	output := tree.K_element(5)
	if output != 15 {
		t.Fatalf("Expected 15, received %v", output)
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

func binary_tree_5() *tree.IntBinaryTree {
	leaf_80 := tree.IntNode{
		Value: 80,
	}
	leaf_90 := tree.IntNode{
		Value: 90,
	}
	node_60 := tree.IntNode{
		Value: 60,
		Left:  &leaf_80,
		Right: &leaf_90,
	}
	leaf_95 := tree.IntNode{
		Value: 95,
	}
	leaf_105 := tree.IntNode{
		Value: 105,
	}
	node_15 := tree.IntNode{
		Value: 15,
		Left:  &leaf_95,
		Right: &leaf_105,
	}
	node_20 := tree.IntNode{
		Value: 20,
		Left:  &node_60,
		Right: &node_15,
	}
	leaf_58 := tree.IntNode{
		Value: 58,
	}
	leaf_71 := tree.IntNode{
		Value: 71,
		Left:  &leaf_58,
	}
	leaf_68 := tree.IntNode{
		Value: 68,
	}
	node_45 := tree.IntNode{
		Value: 45,
		Left:  &leaf_71,
		Right: &leaf_68,
	}
	root := tree.IntNode{
		Value: 1,
		Left:  &node_20,
		Right: &node_45,
	}

	return &tree.IntBinaryTree{
		Root: &root,
	}
}
