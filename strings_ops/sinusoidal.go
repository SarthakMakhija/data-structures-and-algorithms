package stringsops

import "math"

func SinusoidalFrom(str string) string {
	if str == "" {
		return ""
	}
	up := ""
	mid := ""
	bottom := ""

	upNextIndex := 1
	bottomNextIndex := 3

	for index, ch := range str {
		if math.Mod(float64(index), 2) == 0 {
			mid = mid + string(ch)
		} else {
			if index == upNextIndex {
				up = up + string(ch)
				upNextIndex = upNextIndex + 4
			}
			if index == bottomNextIndex {
				bottom = bottom + string(ch)
				bottomNextIndex = bottomNextIndex + 4
			}
		}
	}

	return up + mid + bottom
}
