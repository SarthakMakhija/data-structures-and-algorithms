package tree_test

import (
	"github.com/SarthakMakhija/data-structures-and-algorithms/leetcode/tree"
	"testing"
)

func TestIsValid1(t *testing.T) {
	searchTree := binarySearchTree1()
	isValid := searchTree.IsValid()

	if isValid != true {
		t.Fatalf("Expected true, recevied %v", isValid)
	}
}

func TestIsValid2(t *testing.T) {
	node2 := tree.IntNode{
		Value: 2,
		Left: &tree.IntNode{
			Value: 1,
		},
		Right: &tree.IntNode{
			Value: 3,
		},
	}
	searchTree := &tree.IntBinarySearchTree{
		Root: &node2,
	}

	isValid := searchTree.IsValid()

	if isValid != true {
		t.Fatalf("Expected true, recevied %v", isValid)
	}
}

func TestIsValid3(t *testing.T) {
	node4 := tree.IntNode{
		Value: 4,
		Left: &tree.IntNode{
			Value: 3,
		},
		Right: &tree.IntNode{
			Value: 6,
		},
	}
	node5 := tree.IntNode{
		Value: 5,
		Left: &tree.IntNode{
			Value: 1,
		},
		Right: &node4,
	}
	searchTree := &tree.IntBinarySearchTree{
		Root: &node5,
	}

	isValid := searchTree.IsValid()

	if isValid != false {
		t.Fatalf("Expected false, recevied %v", isValid)
	}
}

func TestIsValid4(t *testing.T) {
	node6 := tree.IntNode{
		Value: 6,
		Left: &tree.IntNode{
			Value: 4,
		},
		Right: &tree.IntNode{
			Value: 8,
		},
	}
	node5 := tree.IntNode{
		Value: 5,
		Left: &tree.IntNode{
			Value: 1,
		},
		Right: &node6,
	}
	searchTree := &tree.IntBinarySearchTree{
		Root: &node5,
	}

	isValid := searchTree.IsValid()

	if isValid != false {
		t.Fatalf("Expected false, recevied %v", isValid)
	}
}

func binarySearchTree1() *tree.IntBinarySearchTree {
	node10 := tree.IntNode{
		Value: 10,
	}
	node21 := tree.IntNode{
		Value: 21,
	}
	node20 := tree.IntNode{
		Value: 20,
		Left:  &node10,
		Right: &node21,
	}
	node40 := tree.IntNode{
		Value: 40,
	}
	node30 := tree.IntNode{
		Value: 30,
		Left:  &node20,
		Right: &node40,
	}
	return &tree.IntBinarySearchTree{
		Root: &node30,
	}
}
