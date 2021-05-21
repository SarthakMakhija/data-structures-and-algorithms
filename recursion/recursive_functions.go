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
	var reverse_internal func(int) []int
	var reversed_slice []int

	reverse_internal = func(index int) []int {
		if index < len(elements) {
			reverse_internal(index + 1)
			reversed_slice = append(reversed_slice, elements[index])
		}
		return reversed_slice
	}
	return reverse_internal(0)
}

func Power_1(x int, y int) int {
	initial := 1
	var power_inner func(int, int) int

	power_inner = func(output, y int) int {
		if y == 0 {
			return output * 1
		}
		return power_inner(output*x, y-1)
	}
	return power_inner(initial, y)
}

func Power_2(x int, y int) int {
	if y == 0 {
		return 1
	}
	return Power_2(x, y-1) * x
}

func Factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * Factorial(n-1)
}

func Is_Palindrome_1(str string) bool {
	if len(str) == 0 {
		return false
	}

	var palindrome_inner func(int) bool
	var first_half []rune
	var other_half string

	mid := len(str) / 2
	is_even_length := math.Mod(float64(len(str)), 2) == 0

	if is_even_length {
		other_half = string(str[mid:])
	} else {
		other_half = string(str[mid+1:])
	}

	palindrome_inner = func(index int) bool {
		if index < mid {
			palindrome_inner(index + 1)
			first_half = append(first_half, rune(str[index]))
		}
		return reflect.DeepEqual(string(first_half), other_half)
	}
	return palindrome_inner(0)
}

func Is_Palindrome_2(str string) bool {
	if len(str) == 0 {
		return false
	}

	var palindrome_inner func(int) bool
	var is_palindrome bool = false

	length := len(str)
	mid := length / 2

	palindrome_inner = func(index int) bool {
		if index < mid {
			if str[index] == str[length-1-index] {
				is_palindrome = true
				return palindrome_inner(index + 1)
			} else {
				is_palindrome = false
				return is_palindrome
			}
		}
		return is_palindrome
	}
	return palindrome_inner(0)
}

func Permutate(str string) []string {

	var permutate_inner func(string) []string

	swap := func(str string, first_index, second_index int) string {

		characters := []rune(str)
		var copied = make([]rune, len(characters))

		copy(copied, characters)

		copied[first_index], copied[second_index] = copied[second_index], copied[first_index]
		return string(copied)
	}

	permutate_inner = func(partial string) []string {
		if len(partial) == 1 {
			return []string{partial}
		}
		if len(partial) == 2 {
			swapped := swap(partial, 0, 1)
			return []string{partial, swapped}
		} else {
			var final_permutations []string
			var swapped string

			for count := 0; count < len(partial); count++ {
				if count == 0 {
					swapped = partial
				} else {
					swapped = swap(partial, 0, count)
				}
				permutations := permutate_inner(swapped[1:])
				for p := 0; p < len(permutations); p++ {
					final_permutations = append(final_permutations, string(swapped[0])+permutations[p])
				}
			}
			return final_permutations
		}
	}
	return permutate_inner(str)
}
