package stack

import (
	"github.com/SarthakMakhija/data-structures-and-algorithms/stack"
)

func SimplifyPath(path string) string {
	simplifiedPathStack := stack.StringStack{}
	pathPart := ""
	for index := 0; index < len(path); {
		ch := path[index]
		incoming := string(ch)

		if incoming == "/" {
			if simplifiedPathStack.IsEmpty() {
				simplifiedPathStack.Push(incoming)
			} else if pathPart == "." {
				simplifiedPathStack.Pop()
				simplifiedPathStack.Push(incoming)
			} else if pathPart == ".." {
				simplifiedPathStack.Pop()
				simplifiedPathStack.Pop()
				simplifiedPathStack.Push(incoming)
			} else {
				simplifiedPathStack.Push(pathPart)
				part := simplifiedPathStack.Pop()
				separator := simplifiedPathStack.Pop()
				simplifiedPathStack.Push(separator + part)
				simplifiedPathStack.Push(incoming)
			}
			pathPart = ""
			index = index + 1
		} else {
			pathPart = pathPart + incoming
			var readIndex int
			for readIndex = index + 1; readIndex < len(path); readIndex++ {
				ch = path[readIndex]
				incoming = string(ch)
				if incoming == "/" {
					index = readIndex
					break
				}
				pathPart = pathPart + incoming
			}
			index = readIndex
		}
	}

	result := ""
	if pathPart == ".." {
		simplifiedPathStack.Pop()
		simplifiedPathStack.Pop()
	}
	if simplifiedPathStack.Top() == "/" {
		simplifiedPathStack.Pop()
	}
	for !simplifiedPathStack.IsEmpty() {
		result = simplifiedPathStack.Pop() + result
	}
	return result
}
