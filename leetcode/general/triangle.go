package general

import (
	"fmt"
	"math"
)

//MinimumTriangleSum
//elements is a triangle shaped array
//Which means number of elements at row[i] = i + 1, i begins from 0
func MinimumTriangleSum(elements [][]int) int {
	totalsRows := len(elements)
	minimumSum := math.MaxInt32

	fmt.Println("totalRows ", totalsRows)

	var minimumTriangleSumInner func(int, int, int)
	minimumTriangleSumInner = func(rowIndex int, elementIndex int, sum int) {
		//fmt.Println(rowIndex)
		//fmt.Println(sum)
		if rowIndex == totalsRows-1 {
			if sum < minimumSum {
				minimumSum = sum
			}
			return
		}
		minimumTriangleSumInner(rowIndex+1, elementIndex, sum+elements[rowIndex+1][elementIndex])
		minimumTriangleSumInner(rowIndex+1, elementIndex+1, sum+elements[rowIndex+1][elementIndex+1])
	}
	minimumTriangleSumInner(0, 0, elements[0][0])
	return minimumSum
}
