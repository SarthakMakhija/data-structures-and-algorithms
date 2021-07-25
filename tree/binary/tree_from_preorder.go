package binarytree

func FromPreOrderTraversalWithMarkers(nodeValues []string) *StringBinaryTree {
	if len(nodeValues) == 0 {
		return nil
	}

	stack := NodeStatusStack{}
	createNode := func(nodeValue string) *StringNode {
		return &StringNode{
			Value: nodeValue,
		}
	}
	isAMarker := func(nodeValue string) bool {
		if nodeValue == "" {
			return true
		}
		return false
	}
	newNodeOrNil := func(nodeValue string) *StringNode {
		if isAMarker(nodeValue) {
			return nil
		} else {
			return createNode(nodeValue)
		}
	}
	push := func(node *StringNode, status status) {
		stack.Push(NodeStatus{
			node:   node,
			status: status,
		})
	}
	pop := func() NodeStatus {
		nodeStatus := stack.Pop()
		for nodeStatus.status == Done {
			nodeStatus = stack.Pop()
		}
		return nodeStatus
	}

	root := createNode(nodeValues[0]) //assume root is not null or is not a marker
	push(root, LeftPending)

	for index := 1; index < len(nodeValues); index++ {
		nodeStatus := pop()
		//fmt.Println(nodeStatus.node.Value)
		newNode := newNodeOrNil(nodeValues[index])

		switch nodeStatus.status {
		case LeftPending:
			nodeStatus.node.Left = newNode
			push(nodeStatus.node, RightPending)
		case RightPending:
			nodeStatus.node.Right = newNode
			push(nodeStatus.node, Done)
		}
		if newNode != nil {
			push(newNode, LeftPending)
		}
	}

	return &StringBinaryTree{
		Root: root,
	}
}

type status int

const (
	LeftPending status = iota
	RightPending
	Done
)

type NodeStatus struct {
	node   *StringNode
	status status
}

type NodeStatusStack struct {
	stack    []NodeStatus
	stackTop int
}

func (s *NodeStatusStack) Push(element NodeStatus) {
	s.stack = append(s.stack, element)
	s.stackTop = s.stackTop + 1
}

func (s *NodeStatusStack) Pop() NodeStatus {
	s.stackTop = s.stackTop - 1
	if s.stackTop < 0 {
		return NodeStatus{}
	}
	top := s.stack[s.stackTop]
	s.stack = s.stack[0:s.stackTop]
	return top
}
