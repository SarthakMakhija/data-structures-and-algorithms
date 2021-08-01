package stringsops

import (
	"math"
	"strings"
)

func EncodeExcelColumn(column string) int {
	if column == "" {
		return -1
	}
	upperCased := strings.ToUpper(column)
	encoded := 0
	for index := len(upperCased) - 1; index >= 0; index-- {
		raiseTo := (len(upperCased) - 1) - index
		offset := upperCased[index] - 65 + 1
		encoded = encoded + int(math.Pow(float64(26), float64(raiseTo))*float64(offset))
	}
	return encoded
}
