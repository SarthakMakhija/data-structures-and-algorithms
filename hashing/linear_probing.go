package hashing

import (
	"errors"
	"math"
)

const size int = 10

type FixedSizeBucket struct {
	entries [size]*int
}

func (f *FixedSizeBucket) Add(element int) (bool, error) {
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

func (f *FixedSizeBucket) Contains(element int) bool {
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

func (f *FixedSizeBucket) nextSlotAvailableAfter(index int) (int, error) {
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

func (f *FixedSizeBucket) indexAfter(index int) int {
	return int(math.Mod(float64(index+1), float64(f.size())))
}

func (f *FixedSizeBucket) indexOf(key int) int {
	return int(math.Mod(float64(key), float64(f.size())))
}

func (f *FixedSizeBucket) size() int {
	return len(f.entries)
}
