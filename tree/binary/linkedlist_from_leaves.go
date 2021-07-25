package binarytree

import (
	linkedlist "github.com/SarthakMakhija/data-structures-and-algorithms/linked_list"
)

func (t IntBinaryTree) ToLinkedListFromLeaves() *linkedlist.LinkedList {
	if t.Root == nil {
		return nil
	}

	var toLinkedListFromLeavesInner func(node *IntNode, list *linkedlist.LinkedList)
	toLinkedListFromLeavesInner = func(node *IntNode, list *linkedlist.LinkedList) {
		if node == nil {
			return
		} else if node.Left == nil && node.Right == nil {
			list.Add(node.Value)
			return
		} else {
			toLinkedListFromLeavesInner(node.Left, list)
			toLinkedListFromLeavesInner(node.Right, list)
			return
		}
	}

	list := linkedlist.LinkedList{}
	toLinkedListFromLeavesInner(t.Root, &list)
	return &list
}
