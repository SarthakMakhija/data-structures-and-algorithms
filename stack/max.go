package stack

type MaxStack struct {
	stack    []*ElementWithMax
	stackTop int
	peekTop  int
}

type ElementWithMax struct {
	element int
	Max     int
}

func (s *MaxStack) Push(element int) {
	elementWithMax := s.Peek()

	max := element
	if elementWithMax == nil || element > elementWithMax.Max {
		max = element
	} else {
		max = elementWithMax.Max
	}
	s.stack = append(s.stack, &ElementWithMax{
		element: element,
		Max:     max,
	})
	s.stackTop = s.stackTop + 1
	s.peekTop = s.stackTop
}

func (s *MaxStack) Pop() *ElementWithMax {
	s.stackTop = s.stackTop - 1
	s.peekTop = s.stackTop
	if s.stackTop < 0 {
		return nil
	}
	top := s.stack[s.stackTop]
	s.stack = s.stack[0:s.stackTop]
	return top
}

func (s *MaxStack) Peek() *ElementWithMax {
	s.peekTop = s.peekTop - 1
	if s.peekTop < 0 {
		return nil
	}
	top := s.stack[s.peekTop]
	return top
}

func (s *MaxStack) Max() *ElementWithMax {
	return s.Peek()
}
