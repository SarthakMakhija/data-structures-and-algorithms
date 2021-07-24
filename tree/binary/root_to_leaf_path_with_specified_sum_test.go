package binarytree_test

import (
	binarytree "github.com/SarthakMakhija/data-structures-and-algorithms/tree/binary"
	"testing"
)

func TestExistsRootToLeafPathWithSum1(t *testing.T) {
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

	tree := binarytree.IntBinaryTree{
		Root: &root,
	}
	exists := tree.ExistsRootToLeafPathWithSum(591)
	if exists != true {
		t.Fatalf("Expected true, received %v", exists)
	}
}

func TestExistsRootToLeafPathWithSum2(t *testing.T) {
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

	tree := binarytree.IntBinaryTree{
		Root: &root,
	}
	exists := tree.ExistsRootToLeafPathWithSum(590)
	if exists != false {
		t.Fatalf("Expected false, received %v", exists)
	}
}
