package binarytree_test

import (
	binarytree "github.com/SarthakMakhija/data-structures-and-algorithms/tree/binary"
	"reflect"
	"testing"
)

func TestSetRightSibling1(t *testing.T) {
	root := binarytree.StringNodeWithLevelNext{
		Value: "A",
		Left: &binarytree.StringNodeWithLevelNext{
			Value: "B",
		},
		Right: &binarytree.StringNodeWithLevelNext{
			Value: "C",
		},
	}
	tree := binarytree.StringBinaryTreeWithLevelNext{
		Root: &root,
	}

	tree.SetSibling()

	nodeSiblingPairs := tree.LevelOrderTraversal()
	expected := []binarytree.NodeSiblingPair{
		{
			Node:    "B",
			Sibling: "C",
		},
		{
			Node:    "C",
			Sibling: "NA",
		},
	}

	if !reflect.DeepEqual(nodeSiblingPairs, expected) {
		t.Fatalf("Expected %v, received %v", expected, nodeSiblingPairs)
	}
}

func TestSetRightSibling2(t *testing.T) {
	root := binarytree.StringNodeWithLevelNext{
		Value: "A",
		Left: &binarytree.StringNodeWithLevelNext{
			Value: "B",
			Left: &binarytree.StringNodeWithLevelNext{
				Value: "C",
			},
		},
		Right: &binarytree.StringNodeWithLevelNext{
			Value: "D",
			Left: &binarytree.StringNodeWithLevelNext{
				Value: "E",
				Left: &binarytree.StringNodeWithLevelNext{
					Value: "F",
				},
			},
			Right: &binarytree.StringNodeWithLevelNext{
				Value: "G",
				Left: &binarytree.StringNodeWithLevelNext{
					Value: "H",
				},
				Right: &binarytree.StringNodeWithLevelNext{
					Value: "I",
				},
			},
		},
	}
	tree := binarytree.StringBinaryTreeWithLevelNext{
		Root: &root,
	}

	tree.SetSibling()

	nodeSiblingPairs := tree.LevelOrderTraversal()
	expected := []binarytree.NodeSiblingPair{
		{
			Node:    "B",
			Sibling: "D",
		},
		{
			Node:    "D",
			Sibling: "NA",
		},
		{
			Node:    "C",
			Sibling: "E",
		},
		{
			Node:    "E",
			Sibling: "G",
		},
		{
			Node:    "G",
			Sibling: "NA",
		},
		{
			Node:    "F",
			Sibling: "H",
		},
		{
			Node:    "H",
			Sibling: "I",
		},

		{
			Node:    "I",
			Sibling: "NA",
		},
	}

	if !reflect.DeepEqual(nodeSiblingPairs, expected) {
		t.Fatalf("Expected %v, received %v", expected, nodeSiblingPairs)
	}
}

func TestSetRightSibling3(t *testing.T) {
	root := binarytree.StringNodeWithLevelNext{
		Value: "A",
		Left: &binarytree.StringNodeWithLevelNext{
			Value: "B",
			Left: &binarytree.StringNodeWithLevelNext{
				Value: "D",
				Left: &binarytree.StringNodeWithLevelNext{
					Value: "P",
				},
				Right: &binarytree.StringNodeWithLevelNext{
					Value: "Q",
				},
			},
			Right: &binarytree.StringNodeWithLevelNext{
				Value: "O",
				Left: &binarytree.StringNodeWithLevelNext{
					Value: "X",
				},
			},
		},
		Right: &binarytree.StringNodeWithLevelNext{
			Value: "C",
			Left: &binarytree.StringNodeWithLevelNext{
				Value: "E",
				Left: &binarytree.StringNodeWithLevelNext{
					Value: "R",
				},
				Right: &binarytree.StringNodeWithLevelNext{
					Value: "T",
				},
			},
			Right: &binarytree.StringNodeWithLevelNext{
				Value: "G",
				Left: &binarytree.StringNodeWithLevelNext{
					Value: "H",
				},
				Right: &binarytree.StringNodeWithLevelNext{
					Value: "I",
				},
			},
		},
	}
	tree := binarytree.StringBinaryTreeWithLevelNext{
		Root: &root,
	}

	tree.SetSibling()

	nodeSiblingPairs := tree.LevelOrderTraversal()
	expected := []binarytree.NodeSiblingPair{
		{
			Node:    "B",
			Sibling: "C",
		},
		{
			Node:    "C",
			Sibling: "NA",
		},
		{
			Node:    "D",
			Sibling: "O",
		},
		{
			Node:    "O",
			Sibling: "E",
		},
		{
			Node:    "E",
			Sibling: "G",
		},
		{
			Node:    "G",
			Sibling: "NA",
		},
		{
			Node:    "P",
			Sibling: "Q",
		},
		{
			Node:    "Q",
			Sibling: "X",
		},
		{
			Node:    "X",
			Sibling: "R",
		},
		{
			Node:    "R",
			Sibling: "T",
		},
		{
			Node:    "T",
			Sibling: "H",
		},
		{
			Node:    "H",
			Sibling: "I",
		},
		{
			Node:    "I",
			Sibling: "NA",
		},
	}

	if !reflect.DeepEqual(nodeSiblingPairs, expected) {
		t.Fatalf("Expected %v, received %v", expected, nodeSiblingPairs)
	}
}
