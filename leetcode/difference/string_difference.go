package difference

//FindTheDifference
//You are given two strings source and transformed.
//String transformed is generated by random shuffling string source and then add one more letter at a random position.
//Return the letter that was added to transformed.
func FindTheDifference(source string, transformed string) byte {
	characterPresenceIndication := 0
	base := 1

	leftShiftBy := func(number int, by int) int {
		count := 1
		value := number
		for count <= by {
			value = value << 1
			count = count + 1
		}
		return value
	}

	firstPositionOfSetBitIn := func(number int) int {
		position := 0
		for number > 0 {
			number = number >> 1
			position = position + 1
		}
		return position
	}

	for index, transformedRune := range transformed {
		transformedBitIndication := transformedRune - 97
		characterPresenceIndication = characterPresenceIndication ^ leftShiftBy(base, int(transformedBitIndication))

		if index < len(source) {
			sourceBitIndication := source[index] - 97
			characterPresenceIndication = characterPresenceIndication ^ leftShiftBy(base, int(sourceBitIndication))
		}
	}
	return byte(firstPositionOfSetBitIn(characterPresenceIndication) - 1 + 97)
}
