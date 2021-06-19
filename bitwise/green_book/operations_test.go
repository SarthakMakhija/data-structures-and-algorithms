package green_book_test

import (
	"fmt"
	"github.com/SarthakMakhija/data-structures-and-algorithms/bitwise/green_book"
	"testing"
)

func TestBitsToFlipToConvertAToB_1(t *testing.T) {
	aNumber := 15
	targetNumber := 29

	result := green_book.BitsToFlipToConvertAToB(aNumber, targetNumber)
	if result != 2 {
		t.Fatalf("Expected 2, found %v", result)
	}
}

func TestBitsToFlipToConvertAToB_2(t *testing.T) {
	aNumber := 8
	targetNumber := 7

	result := green_book.BitsToFlipToConvertAToB(aNumber, targetNumber)
	if result != 4 {
		t.Fatalf("Expected 4, found %v", result)
	}
}

func TestBitsToFlipToConvertAToB_3(t *testing.T) {
	aNumber := 8
	targetNumber := 9

	result := green_book.BitsToFlipToConvertAToB(aNumber, targetNumber)
	if result != 1 {
		t.Fatalf("Expected 1, found %v", result)
	}
}

func TestBitsToFlipToConvertAToB_4(t *testing.T) {
	aNumber := 1
	targetNumber := 0

	result := green_book.BitsToFlipToConvertAToB(aNumber, targetNumber)
	if result != 1 {
		t.Fatalf("Expected 1, found %v", result)
	}
}

func TestBitsToFlipToConvertAToB_5(t *testing.T) {
	aNumber := 1
	targetNumber := 1

	result := green_book.BitsToFlipToConvertAToB(aNumber, targetNumber)
	if result != 0 {
		t.Fatalf("Expected 0, found %v", result)
	}
}

func TestBitsToFlipToConvertAToB_6(t *testing.T) {
	aNumber := 16
	targetNumber := 25

	result := green_book.BitsToFlipToConvertAToB(aNumber, targetNumber)
	if result != 2 {
		t.Fatalf("Expected 2, found %v", result)
	}
}

func TestName(t *testing.T) {
	e1 := []int{1, 2, 3, 4, 5}
	s1 := e1[0:3]

	s1 = append(s1, 50)
	fmt.Println(e1, s1)
	t.Fail()

}
