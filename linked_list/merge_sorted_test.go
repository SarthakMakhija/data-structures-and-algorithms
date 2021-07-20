package linkedlist_test

import (
	linkedlist "github.com/SarthakMakhija/data-structures-and-algorithms/linked_list"
	"testing"
)

func TestMergeSorted1(t *testing.T) {

	first := linkedlist.LinkedList{}
	first.Add(5)
	first.Add(10)
	first.Add(14)
	first.Add(16)

	second := linkedlist.LinkedList{}
	second.Add(3)
	second.Add(7)
	second.Add(11)
	second.Add(13)
	second.Add(15)
	second.Add(18)

	merged := linkedlist.MergeSorted(first, second)
	output := merged.AllAsString()
	expected := "35710111314151618"

	if output != expected {
		t.Fatalf("Expected %v, received %v", expected, output)
	}
}

func TestMergeSorted2(t *testing.T) {

	first := linkedlist.LinkedList{}
	first.Add(2)
	first.Add(5)
	first.Add(7)

	second := linkedlist.LinkedList{}
	second.Add(3)
	second.Add(9)
	second.Add(10)
	second.Add(12)

	merged := linkedlist.MergeSorted(first, second)
	output := merged.AllAsString()
	expected := "235791012"

	if output != expected {
		t.Fatalf("Expected %v, received %v", expected, output)
	}
}

func TestMergeSorted3(t *testing.T) {

	first := linkedlist.LinkedList{}
	first.Add(4)
	first.Add(5)
	first.Add(6)
	first.Add(7)

	second := linkedlist.LinkedList{}
	second.Add(1)
	second.Add(2)
	second.Add(3)

	merged := linkedlist.MergeSorted(first, second)
	output := merged.AllAsString()
	expected := "1234567"

	if output != expected {
		t.Fatalf("Expected %v, received %v", expected, output)
	}
}
