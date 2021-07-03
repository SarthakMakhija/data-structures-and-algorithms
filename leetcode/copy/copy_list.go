package copy

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}
type NodeValue struct {
	Value       int
	RandomValue int
}

func CopyWithRandomPointer(head *Node) *Node {
	var newFirst, newHead *Node = nil, nil
	sourceFirst := head
	copiedNodeAddressBySourceAddress := make(map[*Node]*Node)

	for sourceFirst != nil {
		newNode, ok := copiedNodeAddressBySourceAddress[sourceFirst]
		if !ok {
			newNode = copyNodeWith(sourceFirst.Val)
		}
		if newHead == nil {
			newHead, newFirst = newNode, newNode
		} else {
			newHead.Next = newNode
			newHead = newHead.Next
		}
		if sourceFirst.Random != nil {
			node := copyNodeWith(sourceFirst.Random.Val)
			newNode.Random = node
			copiedNodeAddressBySourceAddress[sourceFirst.Random] = node
		}
		sourceFirst = sourceFirst.Next
	}
	return newFirst
}

func (node *Node) All() []NodeValue {
	head := node
	var nodeValues []NodeValue

	for head != nil {
		value := NodeValue{
			Value: head.Val,
		}
		if head.Random != nil {
			value.RandomValue = head.Random.Val
		}
		nodeValues = append(nodeValues, value)
		head = head.Next
	}
	return nodeValues
}

func copyNodeWith(value int) *Node {
	return &Node{
		Val: value,
	}
}
