package binarytree

import (
	"math"
)

func (t StringBinaryTree) IsHeightBalanced() bool {

	var isHeightBalancedInner func(node *StringNode) (bool, int)
	isHeightBalancedInner = func(node *StringNode) (bool, int) {
		if node == nil {
			return true, 0
		}
		isHeightBalancedLeft, depthLeft := isHeightBalancedInner(node.Left)
		isHeightBalancedRight, depthRight := isHeightBalancedInner(node.Right)

		if isHeightBalancedLeft && isHeightBalancedRight && depthLeft-depthRight <= 1 {
			if depthLeft >= depthRight {
				return true, depthLeft + 1
			} else {
				return true, depthRight + 1
			}
		} else {
			return false, -1
		}
	}
	isHeightBalanced, _ := isHeightBalancedInner(t.Root)
	return isHeightBalanced
}

func (t StringBinaryTree) NodeWithHeightUnbalancedByK(k int) *StringNode {

	var unbalancedNode *StringNode = nil
	var isHeightBalancedInner func(node *StringNode) (bool, int)
	isHeightBalancedInner = func(node *StringNode) (bool, int) {
		if node == nil {
			return true, 0
		}
		isHeightBalancedLeft, depthLeft := isHeightBalancedInner(node.Left)
		isHeightBalancedRight, depthRight := isHeightBalancedInner(node.Right)

		if isHeightBalancedLeft && isHeightBalancedRight && math.Abs(float64(depthLeft)-float64(depthRight)) < float64(k) {
			if depthLeft >= depthRight {
				return true, depthLeft + 1
			} else {
				return true, depthRight + 1
			}
		} else {
			if unbalancedNode == nil {
				unbalancedNode = node
			}
			return false, -1
		}
	}
	isHeightBalancedInner(t.Root)
	return unbalancedNode
}
