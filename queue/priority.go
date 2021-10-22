package queue

import "errors"

//PriorityQueue
//lower value means higher priority
type PriorityQueue struct {
	elements []int
	rear     int
}

const rearInitialIndex = -1

func NewPriority(size int) *PriorityQueue {
	return &PriorityQueue{
		elements: make([]int, size),
		rear:     rearInitialIndex,
	}
}

func (p *PriorityQueue) Add(element int) (bool, error) {
	if p.isFull() {
		return false, errors.New("overflow")
	}
	if p.isEmpty() || element <= p.elements[p.rear] {
		p.add(element)
		return true, nil
	}

	p.add(element)
	for rearIndex := p.rear; rearIndex > 0; rearIndex-- {
		if p.elements[rearIndex] > p.elements[rearIndex-1] {
			p.elements[rearIndex], p.elements[rearIndex-1] = p.elements[rearIndex-1], p.elements[rearIndex]
		} else {
			break
		}
	}
	return true, nil
}

func (p *PriorityQueue) Get() (int, error) {
	if p.isEmpty() {
		return -1, errors.New("underflow")
	}
	element := p.elements[p.rear]
	p.rear = p.rear - 1
	return element, nil
}

func (p *PriorityQueue) All() []int {
	copied := make([]int, len(p.elements))
	copy(copied, p.elements)
	return copied
}

func (p *PriorityQueue) add(element int) {
	p.rear = p.rear + 1
	p.elements[p.rear] = element
}

func (p *PriorityQueue) isFull() bool {
	return p.rear == len(p.elements)-1
}

func (p *PriorityQueue) isEmpty() bool {
	return p.rear == rearInitialIndex
}
