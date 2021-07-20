package linkedlist

func (l LinkedList) PivotOn(k int) {
	if l.first == nil {
		return
	}

	var pivots []int
	var greater []int
	var smaller []int

	head := l.first

	for head != nil {
		if head.value == k {
			pivots = append(pivots, head.value)
		} else if head.value > k {
			greater = append(greater, head.value)
		} else {
			smaller = append(smaller, head.value)
		}
		head = head.next
	}

	updateListWith := func(head *Node, elements []int) *Node {
		for _, v := range elements {
			head.value = v
			head = head.next
		}
		return head
	}
	head = updateListWith(l.first, smaller)
	head = updateListWith(head, pivots)
	head = updateListWith(head, greater)
}

func (l LinkedList) PivotOnAlternative(k int) LinkedList {
	if l.first == nil {
		return LinkedList{}
	}

	var pivotHead = &Node{}
	var greaterHead = &Node{}
	var smallerHead = &Node{}

	var pivotIterator = pivotHead
	var greaterIterator = greaterHead
	var smallerIterator = smallerHead

	head := l.first

	for head != nil {
		if head.value == k {
			pivotIterator.next = head
			pivotIterator = head
		} else if head.value > k {
			greaterIterator.next = head
			greaterIterator = head
		} else {
			smallerIterator.next = head
			smallerIterator = head
		}
		head = head.next
	}

	greaterIterator.next = nil
	pivotIterator.next = greaterHead.next
	smallerIterator.next = pivotHead.next

	return LinkedList{
		first: smallerHead.next,
	}
}
