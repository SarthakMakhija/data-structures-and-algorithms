package binarytree_test

import (
	binarytree "github.com/SarthakMakhija/data-structures-and-algorithms/tree/binary"
	"testing"
)

func TestIsHeightBalanced1(t *testing.T) {
	root := binarytree.StringNode{
		Value: "A",
		Left: &binarytree.StringNode{
			Value: "B",
			Left: &binarytree.StringNode{
				Value: "C",
				Left: &binarytree.StringNode{
					Value: "D",
					Left: &binarytree.StringNode{
						Value: "E",
					},
					Right: &binarytree.StringNode{
						Value: "F",
					},
				},
				Right: &binarytree.StringNode{
					Value: "G",
				},
			},
			Right: &binarytree.StringNode{
				Value: "H",
				Left: &binarytree.StringNode{
					Value: "I",
				},
				Right: &binarytree.StringNode{
					Value: "J",
				},
			},
		},
		Right: &binarytree.StringNode{
			Value: "K",
			Left: &binarytree.StringNode{
				Value: "L",
				Left: &binarytree.StringNode{
					Value: "M",
				},
				Right: &binarytree.StringNode{
					Value: "N",
				},
			},
			Right: &binarytree.StringNode{
				Value: "O",
			},
		},
	}
	tree := binarytree.StringBinaryTree{
		Root: &root,
	}

	output := tree.IsHeightBalanced()
	if output != true {
		t.Fatalf("Expected true, recevied %v", output)
	}
}

func TestIsHeightBalanced2(t *testing.T) {
	root := binarytree.StringNode{
		Value: "A",
		Left: &binarytree.StringNode{
			Value: "B",
			Left: &binarytree.StringNode{
				Value: "C",
				Left: &binarytree.StringNode{
					Value: "D",
					Left: &binarytree.StringNode{
						Value: "E",
					},
					Right: &binarytree.StringNode{
						Value: "F",
					},
				},
				Right: &binarytree.StringNode{
					Value: "G",
				},
			},
		},
		Right: &binarytree.StringNode{
			Value: "K",
			Left: &binarytree.StringNode{
				Value: "L",
				Left: &binarytree.StringNode{
					Value: "M",
				},
				Right: &binarytree.StringNode{
					Value: "N",
				},
			},
			Right: &binarytree.StringNode{
				Value: "O",
			},
		},
	}
	tree := binarytree.StringBinaryTree{
		Root: &root,
	}

	output := tree.IsHeightBalanced()
	if output != false {
		t.Fatalf("Expected false, recevied %v", output)
	}
}
