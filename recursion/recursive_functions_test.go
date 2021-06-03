package recursion_test

import (
	"reflect"
	"testing"

	"github.com/SarthakMakhija/data-structures-and-algorithms/recursion"
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

	var elements []int
	reversed := recursion.Reverse(elements)

	if len(reversed) != 0 {
		t.Fatalf(`Expected length of reversed to be 0 found %v`, len(reversed))
	}
}

func TestReverseSlice(t *testing.T) {

	var elements = []int{1, 2, 3}
	reversed := recursion.Reverse(elements)
	var expected = []int{3, 2, 1}

	if !reflect.DeepEqual(reversed, expected) {
		t.Fatalf(`Expected reversed to be %v but found %v`, expected, reversed)
	}
}

func TestPowerWithYAsZeroUsingImplementation1(t *testing.T) {

	powered := recursion.Power1(2, 0)
	if powered != 1 {
		t.Fatalf(`Expected 1 but found %v`, powered)
	}
}

func TestPowerWithYAs3UsingImplementation1(t *testing.T) {

	powered := recursion.Power1(2, 3)
	if powered != 8 {
		t.Fatalf(`Expected 8 but found %v`, powered)
	}
}

func TestPowerWithYAs10UsingImplementation2(t *testing.T) {

	powered := recursion.Power2(2, 10)
	if powered != 1024 {
		t.Fatalf(`Expected 1024 but found %v`, powered)
	}
}

func TestPowerWithYAsZeroUsingImplementation2(t *testing.T) {

	powered := recursion.Power2(2, 0)
	if powered != 1 {
		t.Fatalf(`Expected 1 but found %v`, powered)
	}
}

func TestPowerWithYAs3UsingImplementation2(t *testing.T) {

	powered := recursion.Power2(2, 3)
	if powered != 8 {
		t.Fatalf(`Expected 8 but found %v`, powered)
	}
}

func TestPowerWithYAs10UsingImplementation1(t *testing.T) {

	powered := recursion.Power1(2, 10)
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

	isPalindrome := recursion.IsPalindrome1("")
	if isPalindrome != false {
		t.Fatalf(`Expected false but found %v`, isPalindrome)
	}
}

func TestIsPalindromeWithOddCharactersUsingImplementation1(t *testing.T) {

	isPalindrome := recursion.IsPalindrome1("MALAYALAM")
	if isPalindrome != true {
		t.Fatalf(`Expected true but found %v`, isPalindrome)
	}
}

func TestIsPalindromeWithEvenCharactersUsingImplementation1(t *testing.T) {

	isPalindrome := recursion.IsPalindrome1("ABBA")
	if isPalindrome != true {
		t.Fatalf(`Expected true but found %v`, isPalindrome)
	}
}

func TestIsNotPalindromeWithEvenCharactersUsingImplementation1(t *testing.T) {

	isPalindrome := recursion.IsPalindrome1("ABCFBA")
	if isPalindrome != false {
		t.Fatalf(`Expected true but found %v`, isPalindrome)
	}
}

func TestIsPalindromeWithEmptyStringUsingImplementation2(t *testing.T) {

	isPalindrome := recursion.IsPalindrome2("")
	if isPalindrome != false {
		t.Fatalf(`Expected false but found %v`, isPalindrome)
	}
}

func TestIsPalindromeWithOddCharactersUsingImplementation2(t *testing.T) {

	isPalindrome := recursion.IsPalindrome2("MALAYALAM")
	if isPalindrome != true {
		t.Fatalf(`Expected true but found %v`, isPalindrome)
	}
}

func TestIsPalindromeWithEvenCharactersUsingImplementation2(t *testing.T) {

	isPalindrome := recursion.IsPalindrome2("ABBA")
	if isPalindrome != true {
		t.Fatalf(`Expected true but found %v`, isPalindrome)
	}
}

func TestIsNotPalindromeWithEvenCharactersUsingImplementation2(t *testing.T) {

	isPalindrome := recursion.IsPalindrome2("ABCFBA")
	if isPalindrome != false {
		t.Fatalf(`Expected true but found %v`, isPalindrome)
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

func TestBinomialCoefficient_1(t *testing.T) {

	output := recursion.BinomialCoefficient(3, 0)
	if output != 1 {
		t.Fatalf(`Expected 3C0 to be 1 but found  %v`, output)
	}
}

func TestBinomialCoefficient_2(t *testing.T) {

	output := recursion.BinomialCoefficient(3, 3)
	if output != 1 {
		t.Fatalf(`Expected 3C3 to be 1 but found  %v`, output)
	}
}

func TestBinomialCoefficient_3(t *testing.T) {

	output := recursion.BinomialCoefficient(5, 2)
	if output != 10 {
		t.Fatalf(`Expected 5C2 to be 10 but found  %v`, output)
	}
}

func TestBinomialCoefficient_4(t *testing.T) {

	output := recursion.BinomialCoefficient(10, 2)
	if output != 45 {
		t.Fatalf(`Expected 10C2 to be 45 but found  %v`, output)
	}
}

func TestFirstCharacterOccurrence_1(t *testing.T) {

	output := recursion.FirstCharacterOccurrence("sample", 'a')
	if output != 1 {
		t.Fatalf(`Expected index of character a to be 1 but found  %v`, output)
	}
}

func TestFirstCharacterOccurrence_2(t *testing.T) {

	output := recursion.FirstCharacterOccurrence("Hello", 'l')
	if output != 2 {
		t.Fatalf(`Expected index of character l to be 2 but found  %v`, output)
	}
}

func TestFirstCharacterOccurrence_3(t *testing.T) {

	output := recursion.FirstCharacterOccurrence("Fascinated", 'd')
	if output != 9 {
		t.Fatalf(`Expected index of character d to be 9 but found  %v`, output)
	}
}

func TestFirstCharacterOccurrence_4(t *testing.T) {

	output := recursion.FirstCharacterOccurrence("dummy", 'e')
	if output != -1 {
		t.Fatalf(`Expected index of character ea to be -1 but found  %v`, output)
	}
}

func TestSumByPartitioning_1(t *testing.T) {

	elements := []int{1, 2, 3, 4, 5, 6, 7, 8}
	sum := recursion.SumByPartitioning(elements)

	if sum != 36 {
		t.Fatalf(`Expected sum to be 36 but was %v`, sum)
	}
}

func TestSumByPartitioning_2(t *testing.T) {

	elements := []int{1, 2, 3, 4, 5}
	sum := recursion.SumByPartitioning(elements)

	if sum != 15 {
		t.Fatalf(`Expected sum to be 15 but was %v`, sum)
	}
}

func TestSumByPartitioning_3(t *testing.T) {

	elements := []int{1, 2, 3, 4, 5, 7}
	sum := recursion.SumByPartitioning(elements)

	if sum != 22 {
		t.Fatalf(`Expected sum to be 22 but was %v`, sum)
	}
}
