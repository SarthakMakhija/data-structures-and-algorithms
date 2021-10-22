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

func newSentinelNode() *node {
	return &node{}
}

func NewList(towerSize int) *List {
	list := &List{tower: make([]*node, towerSize)}
	list.initializeSentinelNode(towerSize)
	return list
}

func (list *List) GetByKey(key int) (int, int, bool) {
	var getByKeyInner func(*node, int) (int, int, bool)
	getByKeyInner = func(targetNode *node, level int) (int, int, bool) {
		switch {
		case targetNode == nil:
			return -1, -1, false
		case targetNode.key == key:
			return targetNode.value, level, true
		case targetNode.right == nil:
			return getByKeyInner(targetNode.down, level-1)
		case targetNode.right.key <= key:
			return getByKeyInner(targetNode.right, level)
		default:
			return getByKeyInner(targetNode.down, level-1)
		}
	}
	return getByKeyInner(list.tower[len(list.tower)-1], len(list.tower)-1)
}

func (list List) Put(key, value int) {
	var putInner func(*node) (*node, bool)
	putInner = func(targetNode *node) (*node, bool) {
		if targetNode == nil {
			return nil, false
		}
		if targetNode.down == nil {
			lowestLevelNode := targetNode
			for lowestLevelNode.right != nil && lowestLevelNode.right.key <= key {
				lowestLevelNode = lowestLevelNode.right
			}
			node := &node{key: key, value: value}
			node.updateLeft(lowestLevelNode).updateRight(lowestLevelNode.right)
			lowestLevelNode.updateRight(node)
			if node.right != nil {
				node.right.updateLeft(node)
			}
			return node, flipCoin()
		}
		var newNode *node
		var insertAtHigherLevel bool

		if targetNode.right == nil {
			newNode, insertAtHigherLevel = putInner(targetNode.down)
		} else {
			if targetNode.right.key <= key {
				newNode, insertAtHigherLevel = putInner(targetNode.right)
			} else {
				newNode, insertAtHigherLevel = putInner(targetNode.down)
			}
		}
		if insertAtHigherLevel {
			copied := newNode.copy()
			copied.updateRight(targetNode.right).updateLeft(targetNode).updateDown(newNode)
			targetNode.updateRight(copied)
			if copied.right != nil {
				copied.right.updateLeft(copied)
			}
			return copied, flipCoin()
		} else {
			return nil, false
		}
	}
	putInner(list.tower[len(list.tower)-1])
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
