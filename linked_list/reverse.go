package linkedlist

func (l LinkedList) ReverseAlternative() LinkedList {
	if l.first == nil {
		return LinkedList{}
	}

	var newFirst *Node = nil

	var reverseInner func(*Node, *Node)
	reverseInner = func(head *Node, tail *Node) {
		if head == nil {
			return
		}
		reverseInner(head.next, head)
		if head.next == nil {
			newFirst = head
		}
		head.next = tail
	}
	reverseInner(l.first, nil)
	return LinkedList{
		first: newFirst,
	}
}
