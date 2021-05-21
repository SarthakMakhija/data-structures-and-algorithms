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
	var occurrences []CharacterOccurrence = make([]CharacterOccurrence, length)
	var duplicates []string

	is_occurrence_slot_available := func(index int) bool {
		return len(occurrences[index].ch) == 0
	}
	add_occurrence := func(ch byte, index int) {
		occurrences[index] = CharacterOccurrence{
			ch:    string(ch),
			count: 1,
		}
	}
	slot_after := func(index int, ch byte) int {
		existing := occurrences[index]
		if existing.ch == string(ch) {
			return index
		}
		for count := 0; count < length; count++ {
			if is_occurrence_slot_available(count) || occurrences[count].ch == string(ch) {
				return count
			}
		}
		return -1
	}

	for index := 0; index < length; index++ {
		ch := lowercase[index]
		occurrence_index := int(math.Mod(float64(ch), float64(length)))

		if is_occurrence_slot_available(occurrence_index) {
			add_occurrence(ch, occurrence_index)
		} else {
			next_slot := slot_after(occurrence_index, ch)
			occurrences[next_slot] = CharacterOccurrence{
				ch:    string(ch),
				count: occurrences[next_slot].count + 1,
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
