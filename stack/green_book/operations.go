package green_book

type Stack struct {
	elements [10]int
	top      int
	minIndex int
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
