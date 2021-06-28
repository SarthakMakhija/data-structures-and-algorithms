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
	advanceBy2 := func(node *ListNode) *ListNode {
		if node == nil || node.Next == nil {
			return nil
		}
		return node.Next.Next
	}
	valueOf := func(node *ListNode) int {
		if node == nil {
			return -1
		}
		return node.Val
	}
	if s.forward == nil && s.tail == nil {
		s.forward = OriginalStateOfForward(s.head)
		s.tail = OriginalStateOfTail(s.head)
	}
	if s.randomSelectionPointer == Forward && s.forward != nil {
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
