package recursion

import (
	"math"
	"reflect"
)

func Search(elements []int, element int) bool {
	if len(elements) == 0 {
		return false
	}
	if elements[0] == element {
		return true
	}
	return Search(elements[1:], element)
}

func Reverse(elements []int) []int {
	var reverseInternal func(int) []int
	var reversedSlice []int

	reverseInternal = func(index int) []int {
		if index < len(elements) {
			reverseInternal(index + 1)
			reversedSlice = append(reversedSlice, elements[index])
		}
		return reversedSlice
	}
	return reverseInternal(0)
}

func Power1(x int, y int) int {
	initial := 1
	var powerInner func(int, int) int

	powerInner = func(output, y int) int {
		if y == 0 {
			return output * 1
		}
		return powerInner(output*x, y-1)
	}
	return powerInner(initial, y)
}

func Power2(x int, y int) int {
	if y == 0 {
		return 1
	}
	return Power2(x, y-1) * x
}

func Factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * Factorial(n-1)
}

func IsPalindrome1(str string) bool {
	if len(str) == 0 {
		return false
	}

	var palindromeInner func(int) bool
	var firstHalf []rune
	var otherHalf string

	mid := len(str) / 2
	isEvenLength := math.Mod(float64(len(str)), 2) == 0

	if isEvenLength {
		otherHalf = str[mid:]
	} else {
		otherHalf = str[mid+1:]
	}

	palindromeInner = func(index int) bool {
		if index < mid {
			palindromeInner(index + 1)
			firstHalf = append(firstHalf, rune(str[index]))
		}
		return reflect.DeepEqual(string(firstHalf), otherHalf)
	}
	return palindromeInner(0)
}

func IsPalindrome2(str string) bool {
	if len(str) == 0 {
		return false
	}

	var palindromeInner func(int) bool
	var isPalindrome bool = false

	length := len(str)
	mid := length / 2

	palindromeInner = func(index int) bool {
		if index < mid {
			if str[index] == str[length-1-index] {
				isPalindrome = true
				return palindromeInner(index + 1)
			} else {
				isPalindrome = false
				return isPalindrome
			}
		}
		return isPalindrome
	}
	return palindromeInner(0)
}

func Permutate(str string) []string {

	var permutateInner func(string) []string

	swap := func(str string, firstIndex, secondIndex int) string {

		characters := []rune(str)
		var copied = make([]rune, len(characters))

		copy(copied, characters)

		copied[firstIndex], copied[secondIndex] = copied[secondIndex], copied[firstIndex]
		return string(copied)
	}

	permutateInner = func(partial string) []string {
		if len(partial) == 1 {
			return []string{partial}
		}
		if len(partial) == 2 {
			swapped := swap(partial, 0, 1)
			return []string{partial, swapped}
		} else {
			var finalPermutations []string
			var swapped string

			for count := 0; count < len(partial); count++ {
				if count == 0 {
					swapped = partial
				} else {
					swapped = swap(partial, 0, count)
				}
				permutations := permutateInner(swapped[1:])
				for p := 0; p < len(permutations); p++ {
					finalPermutations = append(finalPermutations, string(swapped[0])+permutations[p])
				}
			}
			return finalPermutations
		}
	}
	return permutateInner(str)
}

func BinomialCoefficient(n int, k int) int {
	if k == 0 || k == n {
		return 1
	} else {
		return BinomialCoefficient(n-1, k-1) + BinomialCoefficient(n-1, k)
	}
}

func FirstCharacterOccurrence(source string, char rune) int {
	var firstOccurrenceInner func(string, int) int

	firstOccurrenceInner = func(str string, index int) int {
		if len(str) == 0 {
			return -1
		} else if rune(str[0]) == char {
			return index
		} else {
			index++
			return firstOccurrenceInner(str[1:], index)
		}
	}
	return firstOccurrenceInner(source, 0)
}
