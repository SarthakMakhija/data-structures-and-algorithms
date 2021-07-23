package binarytree

func (t IntBinaryTree) IsSymmetric() bool {
	root := t.Root
	if root == nil {
		return false
	}
	leftSubtreeNode := t.Root.Left
	rightSubtreeNode := t.Root.Right

	for leftSubtreeNode != nil && rightSubtreeNode != nil {
		if leftSubtreeNode.Value != rightSubtreeNode.Value {
			return false
		}
		leftSubtreeNode = leftSubtreeNode.Right
		rightSubtreeNode = rightSubtreeNode.Left
	}
	return leftSubtreeNode == nil && rightSubtreeNode == nil
}
