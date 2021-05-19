package linkedlist_test

import (
	"testing"

	linkedlist "example.com/data-structures-and-algorithms/linked_list"
)

func TestElementIsContainedInSearchOptimisedLinkedList(t *testing.T) {

	list := linkedlist.NewSearchOptimisedLinkedList()

	list.Add("123")
	list.Add("423")
	list.Add("456")
	list.Add("9826")
	list.Add("90176")

	contains := list.Contains("90176")
	if contains != true {
		t.Fatalf("Expected true received %v", contains)
	}
}

func TestElementIsNotContainedInSearchOptimisedLinkedList(t *testing.T) {

	list := linkedlist.NewSearchOptimisedLinkedList()

	list.Add("123")
	list.Add("423")
	list.Add("456")
	list.Add("9826")
	list.Add("90176")

	contains := list.Contains("90170")
	if contains != false {
		t.Fatalf("Expected false received %v", contains)
	}
}
