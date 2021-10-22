package binarytree

import (
	linearQueue "github.com/SarthakMakhija/data-structures-and-algorithms/queue"
)

type StringBinaryTreeWithLevelNext struct {
	Root *StringNodeWithLevelNext
}

type StringNodeWithLevelNext struct {
	Value   string
	Left    *StringNodeWithLevelNext
	Right   *StringNodeWithLevelNext
	Sibling *StringNodeWithLevelNext
}

func (t StringBinaryTreeWithLevelNext) SetSibling() {
	if t.Root == nil {
		return
	}

	potentialSiblingForRightChildOf := func(node *StringNodeWithLevelNext) *StringNodeWithLevelNext {
		potentialSibling := node.Sibling
		switch {
		case potentialSibling == nil:
			return nil
		case potentialSibling.Left != nil:
			return potentialSibling.Left
		case potentialSibling.Right != nil:
			return potentialSibling.Right
		default:
			return nil
		}
	}
	potentialSiblingForLeftChildOf := func(node *StringNodeWithLevelNext) *StringNodeWithLevelNext {
		potentialSibling := node.Right
		if potentialSibling == nil {
			return potentialSiblingForRightChildOf(node)
		}
		return potentialSibling
	}

	var setSiblingRightInner func(
		node *StringNodeWithLevelNext,
		potentialSibling *StringNodeWithLevelNext)

	setSiblingRightInner = func(
		node *StringNodeWithLevelNext,
		potentialSibling *StringNodeWithLevelNext) {

		if node == nil {
			return
		} else {
			node.Sibling = potentialSibling
		}
		//left subtree
		setSiblingRightInner(node.Left, potentialSiblingForLeftChildOf(node))
		//right subtree
		setSiblingRightInner(node.Right, potentialSiblingForRightChildOf(node))
	}
	setSiblingRightInner(t.Root, nil)
}

func (t StringBinaryTreeWithLevelNext) LevelOrderTraversal() []NodeSiblingPair {
	if t.Root == nil {
		return []NodeSiblingPair{}
	}
	queue := linearQueue.NewLinear(100)
	_, _ = queue.Enqueue(t.Root.Left)
	_, _ = queue.Enqueue(t.Root.Right)

	siblingOf := func(node *StringNodeWithLevelNext) string {
		sibling := "NA"
		if node.Sibling != nil {
			sibling = node.Sibling.Value
		}
		return sibling
	}

	var elements []NodeSiblingPair
	for !queue.IsEmpty() {
		node, _ := queue.Dequeue()
		stringNodeWithLevelNext := node.(*StringNodeWithLevelNext)
		elements = append(elements, NodeSiblingPair{
			Node:    stringNodeWithLevelNext.Value,
			Sibling: siblingOf(stringNodeWithLevelNext),
		})
		if stringNodeWithLevelNext.Left != nil {
			_, _ = queue.Enqueue(stringNodeWithLevelNext.Left)
		}
		if stringNodeWithLevelNext.Right != nil {
			_, _ = queue.Enqueue(stringNodeWithLevelNext.Right)
		}
	}
	return elements
}

type NodeSiblingPair struct {
	Node    string
	Sibling string
}
