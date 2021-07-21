package binary_tree

import (
	"errors"
	"github.com/SarthakMakhija/data-structures-and-algorithms/stack"
	binarytree "github.com/SarthakMakhija/data-structures-and-algorithms/tree/binary"
)

func BinaryTreeNodesInOrderOfIncreasingDepth(tree binarytree.IntBinaryTree) []int {
	if tree.Root == nil {
		return []int{}
	}
	queue := Queue{
		Size: 100,
	}
	var nodeValues []int
	_, _ = queue.enqueue(tree.Root)

	for !queue.isEmpty() {
		node, _ := queue.dequeue()
		nodeValues = append(nodeValues, node.Value)

		if node.Left != nil {
			_, _ = queue.enqueue(node.Left)
		}
		if node.Right != nil {
			_, _ = queue.enqueue(node.Right)
		}
	}
	return nodeValues
}

func BinaryTreeNodesBottomUpLeftRight(tree binarytree.IntBinaryTree) []int {
	if tree.Root == nil {
		return []int{}
	}
	queue := Queue{
		Size: 100,
	}
	intStack := stack.IntStack{}

	_, _ = queue.enqueue(tree.Root)
	for !queue.isEmpty() {
		node, _ := queue.dequeue()
		intStack.Push(node.Value)

		if node.Right != nil {
			_, _ = queue.enqueue(node.Right)
		}
		if node.Left != nil {
			_, _ = queue.enqueue(node.Left)
		}
	}

	var nodeValues []int
	for e := intStack.Pop(); e != -1; e = intStack.Pop() {
		nodeValues = append(nodeValues, e)
	}
	return nodeValues
}

func BinaryTreeNodesAverageAtEachLevel(tree binarytree.IntBinaryTree) []float64 {
	if tree.Root == nil {
		return []float64{}
	}
	queue := Queue{
		Size: 100,
	}
	var averageValues []float64
	_, _ = queue.enqueue(tree.Root)

	levelAverageWithNodes := func() (float64, []*binarytree.IntNode) {
		sum := 0.0
		count := 0
		var nodes []*binarytree.IntNode

		for !queue.isEmpty() {
			node, _ := queue.dequeue()
			count = count + 1
			sum = sum + float64(node.Value)
			nodes = append(nodes, node)
		}
		return sum / float64(count), nodes
	}
	for !queue.isEmpty() {
		average, levelNodes := levelAverageWithNodes()
		for _, node := range levelNodes {
			if node.Left != nil {
				_, _ = queue.enqueue(node.Left)
			}
			if node.Right != nil {
				_, _ = queue.enqueue(node.Right)
			}
		}
		averageValues = append(averageValues, average)
	}
	return averageValues
}

type Queue struct {
	elements []*binarytree.IntNode
	front    int
	rear     int
	Size     int
}

func (l *Queue) enqueue(element *binarytree.IntNode) (bool, error) {
	const fixedSize = 5
	if l.isEmpty() {
		var size = l.Size
		if l.Size == 0 {
			size = fixedSize
		}
		l.elements = make([]*binarytree.IntNode, size)
		l.front = 0
		l.rear = -1
		l.Size = size
	} else if l.isFull() {
		return false, errors.New("OVERFLOW")
	}
	l.rear = l.rear + 1
	l.elements[l.rear] = element
	return true, nil
}

func (l *Queue) dequeue() (*binarytree.IntNode, error) {
	if l.isEmpty() {
		return nil, errors.New("EMPTY")
	}
	element := l.elements[l.front]
	l.front = l.front + 1
	return element, nil
}

func (l *Queue) isEmpty() bool {
	return len(l.elements) == 0 || l.front > l.rear
}

func (l *Queue) isFull() bool {
	return l.rear == l.Size-1
}
