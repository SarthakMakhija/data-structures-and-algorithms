package array

import (
	"math"
)

type Array struct {
	Size     int
	elements []int
	index    int
}

type Pair struct {
	Element1 int
	Element2 int
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

func (a *Array) FindDuplicates() []int {
	type Occurrences struct {
		element int
		count   int
	}
	length := len(a.elements)
	var occurrences = make([]*Occurrences, length)

	addOccurrenceAt := func(index int, withElement int) {
		occurrences[index] = &Occurrences{
			element: withElement,
			count:   1,
		}
	}
	increaseOccurrenceAt := func(index int) {
		occurrences[index].count = occurrences[index].count + 1
	}
	findNextSlotFor := func(element int) int {
		emptySlot := -1
		for slot := 0; slot < len(occurrences); slot++ {
			if occurrences[slot] != nil && occurrences[slot].element == element {
				return slot
			}
			if emptySlot == -1 && occurrences[slot] == nil {
				emptySlot = slot
			}
		}
		return emptySlot
	}
	upsertOccurrenceAt := func(index int, element int) {
		if occurrences[index] == nil {
			addOccurrenceAt(index, element)
		} else {
			increaseOccurrenceAt(index)
		}
	}

	for _, element := range a.elements {
		index := int(math.Mod(float64(element), float64(length)))
		if occurrences[index] == nil {
			addOccurrenceAt(index, element)
		} else if occurrences[index].element == element {
			increaseOccurrenceAt(index)
		} else {
			slot := findNextSlotFor(element)
			upsertOccurrenceAt(slot, element)
		}
	}

	var duplicateElements []int
	for _, occurrence := range occurrences {
		if occurrence != nil && occurrence.count > 1 {
			duplicateElements = append(duplicateElements, occurrence.element)
		}
	}
	return duplicateElements
}

//PairWithSumEqualTo1
//Assume array elements are sorted
func (a *Array) PairWithSumEqualTo1(k int) []Pair {

	var pairs []Pair

	minIndex := 0
	maxIndex := len(a.elements) - 1

	for index := maxIndex; index > minIndex; {
		if a.elements[index]+a.elements[minIndex] != k {
			index = index - 1
		} else {
			pairs = append(pairs, Pair{
				Element1: a.elements[minIndex],
				Element2: a.elements[index],
			})
			index = index - 1
			minIndex = minIndex + 1
		}
	}
	return pairs
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
