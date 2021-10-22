package skip

import (
	"fmt"
	"math/rand"
)

type List struct {
	tower []*node
}

type node struct {
	key, value        int
	left, right, down *node
}

func (n *node) copy() *node {
	return &node{key: n.key, value: n.value}
}

func (n *node) updateRight(right *node) *node {
	n.right = right
	return n
}

func (n *node) updateLeft(left *node) *node {
	n.left = left
	return n
}

func (n *node) updateDown(down *node) *node {
	n.down = down
	return n
}

func addNewNode(key, value int, left *node) *node {
	node := &node{key: key, value: value}
	node.updateRight(left.right).updateLeft(left)
	left.updateRight(node)
	if node.right != nil {
		node.right.updateLeft(node)
	}
	return node
}

func newSentinelNode() *node {
	return &node{}
}

func NewList(towerSize int) *List {
	list := &List{tower: make([]*node, towerSize)}
	list.initializeSentinelNode(towerSize)
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

func (list List) PutNew(key, value int) {
	var parents []*node
	targetNode := list.tower[len(list.tower)-1]
	for ; targetNode != nil; targetNode = targetNode.down {
		for targetNode.right != nil && targetNode.right.key <= key {
			targetNode = targetNode.right
		}
		parents = append(parents, targetNode)
	}
	if parents != nil {
		leftSibling := parents[len(parents)-1]
		node := addNewNode(key, value, leftSibling)
		for flipCoin() && len(parents) > 0 { //handle parents empty .. when tower needs to increase
			parents = parents[0 : len(parents)-1]
			if len(parents) > 0 {
				leftSibling = parents[len(parents)-1]
				newNode := addNewNode(key, value, leftSibling)
				newNode.updateDown(node)
				node = newNode
			}
		}
	}
}

func flipCoin() bool {
	return rand.Intn(2) == 1
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

func (list *List) initializeSentinelNode(towerSize int) {
	for index := 0; index < towerSize; index++ {
		sentinelNode := newSentinelNode()
		list.tower[index] = sentinelNode
		if index > 0 {
			sentinelNode.down = list.tower[index-1]
		}
	}
}
