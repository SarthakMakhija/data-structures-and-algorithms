package b_plus

import (
	"sort"
)

type Items []Item
type Item struct {
	Key   int
	Value string
}

func (items Items) InsertionIndexOf(key int) (int, bool) {
	index := sort.Search(
		len(items),
		func(index int) bool {
			return items[index].Key >= key
		})

	if index < len(items) {
		if items[index].Key == key {
			return index, true
		}
		return index, false
	}
	return index, false
}

func (items *Items) InsertAt(index int, item Item) {
	*items = append(*items, Item{})
	if index < len(*items) {
		copy((*items)[index+1:], (*items)[index:])
	}
	(*items)[index] = item
}

func (items *Items) Insert(item Item) {
	index, exists := items.InsertionIndexOf(item.Key)
	if !exists {
		items.InsertAt(index, item)
	}
}

func (items Items) AllKeys() []int {
	var keys []int
	for _, item := range items {
		keys = append(keys, item.Key)
	}
	return keys
}
