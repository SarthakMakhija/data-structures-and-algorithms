package recursion_test

import (
	"reflect"
	"testing"

	"example.com/data-structures-and-algorithms/recursion"
)

func TestSearchElementFound(t *testing.T) {

	elements := []int{1, 2, 3, 4, 5}
	found := recursion.Search(elements, 5)

	if found != true {
		t.Fatalf(`Expected element 5 to be found but was %v`, found)
	}
}

func TestSearchElementNotFound(t *testing.T) {

	elements := []int{1, 2, 3, 4, 5}
	found := recursion.Search(elements, 30)

	if found != false {
		t.Fatalf(`Expected element 30 to be missing but was %v`, found)
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

func TestPowerWithYAsZeroUsingImplementation1(t *testing.T) {

	powered := recursion.Power_1(2, 0)
	if powered != 1 {
		t.Fatalf(`Expected 1 but found %v`, powered)
	}
}

func TestPowerWithYAs3UsingImplementation1(t *testing.T) {

	powered := recursion.Power_1(2, 3)
	if powered != 8 {
		t.Fatalf(`Expected 8 but found %v`, powered)
	}
}

func TestPowerWithYAs10UsingImplementation2(t *testing.T) {

	powered := recursion.Power_2(2, 10)
	if powered != 1024 {
		t.Fatalf(`Expected 1024 but found %v`, powered)
	}
}

func TestPowerWithYAsZeroUsingImplementation2(t *testing.T) {

	powered := recursion.Power_2(2, 0)
	if powered != 1 {
		t.Fatalf(`Expected 1 but found %v`, powered)
	}
}

func TestPowerWithYAs3UsingImplementation2(t *testing.T) {

	powered := recursion.Power_2(2, 3)
	if powered != 8 {
		t.Fatalf(`Expected 8 but found %v`, powered)
	}
}

func TestPowerWithYAs10UsingImplementation1(t *testing.T) {

	powered := recursion.Power_1(2, 10)
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

func TestIsPalindromeWithEmptyStringUsingImplementation1(t *testing.T) {

	is_palindrome := recursion.Is_Palindrome_1("")
	if is_palindrome != false {
		t.Fatalf(`Expected false but found %v`, is_palindrome)
	}
}

func TestIsPalindromeWithOddCharactersUsingImplementation1(t *testing.T) {

	is_palindrome := recursion.Is_Palindrome_1("MALAYALAM")
	if is_palindrome != true {
		t.Fatalf(`Expected true but found %v`, is_palindrome)
	}
}

func TestIsPalindromeWithEvenCharactersUsingImplementation1(t *testing.T) {

	is_palindrome := recursion.Is_Palindrome_1("ABBA")
	if is_palindrome != true {
		t.Fatalf(`Expected true but found %v`, is_palindrome)
	}
}

func TestIsNotPalindromeWithEvenCharactersUsingImplementation1(t *testing.T) {

	is_palindrome := recursion.Is_Palindrome_1("ABCFBA")
	if is_palindrome != false {
		t.Fatalf(`Expected true but found %v`, is_palindrome)
	}
}

func TestIsPalindromeWithEmptyStringUsingImplementation2(t *testing.T) {

	is_palindrome := recursion.Is_Palindrome_2("")
	if is_palindrome != false {
		t.Fatalf(`Expected false but found %v`, is_palindrome)
	}
}

func TestIsPalindromeWithOddCharactersUsingImplementation2(t *testing.T) {

	is_palindrome := recursion.Is_Palindrome_2("MALAYALAM")
	if is_palindrome != true {
		t.Fatalf(`Expected true but found %v`, is_palindrome)
	}
}

func TestIsPalindromeWithEvenCharactersUsingImplementation2(t *testing.T) {

	is_palindrome := recursion.Is_Palindrome_2("ABBA")
	if is_palindrome != true {
		t.Fatalf(`Expected true but found %v`, is_palindrome)
	}
}

func TestIsNotPalindromeWithEvenCharactersUsingImplementation2(t *testing.T) {

	is_palindrome := recursion.Is_Palindrome_2("ABCFBA")
	if is_palindrome != false {
		t.Fatalf(`Expected true but found %v`, is_palindrome)
	}
}

func TestPermutationsOf2CharacterString(t *testing.T) {

	var expected = []string{"12", "21"}
	permutations := recursion.Permutate("12")

	if !reflect.DeepEqual(expected, permutations) {
		t.Fatalf(`Expected %v but found %v`, expected, permutations)
	}
}

func TestPermutationsOf3CharacterString(t *testing.T) {

	var expected = []string{"123", "132", "213", "231", "321", "312"}
	permutations := recursion.Permutate("123")

	if !reflect.DeepEqual(expected, permutations) {
		t.Fatalf(`Expected %v but found %v`, expected, permutations)
	}
}

func TestPermutationsOf4CharacterString(t *testing.T) {

	var expected = []string{"1234", "1243", "1324", "1342", "1432", "1423",
		"2134", "2143", "2314", "2341", "2431", "2413",
		"3214", "3241", "3124", "3142", "3412", "3421",
		"4231", "4213", "4321", "4312", "4132", "4123"}

	permutations := recursion.Permutate("1234")

	if !reflect.DeepEqual(expected, permutations) {
		t.Fatalf(`Expected %v but found %v`, expected, permutations)
	}
}
