package recursion_test

import (
	"reflect"
	"testing"

	"example.com/data-structures-and-algorithms/recursion"
)

func TestSearchElementFound(t *testing.T) {

	elements := []int{1, 2, 3, 4, 5}
	index := recursion.Search(elements, 5, 0)

	if index != 4 {
		t.Fatalf(`Expected index to be 4 found %v`, index)
	}
}

func TestSearchElementNotFound(t *testing.T) {

	elements := []int{1, 2, 3, 4, 5}
	index := recursion.Search(elements, 30, 0)

	if index != -1 {
		t.Fatalf(`Expected index to be -1 found %v`, index)
	}
}

func TestReverseEmptySlice(t *testing.T) {

	elements := []int{}
	reversed := recursion.Reverse(elements)

	if len(reversed) != 0 {
		t.Fatalf(`Expected length of reversed to be 0 found %v`, len(reversed))
	}
}

func TestReverseSlice(t *testing.T) {

	elements := []int{1, 2, 3}
	reversed := recursion.Reverse(elements)
	var expected = []int{3, 2, 1}

	if !reflect.DeepEqual(reversed, expected) {
		t.Fatalf(`Expected reversed to be %v but found %v`, expected, reversed)
	}
}

func TestPowerWithYAsZero(t *testing.T) {

	powered := recursion.Power(2, 0)
	if powered != 1 {
		t.Fatalf(`Expected 1 but found %v`, powered)
	}
}

func TestPowerWithYAs3(t *testing.T) {

	powered := recursion.Power(2, 3)
	if powered != 8 {
		t.Fatalf(`Expected 8 but found %v`, powered)
	}
}

func TestPowerWithYAs10(t *testing.T) {

	powered := recursion.Power(2, 10)
	if powered != 1024 {
		t.Fatalf(`Expected 1024 but found %v`, powered)
	}
}

func TestFactorialZero(t *testing.T) {

	factorial := recursion.Factorial(0)
	if factorial != 1 {
		t.Fatalf(`Expected 1 but found %v`, factorial)
	}
}

func TestFactorialNonZero(t *testing.T) {

	factorial := recursion.Factorial(5)
	if factorial != 120 {
		t.Fatalf(`Expected 120 but found %v`, factorial)
	}
}

func TestIsPalindroneWithEmptyString(t *testing.T) {

	is_palindone := recursion.Is_Palindrome_1("")
	if is_palindone != false {
		t.Fatalf(`Expected false but found %v`, is_palindone)
	}
}

func TestIsPalindroneWithOddCharacters(t *testing.T) {

	is_palindone := recursion.Is_Palindrome_1("MALAYALAM")
	if is_palindone != true {
		t.Fatalf(`Expected true but found %v`, is_palindone)
	}
}

func TestIsPalindroneWithEvenCharacters(t *testing.T) {

	is_palindone := recursion.Is_Palindrome_1("ABBA")
	if is_palindone != true {
		t.Fatalf(`Expected true but found %v`, is_palindone)
	}
}
