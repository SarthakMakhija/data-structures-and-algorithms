package green_book

import (
	"github.com/SarthakMakhija/data-structures-and-algorithms/stack"
)

type Stack struct {
	elements [10]int
	top      int
	minIndex int
}

type SortedStack struct {
	elements [10]int
	top      int
}

//Push
//Assume no overflow
func (s *Stack) Push(element int) {
	s.elements[s.top] = element
	if s.top != 0 {
		if element < s.elements[s.minIndex] {
			s.minIndex = s.top
		}
	}
	s.top = s.top + 1
}

func (s *Stack) Pop() int {
	if s.top == 0 {
		return -1
	}
	element := s.elements[s.top-1]
	s.top = s.top - 1
	return element
}

func (s *Stack) Min() int {
	return s.elements[s.minIndex]
}

//Push
//Assume no overflow
func (s *SortedStack) Push(element int) {
	add := func(e int) {
		s.elements[s.top] = e
		s.top = s.top + 1
	}
	if s.top == 0 || s.elements[s.top-1] > element {
		add(element)
	} else {
		backUpStack := stack.IntStack{}
		for s.top > 0 && s.elements[s.top-1] < element {
			e := s.Pop()
			backUpStack.Push(e)
		}
		add(element)
		e := backUpStack.Pop()
		for e != -1 {
			add(e)
			e = backUpStack.Pop()
		}
	}
}

func (s *SortedStack) Pop() int {
	if s.top == 0 {
		return -1
	}
	element := s.elements[s.top-1]
	s.top = s.top - 1
	return element
}
