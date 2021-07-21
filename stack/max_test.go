package stack_test

import (
	"github.com/SarthakMakhija/data-structures-and-algorithms/stack"
	"testing"
)

func TestStackMax1(t *testing.T) {
	stk := stack.MaxStack{}
	stk.Push(0)
	stk.Push(1)
	stk.Push(15)
	stk.Push(4)

	elementWithMax := stk.Max()
	max := elementWithMax.Max

	if max != 15 {
		t.Fatalf("Expected 15, received %v", max)
	}
}

func TestStackMax2(t *testing.T) {
	stk := stack.MaxStack{}
	stk.Push(0)
	stk.Push(1)
	stk.Push(4)
	stk.Push(15)

	stk.Max()
	stk.Pop()
	elementWithMax := stk.Max()
	max := elementWithMax.Max

	if max != 4 {
		t.Fatalf("Expected 4, received %v", max)
	}
}
