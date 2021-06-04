package queue

import (
	"errors"
)

type LinearQueue struct {
	elements []int
	front    int
	rear     int
	Size     int
}

func (l *LinearQueue) Add(element int) (bool, error) {
	const fixedSize = 5
	if l.IsEmpty() {
		var size = l.Size
		if l.Size == 0 {
			size = fixedSize
		}
		l.elements = make([]int, size)
		l.front = 0
		l.rear = -1
		l.Size = size
	} else if l.IsFull() {
		return false, errors.New("OVERFLOW")
	}
	l.rear = l.rear + 1
	l.elements[l.rear] = element
	return true, nil
}

func (l *LinearQueue) Get() (int, error) {
	if l.IsEmpty() {
		return -1, errors.New("EMPTY")
	}
	element := l.elements[l.front]
	l.front = l.front + 1
	return element, nil
}

func (l *LinearQueue) All() []int {
	var elements []int
	for l.front <= l.rear {
		e, _ := l.Get()
		elements = append(elements, e)
	}
	return elements
}

func (l *LinearQueue) IsEmpty() bool {
	return len(l.elements) == 0
}

func (l *LinearQueue) IsFull() bool {
	return l.rear == l.Size-1
}
