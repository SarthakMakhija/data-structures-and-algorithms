package linkedlist_test

import (
	linkedlist "github.com/SarthakMakhija/data-structures-and-algorithms/linked_list"
	"testing"
)

func TestLinkedListReverseAlternative1(t *testing.T) {

	l := linkedlist.LinkedList{}
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)
	l.Add(5)

	list := l.ReverseAlternative()
	output := list.AllAsString()
	expected := "54321"

	if expected != output {
		t.Fatalf("Expected %v, received %v", expected, output)
	}
}

func TestLinkedListReverseAlternative2(t *testing.T) {

	l := linkedlist.LinkedList{}
	l.Add(1)

	list := l.ReverseAlternative()
	output := list.AllAsString()
	expected := "1"

	if expected != output {
		t.Fatalf("Expected %v, received %v", expected, output)
	}
}

func TestLinkedListReverseAlternative3(t *testing.T) {

	l := linkedlist.LinkedList{}
	l.Add(1)
	l.Add(2)
	l.Add(3)

	list := l.ReverseAlternative()
	output := list.AllAsString()
	expected := "321"

	if expected != output {
		t.Fatalf("Expected %v, received %v", expected, output)
	}
}
