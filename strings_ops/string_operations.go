package stringsops

import (
	"errors"
	"fmt"
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
//Assume characters are english alphabets, needs extra 26 bytes
func Duplicates2(str string) []string {
	var lowerCased = strings.ToLower(str)
	var englishCharactersAppearance [26]bool
	var duplicates []string

	for _, ch := range lowerCased {
		index := int(math.Mod(float64(ch), float64(26)))
		if englishCharactersAppearance[index] != true {
			englishCharactersAppearance[index] = true
		} else {
			duplicates = append(duplicates, string(ch))
		}
	}
	return duplicates
}

//Duplicates3
//Assume characters are english alphabets, needs extra 4 bytes for `occurrence`
func Duplicates3(str string) bool {
	var lowerCased = strings.ToLower(str)

	var occurrence int32 = 0
	var base int32 = 1
	var duplicateExists = false

	for _, ch := range lowerCased {
		var nthBitSet int32 = 0
		nthBit := ch - 97

		if nthBit == 0 {
			nthBitSet = base << (nthBit)
		} else {
			nthBitSet = base << (nthBit - 1)
		}

		if occurrence&nthBitSet == 0 {
			occurrence = occurrence | nthBitSet
		} else {
			duplicateExists = true
			return duplicateExists
		}
	}
	return duplicateExists
}

func ToIntegerFrom(str string) (int32, error) {
	length := len(str)
	if length == 0 {
		return -1, errors.New("EMPTY STRING")
	}
	tenRaiseTo := func(n int) int32 {
		var output int32 = 1
		for count := 1; count <= n; count++ {
			output = output * 10
		}
		return output
	}
	var numericValue int32 = 0
	for index, ch := range str {
		if ch >= 48 && ch <= 57 {
			digit := ch - 48
			numericValue = numericValue + digit*tenRaiseTo(length-index-1)
		} else {
			return -1, errors.New(fmt.Sprintf("CAN NOT CONVERT %v TO INTEGER", str))
		}
	}
	return numericValue, nil
}
