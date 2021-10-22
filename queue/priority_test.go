package queue_test

import (
	"github.com/SarthakMakhija/data-structures-and-algorithms/queue"
	"reflect"
	"testing"
)

func TestPriorityQueueAdd_1(t *testing.T) {
	priorityQueue := queue.NewPriority(2)
	_, _ = priorityQueue.Add(10)
	_, _ = priorityQueue.Add(20)

	element1, _ := priorityQueue.Get()
	element2, _ := priorityQueue.Get()

	elements := []int{element1, element2}
	expected := []int{10, 20}

	if !reflect.DeepEqual(elements, expected) {
		t.Fatalf("Expected %v, received %v", expected, elements)
	}
}

func TestPriorityQueueAdd_2(t *testing.T) {
	priorityQueue := queue.NewPriority(2)
	_, _ = priorityQueue.Add(20)
	_, _ = priorityQueue.Add(10)

	element1, _ := priorityQueue.Get()
	element2, _ := priorityQueue.Get()

	elements := []int{element1, element2}
	expected := []int{10, 20}

	if !reflect.DeepEqual(elements, expected) {
		t.Fatalf("Expected %v, received %v", expected, elements)
	}
}

func TestPriorityQueueAdd_3(t *testing.T) {
	size := 7
	priorityQueue := queue.NewPriority(size)

	_, _ = priorityQueue.Add(9)
	_, _ = priorityQueue.Add(8)
	_, _ = priorityQueue.Add(10)
	_, _ = priorityQueue.Add(5)
	_, _ = priorityQueue.Add(6)
	_, _ = priorityQueue.Add(6)
	_, _ = priorityQueue.Add(2)

	var elements []int
	for count := 1; count <= size; count++ {
		element, _ := priorityQueue.Get()
		elements = append(elements, element)
	}
	expected := []int{2, 5, 6, 6, 8, 9, 10}

	if !reflect.DeepEqual(elements, expected) {
		t.Fatalf("Expected %v, received %v", expected, elements)
	}
}
