package array

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
