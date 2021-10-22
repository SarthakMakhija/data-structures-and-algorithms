package binary_tree

import (
	linearQueue "github.com/SarthakMakhija/data-structures-and-algorithms/queue"
	"github.com/SarthakMakhija/data-structures-and-algorithms/recursion"
	"github.com/SarthakMakhija/data-structures-and-algorithms/stack"
	binaryTree "github.com/SarthakMakhija/data-structures-and-algorithms/tree/binary"
)

func BinaryTreeNodesInOrderOfIncreasingDepth(tree binaryTree.IntBinaryTree) []int {
	if tree.Root == nil {
		return []int{}
	}
	queue := linearQueue.NewLinear(100)
	var nodeValues []int
	_, _ = queue.Enqueue(tree.Root)

	for !queue.IsEmpty() {
		node, _ := queue.Dequeue()
		intNode := node.(*binaryTree.IntNode)
		nodeValues = append(nodeValues, intNode.Value)

		if intNode.Left != nil {
			_, _ = queue.Enqueue(intNode.Left)
		}
		if intNode.Right != nil {
			_, _ = queue.Enqueue(intNode.Right)
		}
	}
	return nodeValues
}

func BinaryTreeNodesTopDownLeftRightAlternating(tree binaryTree.IntBinaryTree) []int {
	if tree.Root == nil {
		return []int{}
	}
	queue := linearQueue.NewLinear(100)
	var nodeValues []int
	direction := 'L'

	nodesInDirection := func(direction rune) ([]int, []*binaryTree.IntNode) {
		var values []int
		var nodes []*binaryTree.IntNode

		for !queue.IsEmpty() {
			node, _ := queue.Dequeue()
			intNode := node.(*binaryTree.IntNode)
			values = append(values, intNode.Value)
			nodes = append(nodes, intNode)
		}
		if direction == 'R' {
			return recursion.Reverse(values), nodes
		}
		return values, nodes
	}

	_, _ = queue.Enqueue(tree.Root)
	for !queue.IsEmpty() {
		values, levelNodes := nodesInDirection(direction)
		nodeValues = append(nodeValues, values...)

		for _, node := range levelNodes {
			if node.Left != nil {
				_, _ = queue.Enqueue(node.Left)
			}
			if node.Right != nil {
				_, _ = queue.Enqueue(node.Right)
			}
		}
		if direction == 'R' {
			direction = 'L'
		} else {
			direction = 'R'
		}
	}
	return nodeValues
}

func BinaryTreeNodesBottomUpLeftRight(tree binaryTree.IntBinaryTree) []int {
	if tree.Root == nil {
		return []int{}
	}
	queue := linearQueue.NewLinear(100)
	intStack := stack.IntStack{}

	_, _ = queue.Enqueue(tree.Root)
	for !queue.IsEmpty() {
		node, _ := queue.Dequeue()
		intNode := node.(*binaryTree.IntNode)
		intStack.Push(intNode.Value)

		if intNode.Right != nil {
			_, _ = queue.Enqueue(intNode.Right)
		}
		if intNode.Left != nil {
			_, _ = queue.Enqueue(intNode.Left)
		}
	}

	var nodeValues []int
	for e := intStack.Pop(); e != -1; e = intStack.Pop() {
		nodeValues = append(nodeValues, e)
	}
	return nodeValues
}

func BinaryTreeNodesAverageAtEachLevel(tree binaryTree.IntBinaryTree) []float64 {
	if tree.Root == nil {
		return []float64{}
	}
	queue := linearQueue.NewLinear(100)
	var averageValues []float64
	_, _ = queue.Enqueue(tree.Root)

	levelAverageWithNodes := func() (float64, []*binaryTree.IntNode) {
		sum := 0.0
		count := 0
		var nodes []*binaryTree.IntNode

		for !queue.IsEmpty() {
			node, _ := queue.Dequeue()
			intNode := node.(*binaryTree.IntNode)
			count = count + 1
			sum = sum + float64(intNode.Value)
			nodes = append(nodes, intNode)
		}
		return sum / float64(count), nodes
	}
	for !queue.IsEmpty() {
		average, levelNodes := levelAverageWithNodes()
		for _, node := range levelNodes {
			if node.Left != nil {
				_, _ = queue.Enqueue(node.Left)
			}
			if node.Right != nil {
				_, _ = queue.Enqueue(node.Right)
			}
		}
		averageValues = append(averageValues, average)
	}
	return averageValues
}
