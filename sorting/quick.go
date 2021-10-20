package sorting

import "math"

type Quick struct {
	Elements []int
}

func (quick *Quick) Sort() {
	if len(quick.Elements) <= 1 {
		return
	}
	quick.Elements = append(quick.Elements, math.MaxInt32)
	quick.sort(quick.Elements, 0, len(quick.Elements)-1)
}

func (quick *Quick) sort(elements []int, low, high int) {
	if low < high {
		partitionPoint := quick.partition(elements, low, high)
		quick.sort(elements, low, partitionPoint)
		quick.sort(elements, partitionPoint+1, high)
	}
}

func (quick *Quick) partition(elements []int, low, high int) int {
	pivot := elements[low]
	front, end := low, high

	for front < end {
		for elements[front] <= pivot {
			front = front + 1
		}
		for elements[end] > pivot {
			end = end - 1
		}
		if front < end {
			elements[front], elements[end] = elements[end], elements[front]
		}
	}
	elements[low], elements[end] = elements[end], elements[low]
	return end
}

func (quick *Quick) AllElements() []int {
	return quick.Elements[0 : len(quick.Elements)-1]
}
