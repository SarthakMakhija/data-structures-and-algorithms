package stringsops

import "strconv"

func LookAndSayValueFor(k int) string {
	if k < 1 {
		return ""
	}
	nextValueGivenPrevious := func(previous string) string {
		currentCharacterConsecutiveOccurrenceCount := 1
		next := ""
		for index := 0; index < len(previous)-1; index++ {
			if previous[index] == previous[index+1] {
				currentCharacterConsecutiveOccurrenceCount = currentCharacterConsecutiveOccurrenceCount + 1
			} else {
				next = next + strconv.Itoa(currentCharacterConsecutiveOccurrenceCount) + string(previous[index])
				currentCharacterConsecutiveOccurrenceCount = 1
			}
		}
		next = next + strconv.Itoa(currentCharacterConsecutiveOccurrenceCount) + string(previous[len(previous)-1])
		return next
	}

	start := "1"
	for iterator := 1; iterator < k; iterator++ {
		start = nextValueGivenPrevious(start)
	}
	return start
}
