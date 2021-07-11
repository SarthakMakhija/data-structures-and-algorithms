package general

import "strconv"

func LargestNumberFrom(elements []int) string {
	var orderedElements []int

	//isValueTopGreaterUsingMostSignificantDigit
	//shortcut for now
	isValueTopGreaterUsingMostSignificantDigit := func(value int, otherValue int) bool {
		result, _ := strconv.Atoi(strconv.Itoa(value) + strconv.Itoa(otherValue))
		otherResult, _ := strconv.Atoi(strconv.Itoa(otherValue) + strconv.Itoa(value))
		if result > otherResult {
			return true
		}
		return false
	}
	reOrder := func(e int) {
		orderedElements = append(orderedElements, e)
		for index := len(orderedElements) - 1; index > 0; index-- {
			if !isValueTopGreaterUsingMostSignificantDigit(orderedElements[index], orderedElements[index-1]) {
				orderedElements[index], orderedElements[index-1] = orderedElements[index-1], orderedElements[index]
			}
		}
	}
	for _, e := range elements {
		reOrder(e)
	}
	largestNumberAsString := ""
	for index := len(orderedElements) - 1; index >= 0; index-- {
		largestNumberAsString = largestNumberAsString + strconv.Itoa(orderedElements[index])
	}
	return largestNumberAsString
}
