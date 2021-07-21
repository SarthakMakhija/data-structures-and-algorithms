package stack

type JumpNode struct {
	Value rune
	Order int
	Jump  *JumpNode
	Next  *JumpNode
}

type JumpNodeStack struct {
	stack    []*JumpNode
	stackTop int
}

func (s *JumpNodeStack) Push(element *JumpNode) {
	s.stack = append(s.stack, element)
	s.stackTop = s.stackTop + 1
}

func (s *JumpNodeStack) Pop() *JumpNode {
	s.stackTop = s.stackTop - 1
	if s.stackTop < 0 {
		return nil
	}
	top := s.stack[s.stackTop]
	s.stack = s.stack[0:s.stackTop]
	return top
}

//TraverseJumpFirstOrder
//Tries to mimic recursion using an iterative approach
func (j *JumpNode) TraverseJumpFirstOrder() []rune {
	var values []rune
	stack := JumpNodeStack{}

	push := func(node *JumpNode) {
		stack.Push(node)
	}
	appendValues := func(v rune) {
		values = append(values, v)
	}
	traverseJumpFirst := func() {
		head := j
		for head != nil && head.Order == -1 {
			push(head)
			appendValues(head.Value)
			head.Order = head.Order + 1
			head = head.Jump
		}
	}
	traverseNext := func() {
		for node := stack.Pop(); node != nil; node = stack.Pop() {
			next := node.Next
			if next != nil && next.Order == -1 {
				appendValues(next.Value)
				next.Order = next.Order + 1
				push(next)
			}
		}
	}

	traverseJumpFirst()
	traverseNext()

	return values
}
