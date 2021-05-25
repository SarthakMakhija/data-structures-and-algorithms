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
	var treeInner func(t *IntNode) bool
	treeInner = func(t *IntNode) bool {
		if t == nil {
			return false
		} else if t.Value == element {
			return true
		} else if element > t.Value {
			return treeInner(t.Right)
		} else {
			return treeInner(t.Left)
		}
	}
	return treeInner(tree.Root)
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

func (tree *IntBinarySearchTree) Insert1(element int) {

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

func (tree *IntBinarySearchTree) Insert2(element int) {

	if tree.Root == nil {
		return
	}
	var insertInner func(t *IntNode) *IntNode
	var node *IntNode

	insertInner = func(t *IntNode) *IntNode {
		if t == nil {
			node = &IntNode{
				Value: element,
			}
			return node
		}
		if element > t.Value {
			t.Right = insertInner(t.Right)
		} else if element < t.Value {
			t.Left = insertInner(t.Left)
		}
		return t
	}
	insertInner(tree.Root)
}
