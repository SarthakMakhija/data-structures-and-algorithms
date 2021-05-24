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

	var traverse_inner func(t *StringNode) string
	traverse_inner = func(t *StringNode) string {
		if t == nil {
			return ""
		} else if t.Left == nil && t.Right == nil {
			return t.Value
		} else {
			return traverse_inner(t.Left) + t.Value + traverse_inner(t.Right)
		}
	}
	return traverse_inner(tree.Root)
}

func (tree *IntBinaryTree) Traverse() string {
	if tree.Root == nil {
		return ""
	}

	var traverse_inner func(t *IntNode) string
	traverse_inner = func(t *IntNode) string {
		if t == nil {
			return ""
		} else if t.Left == nil && t.Right == nil {
			return strconv.Itoa(t.Value)
		} else {
			return traverse_inner(t.Left) + strconv.Itoa(t.Value) + traverse_inner(t.Right)
		}
	}
	return traverse_inner(tree.Root)
}

func (tree *IntBinaryTree) Sum() int {
	if tree.Root == nil {
		return 0
	}

	var traverse_inner func(t *IntNode) int
	traverse_inner = func(t *IntNode) int {
		if t == nil {
			return 0
		} else if t.Left == nil && t.Right == nil {
			return t.Value
		} else {
			return traverse_inner(t.Left) + t.Value + traverse_inner(t.Right)
		}
	}
	return traverse_inner(tree.Root)
}

func (tree *IntBinaryTree) Count_Nodes() int {
	if tree.Root == nil {
		return 0
	}

	var count_inner func(t *IntNode) int
	count_inner = func(t *IntNode) int {
		if t == nil {
			return 0
		} else if t.Left == nil && t.Right == nil {
			return 1
		} else {
			return count_inner(t.Left) + 1 + count_inner(t.Right)
		}
	}
	return count_inner(tree.Root)
}

func (tree *IntBinaryTree) Count_Leaf_Nodes() int {
	if tree.Root == nil {
		return 0
	}

	var count_inner func(t *IntNode) int
	count_inner = func(t *IntNode) int {
		if t == nil {
			return 0
		} else if t.Left == nil && t.Right == nil {
			return 1
		} else {
			return count_inner(t.Left) + count_inner(t.Right)
		}
	}
	return count_inner(tree.Root)
}

func (tree *IntBinaryTree) Height_1() int {
	if tree.Root == nil {
		return 0
	}
	var max_height int = 0
	var height_inner func(*IntNode, int)

	height_inner = func(t *IntNode, height int) {
		if t == nil {
		} else if t.Left == nil && t.Right == nil {
			if height > max_height {
				max_height = height
			}
		} else {
			height_inner(t.Left, height+1)
			height_inner(t.Right, height+1)
		}
	}
	height_inner(tree.Root, 1)
	return max_height
}

func (tree *IntBinaryTree) Height_2() int {
	if tree.Root == nil {
		return 0
	}

	var height_inner func(*IntNode) int

	height_inner = func(t *IntNode) int {
		if t == nil {
			return 0
		} else {
			left_height := height_inner(t.Left)
			right_height := height_inner(t.Right)
			if left_height > right_height {
				return left_height + 1
			} else {
				return right_height + 1
			}
		}
	}
	return height_inner(tree.Root)
}

func (tree *IntBinaryTree) Max_1() int {
	if tree.Root == nil {
		return 0
	}

	var max_inner func(*IntNode)
	var max int

	max_inner = func(t *IntNode) {
		if t == nil {
			return
		} else {
			if t.Value > max {
				max = t.Value
			}
			if t.Left == nil && t.Right == nil {
				return
			} else {
				max_inner(t.Left)
				max_inner(t.Right)
			}
		}
	}
	max_inner(tree.Root)
	return max
}

func (tree *IntBinaryTree) Max_2() int {
	if tree.Root == nil {
		return 0
	}

	var max_inner func(t *IntNode) int

	max_inner = func(t *IntNode) int {
		if t == nil {
			return 0
		}
		if t.Left == nil && t.Right == nil {
			return t.Value
		} else {
			left_x := max_inner(t.Left)
			left_y := max_inner(t.Right)

			if left_x > left_y {
				return left_x
			} else {
				return left_y
			}
		}
	}
	return max_inner(tree.Root)
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

	var search_inner func(t *IntNode) *SearchResult

	search_inner = func(t *IntNode) *SearchResult {
		if t == nil {
			return nil
		}
		search_left := search_inner(t.Left)
		search_right := search_inner(t.Right)

		if search_left != nil && search_left.Contains {
			return search_left
		} else if search_right != nil && search_right.Contains {
			return search_right
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
	return search_inner(tree.Root)
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

func (tree *IntBinaryTree) Insert_Left_Of(value, element int) {

	if tree.Root == nil {
		return
	}
	search_result := tree.Search(value)
	if search_result == nil {
		return
	}

	search_result.Node.Left = &IntNode{
		Value: element,
		Left:  nil,
		Right: nil,
	}
}

func (tree *IntBinaryTree) Delete(value int) {

	if tree.Root == nil {
		return
	}
	search_result := tree.Search(value)
	if search_result == nil {
		return
	}

	if search_result.Parent == nil {
		search_result.Node = nil
	} else if search_result.MatchDirection == Left {
		search_result.Parent.Left = nil
	} else {
		search_result.Parent.Right = nil
	}
}

func (tree *StringBinaryTree) Evaluate_Postfix() (int, error) {
	if tree.Root == nil {
		return -1, errors.New("EMPTY EXPRESSION")
	}

	var evaluate_inner func(*StringNode) (int, error)
	evaluate_inner = func(t *StringNode) (int, error) {

		if t.Left == nil && t.Right == nil {
			return strconv.Atoi(t.Value)
		} else {
			left_operand, _ := evaluate_inner(t.Left)
			operator := t.Value
			right_operand, _ := evaluate_inner(t.Right)

			switch operator {
			case "+":
				return left_operand + right_operand, nil
			case "*":
				return left_operand * right_operand, nil
			case "-":
				return left_operand - right_operand, nil
			case "/":
				return left_operand / right_operand, nil
			default:
				return -1, errors.New("UNSUPPORTED OPERATOR")
			}
		}
	}
	return evaluate_inner(tree.Root)
}

//assume a linked list is converted to a binary tree and we need to find the kth element
func (tree *IntBinaryTree) K_element(position int) int {

	var node_positions []int
	var height int = int(math.Floor(math.Log2(float64(position)))) + 1

	var node_position int = position - int(math.Pow(float64(2), float64(height-1))) + 1
	node_positions = append(node_positions, node_position)

	for height != 1 {
		height = height - 1
		node_position = int(math.Ceil(float64(node_position) / float64(2)))
		node_positions = append(node_positions, node_position)
	}

	head := tree.Root
	for index := len(node_positions) - 2; index >= 0; index-- {
		position_to_visit := node_positions[index]

		if math.Mod(float64(position_to_visit), float64(2)) == 0 {
			head = head.Right
		} else {
			head = head.Left
		}
	}
	return head.Value
}
