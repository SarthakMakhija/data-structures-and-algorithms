package queue_test

import (
	"github.com/SarthakMakhija/data-structures-and-algorithms/queue"
	"reflect"
	"testing"
)

func TestLinearQueueAdd_1(t *testing.T) {
	linearQueue := queue.NewLinear(1)
	_, _ = linearQueue.Enqueue(10)
	element, _ := linearQueue.Dequeue()

	if element != 10 {
		t.Fatalf("Expected front of the linearQueue to be 10, received %v", element)
	}
}

func TestLinearQueueAdd_2(t *testing.T) {
	linearQueue := queue.NewLinear(2)
	_, _ = linearQueue.Enqueue(10)
	_, _ = linearQueue.Enqueue(20)
	element, _ := linearQueue.Dequeue()

	if element != 10 {
		t.Fatalf("Expected front of the linearQueue to be 10, received %v", element)
	}
}

func TestLinearQueueAdd_3(t *testing.T) {
	linearQueue := queue.NewLinear(5)
	_, _ = linearQueue.Enqueue(10)
	_, _ = linearQueue.Enqueue(20)
	_, _ = linearQueue.Enqueue(30)
	_, _ = linearQueue.Enqueue(40)
	_, _ = linearQueue.Enqueue(50)

	var elements []int
	for _, e := range linearQueue.AllElements() {
		elements = append(elements, e.(int))
	}
	expected := []int{10, 20, 30, 40, 50}

	if !reflect.DeepEqual(elements, expected) {
		t.Fatalf("Expected all elements to be %v, received %v", elements, expected)
	}
}

func TestLinearQueueAdd_4(t *testing.T) {
	linearQueue := queue.NewLinear(5)
	_, _ = linearQueue.Enqueue(10)
	_, _ = linearQueue.Enqueue(20)
	_, _ = linearQueue.Enqueue(30)
	_, _ = linearQueue.Enqueue(40)
	_, _ = linearQueue.Enqueue(50)
	_, err := linearQueue.Enqueue(60)

	if err == nil {
		t.Fatalf("Expected overflow while adding an element to a full queue, received none")
	}
}

func TestLinearQueueAdd_5(t *testing.T) {
	linearQueue := queue.NewLinear(1)
	_, err := linearQueue.Dequeue()

	if err == nil {
		t.Fatalf("Expected error while getting an element from empty queue, received none")
	}
}
