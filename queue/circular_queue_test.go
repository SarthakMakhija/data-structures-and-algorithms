package queue_test

import (
	"github.com/SarthakMakhija/data-structures-and-algorithms/queue"
	"reflect"
	"testing"
)

func TestCircularBufferAdd_1(t *testing.T) {
	circularBuffer := queue.NewCircular(3)
	_, _ = circularBuffer.Enqueue(10)
	_, _ = circularBuffer.Enqueue(20)

	element1, _ := circularBuffer.Dequeue()
	element2, _ := circularBuffer.Dequeue()

	elements := []int{element1.(int), element2.(int)}
	expected := []int{10, 20}

	if !reflect.DeepEqual(elements, expected) {
		t.Fatalf("Expected %v, received %v", expected, elements)
	}
}

func TestCircularBufferAdd_2(t *testing.T) {
	circularBuffer := queue.NewCircular(1)

	_, err := circularBuffer.Dequeue()

	if err == nil {
		t.Fatalf("Expected error while deleting from an empty queue recevied %v", err)
	}
}

func TestCircularBufferAdd_3(t *testing.T) {
	circularBuffer := queue.NewCircular(3)
	_, _ = circularBuffer.Enqueue(10)
	_, _ = circularBuffer.Enqueue(20)

	_, err := circularBuffer.Enqueue(30)

	if err == nil {
		t.Fatalf("Expected error while adding to a full queue recevied %v", err)
	}
}

func TestCircularBufferAdd_4(t *testing.T) {
	circularBuffer := queue.NewCircular(3)
	_, _ = circularBuffer.Enqueue(10)
	_, _ = circularBuffer.Enqueue(20)
	_, _ = circularBuffer.Dequeue()
	_, _ = circularBuffer.Enqueue(30)

	element1, _ := circularBuffer.Dequeue()
	element2, _ := circularBuffer.Dequeue()

	elements := []int{element1.(int), element2.(int)}
	expected := []int{20, 30}

	if !reflect.DeepEqual(elements, expected) {
		t.Fatalf("Expected %v, received %v", expected, elements)
	}
}

func TestCircularBufferAdd_5(t *testing.T) {
	circularBuffer := queue.NewCircular(3)
	_, _ = circularBuffer.Enqueue(10)
	_, _ = circularBuffer.Enqueue(20)
	_, _ = circularBuffer.Dequeue()
	_, _ = circularBuffer.Dequeue()
	_, _ = circularBuffer.Enqueue(30)
	_, _ = circularBuffer.Enqueue(40)

	element1, _ := circularBuffer.Dequeue()
	element2, _ := circularBuffer.Dequeue()

	elements := []int{element1.(int), element2.(int)}
	expected := []int{30, 40}

	if !reflect.DeepEqual(elements, expected) {
		t.Fatalf("Expected %v, received %v", expected, elements)
	}
}
