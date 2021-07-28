package b_plus_test

import (
	"github.com/SarthakMakhija/data-structures-and-algorithms/tree/b_plus"
	"reflect"
	"testing"
)

func TestGetOrEmpty1(t *testing.T) {
	left := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   10,
				Value: "Mark",
			},
		},
	}
	mid := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   20,
				Value: "Adam",
			},
			b_plus.Item{
				Key:   25,
				Value: "Alex",
			},
		},
	}
	right := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   30,
				Value: "Mathew",
			},
			b_plus.Item{
				Key:   40,
				Value: "Ryan",
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
			&left, &mid, &right,
		},
	}

	item := root.GetOrEmpty(25)
	if item.Value != "Alex" {
		t.Fatalf("Expected Alex, received %v", item.Value)
	}
}

func TestGetOrEmpty2(t *testing.T) {
	left := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   10,
				Value: "Mark",
			},
		},
	}
	mid := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   20,
				Value: "Adam",
			},
			b_plus.Item{
				Key:   25,
				Value: "Alex",
			},
		},
	}
	right := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   30,
				Value: "Mathew",
			},
			b_plus.Item{
				Key:   40,
				Value: "Ryan",
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
			&left, &mid, &right,
		},
	}

	item := root.GetOrEmpty(40)
	if item.Value != "Ryan" {
		t.Fatalf("Expected Ryan, received %v", item.Value)
	}
}

func TestGetOrEmpty3(t *testing.T) {
	left := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   10,
				Value: "Mark",
			},
		},
	}
	mid := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   20,
				Value: "Adam",
			},
			b_plus.Item{
				Key:   25,
				Value: "Alex",
			},
		},
	}
	right := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   30,
				Value: "Mathew",
			},
			b_plus.Item{
				Key:   40,
				Value: "Ryan",
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
			&left, &mid, &right,
		},
	}

	item := root.GetOrEmpty(10)
	if item.Value != "Mark" {
		t.Fatalf("Expected Mark, received %v", item.Value)
	}
}

func TestSplitAtAnIndexWithItemForRoot(t *testing.T) {
	node := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   20,
				Value: "Adam",
			},
			b_plus.Item{
				Key:   30,
				Value: "John",
			},
			b_plus.Item{
				Key:   40,
				Value: "Jack",
			},
		},
	}
	item, _, _ := node.SplitAt(1, nil)
	if item.Key != 30 {
		t.Fatalf("Expected split at key %v, received %v", 30, item.Key)
	}
}

func TestSplitAtAnIndexWithNewNodeForRoot(t *testing.T) {
	node := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   20,
				Value: "Adam",
			},
			b_plus.Item{
				Key:   30,
				Value: "John",
			},
			b_plus.Item{
				Key:   40,
				Value: "Jack",
			},
		},
	}
	_, newNode, _ := node.SplitAt(1, nil)
	expectedKeysOfNewNode := []int{30, 40}

	if !reflect.DeepEqual(newNode.Items.AllKeys(), expectedKeysOfNewNode) {
		t.Fatalf("Expected all keys of new node to be %v, received %v", expectedKeysOfNewNode, newNode.Items.AllKeys())
	}
}

func TestSplitAtAnIndexWithNewParentForRoot(t *testing.T) {
	node := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   20,
				Value: "Adam",
			},
			b_plus.Item{
				Key:   30,
				Value: "John",
			},
			b_plus.Item{
				Key:   40,
				Value: "Jack",
			},
		},
	}
	_, _, parent := node.SplitAt(1, nil)
	expectedKeysOfParent := []int{30}

	if !reflect.DeepEqual(parent.Items.AllKeys(), expectedKeysOfParent) {
		t.Fatalf("Expected all keys of parent node to be %v, received %v", expectedKeysOfParent, parent.Items.AllKeys())
	}
}

func TestStringBPlusTreeNodeAttemptSplitWithParentKeys(t *testing.T) {
	parent := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   10,
				Value: "Mark",
			},
		},
	}
	leaf := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   20,
				Value: "Adam",
			},
			b_plus.Item{
				Key:   30,
				Value: "John",
			},
			b_plus.Item{
				Key:   40,
				Value: "Jack",
			},
		},
		Parent: &parent,
	}

	parent.ChildNodeReferences = b_plus.ChildNodeReferences{
		nil, &leaf,
	}

	leaf.AttemptSplit(&parent, &b_plus.StringBPlusTree{Root: &parent})
	expectedKeysOfParent := []int{10, 30}

	if !reflect.DeepEqual(parent.Items.AllKeys(), expectedKeysOfParent) {
		t.Fatalf("Expected all keys of parent node to be %v, received %v", expectedKeysOfParent, parent.Items.AllKeys())
	}
}

func TestStringBPlusTreeNodeAttemptSplitWithParent(t *testing.T) {
	parent := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   10,
				Value: "Mark",
			},
		},
	}
	leaf := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   20,
				Value: "Adam",
			},
			b_plus.Item{
				Key:   30,
				Value: "John",
			},
			b_plus.Item{
				Key:   40,
				Value: "Jack",
			},
		},
		Parent: &parent,
	}

	parent.ChildNodeReferences = b_plus.ChildNodeReferences{
		nil, &leaf,
	}

	leaf.AttemptSplit(&parent, &b_plus.StringBPlusTree{Root: &parent})
	expectedKeysOfLeftChild := []int{20}

	if !reflect.DeepEqual(parent.ChildNodeReferences[1].Items.AllKeys(), expectedKeysOfLeftChild) {
		t.Fatalf("Expected all keys of parent node to be %v, received %v", expectedKeysOfLeftChild, parent.ChildNodeReferences[1].Items.AllKeys())
	}
}

func TestStringBPlusTreeNodeAttemptSplitWithLeftChild(t *testing.T) {
	parent := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   10,
				Value: "Mark",
			},
		},
	}
	leaf := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   20,
				Value: "Adam",
			},
			b_plus.Item{
				Key:   30,
				Value: "John",
			},
			b_plus.Item{
				Key:   40,
				Value: "Jack",
			},
		},
		Parent: &parent,
	}

	parent.ChildNodeReferences = b_plus.ChildNodeReferences{
		nil, &leaf,
	}

	leaf.AttemptSplit(&parent, &b_plus.StringBPlusTree{Root: &parent})
	expectedKeysOfLeftChild := []int{20}

	if !reflect.DeepEqual(parent.ChildNodeReferences[1].Items.AllKeys(), expectedKeysOfLeftChild) {
		t.Fatalf("Expected all keys of parent node to be %v, received %v", expectedKeysOfLeftChild, parent.ChildNodeReferences[1].Items.AllKeys())
	}
}

func TestStringBPlusTreeNodeAttemptSplitWithRightChild(t *testing.T) {
	parent := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   10,
				Value: "Mark",
			},
		},
	}
	leaf := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   20,
				Value: "Adam",
			},
			b_plus.Item{
				Key:   30,
				Value: "John",
			},
			b_plus.Item{
				Key:   40,
				Value: "Jack",
			},
		},
		Parent: &parent,
	}

	parent.ChildNodeReferences = b_plus.ChildNodeReferences{
		nil, &leaf,
	}

	leaf.AttemptSplit(&parent, &b_plus.StringBPlusTree{Root: &parent})
	expectedKeysOfRightChild := []int{30, 40}

	if !reflect.DeepEqual(parent.ChildNodeReferences[2].Items.AllKeys(), expectedKeysOfRightChild) {
		t.Fatalf("Expected all keys of parent node to be %v, received %v", expectedKeysOfRightChild, parent.ChildNodeReferences[2].Items.AllKeys())
	}
}

func TestStringBPlusTreeNodeAttemptMultilevelSplit1(t *testing.T) {
	parent := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   30,
				Value: "Mark",
			},
			b_plus.Item{
				Key:   50,
				Value: "John",
			},
		},
	}
	left := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   10,
				Value: "Carl",
			},
			b_plus.Item{
				Key:   20,
				Value: "Cam",
			},
		},
	}
	mid := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   30,
				Value: "Jim",
			},
			b_plus.Item{
				Key:   40,
				Value: "Tim",
			},
		},
	}
	right := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   70,
				Value: "Rom",
			},
			b_plus.Item{
				Key:   80,
				Value: "Tom",
			},
			b_plus.Item{
				Key:   90,
				Value: "Jom",
			},
		},
	}

	parent.ChildNodeReferences = b_plus.ChildNodeReferences{
		&left, &mid, &right,
	}

	newRoot := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   50,
				Value: "John",
			},
		},
	}

	tree := &b_plus.StringBPlusTree{Root: &parent}
	right.AttemptSplit(&parent, tree)

	if !reflect.DeepEqual(tree.Root.Items, newRoot.Items) {
		t.Fatalf("Expected new root to be %v, received %v", newRoot.Items, tree.Root.Items)
	}
}

func TestStringBPlusTreeNodeAttemptMultilevelSplit2(t *testing.T) {
	parent := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   30,
				Value: "Mark",
			},
			b_plus.Item{
				Key:   50,
				Value: "John",
			},
		},
	}
	left := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   10,
				Value: "Carl",
			},
			b_plus.Item{
				Key:   20,
				Value: "Cam",
			},
		},
	}
	mid := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   30,
				Value: "Jim",
			},
			b_plus.Item{
				Key:   40,
				Value: "Tim",
			},
		},
	}
	right := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   70,
				Value: "Rom",
			},
			b_plus.Item{
				Key:   80,
				Value: "Tom",
			},
			b_plus.Item{
				Key:   90,
				Value: "Jom",
			},
		},
	}

	parent.ChildNodeReferences = b_plus.ChildNodeReferences{
		&left, &mid, &right,
	}

	rootLeft := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   30,
				Value: "Jim",
			},
		},
	}

	tree := &b_plus.StringBPlusTree{Root: &parent}
	right.AttemptSplit(&parent, tree)

	if !reflect.DeepEqual(tree.Root.ChildNodeReferences[0].Items.AllKeys(), rootLeft.Items.AllKeys()) {
		t.Fatalf("Expected root's left to be %v, received %v", rootLeft.Items.AllKeys(), tree.Root.ChildNodeReferences[0].Items.AllKeys())
	}
}

func TestStringBPlusTreeNodeAttemptMultilevelSplit3(t *testing.T) {
	parent := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   30,
				Value: "Mark",
			},
			b_plus.Item{
				Key:   50,
				Value: "John",
			},
		},
	}
	left := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   10,
				Value: "Carl",
			},
			b_plus.Item{
				Key:   20,
				Value: "Cam",
			},
		},
	}
	mid := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   30,
				Value: "Jim",
			},
			b_plus.Item{
				Key:   40,
				Value: "Tim",
			},
		},
	}
	right := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   70,
				Value: "Rom",
			},
			b_plus.Item{
				Key:   80,
				Value: "Tom",
			},
			b_plus.Item{
				Key:   90,
				Value: "Jom",
			},
		},
	}

	parent.ChildNodeReferences = b_plus.ChildNodeReferences{
		&left, &mid, &right,
	}

	rootRight := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   50,
				Value: "John",
			},
			b_plus.Item{
				Key:   80,
				Value: "Tom",
			},
		},
	}

	tree := &b_plus.StringBPlusTree{Root: &parent}
	right.AttemptSplit(&parent, tree)

	if !reflect.DeepEqual(tree.Root.ChildNodeReferences[1].Items.AllKeys(), rootRight.Items.AllKeys()) {
		t.Fatalf("Expected root's right to be %v, received %v", rootRight.Items.AllKeys(), tree.Root.ChildNodeReferences[1].Items.AllKeys())
	}
}

func TestStringBPlusTreeNodeAttemptMultilevelSplit4(t *testing.T) {
	parent := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   30,
				Value: "Mark",
			},
			b_plus.Item{
				Key:   50,
				Value: "John",
			},
		},
	}
	left := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   10,
				Value: "Carl",
			},
			b_plus.Item{
				Key:   20,
				Value: "Cam",
			},
		},
	}
	mid := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   30,
				Value: "Jim",
			},
			b_plus.Item{
				Key:   40,
				Value: "Tim",
			},
		},
	}
	right := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   70,
				Value: "Rom",
			},
			b_plus.Item{
				Key:   80,
				Value: "Tom",
			},
			b_plus.Item{
				Key:   90,
				Value: "Jom",
			},
		},
	}

	parent.ChildNodeReferences = b_plus.ChildNodeReferences{
		&left, &mid, &right,
	}

	leafLeft := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   10,
				Value: "Carl",
			},
			b_plus.Item{
				Key:   20,
				Value: "Cam",
			},
		},
	}

	tree := &b_plus.StringBPlusTree{Root: &parent}
	right.AttemptSplit(&parent, tree)

	if !reflect.DeepEqual(tree.Root.ChildNodeReferences[0].ChildNodeReferences[0].Items.AllKeys(), leafLeft.Items.AllKeys()) {
		t.Fatalf("Expected left leaf to be %v, received %v", leafLeft.Items.AllKeys(), tree.Root.ChildNodeReferences[0].ChildNodeReferences[0].Items.AllKeys())
	}
}

func TestStringBPlusTreeNodeAttemptMultilevelSplit5(t *testing.T) {
	parent := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   30,
				Value: "Mark",
			},
			b_plus.Item{
				Key:   50,
				Value: "John",
			},
		},
	}
	left := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   10,
				Value: "Carl",
			},
			b_plus.Item{
				Key:   20,
				Value: "Cam",
			},
		},
	}
	mid := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   30,
				Value: "Jim",
			},
			b_plus.Item{
				Key:   40,
				Value: "Tim",
			},
		},
	}
	right := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   70,
				Value: "Rom",
			},
			b_plus.Item{
				Key:   80,
				Value: "Tom",
			},
			b_plus.Item{
				Key:   90,
				Value: "Jom",
			},
		},
	}

	parent.ChildNodeReferences = b_plus.ChildNodeReferences{
		&left, &mid, &right,
	}

	leafRight := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   30,
				Value: "Jim",
			},
			b_plus.Item{
				Key:   40,
				Value: "Tim",
			},
		},
	}

	tree := &b_plus.StringBPlusTree{Root: &parent}
	right.AttemptSplit(&parent, tree)

	if !reflect.DeepEqual(tree.Root.ChildNodeReferences[0].ChildNodeReferences[1].Items.AllKeys(), leafRight.Items.AllKeys()) {
		t.Fatalf("Expected right leaf to be %v, received %v", leafRight.Items.AllKeys(), tree.Root.ChildNodeReferences[0].ChildNodeReferences[1].Items.AllKeys())
	}
}

func TestStringBPlusTreeNodeAttemptMultilevelSplit6(t *testing.T) {
	parent := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   30,
				Value: "Mark",
			},
			b_plus.Item{
				Key:   50,
				Value: "John",
			},
		},
	}
	left := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   10,
				Value: "Carl",
			},
			b_plus.Item{
				Key:   20,
				Value: "Cam",
			},
		},
	}
	mid := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   30,
				Value: "Jim",
			},
			b_plus.Item{
				Key:   40,
				Value: "Tim",
			},
		},
	}
	right := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   70,
				Value: "Rom",
			},
			b_plus.Item{
				Key:   80,
				Value: "Tom",
			},
			b_plus.Item{
				Key:   90,
				Value: "Jom",
			},
		},
	}

	parent.ChildNodeReferences = b_plus.ChildNodeReferences{
		&left, &mid, &right,
	}

	leafLeft := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   70,
				Value: "Rom",
			},
		},
	}

	tree := &b_plus.StringBPlusTree{Root: &parent}
	right.AttemptSplit(&parent, tree)

	if !reflect.DeepEqual(tree.Root.ChildNodeReferences[1].ChildNodeReferences[0].Items.AllKeys(), leafLeft.Items.AllKeys()) {
		t.Fatalf("Expected left leaf to be %v, received %v", leafLeft.Items.AllKeys(), tree.Root.ChildNodeReferences[1].ChildNodeReferences[0].Items.AllKeys())
	}
}

func TestStringBPlusTreeNodeAttemptMultilevelSplit7(t *testing.T) {
	parent := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   30,
				Value: "Mark",
			},
			b_plus.Item{
				Key:   50,
				Value: "John",
			},
		},
	}
	left := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   10,
				Value: "Carl",
			},
			b_plus.Item{
				Key:   20,
				Value: "Cam",
			},
		},
	}
	mid := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   30,
				Value: "Jim",
			},
			b_plus.Item{
				Key:   40,
				Value: "Tim",
			},
		},
	}
	right := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   70,
				Value: "Rom",
			},
			b_plus.Item{
				Key:   80,
				Value: "Tom",
			},
			b_plus.Item{
				Key:   90,
				Value: "Jom",
			},
		},
	}

	parent.ChildNodeReferences = b_plus.ChildNodeReferences{
		&left, &mid, &right,
	}

	leafRight := b_plus.StringBPlusTreeNode{
		Items: b_plus.Items{
			b_plus.Item{
				Key:   80,
				Value: "Tom",
			},
			b_plus.Item{
				Key:   90,
				Value: "Jom",
			},
		},
	}

	tree := &b_plus.StringBPlusTree{Root: &parent}
	right.AttemptSplit(&parent, tree)

	if !reflect.DeepEqual(tree.Root.ChildNodeReferences[1].ChildNodeReferences[1].Items.AllKeys(), leafRight.Items.AllKeys()) {
		t.Fatalf("Expected right leaf to be %v, received %v", leafRight.Items.AllKeys(), tree.Root.ChildNodeReferences[1].ChildNodeReferences[1].Items.AllKeys())
	}
}
