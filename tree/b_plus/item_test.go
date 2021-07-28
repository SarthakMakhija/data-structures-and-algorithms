package b_plus_test

import (
	"github.com/SarthakMakhija/data-structures-and-algorithms/tree/b_plus"
	"reflect"
	"testing"
)

func TestInsertionIndexOfInTheMiddle(t *testing.T) {
	items := b_plus.Items{
		b_plus.Item{
			Key: 1,
		},
		b_plus.Item{
			Key: 3,
		},
	}

	index, _ := items.InsertionIndexOf(2)

	if index != 1 {
		t.Fatalf("Expected insertion index %v, received %v", 1, index)
	}
}

func TestInsertionIndexOfInTheEnd(t *testing.T) {
	items := b_plus.Items{
		b_plus.Item{
			Key: 1,
		},
		b_plus.Item{
			Key: 2,
		},
		b_plus.Item{
			Key: 3,
		},
	}

	index, _ := items.InsertionIndexOf(4)

	if index != 3 {
		t.Fatalf("Expected insertion index %v, received %v", 3, index)
	}
}

func TestInsertionIndexOfAtTheBeginning(t *testing.T) {
	items := b_plus.Items{
		b_plus.Item{
			Key: 2,
		},
		b_plus.Item{
			Key: 3,
		},
	}

	index, _ := items.InsertionIndexOf(1)

	if index != 0 {
		t.Fatalf("Expected insertion index %v, received %v", 0, index)
	}
}

func TestInsertionIndexOfIfTheKeyAlreadyExists1(t *testing.T) {
	items := b_plus.Items{
		b_plus.Item{
			Key: 1,
		},
		b_plus.Item{
			Key: 2,
		},
	}

	_, exists := items.InsertionIndexOf(2)

	if exists != true {
		t.Fatalf("Expected element to already exist, received %v", exists)
	}
}

func TestInsertionIndexOfIfTheKeyAlreadyExists2(t *testing.T) {
	items := b_plus.Items{
		b_plus.Item{
			Key: 1,
		},
		b_plus.Item{
			Key: 2,
		},
		b_plus.Item{
			Key: 3,
		},
	}

	index, _ := items.InsertionIndexOf(3)

	if index != 2 {
		t.Fatalf("Expected 2, received %v", index)
	}
}

func TestInsertAt1(t *testing.T) {
	items := b_plus.Items{
		b_plus.Item{
			Key: 1,
		},
		b_plus.Item{
			Key: 3,
		},
	}

	items.InsertAt(1, b_plus.Item{
		Key: 2,
	})

	allKeys := items.AllKeys()
	expected := []int{1, 2, 3}

	if !reflect.DeepEqual(allKeys, expected) {
		t.Fatalf("Expected %v, received %v", expected, allKeys)
	}
}

func TestInsertAt2(t *testing.T) {
	items := b_plus.Items{
		b_plus.Item{
			Key: 1,
		},
		b_plus.Item{
			Key: 3,
		},
		b_plus.Item{
			Key: 4,
		},
	}

	index, _ := items.InsertionIndexOf(5)
	items.InsertAt(index, b_plus.Item{
		Key: 5,
	})

	allKeys := items.AllKeys()
	expected := []int{1, 3, 4, 5}

	if !reflect.DeepEqual(allKeys, expected) {
		t.Fatalf("Expected %v, received %v", expected, allKeys)
	}
}
