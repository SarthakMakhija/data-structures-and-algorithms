package tree

func (t IntBinaryTree) Flatten() {
	if t.Root == nil {
		return
	}
	detachRightOf := func(node *IntNode) {
		ptr := node
		for ptr.Left != nil {
			ptr = ptr.Left
		}
		ptr.Left = node.Right
		node.Right = nil
	}
	var flattenInner func(node *IntNode)
	flattenInner = func(node *IntNode) {
		if node == nil {
			return
		} else {
			flattenInner(node.Left)
			flattenInner(node.Right)
			detachRightOf(node)
		}
	}
	flattenInner(t.Root)
}

func (t IntBinaryTree) FlattenedRead() []int {
	var elements []int

	ptr := t.Root
	for ptr != nil {
		elements = append(elements, ptr.Value)
		ptr = ptr.Left
	}
	return elements
}
