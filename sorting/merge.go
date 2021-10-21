package sorting

type Merge struct {
	Elements []int
}

func (merge *Merge) Sort() {
	merge.sort(0, len(merge.Elements)-1)
}

func (merge *Merge) sort(low, high int) {
	if low < high {
		mid := (low + high) / 2
		merge.sort(low, mid)
		merge.sort(mid+1, high)
		merge.MergeSorted(low, mid, high)
	}
}

func (merge *Merge) MergeSorted(low, mid, high int) {
	destination := make([]int, mid-low+1+high-mid)
	aIndex, bIndex, destinationIndex := low, mid+1, 0

	for aIndex <= mid && bIndex <= high {
		if merge.Elements[aIndex] < merge.Elements[bIndex] {
			destination[destinationIndex] = merge.Elements[aIndex]
			destinationIndex, aIndex = destinationIndex+1, aIndex+1
		} else {
			destination[destinationIndex] = merge.Elements[bIndex]
			destinationIndex, bIndex = destinationIndex+1, bIndex+1
		}
	}
	if aIndex <= mid {
		copy(destination[destinationIndex:], merge.Elements[aIndex:])
	} else {
		copy(destination[destinationIndex:], merge.Elements[bIndex:])
	}
	copy(merge.Elements[low:], destination)
}
