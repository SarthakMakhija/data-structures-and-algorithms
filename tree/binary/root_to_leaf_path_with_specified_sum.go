package binarytree

func (t IntBinaryTree) ExistsRootToLeafPathWith(sum int) bool {
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

type Path struct {
	NodeValues []int
}

func (t IntBinaryTree) AllPathsFromRootToLeafWith(sum int) []Path {
	if t.Root == nil {
		return []Path{}
	}

	var finalPaths []Path

	var rootToLeafPathWithSumInner func(node *IntNode, partialSum int, paths []int)
	rootToLeafPathWithSumInner = func(node *IntNode, partialSum int, paths []int) {
		if node == nil {
			return
		} else if node.Left == nil && node.Right == nil && node.Value+partialSum == sum {
			finalPaths = append(finalPaths, Path{
				NodeValues: append(paths, node.Value),
			})
		} else {
			rootToLeafPathWithSumInner(node.Left, partialSum+node.Value, append(paths, node.Value))
			rootToLeafPathWithSumInner(node.Right, partialSum+node.Value, append(paths, node.Value))
			return
		}
	}
	rootToLeafPathWithSumInner(t.Root, 0, []int{})
	return finalPaths
}
