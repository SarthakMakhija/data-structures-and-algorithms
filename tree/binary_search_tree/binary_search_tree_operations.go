package binarysearchtree

import (
	"strconv"
)

type IntBinarySearchTree struct {
	Root *IntNode
}

type IntNode struct {
	Value int
	Left  *IntNode
	Right *IntNode
}

func (tree *IntBinarySearchTree) Search(element int) bool {

	if tree.Root == nil {
		return false
	}
	var tree_inner func(t *IntNode) bool
	tree_inner = func(t *IntNode) bool {
		if t == nil {
			return false
		} else if t.Value == element {
			return true
		} else if element > t.Value {
			return tree_inner(t.Right)
		} else {
			return tree_inner(t.Left)
		}
	}
	return tree_inner(tree.Root)
}

func (tree *IntBinarySearchTree) InOrder_Traversal() string {

	if tree.Root == nil {
		return ""
	}
	var inorder func(t *IntNode) string
	inorder = func(t *IntNode) string {
		if t == nil {
			return ""
		}
		return inorder(t.Left) + strconv.Itoa(t.Value) + inorder(t.Right)
	}
	return inorder(tree.Root)
}

func (tree *IntBinarySearchTree) Insert(element int) {

	if tree.Root == nil {
		return
	}
	var targetNode func(t *IntNode) *IntNode
	targetNode = func(t *IntNode) *IntNode {
		if t == nil {
			return nil
		} else if t.Value == element {
			return t
		} else if element > t.Value {
			target := targetNode(t.Right)
			if target == nil {
				return t
			} else {
				return target
			}
		} else {
			target := targetNode(t.Left)
			if target == nil {
				return t
			} else {
				return target
			}
		}
	}
	target := targetNode(tree.Root)
	if target != nil && target.Value != element {
		node := IntNode{
			Value: element,
		}
		if element > target.Value {
			target.Right = &node
		} else {
			target.Left = &node
		}
	}
}
