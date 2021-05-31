package stringsops

import (
	"math"
	"strings"
)

func Duplicates1(str string) []string {

	length := len(str)
	if length == 0 {
		return nil
	}
	lowercase := strings.ToLower(str)

	type CharacterOccurrence struct {
		ch    string
		count int
	}
	var occurrences = make([]CharacterOccurrence, length)
	var duplicates []string

	isOccurrenceSlotAvailable := func(index int) bool {
		return len(occurrences[index].ch) == 0
	}
	addOccurrence := func(ch byte, index int) {
		occurrences[index] = CharacterOccurrence{
			ch:    string(ch),
			count: 1,
		}
	}
	slotAfter := func(index int, ch byte) int {
		existing := occurrences[index]
		if existing.ch == string(ch) {
			return index
		}
		for count := 0; count < length; count++ {
			if isOccurrenceSlotAvailable(count) || occurrences[count].ch == string(ch) {
				return count
			}
		}
		return -1
	}

	for index := 0; index < length; index++ {
		ch := lowercase[index]
		occurrenceIndex := int(math.Mod(float64(ch), float64(length)))

		if isOccurrenceSlotAvailable(occurrenceIndex) {
			addOccurrence(ch, occurrenceIndex)
		} else {
			nextSlot := slotAfter(occurrenceIndex, ch)
			occurrences[nextSlot] = CharacterOccurrence{
				ch:    string(ch),
				count: occurrences[nextSlot].count + 1,
			}
		}
	}

	for _, occurrence := range occurrences {
		if occurrence.count > 1 {
			duplicates = append(duplicates, occurrence.ch)
		}
	}
	return duplicates
}

//Duplicates2
//Assume characters are english alphabets (Latin)
func Duplicates2(str string) []string {
	var lowerCased = strings.ToLower(str)
	var englishCharacters = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	var duplicates []string

	for _, ch := range lowerCased {
		index := int(math.Mod(float64(ch), float64(26)))
		if englishCharacters[index] != '$' {
			englishCharacters[index] = '$'
		} else {
			duplicates = append(duplicates, string(ch))
		}
	}
	return duplicates
}
