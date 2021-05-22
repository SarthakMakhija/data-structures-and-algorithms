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

func (tree *IntBinaryTree) Count_Nodes() int {
	if tree.Root == nil {
		return 0
	}

	var count_inner func(t *IntNode) int
	count_inner = func(t *IntNode) int {
		if t == nil {
			return 0
		} else if t.Left == nil && t.Right == nil {
			return 1
		} else {
			return count_inner(t.Left) + 1 + count_inner(t.Right)
		}
	}
	return count_inner(tree.Root)
}

func (tree *IntBinaryTree) Count_Leaf_Nodes() int {
	if tree.Root == nil {
		return 0
	}

	var count_inner func(t *IntNode) int
	count_inner = func(t *IntNode) int {
		if t == nil {
			return 0
		} else if t.Left == nil && t.Right == nil {
			return 1
		} else {
			return count_inner(t.Left) + count_inner(t.Right)
		}
	}
	return count_inner(tree.Root)
}

func (tree *IntBinaryTree) Height() int {
	if tree.Root == nil {
		return 0
	}
	var max_height int = 0
	var height_inner func(*IntNode, int)

	height_inner = func(t *IntNode, height int) {
		if t == nil {
		} else if t.Left == nil && t.Right == nil {
			if height > max_height {
				max_height = height
			}
		} else {
			height_inner(t.Left, height+1)
			height_inner(t.Right, height+1)
		}
	}
	height_inner(tree.Root, 0)
	return max_height
}
