package tree

import (
	"strconv"
)

type IntBinaryTree struct {
	Root *IntNode
}

type IntNode struct {
	Value int
	Left  *IntNode
	Right *IntNode
}

func (t *IntBinaryTree) Serialize() string {
	if t.Root == nil {
		return ""
	}
	var serializeInner func(node *IntNode) string
	serializeInner = func(node *IntNode) string {
		if node == nil {
			return "X"
		} else {
			return strconv.Itoa(node.Value) + serializeInner(node.Left) + serializeInner(node.Right)
		}
	}
	return serializeInner(t.Root)
}

func Deserialize(str string) *IntBinaryTree {
	if str == "" {
		return nil
	}

	tree := IntBinaryTree{}

	index := 0
	v, _ := strconv.Atoi(string(str[index]))
	tree.Root = &IntNode{
		Value: v,
	}
	index = index + 1

	var deserializeInner func(*IntNode)
	deserializeInner = func(node *IntNode) {
		if node == nil {
			return
		} else if string(str[index]) == "X" {
			if node.Left == nil {
				index = index + 1
				deserializeInner(node.Left)
			}
			if node.Right == nil {
				index = index + 1
				return
			}
		}
		if string(str[index]) != "X" {
			v, _ := strconv.Atoi(string(str[index]))
			n := &IntNode{
				Value: v,
			}
			if node.Left == nil {
				node.Left = n
				index = index + 1
				deserializeInner(node.Left)
			} else if node.Right == nil {
				node.Right = n
				index = index + 1
				deserializeInner(node.Right)
			} else {
				return
			}
		}
		if index < len(str) {
			deserializeInner(node)
		}
	}
	deserializeInner(tree.Root)
	return &tree
}
