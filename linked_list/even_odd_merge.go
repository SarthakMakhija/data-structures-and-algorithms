package linkedlist

import "math"

func (l LinkedList) EvenOddMerge() LinkedList {
	if l.first == nil {
		return LinkedList{}
	}

	var evenIterator = &Node{}
	var oddIterator = &Node{}

	var evenHead = evenIterator
	var oddHead = oddIterator

	head := l.first
	position := 0
	for head != nil {
		if math.Mod(float64(position), 2) == 0 {
			evenIterator.next = head
			evenIterator = head
		} else {
			oddIterator.next = head
			oddIterator = head
		}
		position = position + 1
		head = head.next
	}
	oddIterator.next = nil
	evenIterator.next = oddHead.next

	return LinkedList{first: evenHead.next}
}
