package tree_test

import (
	serde "github.com/SarthakMakhija/data-structures-and-algorithms/leetcode/tree"
	"reflect"
	"testing"
)

func TestIntBinaryTreeFlatten1(t *testing.T) {
	binaryTree := serde.IntBinaryTree{
		Root: &serde.IntNode{
			Value: 1,
			Left: &serde.IntNode{
				Value: 2,
				Left: &serde.IntNode{
					Value: 3,
				},
				Right: &serde.IntNode{
					Value: 4,
				},
			},
			Right: &serde.IntNode{
				Value: 5,
				Right: &serde.IntNode{
					Value: 6,
				},
			},
		},
	}
	binaryTree.Flatten()

	elements := binaryTree.FlattenedRead()
	expected := []int{1, 2, 3, 4, 5, 6}

	if !reflect.DeepEqual(elements, expected) {
		t.Fatalf("Expected %v, received %v", expected, elements)
	}
}

func TestIntBinaryTreeFlatten2(t *testing.T) {
	binaryTree := serde.IntBinaryTree{
		Root: &serde.IntNode{
			Value: 1,
			Left: &serde.IntNode{
				Value: 2,
				Left: &serde.IntNode{
					Value: 3,
					Left: &serde.IntNode{
						Value: 8,
					},
				},
				Right: &serde.IntNode{
					Value: 4,
					Left: &serde.IntNode{
						Value: 16,
					},
					Right: &serde.IntNode{
						Value: 14,
						Left: &serde.IntNode{
							Value: 20,
						},
						Right: &serde.IntNode{
							Value: 21,
						},
					},
				},
			},
			Right: &serde.IntNode{
				Value: 5,
				Left: &serde.IntNode{
					Value: 7,
				},
				Right: &serde.IntNode{
					Value: 6,
				},
			},
		},
	}
	binaryTree.Flatten()

	elements := binaryTree.FlattenedRead()
	expected := []int{1, 2, 3, 8, 4, 16, 14, 20, 21, 5, 7, 6}

	if !reflect.DeepEqual(elements, expected) {
		t.Fatalf("Expected %v, received %v", expected, elements)
	}
}
