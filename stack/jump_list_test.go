package stack_test

import (
	"github.com/SarthakMakhija/data-structures-and-algorithms/stack"
	"reflect"
	"testing"
)

func TestTraverseJumpFirstOrder1(t *testing.T) {
	nodeD := stack.JumpNode{
		Order: -1,
		Value: 'D',
	}
	nodeB := stack.JumpNode{
		Order: -1,
		Value: 'B',
		Jump:  &nodeD,
	}
	nodeC := stack.JumpNode{
		Order: -1,
		Value: 'C',
		Jump:  &nodeB,
		Next:  &nodeD,
	}
	nodeA := stack.JumpNode{
		Order: -1,
		Value: 'A',
		Jump:  &nodeC,
		Next:  &nodeB,
	}
	nodeB.Next = &nodeC
	nodeD.Jump = &nodeD

	elements := nodeA.TraverseJumpFirstOrder()
	expected := []rune{'A', 'C', 'B', 'D'}

	if !reflect.DeepEqual(expected, elements) {
		t.Fatalf("Expected %v, received %v", expected, elements)
	}
}

func TestTraverseJumpFirstOrder2(t *testing.T) {
	nodeD := stack.JumpNode{
		Order: -1,
		Value: 'D',
	}
	nodeB := stack.JumpNode{
		Order: -1,
		Value: 'B',
		Jump:  &nodeD,
	}
	nodeC := stack.JumpNode{
		Order: -1,
		Value: 'C',
		Jump:  &nodeD,
		Next:  &nodeD,
	}
	nodeA := stack.JumpNode{
		Order: -1,
		Value: 'A',
		Jump:  &nodeC,
		Next:  &nodeB,
	}
	nodeB.Next = &nodeC
	nodeD.Jump = &nodeD

	elements := nodeA.TraverseJumpFirstOrder()
	expected := []rune{'A', 'C', 'D', 'B'}

	if !reflect.DeepEqual(expected, elements) {
		t.Fatalf("Expected %v, received %v", expected, elements)
	}
}

func TestTraverseJumpFirstOrder3(t *testing.T) {
	nodeD := stack.JumpNode{
		Order: -1,
		Value: 'D',
	}
	nodeB := stack.JumpNode{
		Order: -1,
		Value: 'B',
	}
	nodeC := stack.JumpNode{
		Order: -1,
		Value: 'C',
		Next:  &nodeD,
	}
	nodeA := stack.JumpNode{
		Order: -1,
		Value: 'A',
		Jump:  &nodeD,
		Next:  &nodeB,
	}
	nodeB.Next = &nodeC
	nodeD.Jump = &nodeD

	elements := nodeA.TraverseJumpFirstOrder()
	expected := []rune{'A', 'D', 'B', 'C'}

	if !reflect.DeepEqual(expected, elements) {
		t.Fatalf("Expected %v, received %v", expected, elements)
	}
}
