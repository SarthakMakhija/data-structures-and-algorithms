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

func (l *LinkedList) AddWithNext(element int, next *Node) {

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

func (l *LinkedList) AddSorted(element int) {

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

func (l *LinkedList) AllAsString() string {
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

	var maxInner func(*Node) int
	var max int = l.first.value

	maxInner = func(p *Node) int {
		if p.value > max {
			max = p.value
		}
		if p.next == nil {
			return max
		} else {
			return maxInner(p.next)
		}
	}
	return maxInner(l.first), nil
}

func (l *LinkedList) Contains(element int) bool {
	if l.first == nil {
		return false
	}

	var containsInner func(*Node) bool

	containsInner = func(p *Node) bool {
		if p == nil {
			return false
		}
		if p.value == element {
			return true
		} else {
			return containsInner(p.next)
		}
	}
	return containsInner(l.first)
}

func (l *LinkedList) ContainsCycle() bool {

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
	var reverseInner func(head *Node)

	reverseInner = func(head *Node) {
		if head == nil {
			return
		} else {
			reverseInner(head.next)
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
	reverseInner(l.first)
	return &list
}

func (l *LinkedList) RemoveDuplicatesFromSorted() {

	if l.first == nil {
		return
	}

	currentValue := l.first.value
	previous := l.first

	for head := l.first.next; head != nil; head = head.next {

		for head.next != nil && head.value == currentValue {
			head = head.next
		}
		if head.value == currentValue {
			previous.next = head.next
		} else {
			previous.next = head
		}
		previous = previous.next
		currentValue = head.value
	}
}

func (l *LinkedList) MidValue() int {

	if l.first == nil {
		return 0
	}

	follower := l.first
	head := l.first

	for head.next != nil {
		head = head.next
		if head.next != nil {
			head = head.next
			follower = follower.next
		} else {
			break
		}
	}
	return follower.value
}

//Intersection
/**
this implementation relies on an additional property visited to be presented in the list
other option is iterate through both the lists one by one and collect all the elements
(not the element values but nodes itself) in stack1 and stack2
let's consider 2 lists :
	1->2->3->4->5
	6->7->3 (here 3 means the node 3 from list1)
which means once we have created 2 stacks, if there is an intersection we can compare stack top of both the stacks
and keep repeating it until there is no match
This means extra space and extra iterations
**/
func (l *LinkedList) Intersection(other *LinkedList) int {

	head1 := l.first
	head2 := other.first

	var intersectionValue = -1

	for head1 != nil {
		head1.visited = true
		head1 = head1.next
	}
	for head2 != nil {
		if head2.visited {
			intersectionValue = head2.value
			break
		}
		head2.visited = true
		head2 = head2.next
	}

	return intersectionValue
}
