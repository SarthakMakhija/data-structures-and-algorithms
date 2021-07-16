package recursion

func TotalsPathsFromUpperLeftmostToLowerRightmostIn(rows int, columns int) int {
	if rows <= 0 || columns <= 0 {
		return 0
	}

	pathCount := 0
	var totalsPathsFromUpperLeftmostToLowerRightmostInner func(int, int)
	totalsPathsFromUpperLeftmostToLowerRightmostInner = func(xPosition int, yPosition int) {
		if xPosition == rows-1 && yPosition == columns-1 {
			pathCount = pathCount + 1
			return
		} else {
			if xPosition <= rows-1 {
				totalsPathsFromUpperLeftmostToLowerRightmostInner(xPosition+1, yPosition)
			}
			if yPosition <= columns-1 {
				totalsPathsFromUpperLeftmostToLowerRightmostInner(xPosition, yPosition+1)
			}
		}
	}
	totalsPathsFromUpperLeftmostToLowerRightmostInner(0, 0)
	return pathCount
}
