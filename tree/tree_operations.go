package tree

type StringBinaryTree struct {
	Root *StringNode
}

type StringNode struct {
	Value string
	Left  *StringNode
	Right *StringNode
}

type IntBinaryTree struct {
	Root *IntNode
}

type IntNode struct {
	Value int
	Left  *IntNode
	Right *IntNode
}

func (tree *StringBinaryTree) Traverse() string {
	if tree.Root == nil {
		return ""
	}

	var traverse_inner func(t *StringNode) string
	traverse_inner = func(t *StringNode) string {
		if t == nil {
			return ""
		} else if t.Left == nil && t.Right == nil {
			return t.Value
		} else {
			return traverse_inner(t.Left) + t.Value + traverse_inner(t.Right)
		}
	}
	return traverse_inner(tree.Root)
}

func (tree *IntBinaryTree) Sum() int {
	if tree.Root == nil {
		return 0
	}

	var traverse_inner func(t *IntNode) int
	traverse_inner = func(t *IntNode) int {
		if t == nil {
			return 0
		} else if t.Left == nil && t.Right == nil {
			return t.Value
		} else {
			return traverse_inner(t.Left) + t.Value + traverse_inner(t.Right)
		}
	}
	return traverse_inner(tree.Root)
}
