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
	value   int
	next    *Node
	visited bool
}

func (l *LinkedList) Add(element int) *Node {

	node := Node{
		value:   element,
		next:    nil,
		visited: false,
	}
	if l.first == nil {
		l.first = &node
		l.current = l.first
	} else {
		l.current.next = &node
		l.current = &node
	}
	return &node
}

func (l *LinkedList) Add_With_Next(element int, next *Node) {

	node := Node{
		value:   element,
		next:    nil,
		visited: false,
	}
	if l.first == nil {
		l.first = &node
		l.current = l.first
	} else {
		l.current.next = next
		l.current = &node
	}
}

func (l *LinkedList) Add_Sorted(element int) {

	node := Node{
		value:   element,
		next:    nil,
		visited: false,
	}
	if l.first == nil {
		l.first = &node
		l.current = l.first
	} else {
		if element <= l.first.value {
			node.next = l.first
			l.first = &node
		} else if element >= l.current.value {
			l.current.next = &node
			l.current = &node
		} else {
			newFirst := l.first
			var add func(*Node)

			add = func(p *Node) {
				if p == nil || p.next == nil {
					return
				}
				if element >= p.value && element <= p.next.value {
					previous_next := p.next
					p.next = &node
					node.next = previous_next
				} else {
					add(p.next)
				}
			}
			add(newFirst)
		}
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

func (l *LinkedList) Contains_Cycle() bool {

	if l.first == nil {
		return false
	}

	var visit func(*Node) bool

	visit = func(p *Node) bool {
		if p == nil {
			return false
		}
		if p.visited {
			return true
		} else {
			p.visited = true
			return visit(p.next)
		}
	}
	return visit(l.first)
}

func (l *LinkedList) Reverse() *LinkedList {

	if l.first == nil {
		return nil
	}

	var list LinkedList = LinkedList{}
	var reverse_inner func(head *Node)

	reverse_inner = func(head *Node) {
		if head == nil {
			return
		} else {
			reverse_inner(head.next)
			node := Node{
				value:   head.value,
				visited: false,
				next:    nil,
			}
			if list.first == nil {
				list.first = &node
				list.current = &node
			} else {
				list.current.next = &node
				list.current = &node
			}
		}
	}
	reverse_inner(l.first)
	return &list
}
