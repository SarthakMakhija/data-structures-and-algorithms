package sorting

type Shell struct {
	Elements []int
}

func (shell *Shell) Sort() {
	if len(shell.Elements) <= 1 {
		return
	}
	for gap := len(shell.Elements) / 2; gap >= 1; gap = gap / 2 {
		shell.attemptMovingForwardSwap(gap)
	}
}

func (shell *Shell) attemptMovingForwardSwap(gap int) {
	index, nextIndex := 0, 0
	for ; ; index++ {
		nextIndex = index + gap
		if nextIndex >= len(shell.Elements) {
			break
		}
		if shell.Elements[index] > shell.Elements[nextIndex] {
			shell.Elements[index], shell.Elements[nextIndex] = shell.Elements[nextIndex], shell.Elements[index]
			shell.attemptMovingBackwardSwap(index, gap)
		}
	}
}

func (shell *Shell) attemptMovingBackwardSwap(index, gap int) {
	runningPreviousIndex, previousIndex := index-gap, index
	for runningPreviousIndex >= 0 {
		if shell.Elements[previousIndex] < shell.Elements[runningPreviousIndex] {
			shell.Elements[previousIndex], shell.Elements[runningPreviousIndex] =
				shell.Elements[runningPreviousIndex], shell.Elements[previousIndex]
			runningPreviousIndex, previousIndex = runningPreviousIndex-gap, previousIndex-gap
		} else {
			break
		}
	}
}
