package linkedlist

import (
	"errors"
	"fmt"
)

type LinkedList struct {
	first   *Node
	current *Node
}

type Node struct {
	value int
	next  *Node
}

func (l *LinkedList) Add(element int) {

	node := Node{
		value: element,
		next:  nil,
	}
	if l.first == nil {
		l.first = &node
		l.current = l.first
	} else {
		l.current.next = &node
		l.current = &node
	}
}

func (l *LinkedList) All_As_String() string {
	if l.first == nil {
		return ""
	}

	var iterate func(*Node) string
	var output string = ""

	iterate = func(p *Node) string {
		output = output + fmt.Sprint(p.value)
		if p.next == nil {
			return output
		} else {
			return iterate(p.next)
		}
	}
	return iterate(l.first)
}

func (l *LinkedList) Max() (int, error) {
	if l.first == nil {
		return 0, errors.New("EMPTY LIST, NO MAXIMUM")
	}

	var max_inner func(*Node) int
	var max int = l.first.value

	max_inner = func(p *Node) int {
		if p.value > max {
			max = p.value
		}
		if p.next == nil {
			return max
		} else {
			return max_inner(p.next)
		}
	}
	return max_inner(l.first), nil
}

func (l *LinkedList) Contains(element int) bool {
	if l.first == nil {
		return false
	}

	var contains_inner func(*Node) bool

	contains_inner = func(p *Node) bool {
		if p == nil {
			return false
		}
		if p.value == element {
			return true
		} else {
			return contains_inner(p.next)
		}
	}
	return contains_inner(l.first)
}
