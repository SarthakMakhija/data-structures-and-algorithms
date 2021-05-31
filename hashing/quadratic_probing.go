package hashing

import (
	"errors"
	"math"
)

type Bucket struct {
	entries [size]*int
}

func (f *Bucket) Add(element int) (bool, error) {
	index := f.indexOf(element)
	if f.entries[index] == nil {
		f.entries[index] = &element
		return true, nil
	} else {
		nextIndex, err := f.nextSlotAvailableAfter(index)
		//fmt.Println("nextindex ", nextIndex, " for element ", element)
		if err == nil {
			f.entries[nextIndex] = &element
			return true, nil
		} else {
			return false, err
		}
	}
}

func (f *Bucket) Contains(element int) bool {
	initialIndex := f.indexOf(element)
	index := initialIndex
	incrementFactor := 1

	for incrementFactor <= size {
		if f.entries[index] != nil && *f.entries[index] == element {
			return true
		}
		index = int(math.Mod(float64(initialIndex+(incrementFactor*incrementFactor)), float64(f.size())))
		incrementFactor = incrementFactor + 1
	}
	return false
}

func (f *Bucket) nextSlotAvailableAfter(index int) (int, error) {
	slotValue := f.entries[index]
	nextIndex := index
	incrementFactor := 1

	for slotValue != nil && incrementFactor <= size {
		nextIndex = int(math.Mod(float64(index+(incrementFactor*incrementFactor)), float64(f.size())))
		slotValue = f.entries[nextIndex]
		incrementFactor = incrementFactor + 1
	}
	if f.entries[nextIndex] == nil {
		return nextIndex, nil
	} else {
		return -1, errors.New("NO SLOT AVAILABLE")
	}
}

func (f *Bucket) indexOf(key int) int {
	return int(math.Mod(float64(key), float64(f.size())))
}

func (f *Bucket) size() int {
	return len(f.entries)
}