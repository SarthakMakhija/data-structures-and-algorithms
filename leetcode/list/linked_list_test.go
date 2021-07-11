package list_test

import (
	"github.com/SarthakMakhija/data-structures-and-algorithms/leetcode/list"
	"testing"
)

func TestLinkedListReverseII1(t *testing.T) {
	node := &list.Node{
		Value: 1,
		Next: &list.Node{
			Value: 2,
			Next: &list.Node{
				Value: 3,
				Next: &list.Node{
					Value: 4,
					Next: &list.Node{
						Value: 5,
					},
				},
			},
		},
	}
	linkedList := list.LinkedList{First: node}

	linkedList.ReverseII(2, 4)
	elementsAsString := linkedList.AllAsString()
	expected := "14325"

	if elementsAsString != expected {
		t.Fatalf("Expected %v, received %v", expected, elementsAsString)
	}
}

func TestLinkedListReverseII2(t *testing.T) {
	node := &list.Node{
		Value: 1,
		Next: &list.Node{
			Value: 2,
			Next: &list.Node{
				Value: 3,
				Next: &list.Node{
					Value: 4,
					Next: &list.Node{
						Value: 5,
					},
				},
			},
		},
	}
	linkedList := list.LinkedList{First: node}

	linkedList.ReverseII(1, 4)
	elementsAsString := linkedList.AllAsString()
	expected := "43215"

	if elementsAsString != expected {
		t.Fatalf("Expected %v, received %v", expected, elementsAsString)
	}
}

func TestLinkedListReverseII3(t *testing.T) {
	node := &list.Node{
		Value: 1,
		Next: &list.Node{
			Value: 2,
			Next: &list.Node{
				Value: 3,
				Next: &list.Node{
					Value: 4,
					Next: &list.Node{
						Value: 5,
					},
				},
			},
		},
	}
	linkedList := list.LinkedList{First: node}

	linkedList.ReverseII(3, 5)
	elementsAsString := linkedList.AllAsString()
	expected := "12543"

	if elementsAsString != expected {
		t.Fatalf("Expected %v, received %v", expected, elementsAsString)
	}
}
