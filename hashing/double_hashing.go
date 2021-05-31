package hashing

import (
	"errors"
	"math"
)

type DoubleHashingFixedSizedBucket struct {
	entries [fixedBucketSize]*int
}

func (f *DoubleHashingFixedSizedBucket) Add(element int) (bool, error) {
	index := f.indexOf(element)
	if f.entries[index] == nil {
		f.entries[index] = &element
		return true, nil
	} else {
		nextIndex, err := f.nextSlotAvailableAfter(index, element)
		if err == nil {
			f.entries[nextIndex] = &element
			return true, nil
		} else {
			return false, err
		}
	}
}

func (f *DoubleHashingFixedSizedBucket) Contains(element int) bool {
	secondHash := f.secondHash(element)
	initialIndex := f.indexOf(element)

	index := initialIndex
	incrementFactor := 1

	for incrementFactor <= fixedBucketSize {
		if f.entries[index] != nil && *f.entries[index] == element {
			return true
		}
		index = int(math.Mod(float64(initialIndex+(incrementFactor*secondHash)), float64(f.size())))
		incrementFactor = incrementFactor + 1
	}
	return false
}

func (f *DoubleHashingFixedSizedBucket) nextSlotAvailableAfter(index int, element int) (int, error) {
	secondHash := f.secondHash(element)
	slotValue := f.entries[index]

	nextIndex := index
	incrementFactor := 1

	for slotValue != nil && incrementFactor <= fixedBucketSize {
		nextIndex = int(math.Mod(float64(index+(incrementFactor*secondHash)), float64(f.size())))
		slotValue = f.entries[nextIndex]
		incrementFactor = incrementFactor + 1
	}
	if f.entries[nextIndex] == nil {
		return nextIndex, nil
	} else {
		return -1, errors.New("NO SLOT AVAILABLE")
	}
}

func (f *DoubleHashingFixedSizedBucket) indexOf(key int) int {
	return int(math.Mod(float64(key), float64(f.size())))
}

func (f *DoubleHashingFixedSizedBucket) secondHash(element int) int {
	return largestPrimeLessThanBucketSize -
		int(math.Mod(float64(element), float64(largestPrimeLessThanBucketSize)))
}

func (f *DoubleHashingFixedSizedBucket) size() int {
	return len(f.entries)
}
