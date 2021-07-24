package binarytree

import (
	"strconv"
)

func (t IntBinaryTree) BinaryEncodings() []string {
	if t.Root == nil {
		return []string{}
	}
	var encodings []string
	var binaryEncodingsInner func(node *IntNode, bits string)
	binaryEncodingsInner = func(node *IntNode, bits string) {
		if node == nil {
			return
		} else if node.Left == nil && node.Right == nil {
			encodings = append(encodings, bits+strconv.Itoa(node.Value))
			return
		} else {
			binaryEncodingsInner(node.Left, bits+strconv.Itoa(node.Value))
			binaryEncodingsInner(node.Right, bits+strconv.Itoa(node.Value))
			return
		}
	}

	binaryEncodingsInner(t.Root, "")
	return encodings
}

func (t IntBinaryTree) BinaryEncodingsSum() int {
	if t.Root == nil {
		return -1
	}
	sum := 0
	var binaryEncodingsInner func(node *IntNode, bits string)
	binaryEncodingsInner = func(node *IntNode, bits string) {
		if node == nil {
			return
		} else if node.Left == nil && node.Right == nil {
			sum = sum + BinaryToDecimal(bits+strconv.Itoa(node.Value))
			return
		} else {
			binaryEncodingsInner(node.Left, bits+strconv.Itoa(node.Value))
			binaryEncodingsInner(node.Right, bits+strconv.Itoa(node.Value))
			return
		}
	}

	binaryEncodingsInner(t.Root, "")
	return sum
}

func BinaryToDecimal(str string) int {
	if str == "" {
		return 0
	}
	base := 1
	sum := 0
	for index := len(str) - 1; index >= 0; index-- {
		bit := string(str[index])
		if bit == "1" {
			sum = sum + base*1
		} else {
			sum = sum + base*0
		}
		base = base << 1
	}
	return sum
}
