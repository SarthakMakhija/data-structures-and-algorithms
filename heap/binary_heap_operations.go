package heap

import (
	"errors"
)

type BinaryHeap struct {
	elements     []int
	Size         int
	currentIndex int
}

//MaxHeapWith
//Assume Size is always valid and is equal to +1 of the total elements
func (b *BinaryHeap) MaxHeapWith(element int) (bool, error) {
	if b.isEmpty() {
		b.initialize()
	} else if b.isFull() {
		return false, errors.New("NOT ENOUGH SPACE LEFT IN HEAP")
	}
	b.add(element)
	b.reAdjustElementsOnAdding(element)
	b.currentIndex = b.currentIndex + 1
	return true, nil
}

func (b *BinaryHeap) All() []int {
	destination := make([]int, len(b.elements))
	copy(destination, b.elements)
	return destination[1:]
}

func (b *BinaryHeap) reAdjustElementsOnAdding(element int) {
	currentIndex := b.currentIndex
	parentIndex := currentIndex / 2

	for currentIndex > 1 && b.elements[currentIndex] > b.elements[parentIndex] {
		b.elements[currentIndex] = b.elements[parentIndex]
		b.elements[parentIndex] = element

		currentIndex = parentIndex
		parentIndex = currentIndex / 2
	}
}

func (b *BinaryHeap) add(element int) {
	b.elements[b.currentIndex] = element
}

func (b *BinaryHeap) isEmpty() bool {
	return len(b.elements) == 0
}

func (b *BinaryHeap) isFull() bool {
	return b.currentIndex == len(b.elements)
}

func (b *BinaryHeap) initialize() {
	b.currentIndex = 1
	b.elements = make([]int, b.Size)
}
