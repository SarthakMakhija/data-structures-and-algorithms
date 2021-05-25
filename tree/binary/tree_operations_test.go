package binarytree_test

import (
	"testing"

	tree "github.com/SarthakMakhija/data-structures-and-algorithms/tree/binary"
)

func TestTreeTraversal_1(t *testing.T) {
	binaryTree := binaryTree1()
	output := binaryTree.Traverse()
	if output != "DBEAFCG" {
		t.Fatalf("Expected DBEAFCG, received %v", output)
	}
}

func TestTreeTraversal_2(t *testing.T) {
	leafD := tree.StringNode{
		Value: "D",
	}
	binaryTree := tree.StringBinaryTree{
		Root: &leafD,
	}

	output := binaryTree.Traverse()
	if output != "D" {
		t.Fatalf("Expected D, received %v", output)
	}
}

func TestTreeTraversal_3(t *testing.T) {
	binaryTree := binaryTree2()
	output := binaryTree.Traverse()
	if output != "QDBEAFCG" {
		t.Fatalf("Expected QDBEAFCG, received %v", output)
	}
}

func TestTreeSum(t *testing.T) {
	binaryTree := binaryTree3()
	output := binaryTree.Sum()
	if output != 36 {
		t.Fatalf("Expected 36, received %v", output)
	}
}

func TestTreeNodeCount_1(t *testing.T) {
	binaryTree := binaryTree3()
	output := binaryTree.CountNodes()
	if output != 8 {
		t.Fatalf("Expected 8, received %v", output)
	}
}

func TestTreeNodeCount_2(t *testing.T) {
	leafD := tree.IntNode{
		Value: 1,
	}
	binaryTree := tree.IntBinaryTree{
		Root: &leafD,
	}
	output := binaryTree.CountNodes()
	if output != 1 {
		t.Fatalf("Expected 1, received %v", output)
	}
}

func TestTreeLeafNodeCount_1(t *testing.T) {
	binaryTree := binaryTree3()
	output := binaryTree.CountLeafNodes()
	if output != 4 {
		t.Fatalf("Expected 4, received %v", output)
	}
}

func TestTreeLeafNodeCount_2(t *testing.T) {
	leafD := tree.IntNode{
		Value: 1,
	}
	binaryTree := tree.IntBinaryTree{
		Root: &leafD,
	}
	output := binaryTree.CountLeafNodes()
	if output != 1 {
		t.Fatalf("Expected 1, received %v", output)
	}
}

func TestHeight_1(t *testing.T) {
	binaryTree := binaryTree3()
	output := binaryTree.Height1()

	if output != 4 {
		t.Fatalf("Expected 4, received %v", output)
	}
}

func TestHeight_2(t *testing.T) {
	binaryTree := binaryTree4()
	output := binaryTree.Height2()

	if output != 5 {
		t.Fatalf("Expected 5, received %v", output)
	}
}

func TestHeight_3(t *testing.T) {
	leaf2 := tree.IntNode{
		Value: 2,
	}
	leaf3 := tree.IntNode{
		Value: 3,
	}
	leaf1 := tree.IntNode{
		Value: 1,
		Left:  &leaf2,
		Right: &leaf3,
	}
	binaryTree := tree.IntBinaryTree{
		Root: &leaf1,
	}
	output := binaryTree.Height2()

	if output != 2 {
		t.Fatalf("Expected 2, received %v", output)
	}
}

func TestMax_1(t *testing.T) {
	binaryTree := binaryTree3()
	output := binaryTree.Max1()

	if output != 8 {
		t.Fatalf("Expected 8, received %v", output)
	}
}

func TestMax_2(t *testing.T) {
	binaryTree := binaryTree4()
	output := binaryTree.Max1()

	if output != 9 {
		t.Fatalf("Expected 9, received %v", output)
	}
}

func TestMax_3(t *testing.T) {
	leaf2 := tree.IntNode{
		Value: 2,
	}
	leaf3 := tree.IntNode{
		Value: 3,
	}
	leaf1 := tree.IntNode{
		Value: 1,
		Left:  &leaf2,
		Right: &leaf3,
	}
	binaryTree := tree.IntBinaryTree{
		Root: &leaf1,
	}
	output := binaryTree.Max2()

	if output != 3 {
		t.Fatalf("Expected 3, received %v", output)
	}
}

func TestMax_4(t *testing.T) {
	binaryTree := binaryTree3()
	output := binaryTree.Max2()

	if output != 8 {
		t.Fatalf("Expected 8, received %v", output)
	}
}

func TestMax_5(t *testing.T) {
	binaryTree := binaryTree4()
	output := binaryTree.Max2()

	if output != 9 {
		t.Fatalf("Expected 9, received %v", output)
	}
}

func TestContains_1(t *testing.T) {
	binaryTree := binaryTree4()
	output := binaryTree.Search(9)

	if output.Contains != true {
		t.Fatalf("Expected true, received %v", output)
	}
}

func TestContains_2(t *testing.T) {
	binaryTree := binaryTree4()
	output := binaryTree.Search(7)

	if output.Contains != true {
		t.Fatalf("Expected true, received %v", output)
	}
}

func TestContains_3(t *testing.T) {
	binaryTree := binaryTree4()
	output := binaryTree.Search(70)

	if output != nil {
		t.Fatalf("Expected nil, received %v", output)
	}
}

func TestContains_4(t *testing.T) {
	leaf2 := tree.IntNode{
		Value: 2,
	}
	leaf3 := tree.IntNode{
		Value: 3,
	}
	leaf1 := tree.IntNode{
		Value: 1,
		Left:  &leaf2,
		Right: &leaf3,
	}
	binaryTree := tree.IntBinaryTree{
		Root: &leaf1,
	}
	output := binaryTree.Search(3)
	if output.Contains != true {
		t.Fatalf("Expected true, received %v", output)
	}
}

func TestInsert_1(t *testing.T) {
	binaryTree := binaryTree3()
	binaryTree.InsertLeftOf(8, 9)

	output := binaryTree.Traverse()

	if output != "984251637" {
		t.Fatalf("Expected 984251637, received %v", output)
	}
}

func TestInsert_2(t *testing.T) {
	leaf3 := tree.IntNode{
		Value: 3,
	}
	leaf1 := tree.IntNode{
		Value: 1,
		Right: &leaf3,
	}
	binaryTree := tree.IntBinaryTree{
		Root: &leaf1,
	}
	binaryTree.InsertLeftOf(1, 2)

	output := binaryTree.Traverse()

	if output != "213" {
		t.Fatalf("Expected 213, received %v", output)
	}
}

func TestDelete_1(t *testing.T) {
	leaf2 := tree.IntNode{
		Value: 2,
	}
	leaf3 := tree.IntNode{
		Value: 3,
	}
	leaf1 := tree.IntNode{
		Value: 1,
		Left:  &leaf2,
		Right: &leaf3,
	}
	binaryTree := tree.IntBinaryTree{
		Root: &leaf1,
	}
	binaryTree.Delete(2)

	output := binaryTree.Traverse()

	if output != "13" {
		t.Fatalf("Expected 13, received %v", output)
	}
}

func TestDelete_2(t *testing.T) {
	binaryTree := binaryTree3()
	binaryTree.Delete(2)

	output := binaryTree.Traverse()

	if output != "1637" {
		t.Fatalf("Expected 1637, received %v", output)
	}
}

func TestEvaluatePostfix_1(t *testing.T) {
	leaf6 := tree.StringNode{
		Value: "6",
	}
	leaf4 := tree.StringNode{
		Value: "4",
	}
	sum := tree.StringNode{
		Value: "+",
		Left:  &leaf6,
		Right: &leaf4,
	}
	binaryTree := tree.StringBinaryTree{
		Root: &sum,
	}
	result, _ := binaryTree.EvaluatePostfix()
	if result != 10 {
		t.Fatalf("Expected 10, received %v", result)
	}
}

func TestEvaluatePostfix_2(t *testing.T) {
	leaf2 := tree.StringNode{
		Value: "2",
	}
	leaf6 := tree.StringNode{
		Value: "6",
	}
	leaf4 := tree.StringNode{
		Value: "4",
	}
	sum := tree.StringNode{
		Value: "+",
		Left:  &leaf6,
		Right: &leaf4,
	}
	multiply := tree.StringNode{
		Value: "*",
		Left:  &leaf2,
		Right: &sum,
	}
	binaryTree := tree.StringBinaryTree{
		Root: &multiply,
	}
	result, _ := binaryTree.EvaluatePostfix()
	if result != 20 {
		t.Fatalf("Expected 20, received %v", result)
	}
}

func TestEvaluatePostfix_3(t *testing.T) {
	leaf2 := tree.StringNode{
		Value: "2",
	}
	leaf6 := tree.StringNode{
		Value: "6",
	}
	leaf4 := tree.StringNode{
		Value: "4",
	}
	sum := tree.StringNode{
		Value: "+",
		Left:  &leaf6,
		Right: &leaf4,
	}
	minus := tree.StringNode{
		Value: "-",
		Left:  &leaf2,
		Right: &sum,
	}
	binaryTree := tree.StringBinaryTree{
		Root: &minus,
	}
	result, _ := binaryTree.EvaluatePostfix()
	if result != -8 {
		t.Fatalf("Expected -8, received %v", result)
	}
}

func TestEvaluatePostfix_4(t *testing.T) {
	leaf40 := tree.StringNode{
		Value: "40",
	}
	leaf6 := tree.StringNode{
		Value: "6",
	}
	leaf4 := tree.StringNode{
		Value: "4",
	}
	sum := tree.StringNode{
		Value: "+",
		Left:  &leaf6,
		Right: &leaf4,
	}
	divide := tree.StringNode{
		Value: "/",
		Left:  &leaf40,
		Right: &sum,
	}
	binaryTree := tree.StringBinaryTree{
		Root: &divide,
	}
	result, _ := binaryTree.EvaluatePostfix()
	if result != 4 {
		t.Fatalf("Expected 4, received %v", result)
	}
}

func TestKthElement_1(t *testing.T) {
	//Assume there is a linked list which will be converted to a tree represented below
	binaryTree := binaryTree5()
	output := binaryTree.KElement(11)
	if output != 105 {
		t.Fatalf("Expected 105, received %v", output)
	}
}

func TestKthElement_2(t *testing.T) {
	//Assume there is a linked list which will be converted to a tree represented below
	binaryTree := binaryTree5()
	output := binaryTree.KElement(1)
	if output != 1 {
		t.Fatalf("Expected 1, received %v", output)
	}
}

func TestKthElement_3(t *testing.T) {
	//Assume there is a linked list which will be converted to a binaryTree represented below
	binaryTree := binaryTree5()
	output := binaryTree.KElement(9)
	if output != 90 {
		t.Fatalf("Expected 90, received %v", output)
	}
}

func TestKthElement_4(t *testing.T) {
	//Assume there is a linked list which will be converted to a binaryTree represented below
	binaryTree := binaryTree5()
	output := binaryTree.KElement(2)
	if output != 20 {
		t.Fatalf("Expected 20, received %v", output)
	}
}

func TestKthElement_5(t *testing.T) {
	//Assume there is a linked list which will be converted to a binaryTree represented below
	binaryTree := binaryTree5()
	output := binaryTree.KElement(5)
	if output != 15 {
		t.Fatalf("Expected 15, received %v", output)
	}
}

func binaryTree1() *tree.StringBinaryTree {

	leafD := tree.StringNode{
		Value: "D",
	}
	leafE := tree.StringNode{
		Value: "E",
	}
	leafB := tree.StringNode{
		Value: "B",
		Left:  &leafD,
		Right: &leafE,
	}

	leafF := tree.StringNode{
		Value: "F",
	}
	leafG := tree.StringNode{
		Value: "G",
	}
	leafC := tree.StringNode{
		Value: "C",
		Left:  &leafF,
		Right: &leafG,
	}

	leafA := tree.StringNode{
		Value: "A",
		Left:  &leafB,
		Right: &leafC,
	}
	return &tree.StringBinaryTree{
		Root: &leafA,
	}
}

func binaryTree2() *tree.StringBinaryTree {

	leafQ := tree.StringNode{
		Value: "Q",
	}
	leafD := tree.StringNode{
		Value: "D",
		Left:  &leafQ,
	}
	leafE := tree.StringNode{
		Value: "E",
	}
	leafB := tree.StringNode{
		Value: "B",
		Left:  &leafD,
		Right: &leafE,
	}

	leafF := tree.StringNode{
		Value: "F",
	}
	leafG := tree.StringNode{
		Value: "G",
	}
	leafC := tree.StringNode{
		Value: "C",
		Left:  &leafF,
		Right: &leafG,
	}

	leafA := tree.StringNode{
		Value: "A",
		Left:  &leafB,
		Right: &leafC,
	}
	return &tree.StringBinaryTree{
		Root: &leafA,
	}
}

func binaryTree3() *tree.IntBinaryTree {

	leaf8 := tree.IntNode{
		Value: 8,
	}
	leaf4 := tree.IntNode{
		Value: 4,
		Left:  &leaf8,
	}
	leaf5 := tree.IntNode{
		Value: 5,
	}
	leaf2 := tree.IntNode{
		Value: 2,
		Left:  &leaf4,
		Right: &leaf5,
	}

	leaf6 := tree.IntNode{
		Value: 6,
	}
	leaf7 := tree.IntNode{
		Value: 7,
	}
	leaf3 := tree.IntNode{
		Value: 3,
		Left:  &leaf6,
		Right: &leaf7,
	}

	leaf1 := tree.IntNode{
		Value: 1,
		Left:  &leaf2,
		Right: &leaf3,
	}
	return &tree.IntBinaryTree{
		Root: &leaf1,
	}
}

func binaryTree4() *tree.IntBinaryTree {

	leaf9 := tree.IntNode{
		Value: 9,
	}
	leaf8 := tree.IntNode{
		Value: 8,
		Right: &leaf9,
	}
	leaf4 := tree.IntNode{
		Value: 4,
		Left:  &leaf8,
	}
	leaf5 := tree.IntNode{
		Value: 5,
	}
	leaf2 := tree.IntNode{
		Value: 2,
		Left:  &leaf4,
		Right: &leaf5,
	}

	leaf6 := tree.IntNode{
		Value: 6,
	}
	leaf7 := tree.IntNode{
		Value: 7,
	}
	leaf3 := tree.IntNode{
		Value: 3,
		Left:  &leaf6,
		Right: &leaf7,
	}

	leaf1 := tree.IntNode{
		Value: 1,
		Left:  &leaf2,
		Right: &leaf3,
	}
	return &tree.IntBinaryTree{
		Root: &leaf1,
	}
}

func binaryTree5() *tree.IntBinaryTree {
	leaf80 := tree.IntNode{
		Value: 80,
	}
	leaf90 := tree.IntNode{
		Value: 90,
	}
	node60 := tree.IntNode{
		Value: 60,
		Left:  &leaf80,
		Right: &leaf90,
	}
	leaf95 := tree.IntNode{
		Value: 95,
	}
	leaf105 := tree.IntNode{
		Value: 105,
	}
	node15 := tree.IntNode{
		Value: 15,
		Left:  &leaf95,
		Right: &leaf105,
	}
	node20 := tree.IntNode{
		Value: 20,
		Left:  &node60,
		Right: &node15,
	}
	leaf58 := tree.IntNode{
		Value: 58,
	}
	leaf71 := tree.IntNode{
		Value: 71,
		Left:  &leaf58,
	}
	leaf68 := tree.IntNode{
		Value: 68,
	}
	node45 := tree.IntNode{
		Value: 45,
		Left:  &leaf71,
		Right: &leaf68,
	}
	root := tree.IntNode{
		Value: 1,
		Left:  &node20,
		Right: &node45,
	}

	return &tree.IntBinaryTree{
		Root: &root,
	}
}
