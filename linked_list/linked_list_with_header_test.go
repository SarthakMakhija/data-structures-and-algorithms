package linkedlist_test

import (
	linkedlist "github.com/SarthakMakhija/data-structures-and-algorithms/linked_list"
	"reflect"
	"testing"
)

func TestLinkedListWithHeaderSize_1(t *testing.T) {
	list := linkedlist.LinkedListWithHeader{}
	list.Add(10)
	list.Add(20)

	headerView := list.HeaderView()
	if headerView.Size != 2 {
		t.Fatalf("Expected size to be %v, received %v", 2, headerView.Size)
	}
}

func TestLinkedListWithHeaderSize_2(t *testing.T) {
	list := linkedlist.LinkedListWithHeader{}
	list.Add(10)
	list.Add(20)
	list.Add(9)
	list.Add(80)

	headerView := list.HeaderView()
	if headerView.Size != 4 {
		t.Fatalf("Expected size to be %v, received %v", 4, headerView.Size)
	}
}

func TestLinkedListWithHeaderEvenValueCount(t *testing.T) {
	list := linkedlist.LinkedListWithHeader{}
	list.Add(10)
	list.Add(20)

	headerView := list.HeaderView()
	if headerView.EvenValueCount != 2 {
		t.Fatalf("Expected even value count to be %v, received %v", 2, headerView.EvenValueCount)
	}
}

func TestLinkedListWithHeaderOddValueCount(t *testing.T) {
	list := linkedlist.LinkedListWithHeader{}
	list.Add(10)
	list.Add(13)

	headerView := list.HeaderView()
	if headerView.OddValueCount != 1 {
		t.Fatalf("Expected odd value count to be %v, received %v", 1, headerView.OddValueCount)
	}
}

func TestLinkedListWithHeaderAllElements_1(t *testing.T) {
	list := linkedlist.LinkedListWithHeader{}
	list.Add(10)
	list.Add(20)
	list.Add(9)
	list.Add(80)

	expected := []int{10, 20, 9, 80}
	allElements := list.All()

	if !reflect.DeepEqual(allElements, expected) {
		t.Fatalf("Expected all elements to be %v, received %v", expected, allElements)
	}
}
