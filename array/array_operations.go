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

func (a *Array) Max() int {
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

//PairWithSumEqualTo2
//Assume array elements are not sorted
func (a *Array) PairWithSumEqualTo2(k int) []Pair {
	var pairs []Pair
	var potentialSumPairs = make(map[int]*Pair, len(a.elements))
	var EmptyPair = &Pair{}

	for _, element := range a.elements {
		difference := k - element
		if potentialSumPairs[difference] == nil {
			potentialSumPairs[element] = EmptyPair
		} else {
			potentialSumPairs[difference] = &Pair{
				Element1: element,
				Element2: difference,
			}
		}
	}
	for _, value := range potentialSumPairs {
		if value != EmptyPair {
			pairs = append(pairs, Pair{Element1: value.Element2, Element2: value.Element1})
		}
	}
	return pairs
}

func (a *Array) SecondHighestElement() int {
	if len(a.elements) == 0 {
		return 0
	} else if len(a.elements) == 1 {
		return a.elements[0]
	}

	const ElementIndexWithMaxValue = 0
	const ElementIndexWithSecondMaxValue = 1

	var highestElements = make([]int, 2)

	if a.elements[0] > a.elements[1] {
		highestElements[ElementIndexWithMaxValue] = a.elements[0]
		highestElements[ElementIndexWithSecondMaxValue] = a.elements[1]
	} else {
		highestElements[ElementIndexWithMaxValue] = a.elements[1]
		highestElements[ElementIndexWithSecondMaxValue] = a.elements[0]
	}
	for index := 2; index < len(a.elements); index++ {
		if a.elements[index] > highestElements[ElementIndexWithMaxValue] {
			highestElements[ElementIndexWithSecondMaxValue] = highestElements[ElementIndexWithMaxValue]
			highestElements[ElementIndexWithMaxValue] = a.elements[index]
		} else if a.elements[index] > highestElements[ElementIndexWithSecondMaxValue] {
			highestElements[ElementIndexWithSecondMaxValue] = a.elements[index]
		}
	}
	return highestElements[ElementIndexWithSecondMaxValue]
}

func (a *Array) NthHighestElement(n int) int {
	if n > len(a.elements) {
		return 0
	}
	var highestElements = make([]int, n)
	var top = 0

	addHighestElement := func(element int) {
		highestElements[top] = element
		top = top + 1
	}
	swapHighestOrderElements := func() {
		for index := top - 1; index > 0; index-- {
			if highestElements[index] > highestElements[index-1] {
				topElement := highestElements[index]
				highestElements[index] = highestElements[index-1]
				highestElements[index-1] = topElement
			}
		}
	}
	orderHighestElements := func(element int) {
		if top >= n {
			index := top - 1
			if element > highestElements[index] {
				highestElements[index] = element
			}
			swapHighestOrderElements()
		} else {
			highestElements[top] = element
			top = top + 1
			swapHighestOrderElements()
		}
	}
	for _, element := range a.elements {
		if top == 0 {
			addHighestElement(element)
		} else {
			orderHighestElements(element)
		}
	}
	return highestElements[top-1]
}

//Merge
//Assume both are sorted arrays
func (a *Array) Merge(other Array) []int {

	merged := make([]int, len(a.elements)+len(other.elements))

	mergedArrayIndex := 0
	firstArrayIndex, secondArrayIndex := 0, 0

	firstArray := a.elements
	secondArray := other.elements

	pick := func(element int) {
		merged[mergedArrayIndex] = element
		mergedArrayIndex = mergedArrayIndex + 1
	}
	copyRemaining := func(arr []int, fromIndex int) {
		for index := fromIndex; index < len(arr); index++ {
			merged[mergedArrayIndex] = arr[index]
			mergedArrayIndex = mergedArrayIndex + 1
		}
	}
	for firstArrayIndex < len(a.elements) && secondArrayIndex < len(other.elements) {
		if firstArray[firstArrayIndex] < secondArray[secondArrayIndex] {
			pick(firstArray[firstArrayIndex])
			firstArrayIndex = firstArrayIndex + 1
		} else {
			pick(secondArray[secondArrayIndex])
			secondArrayIndex = secondArrayIndex + 1
		}
	}
	if firstArrayIndex >= len(firstArray) {
		copyRemaining(secondArray, secondArrayIndex)
	} else {
		copyRemaining(firstArray, firstArrayIndex)
	}
	return merged
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
