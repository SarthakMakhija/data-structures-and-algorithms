package sorting

type Insertion struct {
	Elements []int
}

func (insertion *Insertion) Sort() {
	if len(insertion.Elements) <= 1 {
		return
	}
	for pass := 1; pass <= len(insertion.Elements)-1; pass++ {
		for index := pass; index > 0; index-- {
			if insertion.Elements[index] < insertion.Elements[index-1] {
				insertion.Elements[index], insertion.Elements[index-1] =
					insertion.Elements[index-1], insertion.Elements[index]
			}
		}
	}
}
