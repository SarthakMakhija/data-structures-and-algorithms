package binary_indexed

type BinaryIndexedTree struct {
	values []int
}

func New(elements []int) *BinaryIndexedTree {
	tree := &BinaryIndexedTree{values: make([]int, len(elements)+1)}
	for index, element := range elements {
		tree.updateAt(uint(index+1), element)
	}
	return tree
}

func (tree *BinaryIndexedTree) SumUptoIndex(n int) int {
	var sumInner func(index uint, sum int) int
	sumInner = func(index uint, sum int) int {
		if index == 0 || index >= uint(len(tree.values)) {
			return sum
		}
		return sumInner(parentIndexOf(index), sum+tree.values[index])
	}
	return sumInner(uint(n+1), 0)
}

func (tree *BinaryIndexedTree) updateAt(index uint, element int) {
	if index >= uint(len(tree.values)) {
		return
	}
	tree.values[index] = tree.values[index] + element
	tree.updateAt(nextIndexOf(index), element)
}

func nextIndexOf(index uint) uint {
	return index + (((^index) + 1) & index)
}

func parentIndexOf(index uint) uint {
	return index - (((^index) + 1) & index)
}
