package linkedlist

import (
	"strconv"
)

type PolynomialList struct {
	first *PolynomialNode
}

type PolynomialNode struct {
	coefficient int
	exponent    int
	next        *PolynomialNode
}

func (p *PolynomialList) Add(coefficient int, exponent int) {
	node := &PolynomialNode{
		coefficient: coefficient,
		exponent:    exponent,
	}
	if p.first == nil {
		p.first = node
	} else {
		if exponent > p.first.exponent {
			node.next = p.first
			p.first = node
		} else if exponent == p.first.exponent {
			p.first.coefficient = p.first.coefficient + coefficient
		} else {
			head := p.first
			previous := head
			for head != nil && exponent < head.exponent {
				previous = head
				head = head.next
			}
			if previous.next == nil || previous.next.exponent != exponent {
				node.next = head
				previous.next = node
			} else {
				previous.next.coefficient = previous.next.coefficient + coefficient
			}
		}
	}
}

func (p *PolynomialList) Plus(other PolynomialList) PolynomialList {
	firstHead := p.first
	secondHead := other.first

	resultantList := PolynomialList{}

	addInResult := func(coefficient int, exponent int) {
		resultantList.Add(coefficient, exponent)
	}
	copyFirst := func() {
		addInResult(firstHead.coefficient, firstHead.exponent)
		firstHead = firstHead.next
	}
	copySecond := func() {
		addInResult(secondHead.coefficient, secondHead.exponent)
		secondHead = secondHead.next
	}
	copyFirstPolynomialExponentsIfApplicable := func() {
		for firstHead.exponent > secondHead.exponent {
			copyFirst()
		}
	}
	copySecondPolynomialExponentsIfApplicable := func() {
		for secondHead.exponent > firstHead.exponent {
			copySecond()
		}
	}
	for firstHead != nil && secondHead != nil {
		if firstHead.exponent == secondHead.exponent {
			resultantList.Add(firstHead.coefficient+secondHead.coefficient, firstHead.exponent)
			firstHead = firstHead.next
			secondHead = secondHead.next
		} else {
			copyFirstPolynomialExponentsIfApplicable()
			copySecondPolynomialExponentsIfApplicable()
		}
	}
	for firstHead != nil {
		copyFirst()
	}
	for secondHead != nil {
		copySecond()
	}
	return resultantList
}

func (p *PolynomialList) Expression() string {
	var output string
	first := p.first
	for first != nil {
		output = output + strconv.Itoa(first.coefficient) + "x" + strconv.Itoa(first.exponent)
		if first.next != nil {
			output = output + " + "
		} else {
			output = output + " "
		}
		first = first.next
	}
	return output
}
