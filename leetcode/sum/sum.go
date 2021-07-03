package sum

import (
	"github.com/SarthakMakhija/data-structures-and-algorithms/stack"
)

//Add
//add 2 integers without using +
func Add(augend, addend int) int {
	bitStack := stack.IntStack{}
	addition := 0
	carryBit := 0
	base := 1

	for augend != 0 || addend != 0 || carryBit != 0 {
		resultBit := ((augend & base) ^ (addend & base)) ^ carryBit
		if (augend&base == 1) && (addend&base == 1) ||
			((augend&base == 1) && carryBit == 1) ||
			((addend&base == 1) && carryBit == 1) {
			carryBit = 1
		} else {
			carryBit = 0
		}
		augend = augend >> 1
		addend = addend >> 1
		bitStack.Push(resultBit)
	}

	bitArray := bitStack.All()
	raiseTo := 1
	for index := len(bitArray) - 1; index >= 0; index-- {
		addition = addition + bitArray[index]*(raiseTo)
		raiseTo = raiseTo << 1
	}
	return addition
}
