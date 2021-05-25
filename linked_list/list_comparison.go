package linkedlist

import (
	"errors"
)

//Compare
//Assume lists are of the same size || One of the naive implementations, other than sorting both and comparing
func Compare(list1 *LinkedList, list2 *LinkedList) bool {

	if list1.first == nil || list2.first == nil {
		return false
	}

	matched := true

	for list1Current := list1.first; list1Current != nil; {
		v1 := list1Current.value
		v1Matched := false

		var v2Matched, v3Matched, v4Matched bool

		v2, err := valueOf(list1Current, 1)
		if err == nil {
			v2Matched = false
		} else {
			v2Matched = true
		}

		v3, err := valueOf(list1Current, 2)
		if err == nil {
			v3Matched = false
		} else {
			v3Matched = true
		}

		v4, err := valueOf(list1Current, 3)
		if err == nil {
			v4Matched = false
		} else {
			v4Matched = true
		}

		for list2Current := list2.first; list2Current != nil; list2Current = list2Current.next {

			if list2Current.value == v1 {
				v1Matched = true
			} else if list2Current.value == v2 {
				v2Matched = true
			} else if list2Current.value == v3 {
				v3Matched = true
			} else if list2Current.value == v4 {
				v4Matched = true
			}
		}

		if v1Matched && v2Matched && v3Matched && v4Matched {
			list1Current = advance(list1Current, 4)
		} else {
			matched = false
			break
		}
	}

	return matched
}

func valueOf(head *Node, by int) (int, error) {
	node := advance(head, by)
	if node == nil {
		return -1, errors.New("NO VALUE POSSIBLE")
	}
	return node.value, nil
}

func advance(head *Node, by int) *Node {
	for count := 1; count <= by; count++ {
		head = head.next
		if head == nil {
			break
		}
	}
	return head
}
