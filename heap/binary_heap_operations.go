package heap

import (
	"errors"
)

type BinaryHeap struct {
	elements     []int
	Size         int
	CurrentIndex int
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
	b.CurrentIndex = b.CurrentIndex + 1
	return true, nil
}

func (b *BinaryHeap) Delete() {
	if b.isEmpty() {
		return
	}
	swap := func(aIndex int, bIndex int) {
		hold := b.elements[aIndex]
		b.elements[aIndex] = b.elements[bIndex]
		b.elements[bIndex] = hold
	}

	rootIndex := 1
	b.elements[rootIndex] = b.elements[b.CurrentIndex-1]
	b.CurrentIndex = b.CurrentIndex - 1

	for rootIndex <= (b.CurrentIndex-1)/2 {
		leftChild := b.elements[2*rootIndex]
		rightChild := b.elements[2*rootIndex+1]

		if leftChild > rightChild {
			if b.elements[rootIndex] < leftChild {
				swap(rootIndex, 2*rootIndex)
				rootIndex = 2 * rootIndex
			} else {
				break
			}
		} else {
			if b.elements[rootIndex] < rightChild {
				swap(rootIndex, 2*rootIndex+1)
				rootIndex = 2*rootIndex + 1
			} else {
				break
			}
		}
	}
}

func (b *BinaryHeap) All() []int {
	destination := make([]int, len(b.elements))
	copy(destination, b.elements)
	return destination[1:b.CurrentIndex]
}

func (b *BinaryHeap) reAdjustElementsOnAdding(element int) {
	currentIndex := b.CurrentIndex
	parentIndex := currentIndex / 2

	for currentIndex > 1 && b.elements[currentIndex] > b.elements[parentIndex] {
		b.elements[currentIndex] = b.elements[parentIndex]
		b.elements[parentIndex] = element

		currentIndex = parentIndex
		parentIndex = currentIndex / 2
	}
}

func (b *BinaryHeap) add(element int) {
	b.elements[b.CurrentIndex] = element
}

func (b *BinaryHeap) isEmpty() bool {
	return len(b.elements) == 0
}

func (b *BinaryHeap) isFull() bool {
	return b.CurrentIndex == len(b.elements)
}

func (b *BinaryHeap) initialize() {
	b.CurrentIndex = 1
	b.elements = make([]int, b.Size)
}
