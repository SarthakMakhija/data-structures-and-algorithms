package binary_indexed

type BinaryIndexedTree struct {
	Tree []int
}

func New(elements []int) BinaryIndexedTree {
	var tree = make([]int, len(elements)+1)
	var updateNextIndexOfWith func(int, int)

	nextIndexOf := func(index int) int {
		return index + index&(-index)
	}
	updateNextIndexOfWith = func(index int, element int) {
		treeIndex := nextIndexOf(index)
		if index >= len(tree) || treeIndex >= len(tree) {
			return
		}
		tree[treeIndex] = tree[treeIndex] + element
		updateNextIndexOfWith(treeIndex, element)
	}
	for index := 0; index < len(elements); index++ {
		element := elements[index]
		indexInTree := index + 1
		tree[indexInTree] = tree[indexInTree] + element
		updateNextIndexOfWith(indexInTree, element)
	}
	return BinaryIndexedTree{
		Tree: tree,
	}
}

func (t BinaryIndexedTree) FindSumOfFirst(n int) int {
	if n >= len(t.Tree) {
		return -1
	}
	sum := 0
	index := n
	parentIndexOf := func(index int) int {
		x := index - index&(-index)
		return x
	}
	for index > 0 {
		sum = sum + t.Tree[index]
		index = parentIndexOf(index)
	}
	return sum
}
