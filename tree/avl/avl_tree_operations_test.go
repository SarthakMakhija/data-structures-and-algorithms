package avl_test

import (
	"github.com/SarthakMakhija/data-structures-and-algorithms/tree/avl"
	"testing"
)

func TestBalanceFactor_1(t *testing.T) {
	tree := notAnAvlTree()
	tree.UpdateWithBalanceFactor()
	nodeViews := tree.Traverse()

	nodeView := find(nodeViews, 40)
	if nodeView.BalanceFactor != 2 {
		t.Fatalf("Expected balance factor 2 but received %v", nodeView.BalanceFactor)
	}
}

func TestBalanceFactor_2(t *testing.T) {
	tree := notAnAvlTree()
	tree.UpdateWithBalanceFactor()
	nodeViews := tree.Traverse()

	nodeView := find(nodeViews, 60)
	if nodeView.BalanceFactor != 0 {
		t.Fatalf("Expected balance factor 0 but received %v", nodeView.BalanceFactor)
	}
}

func TestBalanceFactor_3(t *testing.T) {
	tree := notAnAvlTree()
	tree.UpdateWithBalanceFactor()
	nodeViews := tree.Traverse()

	nodeView := find(nodeViews, 20)
	if nodeView.BalanceFactor != -1 {
		t.Fatalf("Expected balance factor -1 but received %v", nodeView.BalanceFactor)
	}
}

func TestBalanceFactor_4(t *testing.T) {
	tree := notAnAvlTree()
	tree.UpdateWithBalanceFactor()
	nodeViews := tree.Traverse()

	nodeView := find(nodeViews, 25)
	if nodeView.BalanceFactor != -1 {
		t.Fatalf("Expected balance factor -1 but received %v", nodeView.BalanceFactor)
	}
}

func find(nodeViews []*avl.AvlNodeView, v int) *avl.AvlNodeView {
	for _, nodeView := range nodeViews {
		if nodeView.Value == v {
			return nodeView
		}
	}
	return nil
}

func notAnAvlTree() *avl.AvlTree {
	node60 := avl.AvlNode{
		Value: 60,
	}
	node50 := avl.AvlNode{
		Value: 50,
		Right: &node60,
	}
	node8 := avl.AvlNode{
		Value: 8,
	}
	node10 := avl.AvlNode{
		Value: 10,
		Left:  &node8,
	}
	node36 := avl.AvlNode{
		Value: 36,
	}
	node27 := avl.AvlNode{
		Value: 27,
	}
	node25 := avl.AvlNode{
		Value: 25,
		Right: &node27,
	}
	node30 := avl.AvlNode{
		Value: 30,
		Left:  &node25,
		Right: &node36,
	}
	node20 := avl.AvlNode{
		Value: 20,
		Left:  &node10,
		Right: &node30,
	}
	node40 := avl.AvlNode{
		Value: 40,
		Left:  &node20,
		Right: &node50,
	}
	return &avl.AvlTree{
		Root: &node40,
	}
}
