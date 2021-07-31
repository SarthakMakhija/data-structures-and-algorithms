package stringsops

import (
	"strings"
)

func IsPalindromic(str string) bool {
	if str == "" {
		return false
	}

	isAlphanumeric := func(char rune) bool {
		if (char >= 0 && char <= 9) ||
			(char >= 65 && char <= 90) ||
			(char >= 97 && char <= 122) {
			return true
		}
		return false
	}

	lowerCased := strings.ToLower(str)
	for forwardIndex, backwardIndex := 0, len(lowerCased)-1; forwardIndex < len(lowerCased) && backwardIndex >= 0; {
		forwardValue := rune(lowerCased[forwardIndex])
		backwardValue := rune(lowerCased[backwardIndex])

		if !isAlphanumeric(forwardValue) {
			forwardIndex++
			continue
		}
		if !isAlphanumeric(backwardValue) {
			backwardIndex--
			continue
		}
		if forwardValue == backwardValue {
			forwardIndex++
			backwardIndex--
		} else {
			return false
		}
	}
	return true
}
