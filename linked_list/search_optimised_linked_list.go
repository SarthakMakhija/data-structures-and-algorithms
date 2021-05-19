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

	bucket_index := int(math.Mod(float64(hash(element)), 10))
	if l.buckets[bucket_index].first == nil {
		bucket := SearchOptimisedBucket{
			first:   &node,
			current: &node,
		}
		l.buckets[bucket_index] = bucket
	} else {
		l.buckets[bucket_index].current.next = &node
		l.buckets[bucket_index].current = &node
	}
}

func (l *SearchOptimisedLinkedList) Contains(element string) bool {

	bucket_index := int(math.Mod(float64(hash(element)), 10))
	target_bucket := l.buckets[bucket_index]

	var contains_inner func(*SearchOptimisedNode) bool

	contains_inner = func(p *SearchOptimisedNode) bool {
		if p == nil {
			return false
		}
		if p.value == element {
			return true
		} else {
			return contains_inner(p.next)
		}
	}
	return contains_inner(target_bucket.first)
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
