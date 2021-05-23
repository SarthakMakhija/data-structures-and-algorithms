package binarysearchtree

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
