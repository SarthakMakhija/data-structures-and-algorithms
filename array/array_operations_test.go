package array_test

import (
	"github.com/SarthakMakhija/data-structures-and-algorithms/array"
	"reflect"
	"testing"
)

func TestInsert_1(t *testing.T) {
	arr := array.Array{
		Size: 2,
	}
	arr.Insert(10)
	arr.Insert(20)
	arr.Insert(30)

	expected := []int{10, 20, 30}
	if !reflect.DeepEqual(arr.All(), expected) {
		t.Fatalf("Expected %v but received %v", expected, arr.All())
	}
}

func TestInsert_2(t *testing.T) {
	arr := array.Array{
		Size: 1,
	}
	arr.Insert(10)
	arr.Insert(20)

	expected := []int{10, 20}
	if !reflect.DeepEqual(arr.All(), expected) {
		t.Fatalf("Expected %v but received %v", expected, arr.All())
	}
}

func TestDelete_1(t *testing.T) {
	arr := array.Array{
		Size: 5,
	}
	arr.Insert(10)
	arr.Insert(1)
	arr.Insert(5)
	arr.Insert(6)
	arr.Insert(20)

	arr.DeleteAt(0)
	expected := []int{1, 5, 6, 20}
	if !reflect.DeepEqual(arr.All(), expected) {
		t.Fatalf("Expected %v but received %v", expected, arr.All())
	}
}

func TestDelete_2(t *testing.T) {
	arr := array.Array{
		Size: 5,
	}
	arr.Insert(10)
	arr.Insert(1)
	arr.Insert(5)
	arr.Insert(6)
	arr.Insert(20)

	arr.DeleteAt(2)
	expected := []int{10, 1, 6, 20}
	if !reflect.DeepEqual(arr.All(), expected) {
		t.Fatalf("Expected %v but received %v", expected, arr.All())
	}
}

func TestDelete_3(t *testing.T) {
	arr := array.Array{
		Size: 5,
	}
	arr.Insert(10)
	arr.Insert(1)
	arr.Insert(5)
	arr.Insert(6)
	arr.Insert(20)

	arr.DeleteAt(4)
	expected := []int{10, 1, 5, 6}
	if !reflect.DeepEqual(arr.All(), expected) {
		t.Fatalf("Expected %v but received %v", expected, arr.All())
	}
}

func TestBinarySearch_1(t *testing.T) {
	arr := array.Array{
		Size: 5,
	}
	arr.Insert(1)
	arr.Insert(2)
	arr.Insert(3)
	arr.Insert(4)
	arr.Insert(5)

	contains := arr.BinarySearch(4)
	if contains != true {
		t.Fatalf("Expected true but received %v", contains)
	}
}

func TestBinarySearch_2(t *testing.T) {
	arr := array.Array{
		Size: 10,
	}
	arr.Insert(1)
	arr.Insert(2)
	arr.Insert(3)
	arr.Insert(4)
	arr.Insert(5)
	arr.Insert(6)
	arr.Insert(7)
	arr.Insert(8)
	arr.Insert(9)
	arr.Insert(10)

	contains := arr.BinarySearch(9)
	if contains != true {
		t.Fatalf("Expected true but received %v", contains)
	}
}

func TestBinarySearch_3(t *testing.T) {
	arr := array.Array{
		Size: 10,
	}
	arr.Insert(1)
	arr.Insert(2)
	arr.Insert(3)
	arr.Insert(4)
	arr.Insert(5)
	arr.Insert(6)
	arr.Insert(7)
	arr.Insert(8)
	arr.Insert(9)
	arr.Insert(10)

	contains := arr.BinarySearch(90)
	if contains != false {
		t.Fatalf("Expected false but received %v", contains)
	}
}

func TestMax_1(t *testing.T) {
	arr := array.Array{
		Size: 10,
	}
	arr.Insert(1)
	arr.Insert(2)
	arr.Insert(3)
	arr.Insert(4)
	arr.Insert(5)
	arr.Insert(6)
	arr.Insert(7)
	arr.Insert(8)
	arr.Insert(9)
	arr.Insert(10)

	max := arr.Max()
	if max != 10 {
		t.Fatalf("Expected %v but received %v", 10, max)
	}
}

func TestMax_2(t *testing.T) {
	arr := array.Array{
		Size: 1,
	}
	arr.Insert(1)

	max := arr.Max()
	if max != 1 {
		t.Fatalf("Expected %v but received %v", 1, max)
	}
}

func TestRotate_1(t *testing.T) {
	arr := array.Array{
		Size: 7,
	}
	arr.Insert(1)
	arr.Insert(2)
	arr.Insert(3)
	arr.Insert(4)
	arr.Insert(5)
	arr.Insert(6)
	arr.Insert(7)

	arr.Rotate(2)
	expected := []int{3, 4, 5, 6, 7, 1, 2}
	if !reflect.DeepEqual(arr.All(), expected) {
		t.Fatalf("Expected %v but received %v", expected, arr.All())
	}
}

func TestRotate_2(t *testing.T) {
	arr := array.Array{
		Size: 7,
	}
	arr.Insert(1)
	arr.Insert(2)
	arr.Insert(3)
	arr.Insert(4)
	arr.Insert(5)
	arr.Insert(6)
	arr.Insert(7)

	arr.Rotate(1)
	expected := []int{2, 3, 4, 5, 6, 7, 1}
	if !reflect.DeepEqual(arr.All(), expected) {
		t.Fatalf("Expected %v but received %v", expected, arr.All())
	}
}

func TestRotate_3(t *testing.T) {
	arr := array.Array{
		Size: 7,
	}
	arr.Insert(1)
	arr.Insert(2)
	arr.Insert(3)
	arr.Insert(4)
	arr.Insert(5)
	arr.Insert(6)
	arr.Insert(7)

	arr.Rotate(0)
	expected := []int{1, 2, 3, 4, 5, 6, 7}
	if !reflect.DeepEqual(arr.All(), expected) {
		t.Fatalf("Expected %v but received %v", expected, arr.All())
	}
}

func TestRotate_4(t *testing.T) {
	arr := array.Array{
		Size: 7,
	}
	arr.Insert(1)
	arr.Insert(2)
	arr.Insert(3)
	arr.Insert(4)
	arr.Insert(5)
	arr.Insert(6)
	arr.Insert(7)

	arr.Rotate(5)
	expected := []int{6, 7, 1, 2, 3, 4, 5}
	if !reflect.DeepEqual(arr.All(), expected) {
		t.Fatalf("Expected %v but received %v", expected, arr.All())
	}
}

func TestDuplicate_1(t *testing.T) {
	arr := array.Array{
		Size: 7,
	}
	arr.Insert(1)
	arr.Insert(2)
	arr.Insert(3)
	arr.Insert(4)
	arr.Insert(5)
	arr.Insert(6)
	arr.Insert(7)

	duplicates := arr.FindDuplicates()
	if len(duplicates) != 0 {
		t.Fatalf("Expected no duplicates but received %v", duplicates)
	}
}

func TestDuplicate_2(t *testing.T) {
	arr := array.Array{
		Size: 7,
	}
	arr.Insert(1)
	arr.Insert(2)
	arr.Insert(2)
	arr.Insert(4)
	arr.Insert(5)
	arr.Insert(5)
	arr.Insert(7)

	expected := []int{2, 5}
	duplicates := arr.FindDuplicates()
	if !reflect.DeepEqual(duplicates, expected) {
		t.Fatalf("Expected %v but received %v", expected, duplicates)
	}
}

func TestDuplicate_3(t *testing.T) {
	arr := array.Array{
		Size: 10,
	}
	arr.Insert(10)
	arr.Insert(45)
	arr.Insert(55)
	arr.Insert(45)
	arr.Insert(65)
	arr.Insert(65)
	arr.Insert(70)
	arr.Insert(80)
	arr.Insert(80)
	arr.Insert(90)

	expected := []int{65, 80, 45}
	duplicates := arr.FindDuplicates()
	if !reflect.DeepEqual(duplicates, expected) {
		t.Fatalf("Expected %v but received %v", expected, duplicates)
	}
}

func TestPairWithSumEqualToKForSortedArray_1(t *testing.T) {
	arr := array.Array{
		Size: 4,
	}
	arr.Insert(1)
	arr.Insert(2)
	arr.Insert(3)
	arr.Insert(4)

	expected := []array.Pair{
		{
			Element1: 3,
			Element2: 4,
		},
	}
	pairs := arr.PairWithSumEqualTo1(7)
	if !reflect.DeepEqual(pairs, expected) {
		t.Fatalf("Expected %v but received %v", expected, pairs)
	}
}

func TestPairWithSumEqualToKForSortedArray_2(t *testing.T) {
	arr := array.Array{
		Size: 4,
	}
	arr.Insert(1)
	arr.Insert(2)
	arr.Insert(3)
	arr.Insert(4)

	pairs := arr.PairWithSumEqualTo1(100)
	var expected []array.Pair
	if len(pairs) != 0 {
		t.Fatalf("Expected %v received %v", expected, pairs)
	}
}

func TestPairWithSumEqualToKForSortedArray_3(t *testing.T) {
	arr := array.Array{
		Size: 7,
	}
	arr.Insert(1)
	arr.Insert(2)
	arr.Insert(3)
	arr.Insert(4)
	arr.Insert(5)
	arr.Insert(6)
	arr.Insert(7)

	expected := []array.Pair{
		{
			Element1: 1,
			Element2: 6,
		},
		{
			Element1: 2,
			Element2: 5,
		},
		{
			Element1: 3,
			Element2: 4,
		},
	}
	pairs := arr.PairWithSumEqualTo1(7)
	if !reflect.DeepEqual(pairs, expected) {
		t.Fatalf("Expected %v but received %v", expected, pairs)
	}
}

func TestPairWithSumEqualToKForSortedArray_4(t *testing.T) {
	arr := array.Array{
		Size: 3,
	}
	arr.Insert(10)
	arr.Insert(20)
	arr.Insert(30)

	expected := []array.Pair{
		{
			Element1: 10,
			Element2: 20,
		},
	}
	pairs := arr.PairWithSumEqualTo1(30)
	if !reflect.DeepEqual(pairs, expected) {
		t.Fatalf("Expected %v but received %v", expected, pairs)
	}
}

func TestPairWithSumEqualToKForSortedArray_5(t *testing.T) {
	arr := array.Array{
		Size: 5,
	}
	arr.Insert(-1)
	arr.Insert(0)
	arr.Insert(1)
	arr.Insert(2)
	arr.Insert(3)

	expected := []array.Pair{
		{
			Element1: -1,
			Element2: 3,
		},
		{
			Element1: 0,
			Element2: 2,
		},
	}
	pairs := arr.PairWithSumEqualTo1(2)
	if !reflect.DeepEqual(pairs, expected) {
		t.Fatalf("Expected %v but received %v", expected, pairs)
	}
}

func TestPairWithSumEqualToKForSortedArray_6(t *testing.T) {
	arr := array.Array{
		Size: 8,
	}
	arr.Insert(1)
	arr.Insert(2)
	arr.Insert(2)
	arr.Insert(3)
	arr.Insert(4)
	arr.Insert(4)
	arr.Insert(5)
	arr.Insert(6)

	expected := []array.Pair{
		{
			Element1: 1,
			Element2: 5,
		},
		{
			Element1: 2,
			Element2: 4,
		},
		{
			Element1: 2,
			Element2: 4,
		},
	}
	pairs := arr.PairWithSumEqualTo1(6)
	if !reflect.DeepEqual(pairs, expected) {
		t.Fatalf("Expected %v but received %v", expected, pairs)
	}
}

func TestPairWithSumEqualToKForSortedArray_7(t *testing.T) {
	arr := array.Array{
		Size: 5,
	}
	arr.Insert(5)
	arr.Insert(10)
	arr.Insert(15)
	arr.Insert(20)
	arr.Insert(25)

	expected := []array.Pair{
		{
			Element1: 5,
			Element2: 20,
		},
		{
			Element1: 10,
			Element2: 15,
		},
	}
	pairs := arr.PairWithSumEqualTo1(25)
	if !reflect.DeepEqual(pairs, expected) {
		t.Fatalf("Expected %v but received %v", expected, pairs)
	}
}

func TestPairWithSumEqualToKForSortedArray_8(t *testing.T) {
	arr := array.Array{
		Size: 7,
	}
	arr.Insert(1)
	arr.Insert(2)
	arr.Insert(3)
	arr.Insert(4)
	arr.Insert(5)
	arr.Insert(6)
	arr.Insert(7)

	expected := []array.Pair{
		{
			Element1: 2,
			Element2: 7,
		},
		{
			Element1: 3,
			Element2: 6,
		},
		{
			Element1: 4,
			Element2: 5,
		},
	}
	pairs := arr.PairWithSumEqualTo1(9)
	if !reflect.DeepEqual(pairs, expected) {
		t.Fatalf("Expected %v but received %v", expected, pairs)
	}
}

func TestPairWithSumEqualToKForUnSortedArray_1(t *testing.T) {
	arr := array.Array{
		Size: 7,
	}
	arr.Insert(1)
	arr.Insert(3)
	arr.Insert(2)
	arr.Insert(5)
	arr.Insert(4)
	arr.Insert(6)
	arr.Insert(7)

	expected := []array.Pair{
		{
			Element1: 1,
			Element2: 6,
		},
		{
			Element1: 2,
			Element2: 5,
		},
		{
			Element1: 3,
			Element2: 4,
		},
	}
	pairs := arr.PairWithSumEqualTo2(7)
	if len(pairs) == 0 {
		t.Fatalf("Expected %v but received %v", expected, pairs)
	}
	for _, element := range pairs {
		if !((element.Element1 == 1 && element.Element2 == 6) ||
			(element.Element1 == 2 && element.Element2 == 5) ||
			(element.Element1 == 3 && element.Element2 == 4)) {
			t.Fatalf("Expected %v but received %v", expected, pairs)
		}
	}
}

func TestPairWithSumEqualToKForUnSortedArray_2(t *testing.T) {
	arr := array.Array{
		Size: 3,
	}
	arr.Insert(10)
	arr.Insert(30)
	arr.Insert(20)

	expected := []array.Pair{
		{
			Element1: 10,
			Element2: 20,
		},
	}
	pairs := arr.PairWithSumEqualTo2(30)
	if len(pairs) == 0 {
		t.Fatalf("Expected %v but received %v", expected, pairs)
	}
	for _, element := range pairs {
		if !(element.Element1 == 10 && element.Element2 == 20) {
			t.Fatalf("Expected %v but received %v", expected, pairs)
		}
	}
}

func TestPairWithSumEqualToKForUnSortedArray_3(t *testing.T) {
	arr := array.Array{
		Size: 5,
	}
	arr.Insert(0)
	arr.Insert(-1)
	arr.Insert(2)
	arr.Insert(1)
	arr.Insert(3)

	expected := []array.Pair{
		{
			Element1: -1,
			Element2: 3,
		},
		{
			Element1: 0,
			Element2: 2,
		},
	}
	pairs := arr.PairWithSumEqualTo2(2)
	if len(pairs) == 0 {
		t.Fatalf("Expected %v but received %v", expected, pairs)
	}
	for _, element := range pairs {
		if !((element.Element1 == -1 && element.Element2 == 3) ||
			(element.Element1 == 0 && element.Element2 == 2)) {
			t.Fatalf("Expected %v but received %v", expected, pairs)
		}
	}
}

func TestPairWithSumEqualToKForUnSortedArray_4(t *testing.T) {
	arr := array.Array{
		Size: 8,
	}
	arr.Insert(6)
	arr.Insert(1)
	arr.Insert(2)
	arr.Insert(4)
	arr.Insert(5)
	arr.Insert(3)
	arr.Insert(2)
	arr.Insert(4)

	expected := []array.Pair{
		{
			Element1: 1,
			Element2: 5,
		},
		{
			Element1: 2,
			Element2: 4,
		},
		{
			Element1: 2,
			Element2: 4,
		},
	}
	pairs := arr.PairWithSumEqualTo2(6)
	if len(pairs) == 0 {
		t.Fatalf("Expected %v but received %v", expected, pairs)
	}
	for _, element := range pairs {
		if !((element.Element1 == 1 && element.Element2 == 5) ||
			(element.Element1 == 2 && element.Element2 == 4)) {
			t.Fatalf("Expected %v but received %v", expected, pairs)
		}
	}
}

func TestSecondHighestElement_1(t *testing.T) {
	arr := array.Array{
		Size: 5,
	}
	arr.Insert(1)
	arr.Insert(2)
	arr.Insert(3)
	arr.Insert(4)
	arr.Insert(5)

	secondHighest := arr.SecondHighestElement()
	if secondHighest != 4 {
		t.Fatalf("Expected 4 but received %v", secondHighest)
	}
}

func TestSecondHighestElement_2(t *testing.T) {
	arr := array.Array{
		Size: 5,
	}
	arr.Insert(10)
	arr.Insert(70)
	arr.Insert(-3)
	arr.Insert(30)
	arr.Insert(50)

	secondHighest := arr.SecondHighestElement()
	if secondHighest != 50 {
		t.Fatalf("Expected 50 but received %v", secondHighest)
	}
}

func TestSecondHighestElement_3(t *testing.T) {
	arr := array.Array{
		Size: 2,
	}
	arr.Insert(10)
	arr.Insert(70)

	secondHighest := arr.SecondHighestElement()
	if secondHighest != 10 {
		t.Fatalf("Expected 10 but received %v", secondHighest)
	}
}

func TestSecondHighestElement_4(t *testing.T) {
	arr := array.Array{
		Size: 1,
	}
	arr.Insert(10)

	secondHighest := arr.SecondHighestElement()
	if secondHighest != 10 {
		t.Fatalf("Expected 10 but received %v", secondHighest)
	}
}

func TestNthHighestElement_1(t *testing.T) {
	arr := array.Array{
		Size: 5,
	}
	arr.Insert(1)
	arr.Insert(2)
	arr.Insert(3)
	arr.Insert(4)
	arr.Insert(5)

	secondHighest := arr.NthHighestElement(2)
	if secondHighest != 4 {
		t.Fatalf("Expected 4 but received %v", secondHighest)
	}
}

func TestNthHighestElement_2(t *testing.T) {
	arr := array.Array{
		Size: 5,
	}
	arr.Insert(1)
	arr.Insert(2)
	arr.Insert(3)
	arr.Insert(4)
	arr.Insert(5)

	secondHighest := arr.NthHighestElement(3)
	if secondHighest != 3 {
		t.Fatalf("Expected 3 but received %v", secondHighest)
	}
}

func TestNthHighestElement_3(t *testing.T) {
	arr := array.Array{
		Size: 5,
	}
	arr.Insert(1)
	arr.Insert(2)
	arr.Insert(3)
	arr.Insert(4)
	arr.Insert(5)

	secondHighest := arr.NthHighestElement(5)
	if secondHighest != 1 {
		t.Fatalf("Expected 1 but received %v", secondHighest)
	}
}

func TestNthHighestElement_4(t *testing.T) {
	arr := array.Array{
		Size: 5,
	}
	arr.Insert(-1)
	arr.Insert(20)
	arr.Insert(-3)
	arr.Insert(40)
	arr.Insert(-5)

	secondHighest := arr.NthHighestElement(3)
	if secondHighest != -1 {
		t.Fatalf("Expected -1 but received %v", secondHighest)
	}
}

func TestMerge_1(t *testing.T) {
	arr := array.Array{
		Size: 5,
	}
	other := array.Array{
		Size: 5,
	}
	arr.Insert(10)
	arr.Insert(20)
	arr.Insert(30)
	arr.Insert(40)
	arr.Insert(50)

	other.Insert(5)
	other.Insert(9)
	other.Insert(11)
	other.Insert(31)
	other.Insert(37)

	merged := arr.Merge(other)
	expected := []int{5, 9, 10, 11, 20, 30, 31, 37, 40, 50}
	if !reflect.DeepEqual(merged, expected) {
		t.Fatalf("Expected %v but received %v", expected, merged)
	}
}

func TestMerge_2(t *testing.T) {
	arr := array.Array{
		Size: 5,
	}
	other := array.Array{
		Size: 5,
	}
	arr.Insert(10)
	arr.Insert(20)
	arr.Insert(30)
	arr.Insert(40)
	arr.Insert(50)

	other.Insert(60)
	other.Insert(70)
	other.Insert(80)
	other.Insert(90)
	other.Insert(100)

	merged := arr.Merge(other)
	expected := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	if !reflect.DeepEqual(merged, expected) {
		t.Fatalf("Expected %v but received %v", expected, merged)
	}
}

func TestMerge_3(t *testing.T) {
	arr := array.Array{
		Size: 5,
	}
	other := array.Array{
		Size: 3,
	}
	arr.Insert(10)
	arr.Insert(20)
	arr.Insert(30)
	arr.Insert(40)
	arr.Insert(50)

	other.Insert(0)
	other.Insert(5)
	other.Insert(100)

	merged := arr.Merge(other)
	expected := []int{0, 5, 10, 20, 30, 40, 50, 100}
	if !reflect.DeepEqual(merged, expected) {
		t.Fatalf("Expected %v but received %v", expected, merged)
	}
}

func TestIntersect_1(t *testing.T) {
	arr := array.Array{
		Size: 2,
	}
	other := array.Array{
		Size: 1,
	}
	arr.Insert(10)
	arr.Insert(20)

	other.Insert(10)

	expected := []int{10}
	intersected := arr.IntersectWith(other)
	if !reflect.DeepEqual(intersected, expected) {
		t.Fatalf("Expected %v but received %v", expected, intersected)
	}
}

func TestIntersect_2(t *testing.T) {
	arr := array.Array{
		Size: 5,
	}
	other := array.Array{
		Size: 5,
	}
	arr.Insert(10)
	arr.Insert(30)
	arr.Insert(20)
	arr.Insert(50)
	arr.Insert(40)

	other.Insert(100)
	other.Insert(200)
	other.Insert(300)
	other.Insert(40)
	other.Insert(50)

	expected := []int{40, 50}
	intersected := arr.IntersectWith(other)
	if !reflect.DeepEqual(intersected, expected) {
		t.Fatalf("Expected %v but received %v", expected, intersected)
	}
}
