package array

import (
	"math"
)

type Array struct {
	Size     int
	elements []int
	index    int
}

func (a *Array) Insert(element int) {
	if len(a.elements) == 0 {
		a.elements = make([]int, a.Size)
	}
	a.ensureCapacity()
	a.elements[a.index] = element
	a.index = a.index + 1
}

func (a *Array) DeleteAt(index int) {
	compact := func(fromIndex int) {
		for index := fromIndex; index < len(a.elements); index++ {
			a.elements[index-1] = a.elements[index]
		}
	}
	compact(index + 1)
	a.index = len(a.elements) - 1
}

//BinarySearch
//Assume Array is sorted
func (a *Array) BinarySearch(element int) bool {
	var binarySearchInner func(e []int) bool

	binarySearchInner = func(elements []int) bool {
		if len(elements) == 0 {
			return false
		}
		pivot := int(math.Floor(float64(len(elements))) / 2)
		pivotElement := elements[pivot]

		if element < pivotElement {
			return binarySearchInner(elements[0:pivot])
		} else if element > pivotElement {
			return binarySearchInner(elements[pivot+1:])
		} else {
			return true
		}
	}
	return binarySearchInner(a.elements)
}

func (a Array) Max() int {
	var maxInner func([]int, int) int

	maxInner = func(elements []int, max int) int {
		if len(elements) == 0 {
			return max
		} else if elements[0] > max {
			return maxInner(elements[1:], elements[0])
		} else {
			return maxInner(elements[1:], max)
		}
	}
	return maxInner(a.elements, a.elements[0])
}

//Rotate
//Assume by is less than the size of the array
func (a *Array) Rotate(by int) {
	if by == 0 {
		return
	}
	var backup = make([]int, by)
	var element int

	length := len(a.elements)
	for index := 0; index < length; index++ {
		if index < by {
			backup[index] = a.elements[index]
		}
		targetIndex := index + by
		if targetIndex >= length {
			targetIndex = targetIndex - length
			element = backup[targetIndex]
		} else {
			element = a.elements[targetIndex]
		}
		a.elements[index] = element
	}
}

func (a *Array) All() []int {
	elements := make([]int, a.index)
	for count := 0; count < a.index; count++ {
		elements[count] = a.elements[count]
	}
	return elements
}

func (a *Array) ensureCapacity() {
	if a.index >= len(a.elements) {
		newSize := a.Size * 2
		copied := make([]int, newSize)
		a.copyElementTo(copied)
		a.Size = newSize
		a.elements = copied
	}
}

func (a *Array) copyElementTo(destination []int) {
	for index, element := range a.elements {
		destination[index] = element
	}
}
