package binarytree_test

import (
	binarytree "github.com/SarthakMakhija/data-structures-and-algorithms/tree/binary"
	"testing"
)

func TestToLinkedListFromLeaves(t *testing.T) {
	root := binarytree.IntNode{
		Value: 314,
		Left: &binarytree.IntNode{
			Value: 6,
			Left: &binarytree.IntNode{
				Value: 271,
				Left: &binarytree.IntNode{
					Value: 28,
				},
				Right: &binarytree.IntNode{
					Value: 0,
				},
			},
			Right: &binarytree.IntNode{
				Value: 561,
				Right: &binarytree.IntNode{
					Value: 3,
					Left: &binarytree.IntNode{
						Value: 17,
					},
				},
			},
		},
		Right: &binarytree.IntNode{
			Value: 6,
			Left: &binarytree.IntNode{
				Value: 2,
				Right: &binarytree.IntNode{
					Value: 1,
					Left: &binarytree.IntNode{
						Value: 401,
						Right: &binarytree.IntNode{
							Value: 641,
						},
					},
					Right: &binarytree.IntNode{
						Value: 257,
					},
				},
			},
			Right: &binarytree.IntNode{
				Value: 271,
				Right: &binarytree.IntNode{
					Value: 28,
				},
			},
		},
	}
	bTree := binarytree.IntBinaryTree{
		Root: &root,
	}

	list := bTree.ToLinkedListFromLeaves()
	allElements := list.AllAsString()
	expected := "2801764125728"

	if allElements != expected {
		t.Fatalf("Expected %v, received %v", expected, allElements)
	}
}
