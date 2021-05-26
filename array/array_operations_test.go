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
