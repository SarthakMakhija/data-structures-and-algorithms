package recursion

func PossibleWaysToTravelEndOfGridSized(rows int, columns int) int {

	type Position struct {
		row int
		col int
	}
	endRowPosition, endColumnPosition := rows-1, columns-1
	waysToTravelByPosition := make(map[Position]int)

	isRightMoveAllowed := func(columnPosition int) bool {
		return columnPosition < columns-1
	}
	isBottomMoveAllowed := func(rowPosition int) bool {
		return rowPosition < rows-1
	}

	var waysPossibleToTravelGridSizedInner func(int, int) int
	waysPossibleToTravelGridSizedInner = func(rowPosition int, columnPosition int) int {
		waysPossible, isPresent := waysToTravelByPosition[Position{row: rowPosition, col: columnPosition}]
		if isPresent {
			return waysPossible
		}
		if rowPosition == endRowPosition && columnPosition == endColumnPosition {
			return 1
		}
		if isRightMoveAllowed(columnPosition) && isBottomMoveAllowed(rowPosition) {
			totalWays := waysPossibleToTravelGridSizedInner(rowPosition, columnPosition+1) +
				waysPossibleToTravelGridSizedInner(rowPosition+1, columnPosition)
			waysToTravelByPosition[Position{row: rowPosition, col: columnPosition}] = totalWays
			return totalWays
		} else if isRightMoveAllowed(columnPosition) {
			return waysPossibleToTravelGridSizedInner(rowPosition, columnPosition+1)
		} else {
			return waysPossibleToTravelGridSizedInner(rowPosition+1, columnPosition)
		}
	}
	totalWays := waysPossibleToTravelGridSizedInner(0, 0)
	return totalWays
}
