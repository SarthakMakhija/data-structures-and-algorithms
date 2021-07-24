package binarytree_test

import (
	binarytree "github.com/SarthakMakhija/data-structures-and-algorithms/tree/binary"
	"reflect"
	"testing"
)

func TestInorderWithoutRecursionWithO1Space1(t *testing.T) {
	root := binarytree.IntNodeWithParent{
		Value: 10,
	}
	node20 := binarytree.IntNodeWithParent{
		Value:  20,
		Parent: &root,
	}
	node30 := binarytree.IntNodeWithParent{
		Value:  30,
		Parent: &root,
	}
	root.Left = &node20
	root.Right = &node30

	tree := binarytree.IntBinaryTreeWithParent{
		Root: &root,
	}
	output := tree.InorderWithoutRecursionWithO1Space()
	expected := []int{20, 10, 30}

	if !reflect.DeepEqual(output, expected) {
		t.Fatalf("Expected %v, received %v", expected, output)
	}
}

func TestInorderWithoutRecursionWithO1Space2(t *testing.T) {
	root := binarytree.IntNodeWithParent{
		Value: 1,
	}
	node2 := binarytree.IntNodeWithParent{
		Value:  2,
		Parent: &root,
	}
	node3 := binarytree.IntNodeWithParent{
		Value:  3,
		Parent: &root,
	}
	node7 := binarytree.IntNodeWithParent{
		Value:  7,
		Parent: &node3,
	}
	node4 := binarytree.IntNodeWithParent{
		Value:  4,
		Parent: &node2,
	}
	node5 := binarytree.IntNodeWithParent{
		Value:  5,
		Parent: &node2,
	}
	node11 := binarytree.IntNodeWithParent{
		Value:  11,
		Parent: &node5,
	}
	root.Left = &node2
	root.Right = &node3
	node2.Left = &node4
	node2.Right = &node5
	node5.Left = &node11
	node3.Right = &node7

	tree := binarytree.IntBinaryTreeWithParent{
		Root: &root,
	}
	output := tree.InorderWithoutRecursionWithO1Space()
	expected := []int{4, 2, 11, 5, 1, 3, 7}

	if !reflect.DeepEqual(output, expected) {
		t.Fatalf("Expected %v, received %v", expected, output)
	}
}

func TestInorderWithoutRecursionWithO1Space3(t *testing.T) {
	root := binarytree.IntNodeWithParent{
		Value: 314,
	}
	node6Left := binarytree.IntNodeWithParent{
		Value:  6,
		Parent: &root,
	}
	node6Right := binarytree.IntNodeWithParent{
		Value:  6,
		Parent: &root,
	}
	node271Left := binarytree.IntNodeWithParent{
		Value:  271,
		Parent: &node6Left,
	}
	node28Left := binarytree.IntNodeWithParent{
		Value:  28,
		Parent: &node271Left,
	}
	node0 := binarytree.IntNodeWithParent{
		Value:  0,
		Parent: &node271Left,
	}
	node561 := binarytree.IntNodeWithParent{
		Value:  561,
		Parent: &node6Left,
	}
	node3 := binarytree.IntNodeWithParent{
		Value:  3,
		Parent: &node561,
	}
	node17 := binarytree.IntNodeWithParent{
		Value:  17,
		Parent: &node3,
	}
	node2 := binarytree.IntNodeWithParent{
		Value:  2,
		Parent: &node6Right,
	}
	node271Right := binarytree.IntNodeWithParent{
		Value:  271,
		Parent: &node6Right,
	}
	node28Right := binarytree.IntNodeWithParent{
		Value:  28,
		Parent: &node271Right,
	}
	node1 := binarytree.IntNodeWithParent{
		Value:  1,
		Parent: &node2,
	}
	node401 := binarytree.IntNodeWithParent{
		Value:  401,
		Parent: &node1,
	}
	node641 := binarytree.IntNodeWithParent{
		Value:  641,
		Parent: &node401,
	}
	node257 := binarytree.IntNodeWithParent{
		Value:  257,
		Parent: &node1,
	}
	root.Left = &node6Left
	root.Right = &node6Right

	node6Left.Left = &node271Left
	node6Left.Right = &node561

	node271Left.Left = &node28Left
	node271Left.Right = &node0

	node561.Right = &node3
	node3.Left = &node17

	node6Right.Left = &node2
	node2.Right = &node1
	node1.Left = &node401
	node1.Right = &node257
	node401.Right = &node641

	node6Right.Right = &node271Right
	node271Right.Right = &node28Right

	tree := binarytree.IntBinaryTreeWithParent{
		Root: &root,
	}
	output := tree.InorderWithoutRecursionWithO1Space()
	expected := []int{28, 271, 0, 6, 561, 17, 3, 314, 2, 401, 641, 1, 257, 6, 271, 28}

	if !reflect.DeepEqual(output, expected) {
		t.Fatalf("Expected %v, received %v", expected, output)
	}
}
