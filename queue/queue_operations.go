package queue

import (
	"errors"
)

type LinearQueue struct {
	elements []interface{}
	front    int
	rear     int
	read     int
}

func NewLinear(size int) *LinearQueue {
	return &LinearQueue{
		elements: make([]interface{}, size),
		front:    -1,
		rear:     -1,
		read:     -1,
	}
}

func (l *LinearQueue) Enqueue(element interface{}) (bool, error) {
	if l.isFull() {
		return false, errors.New("overflow")
	}
	l.rear = l.rear + 1
	l.elements[l.rear] = element
	return true, nil
}

func (l *LinearQueue) Dequeue() (interface{}, error) {
	if l.isEmpty() {
		return false, errors.New("underflow")
	}
	l.front = l.front + 1
	return l.elements[l.front], nil
}

func (l *LinearQueue) AllElements() []interface{} {
	var elements []interface{}
	for l.read < l.rear {
		e, _ := l.get()
		elements = append(elements, e)
	}
	return elements
}

func (l *LinearQueue) get() (interface{}, error) {
	if l.isEmpty() {
		return false, errors.New("underflow")
	}
	l.read = l.read + 1
	return l.elements[l.read], nil
}

func (l *LinearQueue) isFull() bool {
	return l.rear == len(l.elements)-1
}

func (l *LinearQueue) isEmpty() bool {
	return l.rear == l.front
}
