package stack_based

import (
	"github.com/SarthakMakhija/data-structures-and-algorithms/stack"
)

type StackBasedQueue struct {
	WriteStack *stack.IntStack
	ReadStack  *stack.IntStack
	front      int
}

func (s *StackBasedQueue) Add(element int) {
	s.WriteStack.Push(element)
}

//Get
//Requires entire stack to be copied, we can optimize by copying uptil front position, but copy is needed
func (s *StackBasedQueue) Get() int {
	s.ReadStack.Clear()
	writtenElements := s.WriteStack.All()

	for index := 0; index < len(writtenElements)-s.front; index++ {
		s.ReadStack.Push(writtenElements[index])
	}
	s.front = s.front + 1
	return s.ReadStack.Pop()
}
