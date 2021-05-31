package hashing

import "math"

type Node struct {
	key   int
	value int
	next  *Node
}

const fixedSize int = 10

type FixedSizedSortedMap struct {
	entries []*Node
}

func (m *FixedSizedSortedMap) Add(key int, value int) {
	if m.size() == 0 {
		m.entries = make([]*Node, fixedSize)
	}
	addAt := func(index int, node *Node) {
		head := m.entries[index]
		if head == nil || (node.value < head.value) {
			node.next = head
			m.entries[index] = node
		} else {
			tail := head
			for head != nil && node.value > head.value {
				tail = head
				head = head.next
			}
			tail.next = node
			node.next = head
		}
	}
	index := m.indexOf(key)
	addAt(index, &Node{
		key:   key,
		value: value,
		next:  nil,
	})
}

func (m *FixedSizedSortedMap) FindFirstValueBy(key int) int {
	index := m.indexOf(key)
	head := m.entries[index]
	return head.value
}

func (m *FixedSizedSortedMap) FindAllValuesBy(key int) []int {
	var values []int
	index := m.indexOf(key)
	head := m.entries[index]
	for head != nil {
		values = append(values, head.value)
		head = head.next
	}
	return values
}

func (m *FixedSizedSortedMap) Contains(key int, value int) bool {
	index := m.indexOf(key)
	head := m.entries[index]
	for head != nil {
		if head.value == value {
			return true
		} else if head.value > value { //Because the nodes are sorted by value, we don't need to proceed further
			break
		}
		head = head.next
	}
	return false
}

func (m *FixedSizedSortedMap) indexOf(key int) int {
	return int(math.Mod(float64(key), float64(m.size())))
}

func (m *FixedSizedSortedMap) size() int {
	return len(m.entries)
}
