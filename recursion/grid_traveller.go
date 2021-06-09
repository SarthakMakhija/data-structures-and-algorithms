package recursion

func PossibleWaysToTravelEndOfGridSized(rows int, columns int) int {

	endRowPosition, endColumnPosition := rows-1, columns-1
	possibleWays := 0

	isRightMoveAllowed := func(columnPosition int) bool {
		return columnPosition < columns-1
	}
	isBottomMoveAllowed := func(rowPosition int) bool {
		return rowPosition < rows-1
	}

	var waysPossibleToTravelGridSizedInner func(int, int)
	waysPossibleToTravelGridSizedInner = func(rowPosition int, columnPosition int) {
		if rowPosition == endRowPosition && columnPosition == endColumnPosition {
			possibleWays = possibleWays + 1
		}
		if isRightMoveAllowed(columnPosition) {
			waysPossibleToTravelGridSizedInner(rowPosition, columnPosition+1)
		}
		if isBottomMoveAllowed(rowPosition) {
			waysPossibleToTravelGridSizedInner(rowPosition+1, columnPosition)
		}
	}
	waysPossibleToTravelGridSizedInner(0, 0)
	return possibleWays
}
