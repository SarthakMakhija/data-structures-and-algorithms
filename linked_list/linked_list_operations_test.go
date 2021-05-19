package linkedlist_test

import (
	"testing"

	linkedlist "example.com/data-structures-and-algorithms/linked_list"
)

func TestEmptyLinkedListElements(t *testing.T) {

	list := linkedlist.LinkedList{}

	output := list.All_As_String()
	if output != "" {
		t.Fatalf("Expected \"\" received %v", output)
	}
}

func TestNonEmptyLinkedListElements(t *testing.T) {

	list := linkedlist.LinkedList{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)

	output := list.All_As_String()
	if output != "1234" {
		t.Fatalf("Expected 1234 received %v", output)
	}
}

func TestMaxElementInNonEmptyList(t *testing.T) {

	list := linkedlist.LinkedList{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)

	max, _ := list.Max()
	if max != 4 {
		t.Fatalf("Expected 4 received %v", max)
	}
}

func TestElementIsContained(t *testing.T) {

	list := linkedlist.LinkedList{}

	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)

	contains := list.Contains(3)
	if contains != true {
		t.Fatalf("Expected true received %v", contains)
	}
}

func TestElementIsNotContained(t *testing.T) {

	list := linkedlist.LinkedList{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)

	contains := list.Contains(30)
	if contains != false {
		t.Fatalf("Expected false received %v", contains)
	}
}

func TestNonEmptyLinkedListElementsInSortedManner_1(t *testing.T) {

	list := linkedlist.LinkedList{}
	list.Add_Sorted(10)
	list.Add_Sorted(20)
	list.Add_Sorted(3)
	list.Add_Sorted(4)
	list.Add_Sorted(1)

	output := list.All_As_String()
	if output != "1341020" {
		t.Fatalf("Expected 1341020 received %v", output)
	}
}

func TestNonEmptyLinkedListElementsInSortedManner_2(t *testing.T) {

	list := linkedlist.LinkedList{}
	list.Add_Sorted(10)
	list.Add_Sorted(20)
	list.Add_Sorted(30)
	list.Add_Sorted(40)

	output := list.All_As_String()
	if output != "10203040" {
		t.Fatalf("Expected 10203040 received %v", output)
	}
}

func TestNonEmptyLinkedListElementsInSortedManner_3(t *testing.T) {

	list := linkedlist.LinkedList{}
	list.Add_Sorted(10)
	list.Add_Sorted(20)
	list.Add_Sorted(30)
	list.Add_Sorted(40)
	list.Add_Sorted(0)

	output := list.All_As_String()
	if output != "010203040" {
		t.Fatalf("Expected 010203040 received %v", output)
	}
}
