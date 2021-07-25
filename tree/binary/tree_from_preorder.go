package binarytree

func FromPreOrderTraversalWithMarkers(nodeValues []string) *StringBinaryTree {
	if len(nodeValues) == 0 {
		return nil
	}
	stack := nodeStatusStack{}

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
		stack.push(nodeStatus{
			node:   node,
			status: status,
		})
	}
	pop := func() nodeStatus {
		poppedNodeStatus := stack.pop()
		for poppedNodeStatus.status == Done {
			poppedNodeStatus = stack.pop()
		}
		return poppedNodeStatus
	}

	root := createNode(nodeValues[0]) //assume root is not null or is not a marker
	push(root, LeftPending)

	for index := 1; index < len(nodeValues); index++ {
		poppedNodeStatus := pop()
		newNode := newNodeOrNil(nodeValues[index])

		switch poppedNodeStatus.status {
		case LeftPending:
			poppedNodeStatus.node.Left = newNode
			push(poppedNodeStatus.node, RightPending)
		case RightPending:
			poppedNodeStatus.node.Right = newNode
			push(poppedNodeStatus.node, Done)
		}
		if newNode != nil {
			push(newNode, LeftPending)
		}
	}
	stack.clear()
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

type nodeStatus struct {
	node   *StringNode
	status status
}

type nodeStatusStack struct {
	stack    []nodeStatus
	stackTop int
}

func (s *nodeStatusStack) push(element nodeStatus) {
	s.stack = append(s.stack, element)
	s.stackTop = s.stackTop + 1
}

func (s *nodeStatusStack) pop() nodeStatus {
	s.stackTop = s.stackTop - 1
	if s.stackTop < 0 {
		return nodeStatus{}
	}
	top := s.stack[s.stackTop]
	s.stack = s.stack[0:s.stackTop]
	return top
}

func (s *nodeStatusStack) clear() {
	s.stack = []nodeStatus{}
	s.stackTop = 0
}
