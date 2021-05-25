package avl

import (
	"math"
)

type AvlTree struct {
	Root *AvlNode
}

type AvlNode struct {
	Value                 int
	Left                  *AvlNode
	Right                 *AvlNode
	LeftSubtreeEdgeCount  int
	RightSubtreeEdgeCount int
	BalanceFactor         int
}

type AvlNodeView struct {
	Value                 int
	LeftSubtreeEdgeCount  int
	RightSubtreeEdgeCount int
	BalanceFactor         int
}

func (node *AvlNode) toAvlNodeView() *AvlNodeView {
	return &AvlNodeView{
		Value:                 node.Value,
		LeftSubtreeEdgeCount:  node.LeftSubtreeEdgeCount,
		RightSubtreeEdgeCount: node.RightSubtreeEdgeCount,
		BalanceFactor:         node.BalanceFactor,
	}
}

func (tree *AvlTree) UpdateWithBalanceFactor() {
	if tree.Root == nil {
		return
	}
	var updateWithBalanceFactorInner func(*AvlNode)

	updateWithBalanceFactorInner = func(t *AvlNode) {
		if t == nil {
		} else if t.Left == nil && t.Right == nil {
			t.LeftSubtreeEdgeCount = 0
			t.RightSubtreeEdgeCount = 0
			t.BalanceFactor = 0
		} else {
			updateWithBalanceFactorInner(t.Left)
			updateWithBalanceFactorInner(t.Right)
			tree.updateBalanceFactor(t)
		}
	}
	updateWithBalanceFactorInner(tree.Root)
}

func (tree *AvlTree) Traverse() []*AvlNodeView {
	if tree.Root == nil {
		return nil
	}
	var nodeViews []*AvlNodeView
	var traverseInner func(t *AvlNode)

	traverseInner = func(t *AvlNode) {
		if t == nil {
		} else if t.Left == nil && t.Right == nil {
			nodeViews = append(nodeViews, t.toAvlNodeView())
		} else {
			traverseInner(t.Left)
			nodeViews = append(nodeViews, t.toAvlNodeView())
			traverseInner(t.Right)
		}
	}
	traverseInner(tree.Root)
	return nodeViews
}

func (tree *AvlTree) updateBalanceFactor(t *AvlNode) {
	tree.updateEdgeCountsOf(t)
	t.BalanceFactor = t.LeftSubtreeEdgeCount - t.RightSubtreeEdgeCount
}

func (tree *AvlTree) updateEdgeCountsOf(t *AvlNode) {
	connectingEdgeCount := 1

	if t.Right != nil {
		t.RightSubtreeEdgeCount = int(math.Max(
			float64(t.Right.LeftSubtreeEdgeCount), float64(t.Right.RightSubtreeEdgeCount))) + connectingEdgeCount
	}
	if t.Left != nil {
		t.LeftSubtreeEdgeCount = int(math.Max(
			float64(t.Left.LeftSubtreeEdgeCount), float64(t.Left.RightSubtreeEdgeCount))) + connectingEdgeCount
	}
}
