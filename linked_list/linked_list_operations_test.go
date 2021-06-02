package linkedlist_test

import (
	"testing"

	linkedlist "github.com/SarthakMakhija/data-structures-and-algorithms/linked_list"
)

func TestEmptyLinkedListElements(t *testing.T) {

	list := linkedlist.LinkedList{}

	output := list.AllAsString()
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

	output := list.AllAsString()
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
	list.AddSorted(10)
	list.AddSorted(20)
	list.AddSorted(3)
	list.AddSorted(4)
	list.AddSorted(1)

	output := list.AllAsString()
	if output != "1341020" {
		t.Fatalf("Expected 1341020 received %v", output)
	}
}

func TestNonEmptyLinkedListElementsInSortedManner_2(t *testing.T) {

	list := linkedlist.LinkedList{}
	list.AddSorted(10)
	list.AddSorted(20)
	list.AddSorted(30)
	list.AddSorted(40)

	output := list.AllAsString()
	if output != "10203040" {
		t.Fatalf("Expected 10203040 received %v", output)
	}
}

func TestNonEmptyLinkedListElementsInSortedManner_3(t *testing.T) {

	list := linkedlist.LinkedList{}
	list.AddSorted(10)
	list.AddSorted(20)
	list.AddSorted(30)
	list.AddSorted(40)
	list.AddSorted(0)

	output := list.AllAsString()
	if output != "010203040" {
		t.Fatalf("Expected 010203040 received %v", output)
	}
}

func TestLinkedListDoesNotContainCycle(t *testing.T) {

	list := linkedlist.LinkedList{}
	list.Add(10)
	list.Add(20)
	list.Add(30)
	list.Add(40)
	list.Add(0)

	containsCycle := list.ContainsCycle()
	if containsCycle != false {
		t.Fatalf("Expected false received %v", containsCycle)
	}
}

func TestLinkedListContainsCycle(t *testing.T) {

	list := linkedlist.LinkedList{}

	list.Add(10)
	node2 := list.Add(20)
	list.AddWithNext(30, node2)

	containsCycle := list.ContainsCycle()
	if containsCycle != true {
		t.Fatalf("Expected true received %v", containsCycle)
	}
}

func TestLinkedListReverse(t *testing.T) {

	list := linkedlist.LinkedList{}

	list.Add(10)
	list.Add(20)
	list.Add(30)

	reversed := list.Reverse()
	output := reversed.AllAsString()
	if output != "302010" {
		t.Fatalf("Expected 302010 received %v", output)
	}
}

func TestLinkedListRemoveDuplicatesIfThereAreNoDuplicates(t *testing.T) {

	list := linkedlist.LinkedList{}

	list.Add(10)
	list.Add(20)
	list.Add(30)

	list.RemoveDuplicatesFromSorted()
	output := list.AllAsString()
	if output != "102030" {
		t.Fatalf("Expected 102030 received %v", output)
	}
}

func TestLinkedListRemoveDuplicatesIfThereAreDuplicates_1(t *testing.T) {

	list := linkedlist.LinkedList{}

	list.Add(10)
	list.Add(10)
	list.Add(10)
	list.Add(10)
	list.Add(10)
	list.Add(20)
	list.Add(30)

	list.RemoveDuplicatesFromSorted()
	output := list.AllAsString()
	if output != "102030" {
		t.Fatalf("Expected 102030 received %v", output)
	}
}

func TestLinkedListRemoveDuplicatesIfThereAreDuplicates_2(t *testing.T) {

	list := linkedlist.LinkedList{}

	list.Add(10)
	list.Add(20)
	list.Add(30)
	list.Add(30)
	list.Add(30)
	list.Add(30)

	list.RemoveDuplicatesFromSorted()
	output := list.AllAsString()
	if output != "102030" {
		t.Fatalf("Expected 102030 received %v", output)
	}
}

func TestLinkedListMidElement_1(t *testing.T) {

	list := linkedlist.LinkedList{}

	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)
	list.Add(6)
	list.Add(7)

	midValue := list.MidValue()

	if midValue != 4 {
		t.Fatalf("Expected 4 received %v", midValue)
	}
}

func TestLinkedListMidElement_2(t *testing.T) {

	list := linkedlist.LinkedList{}

	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)

	midValue := list.MidValue()

	if midValue != 2 {
		t.Fatalf("Expected 2 received %v", midValue)
	}
}

func TestNoIntersectionBetweenLinkedList(t *testing.T) {

	list1 := linkedlist.LinkedList{}
	list2 := linkedlist.LinkedList{}

	list1.Add(1)
	list1.Add(2)
	list1.Add(3)
	list1.Add(4)

	list2.Add(1)
	list2.Add(2)

	intersectionValue := list1.Intersection(&list2)

	if intersectionValue != -1 {
		t.Fatalf("Expected -1 received %v", intersectionValue)
	}
}

func TestIntersectionBetweenLinkedList_1(t *testing.T) {

	list1 := linkedlist.LinkedList{}
	list2 := linkedlist.LinkedList{}

	list1.Add(1)
	list1.Add(2)
	list1.Add(3)
	node := list1.Add(4)
	list1.Add(5)
	list1.Add(6)
	list1.Add(7)

	list2.Add(50)
	list2.Add(40)
	list2.Add(10)
	list2.Add(2)
	list2.Add(8)
	list2.AddWithNext(4, node)

	intersectionValue := list1.Intersection(&list2)

	if intersectionValue != 4 {
		t.Fatalf("Expected 4 received %v", intersectionValue)
	}
}

func TestIntersectionBetweenLinkedList_2(t *testing.T) {

	list1 := linkedlist.LinkedList{}
	list2 := linkedlist.LinkedList{}

	list1.Add(1)
	list1.Add(2)
	list1.Add(3)
	node := list1.Add(4)

	list2.Add(50)
	list2.Add(40)
	list2.Add(10)
	list2.AddWithNext(4, node)

	intersectionValue := list1.Intersection(&list2)

	if intersectionValue != 4 {
		t.Fatalf("Expected 4 received %v", intersectionValue)
	}
}

func TestLinkedListRemoveDuplicates_1(t *testing.T) {

	list := linkedlist.LinkedList{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)

	list.RemoveDuplicates()
	output := list.AllAsString()
	if output != "1234" {
		t.Fatalf("Expected 1234 received %v", output)
	}
}

func TestLinkedListRemoveDuplicates_2(t *testing.T) {

	list := linkedlist.LinkedList{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(3)
	list.Add(4)
	list.Add(4)

	list.RemoveDuplicates()
	output := list.AllAsString()
	if output != "1234" {
		t.Fatalf("Expected 1234 received %v", output)
	}
}

func TestLinkedListRemoveDuplicates_3(t *testing.T) {

	list := linkedlist.LinkedList{}
	list.Add(1)
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(3)
	list.Add(4)
	list.Add(4)

	list.RemoveDuplicates()
	output := list.AllAsString()
	if output != "1234" {
		t.Fatalf("Expected 1234 received %v", output)
	}
}

func TestLinkedListRemoveDuplicates_4(t *testing.T) {

	list := linkedlist.LinkedList{}
	list.Add(1)
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(3)
	list.Add(4)
	list.Add(3)
	list.Add(4)
	list.Add(3)

	list.RemoveDuplicates()
	output := list.AllAsString()
	if output != "1234" {
		t.Fatalf("Expected 1234 received %v", output)
	}
}

func TestLinkedListIsPalindrome_1(t *testing.T) {

	list := linkedlist.LinkedList{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(2)
	list.Add(1)

	output := list.IsPalindrome()
	if output != true {
		t.Fatalf("Expected true received %v", output)
	}
}

func TestLinkedListIsPalindrome_2(t *testing.T) {

	list := linkedlist.LinkedList{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(3)
	list.Add(2)
	list.Add(1)

	output := list.IsPalindrome()
	if output != true {
		t.Fatalf("Expected true received %v", output)
	}
}

func TestLinkedListIsPalindrome_3(t *testing.T) {

	list := linkedlist.LinkedList{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)

	output := list.IsPalindrome()
	if output != false {
		t.Fatalf("Expected false received %v", output)
	}
}

func TestLinkedListIsPalindrome_4(t *testing.T) {

	list := linkedlist.LinkedList{}
	list.Add(1)
	list.Add(4)
	list.Add(3)
	list.Add(4)
	list.Add(5)

	output := list.IsPalindrome()
	if output != false {
		t.Fatalf("Expected false received %v", output)
	}
}

func TestLinkedListNthNodeFromEnd_1(t *testing.T) {

	list := linkedlist.LinkedList{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)
	list.Add(6)
	list.Add(7)
	list.Add(8)
	list.Add(9)
	list.Add(10)

	output := list.NthNodeFromEnd(8)
	if output != 3 {
		t.Fatalf("Expected 3 received %v", output)
	}
}

func TestLinkedListNthNodeFromEnd_2(t *testing.T) {

	list := linkedlist.LinkedList{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)
	list.Add(6)
	list.Add(7)
	list.Add(8)
	list.Add(9)
	list.Add(10)

	output := list.NthNodeFromEnd(7)
	if output != 4 {
		t.Fatalf("Expected 4 received %v", output)
	}
}

func TestLinkedListNthNodeFromEnd_3(t *testing.T) {

	list := linkedlist.LinkedList{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)
	list.Add(6)
	list.Add(7)
	list.Add(8)
	list.Add(9)
	list.Add(10)

	output := list.NthNodeFromEnd(5)
	if output != 6 {
		t.Fatalf("Expected 6 received %v", output)
	}
}

func TestLinkedListNthNodeFromEnd_4(t *testing.T) {

	list := linkedlist.LinkedList{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)
	list.Add(6)
	list.Add(7)
	list.Add(8)
	list.Add(9)
	list.Add(10)

	output := list.NthNodeFromEnd(3)
	if output != 8 {
		t.Fatalf("Expected 8 received %v", output)
	}
}

func TestLinkedListNthNodeFromEnd_5(t *testing.T) {

	list := linkedlist.LinkedList{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)
	list.Add(6)
	list.Add(7)
	list.Add(8)
	list.Add(9)
	list.Add(10)

	output := list.NthNodeFromEnd(2)
	if output != 9 {
		t.Fatalf("Expected 9 received %v", output)
	}
}

func TestLinkedListNthNodeFromEnd_6(t *testing.T) {

	list := linkedlist.LinkedList{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)
	list.Add(6)
	list.Add(7)
	list.Add(8)
	list.Add(9)
	list.Add(10)

	output := list.NthNodeFromEnd(1)
	if output != 10 {
		t.Fatalf("Expected 10 received %v", output)
	}
}

func TestLinkedListNthNodeFromEnd_7(t *testing.T) {

	list := linkedlist.LinkedList{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)
	list.Add(6)
	list.Add(7)
	list.Add(8)
	list.Add(9)
	list.Add(10)

	output := list.NthNodeFromEnd(10)
	if output != 1 {
		t.Fatalf("Expected 1 received %v", output)
	}
}

func TestLinkedListNthNodeFromEnd_8(t *testing.T) {

	list := linkedlist.LinkedList{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)
	list.Add(6)
	list.Add(7)
	list.Add(8)
	list.Add(9)
	list.Add(10)

	output := list.NthNodeFromEnd(4)
	if output != 7 {
		t.Fatalf("Expected 7 received %v", output)
	}
}

func TestLinkedListNthNodeFromEnd_9(t *testing.T) {

	list := linkedlist.LinkedList{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)
	list.Add(6)
	list.Add(7)
	list.Add(8)
	list.Add(9)

	output := list.NthNodeFromEnd(4)
	if output != 6 {
		t.Fatalf("Expected 6 received %v", output)
	}
}

func TestLinkedListNthNodeFromEnd_10(t *testing.T) {

	list := linkedlist.LinkedList{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)
	list.Add(6)
	list.Add(7)
	list.Add(8)
	list.Add(9)

	output := list.NthNodeFromEnd(9)
	if output != 1 {
		t.Fatalf("Expected 1 received %v", output)
	}
}

func TestLinkedListNthNodeFromEnd_11(t *testing.T) {

	list := linkedlist.LinkedList{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)
	list.Add(6)
	list.Add(7)
	list.Add(8)
	list.Add(9)

	output := list.NthNodeFromEnd(1)
	if output != 9 {
		t.Fatalf("Expected 9 received %v", output)
	}
}

func TestLinkedListNthNodeFromEnd_12(t *testing.T) {

	list := linkedlist.LinkedList{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)
	list.Add(6)
	list.Add(7)
	list.Add(8)
	list.Add(9)

	output := list.NthNodeFromEnd(4)
	if output != 6 {
		t.Fatalf("Expected 6 received %v", output)
	}
}
