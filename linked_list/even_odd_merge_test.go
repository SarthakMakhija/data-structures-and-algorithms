package linkedlist_test

import (
	linkedlist "github.com/SarthakMakhija/data-structures-and-algorithms/linked_list"
	"testing"
)

func TestEvenOddMerge1(t *testing.T) {
	l := linkedlist.LinkedList{}
	l.Add(3)
	l.Add(2)
	l.Add(12)
	l.Add(11)
	l.Add(5)
	list := l.EvenOddMerge()

	output := list.AllAsString()
	expected := "3125211"

	if expected != output {
		t.Fatalf("Expected %v, received %v", expected, output)
	}
}

func TestEvenOddMerge2(t *testing.T) {
	l := linkedlist.LinkedList{}
	l.Add(3)
	l.Add(2)
	l.Add(12)
	l.Add(11)
	l.Add(5)
	l.Add(6)
	list := l.EvenOddMerge()

	output := list.AllAsString()
	expected := "31252116"

	if expected != output {
		t.Fatalf("Expected %v, received %v", expected, output)
	}
}
