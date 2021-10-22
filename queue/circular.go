package queue

import (
	"errors"
)

type CircularQueue struct {
	Size   int
	front  int
	rear   int
	buffer []interface{}
}

func NewCircular(size int) *CircularQueue {
	return &CircularQueue{
		buffer: make([]interface{}, size),
		front:  0,
		rear:   0,
	}
}

func (c *CircularQueue) Enqueue(element interface{}) (bool, error) {
	if c.isFull() {
		return false, errors.New("overflow")
	}
	c.rear = c.rear + 1
	c.rear = c.rear % len(c.buffer)
	c.buffer[c.rear] = element
	return true, nil
}

func (c *CircularQueue) Dequeue() (interface{}, error) {
	if c.isEmpty() {
		return false, errors.New("underflow")
	}
	c.front = c.front + 1
	c.front = c.front % len(c.buffer)
	return c.buffer[c.front], nil
}

func (c *CircularQueue) isFull() bool {
	return (c.rear+1)%len(c.buffer) == c.front
}

func (c *CircularQueue) isEmpty() bool {
	return c.rear == c.front
}
