package array

import "math"

func MaxAmountToBeRobbed(elements []int) int {
	evenSum := 0
	oddSum := 0
	for index := 1; index <= len(elements); index++ {
		if math.Mod(float64(index), float64(2)) != 0 {
			oddSum = oddSum + elements[index-1]
		} else {
			evenSum = evenSum + elements[index-1]
		}
	}
	if evenSum > oddSum {
		return evenSum
	}
	return oddSum
}
