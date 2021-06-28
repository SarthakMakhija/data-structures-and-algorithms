package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

type randomSelectionPointerType int

const (
	Forward randomSelectionPointerType = iota
	Tail
)

type Solution struct {
	head                   *ListNode
	forward                *ListNode
	tail                   *ListNode
	randomSelectionPointer randomSelectionPointerType
}

func Constructor(head *ListNode) Solution {
	return Solution{
		head:                   head,
		forward:                OriginalStateOfForward(head),
		tail:                   OriginalStateOfTail(head),
		randomSelectionPointer: Forward,
	}
}

func OriginalStateOfForward(head *ListNode) *ListNode {
	return head.Next
}

func OriginalStateOfTail(head *ListNode) *ListNode {
	return head
}

func (s *Solution) GetRandom() int {
	isNil := func(node *ListNode) bool {
		return node == nil
	}
	advanceBy2 := func(node *ListNode) *ListNode {
		if isNil(node) || isNil(node.Next) {
			return nil
		}
		return node.Next.Next
	}
	valueOf := func(node *ListNode) int {
		return node.Val
	}
	if isNil(s.forward) && isNil(s.tail) {
		s.forward = OriginalStateOfForward(s.head)
		s.tail = OriginalStateOfTail(s.head)
	}
	if s.randomSelectionPointer == Forward && !isNil(s.forward) {
		v := valueOf(s.forward)
		s.forward = advanceBy2(s.forward)
		s.randomSelectionPointer = Tail
		return v
	} else {
		v := valueOf(s.tail)
		s.tail = advanceBy2(s.tail)
		s.randomSelectionPointer = Forward
		return v
	}
}
