package b_plus

const degree = 4

type ChildNodeReferences []*StringBPlusTreeNode

type StringBPlusTreeNode struct {
	Parent              *StringBPlusTreeNode
	ChildNodeReferences ChildNodeReferences
	Items               Items
}

func (node *StringBPlusTreeNode) GetOrEmpty(key int) Item {
	index, found := node.Items.InsertionIndexOf(key)
	if found {
		return node.Items[index]
	} else if len(node.ChildNodeReferences) > 0 {
		return node.ChildNodeReferences[index].GetOrEmpty(key)
	}
	return Item{}
}

func (node *StringBPlusTreeNode) Add(item Item) *StringBPlusTreeNode {
	var addInner func(node *StringBPlusTreeNode) *StringBPlusTreeNode
	addInner = func(node *StringBPlusTreeNode) *StringBPlusTreeNode {
		index, exists := node.Items.InsertionIndexOf(item.Key)
		if exists {
			return nil
		}
		if len(node.ChildNodeReferences) == 0 {
			node.Items.InsertAt(index, item)
			return node
		}
		return addInner(node.ChildNodeReferences[index])
	}
	return addInner(node)
}

func (node *StringBPlusTreeNode) AttemptSplit(parent *StringBPlusTreeNode, tree *StringBPlusTree) {
	nodeToInspect := node
	parentNode := parent

	for len(nodeToInspect.Items) >= tree.Root.maxItemsPerNode() {
		_, _, parentToInspect := nodeToInspect.SplitAt(tree.Root.maxItemsPerNode()/2, parentNode)
		if parentNode == nil {
			tree.Root = parentToInspect
		}
		nodeToInspect = parentToInspect
		parentNode = nodeToInspect.Parent
	}
}

func (node *StringBPlusTreeNode) SplitAt(index int, parent *StringBPlusTreeNode) (Item, *StringBPlusTreeNode, *StringBPlusTreeNode) {
	item := node.Items[index]
	newNode := emptyNewNode()
	newNode.Items = append(newNode.Items, node.Items[index:]...)
	node.Items = node.Items[:index]

	if len(node.ChildNodeReferences) > 0 {
		newNode.ChildNodeReferences = append(newNode.ChildNodeReferences, node.ChildNodeReferences[index+1:]...)
		node.ChildNodeReferences = node.ChildNodeReferences[:index+1]
	}

	if parent == nil {
		oldRoot := node
		newRoot := emptyNewNode()
		newRoot.Items = append(newRoot.Items, item)

		newRoot.ChildNodeReferences = append(newRoot.ChildNodeReferences, oldRoot, newNode)
		oldRoot.Parent = newRoot
		newNode.Parent = newRoot

		return item, newNode, newNode.Parent
	} else {
		newNode.Parent = parent
		parent.Items.Insert(item)
		childIndex, _ := parent.Items.InsertionIndexOf(newNode.Items[0].Key)
		parent.ChildNodeReferences.InsertAt(childIndex+1, newNode)

		return item, newNode, parent
	}
}

func (node *StringBPlusTreeNode) maxItemsPerNode() int {
	return degree - 1
}

func emptyNewNode() *StringBPlusTreeNode {
	return &StringBPlusTreeNode{}
}
