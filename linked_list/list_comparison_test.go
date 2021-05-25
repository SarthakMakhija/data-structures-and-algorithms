package linkedlist_test

import (
	"testing"

	linkedlist "github.com/SarthakMakhija/data-structures-and-algorithms/linked_list"
)

func TestCompareListsWithSameElementsInDifferentOrder_1(t *testing.T) {

	list1 := linkedlist.LinkedList{}
	list2 := linkedlist.LinkedList{}

	list1.Add(1)
	list1.Add(2)
	list1.Add(3)

	list2.Add(3)
	list2.Add(1)
	list2.Add(2)

	matched := linkedlist.Compare(&list1, &list2)
	if matched != true {
		t.Fatalf("Expected true but found %v", matched)
	}
}

func TestCompareListsWithSameElementsInDifferentOrder_2(t *testing.T) {

	list1 := linkedlist.LinkedList{}
	list2 := linkedlist.LinkedList{}

	list1.Add(1)
	list1.Add(2)
	list1.Add(10)
	list1.Add(9)
	list1.Add(8)

	list2.Add(10)
	list2.Add(9)
	list2.Add(8)
	list2.Add(1)
	list2.Add(2)

	matched := linkedlist.Compare(&list1, &list2)
	if matched != true {
		t.Fatalf("Expected true but found %v", matched)
	}
}

func TestCompareListsWithDifferentElements_1(t *testing.T) {

	list1 := linkedlist.LinkedList{}
	list2 := linkedlist.LinkedList{}

	list1.Add(1)
	list1.Add(2)
	list1.Add(10)
	list1.Add(9)
	list1.Add(8)

	list2.Add(10)
	list2.Add(90)
	list2.Add(8)
	list2.Add(1)
	list2.Add(2)

	matched := linkedlist.Compare(&list1, &list2)
	if matched != false {
		t.Fatalf("Expected false but found %v", matched)
	}
}

func TestCompareListsWithDifferentElements_2(t *testing.T) {

	list1 := linkedlist.LinkedList{}
	list2 := linkedlist.LinkedList{}

	list1.Add(1)
	list1.Add(2)
	list1.Add(2)

	list2.Add(1)
	list2.Add(2)
	list2.Add(3)

	matched := linkedlist.Compare(&list1, &list2)
	if matched != false {
		t.Fatalf("Expected false but found %v", matched)
	}
}

func TestCompareListsWithSameElementsInDifferentOrder_3(t *testing.T) {

	list1 := linkedlist.LinkedList{}
	list2 := linkedlist.LinkedList{}

	list1.Add(1)
	list1.Add(2)
	list1.Add(10)
	list1.Add(9)
	list1.Add(8)
	list1.Add(7)
	list1.Add(6)

	list2.Add(10)
	list2.Add(9)
	list2.Add(8)
	list2.Add(6)
	list2.Add(7)
	list2.Add(1)
	list2.Add(2)

	matched := linkedlist.Compare(&list1, &list2)
	if matched != true {
		t.Fatalf("Expected true but found %v", matched)
	}
}
