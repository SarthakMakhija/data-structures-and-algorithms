package binarytree_test

import (
	binarytree "github.com/SarthakMakhija/data-structures-and-algorithms/tree/binary"
	"reflect"
	"testing"
)

func TestBinaryEncodings1(t *testing.T) {
	root := binarytree.IntNode{
		Value: 1,
		Left: &binarytree.IntNode{
			Value: 0,
			Left: &binarytree.IntNode{
				Value: 0,
				Right: &binarytree.IntNode{
					Value: 1,
				},
			},
			Right: &binarytree.IntNode{
				Value: 1,
				Left: &binarytree.IntNode{
					Value: 0,
					Right: &binarytree.IntNode{
						Value: 1,
					},
				},
				Right: &binarytree.IntNode{
					Value: 1,
					Right: &binarytree.IntNode{
						Value: 0,
					},
				},
			},
		},
	}
	tree := binarytree.IntBinaryTree{
		Root: &root,
	}
	encodings := tree.BinaryEncodings()
	expected := []string{"1001", "10101", "10110"}

	if !reflect.DeepEqual(encodings, expected) {
		t.Fatalf("Expected %v, received %v", expected, encodings)
	}
}

func TestBinaryEncodings2(t *testing.T) {
	root := binarytree.IntNode{
		Value: 1,
		Left: &binarytree.IntNode{
			Value: 0,
			Left: &binarytree.IntNode{
				Value: 0,
				Left: &binarytree.IntNode{
					Value: 0,
				},
				Right: &binarytree.IntNode{
					Value: 1,
				},
			},
			Right: &binarytree.IntNode{
				Value: 1,
				Right: &binarytree.IntNode{
					Value: 1,
					Left: &binarytree.IntNode{
						Value: 0,
					},
				},
			},
		},
		Right: &binarytree.IntNode{
			Value: 1,
			Left: &binarytree.IntNode{
				Value: 0,
				Right: &binarytree.IntNode{
					Value: 0,
					Left: &binarytree.IntNode{
						Value: 1,
						Right: &binarytree.IntNode{
							Value: 1,
						},
					},
					Right: &binarytree.IntNode{
						Value: 0,
					},
				},
			},
			Right: &binarytree.IntNode{
				Value: 0,
				Right: &binarytree.IntNode{
					Value: 0,
				},
			},
		},
	}
	tree := binarytree.IntBinaryTree{
		Root: &root,
	}
	encodings := tree.BinaryEncodings()
	expected := []string{"1000", "1001", "10110", "110011", "11000", "1100"}

	if !reflect.DeepEqual(encodings, expected) {
		t.Fatalf("Expected %v, received %v", expected, encodings)
	}
}

func TestBinaryEncodingsSum1(t *testing.T) {
	root := binarytree.IntNode{
		Value: 1,
		Left: &binarytree.IntNode{
			Value: 0,
			Left: &binarytree.IntNode{
				Value: 0,
				Right: &binarytree.IntNode{
					Value: 1,
				},
			},
			Right: &binarytree.IntNode{
				Value: 1,
				Left: &binarytree.IntNode{
					Value: 0,
					Right: &binarytree.IntNode{
						Value: 1,
					},
				},
				Right: &binarytree.IntNode{
					Value: 1,
					Right: &binarytree.IntNode{
						Value: 0,
					},
				},
			},
		},
	}
	tree := binarytree.IntBinaryTree{
		Root: &root,
	}
	sum := tree.BinaryEncodingsSum()
	expected := 52

	if sum != expected {
		t.Fatalf("Expected %v, received %v", expected, sum)
	}
}

func TestBinaryEncodingsSum2(t *testing.T) {
	root := binarytree.IntNode{
		Value: 1,
		Left: &binarytree.IntNode{
			Value: 0,
			Left: &binarytree.IntNode{
				Value: 0,
				Left: &binarytree.IntNode{
					Value: 0,
				},
				Right: &binarytree.IntNode{
					Value: 1,
				},
			},
			Right: &binarytree.IntNode{
				Value: 1,
				Right: &binarytree.IntNode{
					Value: 1,
					Left: &binarytree.IntNode{
						Value: 0,
					},
				},
			},
		},
		Right: &binarytree.IntNode{
			Value: 1,
			Left: &binarytree.IntNode{
				Value: 0,
				Right: &binarytree.IntNode{
					Value: 0,
					Left: &binarytree.IntNode{
						Value: 1,
						Right: &binarytree.IntNode{
							Value: 1,
						},
					},
					Right: &binarytree.IntNode{
						Value: 0,
					},
				},
			},
			Right: &binarytree.IntNode{
				Value: 0,
				Right: &binarytree.IntNode{
					Value: 0,
				},
			},
		},
	}
	tree := binarytree.IntBinaryTree{
		Root: &root,
	}
	sum := tree.BinaryEncodingsSum()
	expected := 126

	if sum != expected {
		t.Fatalf("Expected %v, received %v", expected, sum)
	}
}

func TestBinaryToDecimal1(t *testing.T) {
	binary := "1001"
	expected := 9
	output := binarytree.BinaryToDecimal(binary)

	if output != expected {
		t.Fatalf("Expected %v, received %v", expected, output)
	}
}

func TestBinaryToDecimal2(t *testing.T) {
	binary := "10011"
	expected := 19
	output := binarytree.BinaryToDecimal(binary)

	if output != expected {
		t.Fatalf("Expected %v, received %v", expected, output)
	}
}
