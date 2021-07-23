package binarytree

import (
	"github.com/SarthakMakhija/data-structures-and-algorithms/stack"
)

func (t StringBinaryTree) LeastCommonAncestorFor(value, anotherValue string) string {

	parentsOfValue := stack.StringStack{}
	parentsOfAnotherValue := stack.StringStack{}

	var parentsOf func(value string, node *StringNode, stack *stack.StringStack) bool
	parentsOf = func(value string, node *StringNode, stack *stack.StringStack) bool {
		if node == nil {
			return false
		} else if node.Value == value {
			return true
		} else {
			matched := parentsOf(value, node.Left, stack) || parentsOf(value, node.Right, stack)
			if matched {
				stack.Push(node.Value)
			}
			return matched
		}
	}

	lca := func() string {
		lca := ""
		for !parentsOfValue.IsEmpty() && !parentsOfAnotherValue.IsEmpty() {
			parentOfValue := parentsOfValue.Pop()
			parentOfAnotherValue := parentsOfAnotherValue.Pop()

			if parentOfValue == parentOfAnotherValue {
				lca = parentOfValue
			} else {
				break
			}
		}
		return lca
	}

	parentsOf(value, t.Root, &parentsOfValue)
	parentsOf(anotherValue, t.Root, &parentsOfAnotherValue)

	return lca()
}
