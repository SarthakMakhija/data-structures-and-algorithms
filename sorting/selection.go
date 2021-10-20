package sorting

type Selection struct {
	Elements []int
}

func (selection *Selection) Sort() {
	if len(selection.Elements) <= 1 {
		return
	}
	elementPosition := 0
	for pass := 1; pass <= len(selection.Elements)-1; pass++ {
		minElementPosition := elementPosition
		for index := elementPosition; index < len(selection.Elements); index++ {
			if selection.Elements[index] < selection.Elements[minElementPosition] {
				minElementPosition = index
			}
		}
		selection.Elements[elementPosition], selection.Elements[minElementPosition] =
			selection.Elements[minElementPosition], selection.Elements[elementPosition]
		elementPosition = elementPosition + 1
	}
}
