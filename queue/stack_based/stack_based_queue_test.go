package stack_based_test

import (
	"github.com/SarthakMakhija/data-structures-and-algorithms/queue/stack_based"
	"github.com/SarthakMakhija/data-structures-and-algorithms/stack"
	"reflect"
	"testing"
)

func TestStackBasedQueueAdd_1(t *testing.T) {
	stackBasedQueue := stack_based.StackBasedQueue{WriteStack: &stack.IntStack{}, ReadStack: &stack.IntStack{}}
	stackBasedQueue.Add(10)
	stackBasedQueue.Add(20)
	stackBasedQueue.Add(30)
	element := stackBasedQueue.Get()

	if element != 10 {
		t.Fatalf("Expected front of the linearQueue to be 10, received %v", element)
	}
}

func TestStackBasedQueueAdd_2(t *testing.T) {
	stackBasedQueue := stack_based.StackBasedQueue{WriteStack: &stack.IntStack{}, ReadStack: &stack.IntStack{}}
	stackBasedQueue.Add(10)
	stackBasedQueue.Add(20)
	stackBasedQueue.Add(30)

	element1 := stackBasedQueue.Get()
	element2 := stackBasedQueue.Get()
	element3 := stackBasedQueue.Get()

	elements := []int{element1, element2, element3}
	expected := []int{10, 20, 30}

	if !reflect.DeepEqual(elements, expected) {
		t.Fatalf("Expected %v, received %v", expected, elements)
	}
}

func TestStackBasedQueueAdd_3(t *testing.T) {
	stackBasedQueue := stack_based.StackBasedQueue{WriteStack: &stack.IntStack{}, ReadStack: &stack.IntStack{}}
	stackBasedQueue.Add(10)
	stackBasedQueue.Add(20)
	stackBasedQueue.Add(30)

	element1 := stackBasedQueue.Get()
	stackBasedQueue.Add(40)
	element2 := stackBasedQueue.Get()

	if element1 != 10 {
		t.Fatalf("Expected element1 %v, received %v", 10, element1)
	}
	if element2 != 20 {
		t.Fatalf("Expected element2 %v, received %v", 10, element2)
	}
}

func TestStackBasedQueueAdd_4(t *testing.T) {
	stackBasedQueue := stack_based.StackBasedQueue{WriteStack: &stack.IntStack{}, ReadStack: &stack.IntStack{}}
	stackBasedQueue.Add(10)
	stackBasedQueue.Add(20)
	stackBasedQueue.Add(30)

	element1 := stackBasedQueue.Get()
	stackBasedQueue.Add(40)
	stackBasedQueue.Add(50)
	stackBasedQueue.Add(60)

	element2 := stackBasedQueue.Get()
	element3 := stackBasedQueue.Get()
	element4 := stackBasedQueue.Get()
	element5 := stackBasedQueue.Get()
	element6 := stackBasedQueue.Get()

	elements := []int{element2, element3, element4, element5, element6}
	expected := []int{20, 30, 40, 50, 60}

	if element1 != 10 {
		t.Fatalf("Expected element1 %v, received %v", 10, element1)
	}
	if !reflect.DeepEqual(elements, expected) {
		t.Fatalf("Expected %v, received %v", expected, elements)
	}
}
