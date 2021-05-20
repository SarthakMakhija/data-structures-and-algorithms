package stringsops

import (
	"fmt"
	"math"
	"strings"
)

func Duplicates(str string) []string {

	length := len(str)
	lowercase := strings.ToLower(str)

	if length == 0 {
		return nil
	}

	type CharacterOccurrence struct {
		ch         string
		occurrence int
	}
	var occurrences []CharacterOccurrence = make([]CharacterOccurrence, length)
	var duplicates []string

	for index := 0; index < length; index++ {
		ch := lowercase[index]
		occurrence_index := int(math.Mod(float64(ch), float64(length)))

		if len(occurrences[occurrence_index].ch) == 0 {
			occurrences[occurrence_index] = CharacterOccurrence{
				ch:         string(ch),
				occurrence: 1,
			}
		} else {
			existing := occurrences[occurrence_index]

			if existing.ch == string(ch) {
				occurrences[occurrence_index].occurrence = occurrences[occurrence_index].occurrence + 1
			} else {
				for index := 0; index < length; index++ {
					if len(occurrences[index].ch) == 0 || occurrences[index].ch == string(ch) {
						occurrences[index] = CharacterOccurrence{
							ch:         string(ch),
							occurrence: occurrences[index].occurrence + 1,
						}
						break
					}
				}
			}
		}
	}

	fmt.Println(occurrences)
	for index := 0; index < length; index++ {
		if occurrences[index].occurrence > 1 {
			duplicates = append(duplicates, occurrences[index].ch)
		}
	}
	return duplicates
}
