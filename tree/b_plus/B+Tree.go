package b_plus

type StringBPlusTree struct {
	Root *StringBPlusTreeNode
}

func (tree *StringBPlusTree) Add(item Item) {
	if tree.Root == nil {
		tree.Root = emptyNewNode()
		tree.Root.Items = append(tree.Root.Items, item)
		return
	}
	insertedAtNode := tree.Root.Add(item)
	insertedAtNode.AttemptSplit(insertedAtNode.Parent, tree)
	return
}

func (tree *StringBPlusTree) GetOrEmpty(key int) Item {
	if tree.Root == nil {
		return Item{}
	}
	return tree.Root.GetOrEmpty(key)
}
