package leetcode

func RotateRight(l *ListNode, k int) *ListNode {

	var newFirst *ListNode = nil
	var parentNewFirst *ListNode = nil
	var newHead *ListNode = nil

	head, tail := l, l
	headSteps, tailSteps := 1, 1

	advanceHead := func() {
		for head.Next != nil {
			head = head.Next
			headSteps = headSteps + 1
		}
	}
	advanceTailMaintainingGapOfKSteps := func() {
		attemptTailAdvancements := func(k int) {
			for headSteps-tailSteps > k {
				tail = tail.Next
				tailSteps = tailSteps + 1
			}
		}
		attemptTailAdvancements(k)
		if tailSteps == 1 && k-headSteps >= 0 {
			actualRotation := k - headSteps
			attemptTailAdvancements(actualRotation)
		}
	}
	newNodes := func() {
		var newListFirst *ListNode = nil
		for newFirst != nil {
			node := &ListNode{
				Val: newFirst.Val,
			}
			if newListFirst == nil {
				newListFirst = node
				newHead = newListFirst
			} else {
				newListFirst.Next = node
				newListFirst = newListFirst.Next
			}
			newFirst = newFirst.Next
		}
		head = l
		for head != parentNewFirst {
			node := &ListNode{
				Val: head.Val,
			}
			newListFirst.Next = node
			newListFirst = newListFirst.Next
			head = head.Next
		}
	}

	advanceHead()
	advanceTailMaintainingGapOfKSteps()

	if tail == head {
		return l
	} else {
		newFirst = tail.Next
		parentNewFirst = newFirst
		newNodes()
		return newHead
	}
}

func (l *ListNode) All() []int {

	var elements []int
	head := l

	for head != nil {
		elements = append(elements, head.Val)
		head = head.Next
	}
	return elements
}
