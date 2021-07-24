package binarytree

import "strconv"

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
