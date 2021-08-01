package stringsops

import "strconv"

func LookAndSayValueFor(k int) string {
	if k < 1 {
		return ""
	}
	nextValueGivePrevious := func(previous string) string {
		currentCharacterCount := 1
		next := ""
		for index := 0; index < len(previous)-1; index++ {
			if previous[index] == previous[index+1] {
				currentCharacterCount = currentCharacterCount + 1
			} else {
				next = next + strconv.Itoa(currentCharacterCount) + string(previous[index])
				currentCharacterCount = 1
			}
		}
		next = next + strconv.Itoa(currentCharacterCount) + string(previous[len(previous)-1])
		return next
	}

	start := "1"
	for iterator := 1; iterator < k; iterator++ {
		start = nextValueGivePrevious(start)
	}
	return start
}
