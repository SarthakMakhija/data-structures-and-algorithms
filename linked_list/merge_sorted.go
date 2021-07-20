package linkedlist

func MergeSorted(first LinkedList, second LinkedList) LinkedList {
	if first.first == nil || second.first == nil {
		if first.first == nil {
			return LinkedList{first: second.first}
		} else {
			return LinkedList{first: first.first}
		}
	}

	start := &Node{}
	firstHead := first.first
	secondHead := second.first

	var ptrToGreater *Node
	if secondHead.value < firstHead.value {
		start.next = secondHead
		ptrToGreater = firstHead
	} else {
		start.next = firstHead
		ptrToGreater = secondHead
	}
	ptrToSmaller := start.next
	for ptrToSmaller.next != nil {
		if ptrToSmaller.next.value > ptrToGreater.value {
			originalNext := ptrToSmaller.next
			ptrToSmaller.next = ptrToGreater
			ptrToSmaller = ptrToSmaller.next
			ptrToGreater = originalNext
		} else {
			ptrToSmaller = ptrToSmaller.next
		}
	}
	ptrToSmaller.next = ptrToGreater
	return LinkedList{first: start.next}
}
