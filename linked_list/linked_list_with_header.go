package linkedlist

import "math"

type LinkedListWithHeader struct {
	header     *HeaderNode
	firstValue *ValueNode
}

type HeaderNode struct {
	size           int
	evenValueCount int
	oddValueCount  int
	next           *ValueNode
}

type ValueNode struct {
	value int
	next  *ValueNode
}

type HeaderView struct {
	Size           int
	EvenValueCount int
	OddValueCount  int
}

func (h *HeaderNode) toHeaderView() HeaderView {
	return HeaderView{
		Size:           h.size,
		EvenValueCount: h.evenValueCount,
		OddValueCount:  h.oddValueCount,
	}
}

func (l *LinkedListWithHeader) Add(value int) {
	addHeaderNode := func() {
		l.header = &HeaderNode{
			size: 0,
		}
	}
	attachValueNode := func(valueNode *ValueNode) {
		l.header.next = valueNode
	}
	addValueNode := func() *ValueNode {
		valueNode := &ValueNode{
			value: value,
			next:  nil,
		}
		first := l.firstValue
		if first == nil {
			l.firstValue = valueNode
		} else {
			for first.next != nil {
				first = first.next
			}
			first.next = valueNode
		}
		return valueNode
	}
	updateHeaderContent := func() {
		l.header.size = l.header.size + 1
		if int(math.Mod(float64(value), float64(2))) == 0 {
			l.header.evenValueCount = l.header.evenValueCount + 1
		} else {
			l.header.oddValueCount = l.header.oddValueCount + 1
		}
	}
	valueNode := addValueNode()
	if l.header == nil {
		addHeaderNode()
		attachValueNode(valueNode)
	}
	updateHeaderContent()
}

func (l *LinkedListWithHeader) HeaderView() HeaderView {
	return l.header.toHeaderView()
}

func (l *LinkedListWithHeader) All() []int {
	var allElements []int

	first := l.firstValue
	for first != nil {
		allElements = append(allElements, first.value)
		first = first.next
	}
	return allElements
}
