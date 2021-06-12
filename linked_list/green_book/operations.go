package green_book

import (
	"fmt"
	"github.com/SarthakMakhija/data-structures-and-algorithms/stack"
	"math"
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
