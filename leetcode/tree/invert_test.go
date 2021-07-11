package tree_test

import (
	serde "github.com/SarthakMakhija/data-structures-and-algorithms/leetcode/tree"
	"testing"
)

func TestIntBinaryTreeInvert1(t *testing.T) {
	tree := binaryTree()

	tree.Invert()
	output := tree.Traverse()
	expected := "9764321"

	if output != expected {
		t.Fatalf("Expected %v, recevied %v", expected, output)
	}
}

func TestIntBinaryTreeInvert2(t *testing.T) {
	tree := anotherBinaryTree()

	tree.Invert()
	output := tree.Traverse()
	expected := "1917161453021"

	if output != expected {
		t.Fatalf("Expected %v, recevied %v", expected, output)
	}
}

func binaryTree() *serde.IntBinaryTree {

	node1 := serde.IntNode{
		Value: 1,
	}
	node3 := serde.IntNode{
		Value: 3,
	}
	node2 := serde.IntNode{
		Value: 2,
		Left:  &node1,
		Right: &node3,
	}

	node6 := serde.IntNode{
		Value: 6,
	}
	node9 := serde.IntNode{
		Value: 9,
	}
	node7 := serde.IntNode{
		Value: 7,
		Left:  &node6,
		Right: &node9,
	}

	node4 := serde.IntNode{
		Value: 4,
		Left:  &node2,
		Right: &node7,
	}
	return &serde.IntBinaryTree{
		Root: &node4,
	}
}

func anotherBinaryTree() *serde.IntBinaryTree {

	node1 := serde.IntNode{
		Value: 1,
	}
	node3 := serde.IntNode{
		Value: 3,
		Left: &serde.IntNode{
			Value: 0,
		},
		Right: &serde.IntNode{
			Value: 5,
		},
	}
	node2 := serde.IntNode{
		Value: 2,
		Left:  &node1,
		Right: &node3,
	}

	node16 := serde.IntNode{
		Value: 16,
	}
	node19 := serde.IntNode{
		Value: 19,
	}
	node17 := serde.IntNode{
		Value: 17,
		Left:  &node16,
		Right: &node19,
	}

	node4 := serde.IntNode{
		Value: 14,
		Left:  &node2,
		Right: &node17,
	}
	return &serde.IntBinaryTree{
		Root: &node4,
	}
}
