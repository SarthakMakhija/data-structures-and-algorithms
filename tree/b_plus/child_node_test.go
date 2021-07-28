package b_plus_test

import (
	"github.com/SarthakMakhija/data-structures-and-algorithms/tree/b_plus"
	"testing"
)

func TestChildNodeReferencesInsertAt(t *testing.T) {
	left := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   10,
				Value: "Mark",
			},
		},
	}
	right := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   30,
				Value: "Mathew",
			},
		},
	}
	root := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   20,
				Value: "Adam",
			},
			b_plus.Item{
				Key:   30,
				Value: "John",
			},
		},
		ChildNodeReferences: b_plus.ChildNodeReferences{
			&left, &right,
		},
	}

	root.ChildNodeReferences.InsertAt(1, &b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   50,
				Value: "Kathy",
			},
		},
	})
	childNodeReference := root.ChildNodeReferences[1]
	key := childNodeReference.Items[0].Key

	if key != 50 {
		t.Fatalf("Expected key as 50, received %v", key)
	}
}

func TestChildNodeReferencesInsertAtForAllReferences(t *testing.T) {
	left := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   10,
				Value: "Mark",
			},
		},
	}
	right := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   30,
				Value: "Mathew",
			},
		},
	}
	root := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   20,
				Value: "Adam",
			},
			b_plus.Item{
				Key:   30,
				Value: "John",
			},
		},
		ChildNodeReferences: b_plus.ChildNodeReferences{
			&left, &right,
		},
	}

	root.ChildNodeReferences.InsertAt(1, &b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   50,
				Value: "Kathy",
			},
		},
	})
	keyLeft := root.ChildNodeReferences[0].Items[0].Key
	keyMid := root.ChildNodeReferences[1].Items[0].Key
	keyRight := root.ChildNodeReferences[2].Items[0].Key

	if keyLeft != 10 {
		t.Fatalf("Expected key as 10, received %v", keyLeft)
	}
	if keyMid != 50 {
		t.Fatalf("Expected key as 50, received %v", keyMid)
	}
	if keyRight != 30 {
		t.Fatalf("Expected key as 30, received %v", keyRight)
	}
}
