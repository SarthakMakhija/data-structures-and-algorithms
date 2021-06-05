package queue_test

import (
	"github.com/SarthakMakhija/data-structures-and-algorithms/queue"
	"reflect"
	"testing"
)

func TestCircularBufferAdd_1(t *testing.T) {
	circularBuffer := queue.CircularBuffer{Size: 2}
	_, _ = circularBuffer.Add(10)
	_, _ = circularBuffer.Add(20)

	element1, _ := circularBuffer.Delete()
	element2, _ := circularBuffer.Delete()

	elements := []int{element1, element2}
	expected := []int{10, 20}

	if !reflect.DeepEqual(elements, expected) {
		t.Fatalf("Expected %v, received %v", expected, elements)
	}
}

func TestCircularBufferAdd_2(t *testing.T) {
	circularBuffer := queue.CircularBuffer{Size: 1}

	_, err := circularBuffer.Delete()

	if err == nil {
		t.Fatalf("Expected error while deleting from an empty queue recevied %v", err)
	}
}

func TestCircularBufferAdd_3(t *testing.T) {
	circularBuffer := queue.CircularBuffer{Size: 2}
	_, _ = circularBuffer.Add(10)
	_, _ = circularBuffer.Add(20)

	_, err := circularBuffer.Add(30)

	if err == nil {
		t.Fatalf("Expected error while adding to a full queue recevied %v", err)
	}
}

func TestCircularBufferAdd_4(t *testing.T) {
	circularBuffer := queue.CircularBuffer{Size: 2}
	_, _ = circularBuffer.Add(10)
	_, _ = circularBuffer.Add(20)
	_, _ = circularBuffer.Delete()
	_, _ = circularBuffer.Add(30)

	element1, _ := circularBuffer.Delete()
	element2, _ := circularBuffer.Delete()

	elements := []int{element1, element2}
	expected := []int{20, 30}

	if !reflect.DeepEqual(elements, expected) {
		t.Fatalf("Expected %v, received %v", expected, elements)
	}
}

func TestCircularBufferAdd_5(t *testing.T) {
	circularBuffer := queue.CircularBuffer{Size: 2}
	_, _ = circularBuffer.Add(10)
	_, _ = circularBuffer.Add(20)
	_, _ = circularBuffer.Delete()
	_, _ = circularBuffer.Delete()
	_, _ = circularBuffer.Add(30)
	_, _ = circularBuffer.Add(40)

	element1, _ := circularBuffer.Delete()
	element2, _ := circularBuffer.Delete()

	elements := []int{element1, element2}
	expected := []int{30, 40}

	if !reflect.DeepEqual(elements, expected) {
		t.Fatalf("Expected %v, received %v", expected, elements)
	}
}
