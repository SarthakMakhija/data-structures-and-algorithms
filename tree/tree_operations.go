package tree

type BinaryTree struct {
	Root *Node
}

type Node struct {
	Value string
	Left  *Node
	Right *Node
}

func (tree *BinaryTree) Traverse() string {
	if tree.Root == nil {
		return ""
	}

	var traverse_inner func(t *Node) string
	traverse_inner = func(t *Node) string {
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
