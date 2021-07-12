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

	isValid := func(node *IntNode, parent *IntNode, subtreeNavigation NavigationDirection, direction NavigationDirection) bool {
		if node == nil || (node.Left == nil && node.Right == nil) {
			return true
		}
		if subtreeNavigation == Left {
			if direction == Left {
				if (node.Left == nil || node.Left.Value < node.Value) &&
					(node.Right == nil || ((node.Right.Value > node.Value) && node.Right.Value < parent.Value)) {
					return true
				}
			} else {
				if (node.Left == nil || ((node.Left.Value < node.Value) && node.Left.Value > parent.Value)) &&
					(node.Right == nil || ((node.Right.Value > node.Value) && node.Right.Value < rootValue)) {
					return true
				}
			}
		} else {
			if direction == Left {
				if (node.Left == nil || ((node.Left.Value < node.Value) && node.Left.Value > rootValue)) &&
					(node.Right == nil || ((node.Right.Value > node.Value) && node.Right.Value < parent.Value)) {
					return true
				}
			} else {
				if (node.Left == nil || ((node.Left.Value < node.Value) && node.Left.Value > parent.Value)) &&
					(node.Right == nil || (node.Right.Value > node.Value)) {
					return true
				}
			}
		}

		return false
	}

	var isValidInner func(node *IntNode, parent *IntNode, subtreeNavigation NavigationDirection, nodeDirection NavigationDirection) bool
	isValidInner = func(node *IntNode, parent *IntNode, subtreeNavigation NavigationDirection, nodeDirection NavigationDirection) bool {
		if node == nil {
			return true
		} else {
			return isValidInner(node.Left, node, subtreeNavigation, Left) &&
				isValidInner(node.Right, node, subtreeNavigation, Right) &&
				isValid(node, parent, subtreeNavigation, nodeDirection)
		}
	}
	return isValidInner(t.Root.Left, t.Root, Left, Left) && isValidInner(t.Root.Right, t.Root, Right, Right)
}
