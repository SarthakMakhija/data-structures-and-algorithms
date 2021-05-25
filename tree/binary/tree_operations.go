package binarytree

import (
	"errors"
	"math"
	"strconv"
)

type StringBinaryTree struct {
	Root *StringNode
}

type StringNode struct {
	Value string
	Left  *StringNode
	Right *StringNode
}

type IntBinaryTree struct {
	Root *IntNode
}

type IntNode struct {
	Value int
	Left  *IntNode
	Right *IntNode
}

func (tree *StringBinaryTree) Traverse() string {
	if tree.Root == nil {
		return ""
	}

	var traverseInner func(t *StringNode) string
	traverseInner = func(t *StringNode) string {
		if t == nil {
			return ""
		} else if t.Left == nil && t.Right == nil {
			return t.Value
		} else {
			return traverseInner(t.Left) + t.Value + traverseInner(t.Right)
		}
	}
	return traverseInner(tree.Root)
}

func (tree *IntBinaryTree) Traverse() string {
	if tree.Root == nil {
		return ""
	}

	var traverseInner func(t *IntNode) string
	traverseInner = func(t *IntNode) string {
		if t == nil {
			return ""
		} else if t.Left == nil && t.Right == nil {
			return strconv.Itoa(t.Value)
		} else {
			return traverseInner(t.Left) + strconv.Itoa(t.Value) + traverseInner(t.Right)
		}
	}
	return traverseInner(tree.Root)
}

func (tree *IntBinaryTree) Sum() int {
	if tree.Root == nil {
		return 0
	}

	var traverseInner func(t *IntNode) int
	traverseInner = func(t *IntNode) int {
		if t == nil {
			return 0
		} else if t.Left == nil && t.Right == nil {
			return t.Value
		} else {
			return traverseInner(t.Left) + t.Value + traverseInner(t.Right)
		}
	}
	return traverseInner(tree.Root)
}

func (tree *IntBinaryTree) CountNodes() int {
	if tree.Root == nil {
		return 0
	}

	var countInner func(t *IntNode) int
	countInner = func(t *IntNode) int {
		if t == nil {
			return 0
		} else if t.Left == nil && t.Right == nil {
			return 1
		} else {
			return countInner(t.Left) + 1 + countInner(t.Right)
		}
	}
	return countInner(tree.Root)
}

func (tree *IntBinaryTree) CountLeafNodes() int {
	if tree.Root == nil {
		return 0
	}

	var countInner func(t *IntNode) int
	countInner = func(t *IntNode) int {
		if t == nil {
			return 0
		} else if t.Left == nil && t.Right == nil {
			return 1
		} else {
			return countInner(t.Left) + countInner(t.Right)
		}
	}
	return countInner(tree.Root)
}

func (tree *IntBinaryTree) Height1() int {
	if tree.Root == nil {
		return 0
	}
	var maxHeight = 0
	var heightInner func(*IntNode, int)

	heightInner = func(t *IntNode, height int) {
		if t == nil {
		} else if t.Left == nil && t.Right == nil {
			if height > maxHeight {
				maxHeight = height
			}
		} else {
			heightInner(t.Left, height+1)
			heightInner(t.Right, height+1)
		}
	}
	heightInner(tree.Root, 1)
	return maxHeight
}

func (tree *IntBinaryTree) Height2() int {
	if tree.Root == nil {
		return 0
	}

	var heightInner func(*IntNode) int

	heightInner = func(t *IntNode) int {
		if t == nil {
			return 0
		} else {
			leftHeight := heightInner(t.Left)
			rightHeight := heightInner(t.Right)
			if leftHeight > rightHeight {
				return leftHeight + 1
			} else {
				return rightHeight + 1
			}
		}
	}
	return heightInner(tree.Root)
}

func (tree *IntBinaryTree) Max1() int {
	if tree.Root == nil {
		return 0
	}

	var maxInner func(*IntNode)
	var max int

	maxInner = func(t *IntNode) {
		if t == nil {
			return
		} else {
			if t.Value > max {
				max = t.Value
			}
			if t.Left == nil && t.Right == nil {
				return
			} else {
				maxInner(t.Left)
				maxInner(t.Right)
			}
		}
	}
	maxInner(tree.Root)
	return max
}

func (tree *IntBinaryTree) Max2() int {
	if tree.Root == nil {
		return 0
	}

	var maxInner func(t *IntNode) int

	maxInner = func(t *IntNode) int {
		if t == nil {
			return 0
		}
		if t.Left == nil && t.Right == nil {
			return t.Value
		} else {
			leftX := maxInner(t.Left)
			leftY := maxInner(t.Right)

			if leftX > leftY {
				return leftX
			} else {
				return leftY
			}
		}
	}
	return maxInner(tree.Root)
}

func (tree *IntBinaryTree) Search(v int) *SearchResult {

	if tree.Root == nil {
		return nil
	}
	if tree.Root.Value == v {
		return &SearchResult{
			Contains: true,
			Parent:   nil,
			Node:     tree.Root,
		}
	}

	var searchInner func(t *IntNode) *SearchResult

	searchInner = func(t *IntNode) *SearchResult {
		if t == nil {
			return nil
		}
		searchLeft := searchInner(t.Left)
		searchRight := searchInner(t.Right)

		if searchLeft != nil && searchLeft.Contains {
			return searchLeft
		} else if searchRight != nil && searchRight.Contains {
			return searchRight
		} else {
			if t.Left != nil && v == t.Left.Value {
				return &SearchResult{
					Contains:       true,
					Parent:         t,
					Node:           t.Left,
					MatchDirection: Left,
				}
			} else if t.Right != nil && v == t.Right.Value {
				return &SearchResult{
					Contains:       true,
					Parent:         t,
					Node:           t.Right,
					MatchDirection: Right,
				}
			} else {
				return nil
			}
		}
	}
	return searchInner(tree.Root)
}

type SearchResult struct {
	Contains       bool
	Parent         *IntNode
	Node           *IntNode
	MatchDirection MatchDirection
}

type MatchDirection int

const (
	Left MatchDirection = iota
	Right
)

func (tree *IntBinaryTree) InsertLeftOf(value, element int) {

	if tree.Root == nil {
		return
	}
	searchResult := tree.Search(value)
	if searchResult == nil {
		return
	}

	searchResult.Node.Left = &IntNode{
		Value: element,
		Left:  nil,
		Right: nil,
	}
}

func (tree *IntBinaryTree) Delete(value int) {

	if tree.Root == nil {
		return
	}
	searchResult := tree.Search(value)
	if searchResult == nil {
		return
	}

	if searchResult.Parent == nil {
		searchResult.Node = nil
	} else if searchResult.MatchDirection == Left {
		searchResult.Parent.Left = nil
	} else {
		searchResult.Parent.Right = nil
	}
}

func (tree *StringBinaryTree) EvaluatePostfix() (int, error) {
	if tree.Root == nil {
		return -1, errors.New("EMPTY EXPRESSION")
	}

	var evaluateInner func(*StringNode) (int, error)
	evaluateInner = func(t *StringNode) (int, error) {

		if t.Left == nil && t.Right == nil {
			return strconv.Atoi(t.Value)
		} else {
			leftOperand, _ := evaluateInner(t.Left)
			operator := t.Value
			rightOperand, _ := evaluateInner(t.Right)

			switch operator {
			case "+":
				return leftOperand + rightOperand, nil
			case "*":
				return leftOperand * rightOperand, nil
			case "-":
				return leftOperand - rightOperand, nil
			case "/":
				return leftOperand / rightOperand, nil
			default:
				return -1, errors.New("UNSUPPORTED OPERATOR")
			}
		}
	}
	return evaluateInner(tree.Root)
}

//KElement
//assume a linked list is converted to a binary tree and we need to find the kth element
func (tree *IntBinaryTree) KElement(position int) int {

	var nodePositions []int
	var height = int(math.Floor(math.Log2(float64(position)))) + 1

	var nodePosition = position - int(math.Pow(float64(2), float64(height-1))) + 1
	nodePositions = append(nodePositions, nodePosition)

	for height != 1 {
		height = height - 1
		nodePosition = int(math.Ceil(float64(nodePosition) / float64(2)))
		nodePositions = append(nodePositions, nodePosition)
	}

	head := tree.Root
	for index := len(nodePositions) - 2; index >= 0; index-- {
		positionToVisit := nodePositions[index]

		if math.Mod(float64(positionToVisit), float64(2)) == 0 {
			head = head.Right
		} else {
			head = head.Left
		}
	}
	return head.Value
}
