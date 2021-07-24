package binarytree

func (t IntBinaryTree) ExistsRootToLeafPathWithSum(sum int) bool {
	if t.Root == nil {
		return false
	}
	var rootToLeafPathWithSumInner func(node *IntNode, partialSum int) bool
	rootToLeafPathWithSumInner = func(node *IntNode, partialSum int) bool {
		if node == nil {
			return false
		} else if node.Left == nil && node.Right == nil && node.Value+partialSum == sum {
			return true
		} else {
			return rootToLeafPathWithSumInner(node.Left, partialSum+node.Value) ||
				rootToLeafPathWithSumInner(node.Right, partialSum+node.Value)
		}
	}
	return rootToLeafPathWithSumInner(t.Root, 0)
}
