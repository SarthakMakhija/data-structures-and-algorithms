package sorting

type Bucket struct {
	Elements []int
}

func (bucket *Bucket) Sort() {
	if len(bucket.Elements) <= 1 {
		return
	}
	bins := make([]*node, maxOf(bucket.Elements)+1)
	for _, element := range bucket.Elements {
		if bins[element] == nil {
			bins[element] = &node{value: element}
		} else {
			head := bins[element]
			for head.next != nil {
				head = head.next
			}
			head.next = &node{value: element}
		}
	}
	sourceIndex := 0
	for _, head := range bins {
		if head != nil {
			var elements []int
			for head != nil {
				elements = append(elements, head.value)
				head = head.next
			}
			copy(bucket.Elements[sourceIndex:], elements)
			sourceIndex = sourceIndex + len(elements)
		}
	}
}

type node struct {
	value int
	next  *node
}
