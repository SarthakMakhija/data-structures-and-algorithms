package linkedlist_test

import (
	linkedlist "github.com/SarthakMakhija/data-structures-and-algorithms/linked_list"
	"testing"
)

func TestPivotOnK1(t *testing.T) {
	l := linkedlist.LinkedList{}
	l.Add(3)
	l.Add(2)
	l.Add(12)
	l.Add(11)
	l.Add(5)
	l.Add(15)
	l.Add(7)
	l.Add(9)
	l.Add(8)
	l.Add(7)
	l.Add(4)

	l.PivotOn(7)
	output := l.AllAsString()
	expected := "32547712111598"

	if output != expected {
		t.Fatalf("Expected %v, recevied %v", expected, output)
	}
}

func TestPivotOnK2(t *testing.T) {
	l := linkedlist.LinkedList{}
	l.Add(3)
	l.Add(2)
	l.Add(12)
	l.Add(11)
	l.Add(5)
	l.Add(15)
	l.Add(7)
	l.Add(9)
	l.Add(8)
	l.Add(7)
	l.Add(4)

	l.PivotOn(11)
	output := l.AllAsString()
	expected := "32579874111215"

	if output != expected {
		t.Fatalf("Expected %v, recevied %v", expected, output)
	}
}
