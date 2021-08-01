package stringsops

import (
	"strconv"
)

func ValidIPAddressesFrom(str string) []string {
	if str == "" || len(str) > 12 {
		return []string{}
	}

	type offsetByPart struct {
		part   string
		offset int
	}
	allValuesFrom := func(m map[string][]offsetByPart) []offsetByPart {
		var values []offsetByPart
		for _, v := range m {
			values = append(values, v...)
		}
		return values
	}
	partNPossibilitiesFrom := func(o offsetByPart) []offsetByPart {
		ipDigitPart := func(startIndex int, endIndex int) *offsetByPart {
			part := offsetByPart{part: str[startIndex:endIndex]}
			part.offset = len(part.part) + o.offset
			numeric, _ := strconv.Atoi(part.part)
			if numeric <= 255 {
				return &part
			}
			return nil
		}

		var ipParts []offsetByPart
		if o.offset <= len(str)-1 {
			part := ipDigitPart(o.offset, o.offset+1)
			if part != nil {
				ipParts = append(ipParts, *part)
			}
		}
		if o.offset+2 <= len(str) {
			part := ipDigitPart(o.offset, o.offset+2)
			if part != nil {
				ipParts = append(ipParts, *part)
			}
		}
		if o.offset+3 <= len(str) {
			part := ipDigitPart(o.offset, o.offset+3)
			if part != nil {
				ipParts = append(ipParts, *part)
			}
		}
		return ipParts
	}
	partNPossibilities := func(offsetByParts []offsetByPart) map[string][]offsetByPart {
		var partByPreviousPartPossibility = make(map[string][]offsetByPart)
		for _, offsetByPart := range offsetByParts {
			partByPreviousPartPossibility[offsetByPart.part] = partNPossibilitiesFrom(offsetByPart)
		}
		return partByPreviousPartPossibility
	}

	part1 := partNPossibilitiesFrom(offsetByPart{offset: 0})
	part2 := partNPossibilities(part1)
	part3 := partNPossibilities(allValuesFrom(part2))
	part4 := partNPossibilities(allValuesFrom(part3))

	var validIps []string
	for _, firstPart := range part1 {
		for _, secondPart := range part2[firstPart.part] {
			for _, thirdPart := range part3[secondPart.part] {
				for _, fourthPart := range part4[thirdPart.part] {
					if len(firstPart.part)+len(secondPart.part)+len(thirdPart.part)+len(fourthPart.part) == len(str) {
						validIps = append(validIps, firstPart.part+"."+secondPart.part+"."+thirdPart.part+"."+fourthPart.part)
					}
				}
			}
		}
	}

	return validIps
}

//192 168

/**
part 1 -> 1, 19, 192 (i=0, i+1, i+2)
part 2 -> (1)[9, 92, 921], (19)[2, 21, 216], (192)[1, 16, 168]
part 3 -> 2, 21, 216, 6, 68, 1, 16, 168, 6, 68, 8,
part 4 -> 1, 16, 168, 6, 8, 68, ...
*/
