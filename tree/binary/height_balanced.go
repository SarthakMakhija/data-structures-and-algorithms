package binarytree

func (t StringBinaryTree) IsHeightBalanced() bool {

	var isHeightBalancedInner func(node *StringNode) (bool, int)
	isHeightBalancedInner = func(node *StringNode) (bool, int) {
		if node == nil {
			return true, 0
		}
		isHeightBalancedLeft, depthLeft := isHeightBalancedInner(node.Left)
		isHeightBalancedRight, depthRight := isHeightBalancedInner(node.Right)

		if isHeightBalancedLeft && isHeightBalancedRight && depthLeft-depthRight <= 1 {
			if depthLeft >= depthRight {
				return true, depthLeft + 1
			} else {
				return true, depthRight + 1
			}
		} else {
			return false, -1
		}
	}
	isHeightBalanced, _ := isHeightBalancedInner(t.Root)
	return isHeightBalanced
}
