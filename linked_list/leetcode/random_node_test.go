package leetcode_test

import (
	"github.com/SarthakMakhija/data-structures-and-algorithms/linked_list/leetcode"
	"reflect"
	"testing"
)

func TestGetRandom1(t *testing.T) {

	lastNode := leetcode.ListNode{Val: 3, Next: nil}
	secondNode := leetcode.ListNode{Val: 2, Next: &lastNode}
	firstNode := leetcode.ListNode{Val: 1, Next: &secondNode}

	solution := leetcode.Constructor(&firstNode)
	output := []int{solution.GetRandom(), solution.GetRandom(), solution.GetRandom()}
	expected := []int{2, 1, 3}

	if !reflect.DeepEqual(output, expected) {
		t.Fatalf("Expected %v, received %v", expected, output)
	}
}

func TestGetRandom2(t *testing.T) {

	lastNode := leetcode.ListNode{Val: 4, Next: nil}
	thirdNode := leetcode.ListNode{Val: 3, Next: &lastNode}
	secondNode := leetcode.ListNode{Val: 2, Next: &thirdNode}
	firstNode := leetcode.ListNode{Val: 1, Next: &secondNode}

	solution := leetcode.Constructor(&firstNode)
	output := []int{solution.GetRandom(), solution.GetRandom(), solution.GetRandom(), solution.GetRandom()}
	expected := []int{2, 1, 4, 3}

	if !reflect.DeepEqual(output, expected) {
		t.Fatalf("Expected %v, received %v", expected, output)
	}
}

func TestGetRandom3(t *testing.T) {

	lastNode := leetcode.ListNode{Val: 3, Next: nil}
	secondNode := leetcode.ListNode{Val: 2, Next: &lastNode}
	firstNode := leetcode.ListNode{Val: 1, Next: &secondNode}

	solution := leetcode.Constructor(&firstNode)
	output := []int{solution.GetRandom(), solution.GetRandom(), solution.GetRandom(), solution.GetRandom(), solution.GetRandom()}
	expected := []int{2, 1, 3, 2, 1}

	if !reflect.DeepEqual(output, expected) {
		t.Fatalf("Expected %v, received %v", expected, output)
	}
}
