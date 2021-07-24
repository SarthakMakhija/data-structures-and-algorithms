package green_book

import (
	"fmt"
	"math"

	"github.com/SarthakMakhija/data-structures-and-algorithms/stack"
)

type LinkedList struct {
	first   *Node
	current *Node
}

type Node struct {
	value int
	next  *Node
}

func (l *LinkedList) Add(element int) *Node {
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
	return &node
}

func (l *LinkedList) AddWithNext(element int, next *Node) {

	node := Node{
		value: element,
		next:  nil,
	}
	if l.first == nil {
		l.first = &node
		l.current = l.first
	} else {
		l.current.next = next
		l.current = &node
	}
}

//AddDigitsOf
//If the numbers (say 617) were represented as 7 -> 1 -> 6, then one of the solutions could have been
//to reverse the list
func (l LinkedList) AddDigitsOf(other LinkedList) *LinkedList {

	carryOver := 0
	additionList := &LinkedList{}

	firstStack := &stack.IntStack{}
	otherStack := &stack.IntStack{}

	listToStack := func(p LinkedList, stack *stack.IntStack) {
		head := p.first
		for head != nil {
			stack.Push(head.value)
			head = head.next
		}
	}
	addNodeToResultantListUsing := func(digit int) {
		node := &Node{
			value: digit,
		}
		if additionList.first == nil {
			additionList.first = node
		} else {
			node.next = additionList.first
			additionList.first = node
		}
	}
	popAndAddResultToList := func() {
		firstValue := firstStack.Pop()
		secondValue := otherStack.Pop()
		for firstValue != -1 || secondValue != -1 {
			if firstValue == -1 {
				firstValue = 0
			}
			if secondValue == -1 {
				secondValue = 0
			}
			sum := firstValue + secondValue + carryOver
			carryOver = sum / 10
			addNodeToResultantListUsing(int(math.Mod(float64(sum), float64(10))))

			firstValue = firstStack.Pop()
			secondValue = otherStack.Pop()
		}
	}
	listToStack(l, firstStack)
	listToStack(other, otherStack)
	popAndAddResultToList()

	if carryOver > 0 {
		addNodeToResultantListUsing(carryOver)
	}
	return additionList
}

func (l *LinkedList) AllAsString() string {
	if l.first == nil {
		return ""
	}
	var iterate func(*Node) string
	var output = ""

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

func (l *LinkedList) HasCycle() bool {
	var tail *Node = nil
	head := l.first

	moveHead2Steps := func() {
		head = head.next
		if head != nil {
			head = head.next
		}
	}
	moveTail1Step := func() {
		if tail == nil {
			tail = l.first
		} else {
			tail = tail.next
		}
	}
	for head != nil {
		moveHead2Steps()
		moveTail1Step()
		if head == nil {
			return false
		}
		if head == tail || head.next == tail {
			return true
		}
	}
	return false
}
