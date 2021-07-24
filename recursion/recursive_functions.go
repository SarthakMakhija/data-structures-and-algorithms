package recursion

import (
	"math"
	"reflect"
	"strconv"
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

func IsPalindrome3(str string) bool {

	length := len(str)
	if length == 0 {
		return false
	}

	var isPalindromeInner func(string, int) bool
	charactersMatchAt := func(s string, firstIndex int, other string, secondIndex int) bool {
		return s[firstIndex] == other[secondIndex]
	}
	isPalindromeInner = func(s string, count int) bool {
		if count == length/2 {
			return charactersMatchAt(s, 0, str, length-count)
		} else {
			result := isPalindromeInner(s[1:], count+1)
			if !result {
				return result
			} else {
				return charactersMatchAt(s, 0, str, length-count)
			}
		}
	}
	return isPalindromeInner(str, 1)
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

func PermutateBruteForce(str string) []string {

	var permutations []string
	var result = make([]byte, len(str))
	var availability = make([]bool, len(str))

	for index := 0; index < len(str); index++ {
		availability[index] = true
	}

	var permutateInner func(int)
	permutateInner = func(level int) {
		if level == len(str) {
			permutations = append(permutations, string(result))
		}
		for index := 0; index < len(str); index++ {
			if availability[index] {
				result[level] = str[index]
				availability[index] = false
				permutateInner(level + 1)
				availability[index] = true
			}
		}
	}
	permutateInner(0)

	return permutations
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

func SumByPartitioning(arr []int) int {
	var sumByPartitioningInner func([]int, []int) int

	sumByPartitioningInner = func(left []int, right []int) int {
		if len(left) == 1 && len(right) == 1 {
			return left[0] + right[0]
		} else if len(left) == 1 && len(right) > 1 {
			return left[0] + sumByPartitioningInner(right[0:len(right)/2], right[len(right)/2:])
		} else if len(right) == 1 && len(left) > 1 {
			return right[0] + sumByPartitioningInner(left[0:len(left)/2], left[len(left)/2:])
		}
		return sumByPartitioningInner(left[0:len(left)/2], left[len(left)/2:]) +
			sumByPartitioningInner(right[0:len(right)/2], right[len(right)/2:])
	}
	return sumByPartitioningInner(arr[0:len(arr)/2], arr[len(arr)/2:])
}

func DecimalToBinary(n int) string {
	var binary = ""
	var decimalInner func(int, int) int

	decimalInner = func(y int, remainder int) int {
		if y == 0 {
			return 1
		} else {
			by2 := y >> 1
			binary = binary + strconv.Itoa(decimalInner(by2, y-(by2*2)))
			return remainder
		}
	}
	decimalInner(n, 0)
	return binary
}
