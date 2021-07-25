package binarytree

import "errors"

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
	queue := queue{
		Size: 100,
	}
	_, _ = queue.enqueue(t.Root.Left)
	_, _ = queue.enqueue(t.Root.Right)

	siblingOf := func(node *StringNodeWithLevelNext) string {
		sibling := "NA"
		if node.Sibling != nil {
			sibling = node.Sibling.Value
		}
		return sibling
	}

	var elements []NodeSiblingPair
	for !queue.isEmpty() {
		node, _ := queue.dequeue()
		elements = append(elements, NodeSiblingPair{
			Node:    node.Value,
			Sibling: siblingOf(node),
		})
		if node.Left != nil {
			_, _ = queue.enqueue(node.Left)
		}
		if node.Right != nil {
			_, _ = queue.enqueue(node.Right)
		}
	}
	return elements
}

type NodeSiblingPair struct {
	Node    string
	Sibling string
}

type queue struct {
	elements []*StringNodeWithLevelNext
	front    int
	rear     int
	Size     int
}

func (l *queue) enqueue(element *StringNodeWithLevelNext) (bool, error) {
	const fixedSize = 5
	if l.isEmpty() {
		var size = l.Size
		if l.Size == 0 {
			size = fixedSize
		}
		l.elements = make([]*StringNodeWithLevelNext, size)
		l.front = 0
		l.rear = -1
		l.Size = size
	} else if l.isFull() {
		return false, errors.New("OVERFLOW")
	}
	l.rear = l.rear + 1
	l.elements[l.rear] = element
	return true, nil
}

func (l *queue) dequeue() (*StringNodeWithLevelNext, error) {
	if l.isEmpty() {
		return nil, errors.New("EMPTY")
	}
	element := l.elements[l.front]
	l.front = l.front + 1
	return element, nil
}

func (l *queue) isEmpty() bool {
	return len(l.elements) == 0 || l.front > l.rear
}

func (l *queue) isFull() bool {
	return l.rear == l.Size-1
}
