package tree

type IntBinarySearchTree struct {
	Root *IntNode
}

type NavigationDirection int

const (
	Left NavigationDirection = iota
	Right
)

func (t IntBinarySearchTree) IsValid() bool {
	if t.Root == nil {
		return false
	}
	rootValue := t.Root.Value

	isValid := func(node *IntNode, direction NavigationDirection) bool {
		if node == nil || (node.Left == nil && node.Right == nil) {
			return true
		}
		if direction == Left &&
			(node.Left == nil || node.Left.Value < node.Value) &&
			(node.Right == nil || ((node.Right.Value > node.Value) && node.Right.Value < rootValue)) {
			return true
		}
		if direction == Right &&
			(node.Left == nil || (node.Left.Value < node.Value && node.Left.Value > rootValue)) &&
			(node.Right == nil || node.Right.Value > node.Value) {
			return true
		}
		return false
	}

	var isValidInner func(node *IntNode, direction NavigationDirection) bool
	isValidInner = func(node *IntNode, direction NavigationDirection) bool {
		if node == nil {
			return true
		} else {
			return isValidInner(node.Left, direction) && isValid(node, direction) && isValidInner(node.Right, direction)
		}
	}
	return isValidInner(t.Root.Left, Left) && isValidInner(t.Root.Right, Right)
}
