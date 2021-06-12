package green_book_test

import (
	"github.com/SarthakMakhija/data-structures-and-algorithms/linked_list/green_book"
	"testing"
)

func TestLinkedListAddDigitsOf_1(t *testing.T) {

	list := green_book.LinkedList{}
	list.Add(6)
	list.Add(1)
	list.Add(7) //617

	otherList := green_book.LinkedList{}
	otherList.Add(2)
	otherList.Add(1)
	otherList.Add(5) //215

	resultantList := list.AddDigitsOf(otherList)
	output := resultantList.AllAsString()

	if output != "832" {
		t.Fatalf("Expected 832, found %v", output)
	}
}

func TestLinkedListAddDigitsOf_2(t *testing.T) {

	list := green_book.LinkedList{}
	list.Add(5)
	list.Add(7)
	list.Add(2)
	list.Add(1)
	list.Add(6) //57216

	otherList := green_book.LinkedList{}
	otherList.Add(5)
	otherList.Add(7)
	otherList.Add(8) //578

	resultantList := list.AddDigitsOf(otherList)
	output := resultantList.AllAsString()

	if output != "57794" {
		t.Fatalf("Expected 57794, found %v", output)
	}
}

func TestLinkedListAddDigitsOf_3(t *testing.T) {

	list := green_book.LinkedList{}
	list.Add(5)
	list.Add(7)
	list.Add(2)
	list.Add(1)
	list.Add(6) //57216

	otherList := green_book.LinkedList{}
	otherList.Add(5)
	otherList.Add(7)
	otherList.Add(2)
	otherList.Add(1)
	otherList.Add(6) //57216

	resultantList := list.AddDigitsOf(otherList)
	output := resultantList.AllAsString()

	if output != "114432" {
		t.Fatalf("Expected 114432, found %v", output)
	}
}
