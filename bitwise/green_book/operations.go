package green_book

func BitsToFlipToConvertAToB(aNumber int, targetNumber int) int {
	if aNumber == targetNumber {
		return 0
	}
	result := aNumber ^ targetNumber
	numberOfOnes := 0
	andWithLastNumber := result & (result - 1)

	for andWithLastNumber != 0 {
		numberOfOnes = numberOfOnes + 1
		result = andWithLastNumber
		andWithLastNumber = result & (result - 1)
	}
	return numberOfOnes + 1
}
