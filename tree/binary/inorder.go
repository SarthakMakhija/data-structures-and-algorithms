package binarytree

type IntBinaryTreeWithParent struct {
	Root *IntNodeWithParent
}

type IntNodeWithParent struct {
	Value  int
	Left   *IntNodeWithParent
	Right  *IntNodeWithParent
	Parent *IntNodeWithParent
}

func (t IntBinaryTreeWithParent) InorderWithoutRecursionWithO1Space() []int {
	if t.Root == nil {
		return []int{}
	}
	var nodeValues []int
	traverse := func(node *IntNodeWithParent, parent *IntNodeWithParent) []int {
		head := node

		moveLeft := func() {
			head = head.Left
		}
		moveRight := func() {
			head = head.Right
		}
		moveUp := func() {
			if head == head.Parent.Left {
				head = head.Parent
			} else {
				for head != node &&
					head != head.Parent.Left {
					head = head.Parent
				}
				head = head.Parent
			}
		}
		var values []int
		for head != t.Root {
			if head.Left != nil {
				moveLeft()
			} else {
				values = append(values, head.Value)
				for head.Right == nil {
					moveUp()
					if head != t.Root {
						values = append(values, head.Value)
					} else {
						break
					}
				}
				if head != t.Root {
					moveRight()
				}
			}
		}
		return values
	}

	nodeValues = append(nodeValues, traverse(t.Root.Left, t.Root)...)
	nodeValues = append(nodeValues, t.Root.Value)
	nodeValues = append(nodeValues, traverse(t.Root.Right, t.Root)...)
	return nodeValues
}
