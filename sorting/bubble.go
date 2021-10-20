package sorting

type Bubble struct {
	Elements []int
}

func (bubble *Bubble) Sort() {
	if len(bubble.Elements) <= 1 {
		return
	}
	for pass := 0; pass < len(bubble.Elements)-1; pass++ {
		for index := 0; index < len(bubble.Elements)-1-pass; index++ {
			if bubble.Elements[index] > bubble.Elements[index+1] {
				bubble.Elements[index], bubble.Elements[index+1] =
					bubble.Elements[index+1], bubble.Elements[index]
			}
		}
	}
}
