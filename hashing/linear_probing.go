package hashing

import (
	"errors"
	"math"
)

type LinearProbingFixedSizedBucket struct {
	entries [fixedBucketSize]*int
}

func (f *LinearProbingFixedSizedBucket) Add(element int) (bool, error) {
	index := f.indexOf(element)
	if f.entries[index] == nil {
		f.entries[index] = &element
		return true, nil
	} else {
		nextIndex, err := f.nextSlotAvailableAfter(index)
		if err == nil {
			f.entries[nextIndex] = &element
			return true, nil
		} else {
			return false, err
		}
	}
}

func (f *LinearProbingFixedSizedBucket) Contains(element int) bool {
	initialIndex := f.indexOf(element)
	index := initialIndex

	for f.entries[index] != nil {
		if *f.entries[index] == element {
			return true
		}
		index = f.indexAfter(index)
		if index == initialIndex {
			return false
		}
	}
	return false
}

func (f *LinearProbingFixedSizedBucket) nextSlotAvailableAfter(index int) (int, error) {
	slotValue := f.entries[index]
	nextIndex := index
	for slotValue != nil {
		nextIndex = f.indexAfter(nextIndex)
		slotValue = f.entries[nextIndex]
		if nextIndex == index {
			return -1, errors.New("NO SLOT AVAILABLE")
		}
	}
	return nextIndex, nil
}

func (f *LinearProbingFixedSizedBucket) indexAfter(index int) int {
	return int(math.Mod(float64(index+1), float64(f.size())))
}

func (f *LinearProbingFixedSizedBucket) indexOf(key int) int {
	return int(math.Mod(float64(key), float64(f.size())))
}

func (f *LinearProbingFixedSizedBucket) size() int {
	return len(f.entries)
}
