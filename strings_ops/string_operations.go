package stringsops

import (
	"math"
	"strings"
)

func Duplicates(str string) []string {

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
