package binarytree_test

import (
	binarytree "github.com/SarthakMakhija/data-structures-and-algorithms/tree/binary"
	"testing"
)

func TestIsSymmetric1(t *testing.T) {
	root := binarytree.IntNode{
		Value: 312,
		Left: &binarytree.IntNode{
			Value: 6,
			Right: &binarytree.IntNode{
				Value: 2,
				Right: &binarytree.IntNode{
					Value: 3,
				},
			},
		},
		Right: &binarytree.IntNode{
			Value: 6,
			Left: &binarytree.IntNode{
				Value: 2,
				Left: &binarytree.IntNode{
					Value: 3,
				},
			},
		},
	}
	tree := binarytree.IntBinaryTree{
		Root: &root,
	}
	result := tree.IsSymmetric()

	if result != true {
		t.Fatalf("Expected true, recevied %v", result)
	}
}

func TestIsSymmetric2(t *testing.T) {
	root := binarytree.IntNode{
		Value: 312,
		Left: &binarytree.IntNode{
			Value: 6,
			Right: &binarytree.IntNode{
				Value: 561,
				Right: &binarytree.IntNode{
					Value: 3,
				},
			},
		},
		Right: &binarytree.IntNode{
			Value: 6,
			Left: &binarytree.IntNode{
				Value: 2,
				Left: &binarytree.IntNode{
					Value: 3,
				},
			},
		},
	}
	tree := binarytree.IntBinaryTree{
		Root: &root,
	}
	result := tree.IsSymmetric()

	if result != false {
		t.Fatalf("Expected false, recevied %v", result)
	}
}

func TestIsSymmetric3(t *testing.T) {
	root := binarytree.IntNode{
		Value: 312,
		Left: &binarytree.IntNode{
			Value: 6,
			Right: &binarytree.IntNode{
				Value: 561,
				Right: &binarytree.IntNode{
					Value: 3,
				},
			},
		},
		Right: &binarytree.IntNode{
			Value: 6,
			Left: &binarytree.IntNode{
				Value: 561,
			},
		},
	}
	tree := binarytree.IntBinaryTree{
		Root: &root,
	}
	result := tree.IsSymmetric()

	if result != false {
		t.Fatalf("Expected false, recevied %v", result)
	}
}
