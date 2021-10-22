package skip

import (
	"fmt"
	"math/rand"
)

type List struct {
	tower []*node
}

type node struct {
	key, value  int
	right, down *node
}

type parentNodes struct {
	nodes []*node
}

func (parentNodes *parentNodes) add(n *node) {
	parentNodes.nodes = append(parentNodes.nodes, n)
}

func (parentNodes *parentNodes) pop() *node {
	n := parentNodes.nodes[len(parentNodes.nodes)-1]
	parentNodes.nodes = parentNodes.nodes[0 : len(parentNodes.nodes)-1]
	return n
}

func (parentNodes *parentNodes) isEmpty() bool {
	return len(parentNodes.nodes) == 0
}

func (n *node) updateRight(right *node) *node {
	n.right = right
	return n
}

func (n *node) updateDown(down *node) *node {
	n.down = down
	return n
}

func addNewNode(key, value int, left *node) *node {
	node := &node{key: key, value: value}
	node.updateRight(left.right)
	left.updateRight(node)
	return node
}

func newSentinelNode() *node {
	return &node{}
}

func NewList(towerSize int) *List {
	list := &List{tower: make([]*node, towerSize)}
	list.initializeWithSentinelNodesOf(towerSize)
	return list
}

func (list *List) GetByKey(key int) (int, bool) {
	targetNode := list.tower[len(list.tower)-1]
	for targetNode != nil {
		for targetNode.right != nil && targetNode.right.key <= key {
			targetNode = targetNode.right
		}
		if targetNode.key == key {
			return targetNode.value, true
		}
		targetNode = targetNode.down
	}
	return -1, false
}

//Put
//assume key does not exist
func (list *List) Put(key, value int) {
	parents := &parentNodes{}
	targetNode := list.tower[len(list.tower)-1]
	for ; targetNode != nil; targetNode = targetNode.down {
		for targetNode.right != nil && targetNode.right.key <= key {
			targetNode = targetNode.right
		}
		parents.add(targetNode)
	}

	left := parents.pop()
	node := addNewNode(key, value, left)
	for flipCoin() {
		if parents.isEmpty() {
			sentinelNode := list.increaseTowerSize()
			parents.add(sentinelNode)
		}
		left = parents.pop()
		newNode := addNewNode(key, value, left)
		newNode.updateDown(node)
		node = newNode
	}
}

func flipCoin() bool {
	return rand.Intn(2) == 1
}

func (list *List) increaseTowerSize() *node {
	sentinelNode := newSentinelNode()
	list.tower = append(list.tower, sentinelNode)
	topIndex := len(list.tower) - 1
	list.tower[topIndex].down = list.tower[topIndex-1].down
	return sentinelNode
}

func (list *List) initializeWithSentinelNodesOf(towerSize int) {
	for index := 0; index < towerSize; index++ {
		sentinelNode := newSentinelNode()
		list.tower[index] = sentinelNode
		if index > 0 {
			sentinelNode.down = list.tower[index-1]
		}
	}
}

func (list *List) Debug() {
	var debug func(*node, int)
	debug = func(node *node, level int) {
		fmt.Print("Level ", level)
		for node != nil {
			fmt.Print(" [", node.key, " ", node.value, "] 	-> ")
			node = node.right
		}
	}
	for index := len(list.tower) - 1; index >= 0; index-- {
		debug(list.tower[index], index)
		fmt.Println()
	}
}
