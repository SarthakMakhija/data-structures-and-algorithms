package linkedlist

import (
	"hash/fnv"
	"math"
)

type SearchOptimisedLinkedList struct {
	buckets []SearchOptimisedBucket
}

type SearchOptimisedBucket struct {
	first   *SearchOptimisedNode
	current *SearchOptimisedNode
}

type SearchOptimisedNode struct {
	value string
	next  *SearchOptimisedNode
}

func NewSearchOptimisedLinkedList() SearchOptimisedLinkedList {
	return SearchOptimisedLinkedList{
		buckets: make([]SearchOptimisedBucket, 10),
	}
}

func (l *SearchOptimisedLinkedList) Add(element string) {

	node := SearchOptimisedNode{
		value: element,
		next:  nil,
	}

	bucketIndex := int(math.Mod(float64(hash(element)), 10))
	if l.buckets[bucketIndex].first == nil {
		bucket := SearchOptimisedBucket{
			first:   &node,
			current: &node,
		}
		l.buckets[bucketIndex] = bucket
	} else {
		l.buckets[bucketIndex].current.next = &node
		l.buckets[bucketIndex].current = &node
	}
}

func (l *SearchOptimisedLinkedList) Contains(element string) bool {

	bucketIndex := int(math.Mod(float64(hash(element)), 10))
	targetBucket := l.buckets[bucketIndex]

	var containsInner func(*SearchOptimisedNode) bool

	containsInner = func(p *SearchOptimisedNode) bool {
		if p == nil {
			return false
		}
		if p.value == element {
			return true
		} else {
			return containsInner(p.next)
		}
	}
	return containsInner(targetBucket.first)
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
