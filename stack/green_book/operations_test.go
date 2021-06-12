package green_book_test

import (
	"github.com/SarthakMakhija/data-structures-and-algorithms/stack/green_book"
	"reflect"
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

func TestSortedStackPush_1(t *testing.T) {
	stack := green_book.SortedStack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	elements := []int{stack.Pop(), stack.Pop(), stack.Pop(), stack.Pop()}
	expected := []int{1, 2, 3, 4}

	if !reflect.DeepEqual(elements, expected) {
		t.Fatalf("Expected %v received %v", expected, elements)
	}
}

func TestSortedStackPush_2(t *testing.T) {
	stack := green_book.SortedStack{}
	stack.Push(40)
	stack.Push(5)
	stack.Push(1)
	stack.Push(50)
	stack.Push(7)
	stack.Push(6)

	elements := []int{stack.Pop(), stack.Pop(), stack.Pop(), stack.Pop(), stack.Pop(), stack.Pop()}
	expected := []int{1, 5, 6, 7, 40, 50}

	if !reflect.DeepEqual(elements, expected) {
		t.Fatalf("Expected %v received %v", expected, elements)
	}
}

func TestSortedStackPush_3(t *testing.T) {
	stack := green_book.SortedStack{}
	stack.Push(4)
	stack.Push(3)
	stack.Push(2)
	stack.Push(1)

	elements := []int{stack.Pop(), stack.Pop(), stack.Pop(), stack.Pop()}
	expected := []int{1, 2, 3, 4}

	if !reflect.DeepEqual(elements, expected) {
		t.Fatalf("Expected %v received %v", expected, elements)
	}
}
