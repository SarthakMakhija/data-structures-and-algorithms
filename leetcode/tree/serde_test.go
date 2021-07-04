package tree_test

import (
	serde "github.com/SarthakMakhija/data-structures-and-algorithms/leetcode/tree"
	"testing"
)

func TestSerialize1(t *testing.T) {
	tree := binaryTree1()
	output := tree.Serialize()
	expected := "12XX34XX5XX"

	if output != expected {
		t.Fatalf("Expected %v, received %v", expected, output)
	}
}

func TestSerialize2(t *testing.T) {
	tree := binaryTree2()
	output := tree.Serialize()
	expected := "1248XXX5XX36XX7XX"

	if output != expected {
		t.Fatalf("Expected %v, received %v", expected, output)
	}
}

func TestDeserialize1(t *testing.T) {
	serialized := "12XX34XX5XX"
	tree := serde.Deserialize(serialized)
	output := tree.Serialize()

	if output != serialized {
		t.Fatalf("Expected %v, received %v", serialized, output)
	}
}

func TestDeserialize2(t *testing.T) {
	serialized := "1248XXX5XX36XX7XX"
	tree := serde.Deserialize(serialized)
	output := tree.Serialize()

	if output != serialized {
		t.Fatalf("Expected %v, received %v", serialized, output)
	}
}

func binaryTree1() *serde.IntBinaryTree {
	node2 := serde.IntNode{
		Value: 2,
	}
	node3 := serde.IntNode{
		Value: 3,
		Left: &serde.IntNode{
			Value: 4,
		},
		Right: &serde.IntNode{
			Value: 5,
		},
	}
	node1 := serde.IntNode{
		Value: 1,
		Left:  &node2,
		Right: &node3,
	}
	return &serde.IntBinaryTree{
		Root: &node1,
	}
}

func binaryTree2() *serde.IntBinaryTree {

	leaf8 := serde.IntNode{
		Value: 8,
	}
	leaf4 := serde.IntNode{
		Value: 4,
		Left:  &leaf8,
	}
	leaf5 := serde.IntNode{
		Value: 5,
	}
	leaf2 := serde.IntNode{
		Value: 2,
		Left:  &leaf4,
		Right: &leaf5,
	}

	leaf6 := serde.IntNode{
		Value: 6,
	}
	leaf7 := serde.IntNode{
		Value: 7,
	}
	leaf3 := serde.IntNode{
		Value: 3,
		Left:  &leaf6,
		Right: &leaf7,
	}

	leaf1 := serde.IntNode{
		Value: 1,
		Left:  &leaf2,
		Right: &leaf3,
	}
	return &serde.IntBinaryTree{
		Root: &leaf1,
	}
}
