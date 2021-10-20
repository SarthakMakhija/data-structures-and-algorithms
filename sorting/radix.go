package sorting

import "math"

type Radix struct {
	Elements []int
}

func (radix *Radix) Sort() {
	if len(radix.Elements) <= 1 {
		return
	}

	totalPasses := totalPassesIn(radix.Elements)
	for pass := 1; pass <= totalPasses; pass++ {
		bins := make([]*node, 10)
		for _, element := range radix.Elements {
			binIndex := (element / int(math.Pow10(pass-1))) % 10
			if bins[binIndex] == nil {
				bins[binIndex] = &node{value: element}
			} else {
				head := bins[binIndex]
				for head.next != nil {
					head = head.next
				}
				head.next = &node{value: element}
			}
		}
		sourceIndex := 0
		for _, head := range bins {
			if head != nil {
				var elements []int
				for head != nil {
					elements = append(elements, head.value)
					head = head.next
				}
				copy(radix.Elements[sourceIndex:], elements)
				sourceIndex = sourceIndex + len(elements)
			}
		}
	}
}

func totalPassesIn(elements []int) int {
	maxDigits := 0
	for _, element := range elements {
		digits := countDigitsIn(element)
		if digits > maxDigits {
			maxDigits = digits
		}
	}
	return maxDigits
}

func countDigitsIn(number int) int {
	if number < 10 {
		return 1
	} else {
		return 1 + countDigitsIn(number/10)
	}
}
