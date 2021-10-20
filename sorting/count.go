package sorting

type Count struct {
	Elements []int
}

func (count *Count) Sort() {
	if len(count.Elements) <= 1 {
		return
	}
	occurrence := make([]int, maxOf(count.Elements)+1)
	for _, element := range count.Elements {
		occurrence[element] = occurrence[element] + 1
	}
	sourceIndex := 0
	for index, times := range occurrence {
		if times != 0 {
			copy(count.Elements[sourceIndex:], buildFixedIntSliceWith(index, times))
			sourceIndex = sourceIndex + times
		}
	}
}

func buildFixedIntSliceWith(e int, size int) []int {
	var elements []int
	for count := 1; count <= size; count++ {
		elements = append(elements, e)
	}
	return elements
}

func maxOf(elements []int) int {
	max := elements[0]
	for index := 1; index < len(elements); index++ {
		if elements[index] > max {
			max = elements[index]
		}
	}
	return max
}
