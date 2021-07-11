package tree

import "strconv"

func (t *IntBinaryTree) Invert() {
	if t.Root == nil {
		return
	}

	var invertInner func(node *IntNode)

	swap := func(node *IntNode) {
		originalLeft := node.Left
		node.Left = node.Right
		node.Right = originalLeft
	}
	invertInner = func(node *IntNode) {
		if node == nil {
			return
		}
		invertInner(node.Left)
		invertInner(node.Right)
		swap(node)
	}
	invertInner(t.Root)
}

func (t *IntBinaryTree) Traverse() string {
	if t.Root == nil {
		return ""
	}

	var traverseInner func(node *IntNode) string
	traverseInner = func(node *IntNode) string {
		if node == nil {
			return ""
		}
		return traverseInner(node.Left) + strconv.Itoa(node.Value) + traverseInner(node.Right)
	}
	return traverseInner(t.Root)
}
