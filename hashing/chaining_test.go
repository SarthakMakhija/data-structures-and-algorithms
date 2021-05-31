package hashing_test

import (
	"github.com/SarthakMakhija/data-structures-and-algorithms/hashing"
	"reflect"
	"testing"
)

func TestFixedSizedMapFindFirstValueBy_1(t *testing.T) {
	fixedSizedMap := hashing.FixedSizedSortedMap{}

	fixedSizedMap.Add(65, 30)
	fixedSizedMap.Add(45, 10)
	fixedSizedMap.Add(55, 20)

	result := fixedSizedMap.FindFirstValueBy(65)
	if result != 10 {
		t.Fatalf("Expected %v, received %v", 10, result)
	}
}

func TestFixedSizedMapFindFirstValueBy_2(t *testing.T) {
	fixedSizedMap := hashing.FixedSizedSortedMap{}

	fixedSizedMap.Add(15, 30)
	fixedSizedMap.Add(15, 10)
	fixedSizedMap.Add(15, 20)

	result := fixedSizedMap.FindFirstValueBy(15)
	if result != 10 {
		t.Fatalf("Expected %v, received %v", 10, result)
	}
}

func TestFixedSizedMapFindFirstValueBy_3(t *testing.T) {
	fixedSizedMap := hashing.FixedSizedSortedMap{}

	fixedSizedMap.Add(14, 30)
	fixedSizedMap.Add(15, 10)
	fixedSizedMap.Add(16, 20)

	result := fixedSizedMap.FindFirstValueBy(16)
	if result != 20 {
		t.Fatalf("Expected %v, received %v", 20, result)
	}
}

func TestFixedSizedMapFindAllValuesBy_1(t *testing.T) {
	fixedSizedMap := hashing.FixedSizedSortedMap{}

	fixedSizedMap.Add(15, 30)
	fixedSizedMap.Add(15, 10)
	fixedSizedMap.Add(15, 20)

	result := fixedSizedMap.FindAllValuesBy(15)
	expected := []int{10, 20, 30}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected %v, received %v", expected, result)
	}
}

func TestFixedSizedMapFindAllValuesBy_2(t *testing.T) {
	fixedSizedMap := hashing.FixedSizedSortedMap{}

	fixedSizedMap.Add(65, 30)
	fixedSizedMap.Add(45, 10)
	fixedSizedMap.Add(55, 20)
	fixedSizedMap.Add(75, 40)
	fixedSizedMap.Add(85, 2)

	result := fixedSizedMap.FindAllValuesBy(65)
	expected := []int{2, 10, 20, 30, 40}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected %v, received %v", expected, result)
	}
}

func TestFixedSizedMapContains_1(t *testing.T) {
	fixedSizedMap := hashing.FixedSizedSortedMap{}

	fixedSizedMap.Add(65, 30)
	fixedSizedMap.Add(45, 10)
	fixedSizedMap.Add(55, 20)
	fixedSizedMap.Add(75, 40)
	fixedSizedMap.Add(85, 2)

	result := fixedSizedMap.Contains(65, 30)

	if result != true {
		t.Fatalf("Expected %v, received %v", true, result)
	}
}

func TestFixedSizedMapContains_2(t *testing.T) {
	fixedSizedMap := hashing.FixedSizedSortedMap{}

	fixedSizedMap.Add(65, 30)

	result := fixedSizedMap.Contains(70, 70)

	if result != false {
		t.Fatalf("Expected %v, received %v", false, result)
	}
}

func TestFixedSizedMapContains_3(t *testing.T) {
	fixedSizedMap := hashing.FixedSizedSortedMap{}

	fixedSizedMap.Add(65, 30)

	result := fixedSizedMap.Contains(65, 40)

	if result != false {
		t.Fatalf("Expected %v, received %v", false, result)
	}
}
