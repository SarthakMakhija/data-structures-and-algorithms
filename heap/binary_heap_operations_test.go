package heap_test

import (
	"github.com/SarthakMakhija/data-structures-and-algorithms/heap"
	"reflect"
	"testing"
)

func TestBinaryHeapMaxHeapWith_1(t *testing.T) {
	binaryHeap := heap.BinaryHeap{Size: 8}
	_, _ = binaryHeap.MaxHeapWith(5)
	_, _ = binaryHeap.MaxHeapWith(10)
	_, _ = binaryHeap.MaxHeapWith(12)
	_, _ = binaryHeap.MaxHeapWith(6)
	_, _ = binaryHeap.MaxHeapWith(20)
	_, _ = binaryHeap.MaxHeapWith(15)
	_, _ = binaryHeap.MaxHeapWith(30)

	elements := binaryHeap.All()
	expected := []int{30, 12, 20, 5, 6, 10, 15}

	if !reflect.DeepEqual(elements, expected) {
		t.Fatalf("Expected %v, received %v", expected, elements)
	}
}

func TestBinaryHeapMaxHeapWith_2(t *testing.T) {
	binaryHeap := heap.BinaryHeap{Size: 4}
	_, _ = binaryHeap.MaxHeapWith(5)
	_, _ = binaryHeap.MaxHeapWith(10)
	_, _ = binaryHeap.MaxHeapWith(12)

	elements := binaryHeap.All()
	expected := []int{12, 5, 10}

	if !reflect.DeepEqual(elements, expected) {
		t.Fatalf("Expected %v, received %v", expected, elements)
	}
}

func TestBinaryHeapDelete_1(t *testing.T) {
	binaryHeap := heap.BinaryHeap{Size: 8}

	_, _ = binaryHeap.MaxHeapWith(5)
	_, _ = binaryHeap.MaxHeapWith(10)
	_, _ = binaryHeap.MaxHeapWith(12)
	_, _ = binaryHeap.MaxHeapWith(6)
	_, _ = binaryHeap.MaxHeapWith(20)
	_, _ = binaryHeap.MaxHeapWith(15)
	_, _ = binaryHeap.MaxHeapWith(30)

	binaryHeap.Delete()

	elements := binaryHeap.All()
	expected := []int{20, 12, 15, 5, 6, 10}

	if !reflect.DeepEqual(elements, expected) {
		t.Fatalf("Expected %v, received %v", expected, elements)
	}
}

func TestBinaryHeapDelete_2(t *testing.T) {
	binaryHeap := heap.BinaryHeap{Size: 8}

	_, _ = binaryHeap.MaxHeapWith(5)
	_, _ = binaryHeap.MaxHeapWith(10)
	_, _ = binaryHeap.MaxHeapWith(12)
	_, _ = binaryHeap.MaxHeapWith(6)
	_, _ = binaryHeap.MaxHeapWith(20)
	_, _ = binaryHeap.MaxHeapWith(15)

	binaryHeap.Delete()

	elements := binaryHeap.All()
	expected := []int{15, 12, 10, 5, 6}

	if !reflect.DeepEqual(elements, expected) {
		t.Fatalf("Expected %v, received %v", expected, elements)
	}
}
