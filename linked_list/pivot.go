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
