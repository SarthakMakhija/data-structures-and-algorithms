package backtracking

func AllPossibleWaysToArrangeChildrenInChairs(allChildren []string) []string {
	if len(allChildren) == 0 {
		return []string{}
	} else if len(allChildren) == 1 {
		return []string{allChildren[0]}
	}

	var arrangements []string
	var allPossibleWaysToArrangeChildrenInChairsInner func(arrangement string, children []string)
	allPossibleWaysToArrangeChildrenInChairsInner = func(arrangement string, children []string) {
		var copied = make([]string, len(children))
		copy(copied, children)

		if len(copied) == 0 {
			arrangements = append(arrangements, arrangement)
			return
		}
		for index := 0; index < len(copied); index++ {
			allPossibleWaysToArrangeChildrenInChairsInner(arrangement+copied[0], copied[1:])
			if index+1 < len(copied) {
				backup := copied[0]
				copied[0] = copied[index+1]
				copied[index+1] = backup
			}
		}
	}

	allPossibleWaysToArrangeChildrenInChairsInner("", allChildren)
	return arrangements
}
