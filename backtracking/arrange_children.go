package backtracking

func AllPossibleWaysToArrangeChildrenInChairs(children []string) []string {
	if len(children) == 0 {
		return []string{}
	} else if len(children) == 1 {
		return []string{children[0]}
	}

	var arrangements []string
	var allPossibleWaysToArrangeChildrenInChairsInner func(arrangement string, parts []string)
	allPossibleWaysToArrangeChildrenInChairsInner = func(arrangement string, parts []string) {
		var copiedParts = make([]string, len(parts))
		copy(copiedParts, parts)

		if len(copiedParts) == 0 {
			arrangements = append(arrangements, arrangement)
			return
		}
		for index := 0; index < len(copiedParts); index++ {
			allPossibleWaysToArrangeChildrenInChairsInner(arrangement+copiedParts[0], copiedParts[1:])
			if index+1 < len(copiedParts) {
				backup := copiedParts[0]
				copiedParts[0] = copiedParts[index+1]
				copiedParts[index+1] = backup
			}
		}
	}

	allPossibleWaysToArrangeChildrenInChairsInner("", children)
	return arrangements
}
