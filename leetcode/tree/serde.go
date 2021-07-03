package tree

import "strconv"

type IntBinaryTree struct {
	Root *IntNode
}

type IntNode struct {
	Value int
	Left  *IntNode
	Right *IntNode
}

func (t *IntBinaryTree) Serialize() string {
	if t.Root == nil {
		return ""
	}
	var serializeInner func(node *IntNode) string
	serializeInner = func(node *IntNode) string {
		if node == nil {
			return "X"
		} else {
			return strconv.Itoa(node.Value) + serializeInner(node.Left) + serializeInner(node.Right)
		}
	}
	return serializeInner(t.Root)
}
