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
