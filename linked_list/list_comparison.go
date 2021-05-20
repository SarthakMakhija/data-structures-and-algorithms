package linkedlist

import (
	"errors"
)

//Assume lists are of the same size || One of the naive implementations, other than sorting both and comparing
func Compare(list1 *LinkedList, list2 *LinkedList) bool {

	if list1.first == nil || list2.first == nil {
		return false
	}

	matched := true

	for list1_current := list1.first; list1_current != nil; {
		v1 := list1_current.value
		v1_matched := false

		var v2_matched, v3_matched, v4_matched bool

		v2, err := value_of(list1_current, 1)
		if err == nil {
			v2_matched = false
		} else {
			v2_matched = true
		}

		v3, err := value_of(list1_current, 2)
		if err == nil {
			v3_matched = false
		} else {
			v3_matched = true
		}

		v4, err := value_of(list1_current, 3)
		if err == nil {
			v4_matched = false
		} else {
			v4_matched = true
		}

		for list2_current := list2.first; list2_current != nil; list2_current = list2_current.next {

			if list2_current.value == v1 {
				v1_matched = true
			} else if list2_current.value == v2 {
				v2_matched = true
			} else if list2_current.value == v3 {
				v3_matched = true
			} else if list2_current.value == v4 {
				v4_matched = true
			}
		}

		if v1_matched && v2_matched && v3_matched && v4_matched {
			list1_current = advance(list1_current, 4)
		} else {
			matched = false
			break
		}
	}

	return matched
}

func value_of(head *Node, by int) (int, error) {
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
