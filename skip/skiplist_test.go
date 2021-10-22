package skip_test

import (
	"github.com/SarthakMakhija/data-structures-and-algorithms/skip"
	"testing"
)

func TestCreatesASkipListAndGetsAValueByKey1(t *testing.T) {
	skipList := skip.NewList(5)
	skipList.PutNew(12, 120)
	skipList.PutNew(17, 170)
	skipList.PutNew(25, 250)
	skipList.PutNew(31, 310)
	skipList.PutNew(20, 200)
	skipList.PutNew(39, 390)

	value, _ := skipList.GetByKey(25)
	if value != 250 {
		t.Fatalf("Expected value to be %v, received %v", 250, value)
	}
}

func TestCreatesASkipListAndGetsAValueByKey2(t *testing.T) {
	skipList := skip.NewList(5)
	skipList.PutNew(12, 120)
	skipList.PutNew(17, 170)
	skipList.PutNew(25, 250)
	skipList.PutNew(31, 310)
	skipList.PutNew(20, 200)
	skipList.PutNew(39, 390)

	value, _ := skipList.GetByKey(12)
	if value != 120 {
		t.Fatalf("Expected value to be %v, received %v", 120, value)
	}
}

func TestCreatesASkipListAndGetsAValueByKey3(t *testing.T) {
	skipList := skip.NewList(5)
	skipList.PutNew(12, 120)
	skipList.PutNew(17, 170)
	skipList.PutNew(25, 250)
	skipList.PutNew(31, 310)
	skipList.PutNew(20, 200)
	skipList.PutNew(39, 390)

	value, _ := skipList.GetByKey(39)
	if value != 390 {
		t.Fatalf("Expected value to be %v, received %v", 390, value)
	}
}

func TestCreatesASkipListAndGetsAValueByKeyGivenKeyDoesNotExist(t *testing.T) {
	skipList := skip.NewList(5)
	skipList.PutNew(12, 120)
	skipList.PutNew(17, 170)
	skipList.PutNew(25, 250)
	skipList.PutNew(31, 310)
	skipList.PutNew(20, 200)
	skipList.PutNew(39, 390)

	value, _ := skipList.GetByKey(90)
	if value != -1 {
		t.Fatalf("Expected value to be %v, received %v", -1, value)
	}
}
