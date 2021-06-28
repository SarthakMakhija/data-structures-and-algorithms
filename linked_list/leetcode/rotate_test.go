package leetcode_test

import (
	"github.com/SarthakMakhija/data-structures-and-algorithms/linked_list/leetcode"
	"reflect"
	"testing"
)

func TestRotateRightBy2(t *testing.T) {

	lastNode := leetcode.ListNode{Val: 5, Next: nil}
	fourthNode := leetcode.ListNode{Val: 4, Next: &lastNode}
	thirdNode := leetcode.ListNode{Val: 3, Next: &fourthNode}
	secondNode := leetcode.ListNode{Val: 2, Next: &thirdNode}
	firstNode := leetcode.ListNode{Val: 1, Next: &secondNode}

	newHead := leetcode.RotateRight(&firstNode, 2)
	elements := newHead.All()
	expected := []int{4, 5, 1, 2, 3}

	if !reflect.DeepEqual(elements, expected) {
		t.Fatalf("Expected %v, received %v", expected, elements)
	}
}

func TestRotateRightBy3(t *testing.T) {

	lastNode := leetcode.ListNode{Val: 5, Next: nil}
	fourthNode := leetcode.ListNode{Val: 4, Next: &lastNode}
	thirdNode := leetcode.ListNode{Val: 3, Next: &fourthNode}
	secondNode := leetcode.ListNode{Val: 2, Next: &thirdNode}
	firstNode := leetcode.ListNode{Val: 1, Next: &secondNode}

	newHead := leetcode.RotateRight(&firstNode, 3)
	elements := newHead.All()
	expected := []int{3, 4, 5, 1, 2}

	if !reflect.DeepEqual(elements, expected) {
		t.Fatalf("Expected %v, received %v", expected, elements)
	}
}

func TestRotateRightBy4(t *testing.T) {

	lastNode := leetcode.ListNode{Val: 5, Next: nil}
	fourthNode := leetcode.ListNode{Val: 4, Next: &lastNode}
	thirdNode := leetcode.ListNode{Val: 3, Next: &fourthNode}
	secondNode := leetcode.ListNode{Val: 2, Next: &thirdNode}
	firstNode := leetcode.ListNode{Val: 1, Next: &secondNode}

	newHead := leetcode.RotateRight(&firstNode, 4)
	elements := newHead.All()
	expected := []int{2, 3, 4, 5, 1}

	if !reflect.DeepEqual(elements, expected) {
		t.Fatalf("Expected %v, received %v", expected, elements)
	}
}

func TestRotateRightBy5(t *testing.T) {

	lastNode := leetcode.ListNode{Val: 5, Next: nil}
	fourthNode := leetcode.ListNode{Val: 4, Next: &lastNode}
	thirdNode := leetcode.ListNode{Val: 3, Next: &fourthNode}
	secondNode := leetcode.ListNode{Val: 2, Next: &thirdNode}
	firstNode := leetcode.ListNode{Val: 1, Next: &secondNode}

	newHead := leetcode.RotateRight(&firstNode, 5)
	elements := newHead.All()
	expected := []int{1, 2, 3, 4, 5}

	if !reflect.DeepEqual(elements, expected) {
		t.Fatalf("Expected %v, received %v", expected, elements)
	}
}

func TestRotateRightBy6(t *testing.T) {

	lastNode := leetcode.ListNode{Val: 5, Next: nil}
	fourthNode := leetcode.ListNode{Val: 4, Next: &lastNode}
	thirdNode := leetcode.ListNode{Val: 3, Next: &fourthNode}
	secondNode := leetcode.ListNode{Val: 2, Next: &thirdNode}
	firstNode := leetcode.ListNode{Val: 1, Next: &secondNode}

	newHead := leetcode.RotateRight(&firstNode, 6)
	elements := newHead.All()
	expected := []int{5, 1, 2, 3, 4}

	if !reflect.DeepEqual(elements, expected) {
		t.Fatalf("Expected %v, received %v", expected, elements)
	}
}
