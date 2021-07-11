package list

import "fmt"

type LinkedList struct {
	First *Node
}

type Node struct {
	Value int
	Next  *Node
}

func (l *LinkedList) ReverseII(start int, end int) {
	advanceHeadToStart := func(head *Node, start int) (*Node, *Node) {
		var follower *Node = nil
		steps := 0
		for steps < start-1 {
			follower = head
			head = head.Next
			steps = steps + 1
		}
		return head, follower
	}

	head := l.First
	head, follower := advanceHeadToStart(head, start)

	expectedPointerMovements := end - start
	for pointerMovements := 1; pointerMovements <= expectedPointerMovements; pointerMovements++ {
		p := head.Next
		head.Next = head.Next.Next
		if follower == nil {
			p.Next = l.First
			l.First = p
		} else {
			p.Next = follower.Next
			follower.Next = p
		}
	}
}

func (l *LinkedList) AllAsString() string {
	if l.First == nil {
		return ""
	}

	var iterate func(*Node) string
	var output = ""

	iterate = func(p *Node) string {
		output = output + fmt.Sprint(p.Value)
		if p.Next == nil {
			return output
		} else {
			return iterate(p.Next)
		}
	}
	return iterate(l.First)
}
