package green_book_test

import (
	"github.com/SarthakMakhija/data-structures-and-algorithms/stack/green_book"
	"testing"
)

func TestStackMin_1(t *testing.T) {
	stack := green_book.Stack{}
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	min := stack.Min()
	if min != 10 {
		t.Fatalf("Expected minimum to be 10 but found %v", min)
	}
}

func TestStackMin_2(t *testing.T) {
	stack := green_book.Stack{}
	stack.Push(30)
	stack.Push(10)
	stack.Push(20)

	min := stack.Min()
	if min != 10 {
		t.Fatalf("Expected minimum to be 10 but found %v", min)
	}
}

func TestStackPop_1(t *testing.T) {
	stack := green_book.Stack{}
	stack.Push(30)
	stack.Push(10)
	stack.Push(20)

	pop := stack.Pop()
	if pop != 20 {
		t.Fatalf("Expected top to be 20 but found %v", pop)
	}
}
