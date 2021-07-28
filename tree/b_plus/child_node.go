package b_plus

func (childNodeReferences *ChildNodeReferences) InsertAt(index int, childNodeReference *StringBPlusTreeNode) {
	*childNodeReferences = append(*childNodeReferences, childNodeReference)
	if index < len(*childNodeReferences) {
		copy((*childNodeReferences)[index+1:], (*childNodeReferences)[index:])
	}
	(*childNodeReferences)[index] = childNodeReference
}
