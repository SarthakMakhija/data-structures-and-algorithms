package queue_test

import (
	"github.com/SarthakMakhija/data-structures-and-algorithms/queue"
	"reflect"
	"testing"
)

func TestLinearQueueAdd_1(t *testing.T) {
	linearQueue := queue.LinearQueue{}
	_, _ = linearQueue.Add(10)
	element, _ := linearQueue.Get()

	if element != 10 {
		t.Fatalf("Expected front of the linearQueue to be 10, received %v", element)
	}
}

func TestLinearQueueAdd_2(t *testing.T) {
	linearQueue := queue.LinearQueue{}
	_, _ = linearQueue.Add(10)
	_, _ = linearQueue.Add(20)
	element, _ := linearQueue.Get()

	if element != 10 {
		t.Fatalf("Expected front of the linearQueue to be 10, received %v", element)
	}
}

func TestLinearQueueAdd_3(t *testing.T) {
	linearQueue := queue.LinearQueue{}
	_, _ = linearQueue.Add(10)
	_, _ = linearQueue.Add(20)
	_, _ = linearQueue.Add(30)
	_, _ = linearQueue.Add(40)
	_, _ = linearQueue.Add(50)

	elements := linearQueue.All()
	expected := []int{10, 20, 30, 40, 50}

	if !reflect.DeepEqual(elements, expected) {
		t.Fatalf("Expected all elements to be %v, received %v", elements, expected)
	}
}

func TestLinearQueueAdd_4(t *testing.T) {
	linearQueue := queue.LinearQueue{}
	_, _ = linearQueue.Add(10)
	_, _ = linearQueue.Add(20)
	_, _ = linearQueue.Add(30)
	_, _ = linearQueue.Add(40)
	_, _ = linearQueue.Add(50)
	_, err := linearQueue.Add(60)

	if err == nil {
		t.Fatalf("Expected overflow while adding an element to a full queue, received none")
	}
}

func TestLinearQueueAdd_5(t *testing.T) {
	linearQueue := queue.LinearQueue{}
	_, err := linearQueue.Get()

	if err == nil {
		t.Fatalf("Expected error while getting an element from empty queue, received none")
	}
}
