package skip_test

import (
	"github.com/SarthakMakhija/data-structures-and-algorithms/skip"
	"testing"
)

func TestCreatesASkipListAndGetsAValueByKey1(t *testing.T) {
	skipList := skip.NewList(5)
	skipList.Put(12, 120)
	skipList.Put(17, 170)
	skipList.Put(25, 250)
	skipList.Put(31, 310)
	skipList.Put(20, 200)
	skipList.Put(39, 390)

	value, _ := skipList.GetByKey(25)
	if value != 250 {
		t.Fatalf("Expected value to be %v, received %v", 250, value)
	}
}

func TestCreatesASkipListAndGetsAValueByKey2(t *testing.T) {
	skipList := skip.NewList(5)
	skipList.Put(12, 120)
	skipList.Put(17, 170)
	skipList.Put(25, 250)
	skipList.Put(31, 310)
	skipList.Put(20, 200)
	skipList.Put(39, 390)

	value, _ := skipList.GetByKey(39)
	if value != 390 {
		t.Fatalf("Expected value to be %v, received %v", 390, value)
	}
}

func TestCreatesASkipListAndGetsAValueByKey3(t *testing.T) {
	skipList := skip.NewList(5)
	skipList.Put(12, 120)
	skipList.Put(17, 170)
	skipList.Put(25, 250)
	skipList.Put(31, 310)
	skipList.Put(20, 200)
	skipList.Put(39, 390)

	value, _ := skipList.GetByKey(90)
	if value != -1 {
		t.Fatalf("Expected value to be %v, received %v", -1, value)
	}
}
