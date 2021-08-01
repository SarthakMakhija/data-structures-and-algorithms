package stringsops

import "strconv"

func RunLengthEncode(str string) string {
	if str == "" {
		return ""
	}

	characterConsecutiveOccurrenceCount := 1
	encoded := ""
	for index := 0; index < len(str)-1; index++ {
		if str[index] == str[index+1] {
			characterConsecutiveOccurrenceCount = characterConsecutiveOccurrenceCount + 1
		} else {
			encoded = encoded + strconv.Itoa(characterConsecutiveOccurrenceCount) + string(str[index])
			characterConsecutiveOccurrenceCount = 1
		}
	}
	encoded = encoded + strconv.Itoa(characterConsecutiveOccurrenceCount) + string(str[len(str)-1])
	return encoded
}

//RunLengthDecode
//Assume a valid encoded string
func RunLengthDecode(encoded string) string {
	if encoded == "" {
		return ""
	}

	repeat := func(character uint8, times int) string {
		result := ""
		for count := 1; count <= times; count++ {
			result = result + string(character)
		}
		return result
	}

	decode := func(indexOfRepetitionCount int, indexOfCharacter int) string {
		characterConsecutiveOccurrenceCount, _ := strconv.Atoi(string(encoded[indexOfRepetitionCount]))
		character := encoded[indexOfCharacter]
		return repeat(character, characterConsecutiveOccurrenceCount)
	}

	decoded := ""
	for index := 0; index < len(encoded)-1; index++ {
		decoded = decoded + decode(index, index+1)
	}
	return decoded
}
