package b_plus_test

import (
	"github.com/SarthakMakhija/data-structures-and-algorithms/tree/b_plus"
	"testing"
)

func TestAddRootNode(t *testing.T) {
	tree := b_plus.StringBPlusTree{}
	tree.Add(b_plus.Item{
		Key:   10,
		Value: "Adam",
	})
	item := tree.GetOrEmpty(10)

	if item.Value != "Adam" {
		t.Fatalf("Expected Adam, recevied %v", item.Value)
	}
}

func TestAdd2ItemsAtRootNode(t *testing.T) {
	tree := b_plus.StringBPlusTree{}
	tree.Add(b_plus.Item{
		Key:   10,
		Value: "Adam",
	})
	tree.Add(b_plus.Item{
		Key:   20,
		Value: "John",
	})
	item := tree.GetOrEmpty(20)

	if item.Value != "John" {
		t.Fatalf("Expected John, recevied %v", item.Value)
	}
}
