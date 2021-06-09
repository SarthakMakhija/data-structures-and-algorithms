package recursion

func PossibleWaysToTravelEndOfGridSized(rows int, columns int) int {

	endRowPosition, endColumnPosition := rows-1, columns-1

	isRightMoveAllowed := func(columnPosition int) bool {
		return columnPosition < columns-1
	}
	isBottomMoveAllowed := func(rowPosition int) bool {
		return rowPosition < rows-1
	}

	var waysPossibleToTravelGridSizedInner func(int, int) int
	waysPossibleToTravelGridSizedInner = func(rowPosition int, columnPosition int) int {
		if rowPosition == endRowPosition && columnPosition == endColumnPosition {
			return 1
		}
		if isRightMoveAllowed(columnPosition) && isBottomMoveAllowed(rowPosition) {
			return waysPossibleToTravelGridSizedInner(rowPosition, columnPosition+1) +
				waysPossibleToTravelGridSizedInner(rowPosition+1, columnPosition)
		} else if isRightMoveAllowed(columnPosition) {
			return waysPossibleToTravelGridSizedInner(rowPosition, columnPosition+1)
		} else {
			return waysPossibleToTravelGridSizedInner(rowPosition+1, columnPosition)
		}
	}
	return waysPossibleToTravelGridSizedInner(0, 0)
}
