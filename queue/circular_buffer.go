package queue

import (
	"errors"
	"math"
)

type CircularBuffer struct {
	Size   int
	buffer []int
	front  int
	rear   int
}

const defaultSize = 10

func (c *CircularBuffer) Add(element int) (bool, error) {
	allocateBufferMemory := func() {
		if c.Size == 0 {
			c.buffer = make([]int, defaultSize)
			c.Size = defaultSize
		} else {
			c.buffer = make([]int, c.Size)
		}
	}
	if c.IsEmpty() {
		c.front, c.rear = -1, -1
		allocateBufferMemory()
	} else if c.IsFull() {
		return false, errors.New("OVERFLOW")
	}
	c.rear = int(math.Mod(float64(c.rear+1), float64(c.Size)))
	c.buffer[c.rear] = element

	return false, nil
}

func (c *CircularBuffer) Delete() (int, error) {
	if c.IsEmpty() {
		return -1, errors.New("UNDERFLOW")
	}
	c.front = int(math.Mod(float64(c.front+1), float64(c.Size)))
	element := c.buffer[c.front]

	//both of them are at len-1
	if (c.front == c.rear) && c.front == (len(c.buffer)-1) {
		c.front, c.rear = -1, -1
	}
	return element, nil
}

func (c *CircularBuffer) IsEmpty() bool {
	return (len(c.buffer) == 0) || (c.front == -1 && c.rear == -1)
}

func (c *CircularBuffer) IsFull() bool {
	return c.rear == c.Size-1 && c.front == -1
}
