package queue_test

import (
	"github.com/SarthakMakhija/data-structures-and-algorithms/queue"
	binarytree "github.com/SarthakMakhija/data-structures-and-algorithms/tree/binary"
	"reflect"
	"testing"
)

func TestBinaryTreeNodesInOrderOfIncreasingDepth1(t *testing.T) {
	root := binarytree.IntNode{
		Value: 10,
		Left: &binarytree.IntNode{
			Value: 20,
			Left: &binarytree.IntNode{
				Value: 40,
			},
			Right: &binarytree.IntNode{
				Value: 50,
			},
		},
		Right: &binarytree.IntNode{
			Value: 30,
			Left: &binarytree.IntNode{
				Value: 60,
			},
			Right: &binarytree.IntNode{
				Value: 70,
			},
		},
	}
	binaryTree := binarytree.IntBinaryTree{
		Root: &root,
	}
	elements := queue.BinaryTreeNodesInOrderOfIncreasingDepth(binaryTree)
	expected := []int{10, 20, 30, 40, 50, 60, 70}

	if !reflect.DeepEqual(elements, expected) {
		t.Fatalf("Expected %v, received %v", expected, elements)
	}
}

func TestBinaryTreeNodesInOrderOfIncreasingDepth2(t *testing.T) {
	root := binarytree.IntNode{
		Value: 10,
		Left: &binarytree.IntNode{
			Value: 20,
			Left: &binarytree.IntNode{
				Value: 40,
			},
			Right: &binarytree.IntNode{
				Value: 50,
			},
		},
		Right: &binarytree.IntNode{
			Value: 30,
			Right: &binarytree.IntNode{
				Value: 70,
			},
		},
	}
	binaryTree := binarytree.IntBinaryTree{
		Root: &root,
	}
	elements := queue.BinaryTreeNodesInOrderOfIncreasingDepth(binaryTree)
	expected := []int{10, 20, 30, 40, 50, 70}

	if !reflect.DeepEqual(elements, expected) {
		t.Fatalf("Expected %v, received %v", expected, elements)
	}
}

func TestBinaryTreeNodesInOrderOfIncreasingDepth3(t *testing.T) {
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
	binaryTree := binarytree.IntBinaryTree{
		Root: &root,
	}
	elements := queue.BinaryTreeNodesInOrderOfIncreasingDepth(binaryTree)
	expected := []int{314, 6, 6, 271, 561, 2, 271, 28, 0, 3, 1, 28, 17, 401, 257, 641}

	if !reflect.DeepEqual(elements, expected) {
		t.Fatalf("Expected %v, received %v", expected, elements)
	}
}

func TestBinaryTreeNodesAverageAtEachLevel1(t *testing.T) {
	root := binarytree.IntNode{
		Value: 10,
		Left: &binarytree.IntNode{
			Value: 20,
			Left: &binarytree.IntNode{
				Value: 40,
			},
			Right: &binarytree.IntNode{
				Value: 50,
			},
		},
		Right: &binarytree.IntNode{
			Value: 30,
			Left: &binarytree.IntNode{
				Value: 60,
			},
			Right: &binarytree.IntNode{
				Value: 70,
			},
		},
	}
	binaryTree := binarytree.IntBinaryTree{
		Root: &root,
	}
	elements := queue.BinaryTreeNodesAverageAtEachLevel(binaryTree)
	expected := []float64{10.0, 25.0, 55.0}

	if !reflect.DeepEqual(elements, expected) {
		t.Fatalf("Expected %v, received %v", expected, elements)
	}
}

func TestBinaryTreeNodesAverageAtEachLevel2(t *testing.T) {
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
	binaryTree := binarytree.IntBinaryTree{
		Root: &root,
	}
	elements := queue.BinaryTreeNodesAverageAtEachLevel(binaryTree)
	expected := []float64{314, 6, 276.25, 12, 225, 641}

	if !reflect.DeepEqual(elements, expected) {
		t.Fatalf("Expected %v, received %v", expected, elements)
	}
}
