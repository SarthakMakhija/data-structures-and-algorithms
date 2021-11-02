package trie

import "fmt"

const endOfStringMarker = "$"

type CompressedTrieNaive struct {
	root *CompressedRoot
}

type CompressedRoot struct {
	nodeByCharacter map[rune]*CompressedNode
}

type CompressedNode struct {
	characters []rune
	next       *CompressedNode
	endOfWord  bool
}

//NewCompressedTrieNaive
//assume str does not contain $
func NewCompressedTrieNaive(str string) *CompressedTrieNaive {
	trie := &CompressedTrieNaive{
		root: &CompressedRoot{
			nodeByCharacter: make(map[rune]*CompressedNode),
		},
	}
	markerTerminated := str + endOfStringMarker
	for phase := 0; phase < len(markerTerminated); phase++ {
		for index := 0; index <= phase; index++ {
			trie.addSuffix(markerTerminated[index : phase+1])
		}
	}
	return trie
}

func (n *CompressedTrieNaive) addSuffix(str string) {
	firstCharacter := rune(str[0])
	node, ok := n.root.nodeByCharacter[firstCharacter]
	if !ok {
		n.root.nodeByCharacter[firstCharacter] = &CompressedNode{characters: append([]rune{}, firstCharacter)}
	} else {
		if string(str[len(str)-1]) == endOfStringMarker {
			if len(str) > len(node.characters) {
				node.endOfWord = true
			} else {
				var originalNodeContent = make([]rune, len(node.characters))
				copy(originalNodeContent, node.characters)
				node.characters, node.endOfWord, node.next = node.characters[:len(str)-1], true, &CompressedNode{
					endOfWord:  true,
					characters: originalNodeContent[len(str)-1:],
					next:       node.next,
				}
			}
		} else if len(str) > len(node.characters) {
			node.characters = append(node.characters, rune(str[len(str)-1]))
		}
	}
}

func (n *CompressedTrieNaive) Contains(pattern string) bool {
	if pattern == "" {
		return false
	}
	node, ok := n.root.nodeByCharacter[rune(pattern[0])]
	if !ok {
		return false
	}
	return node.contains(pattern, 0)
}

func (n *CompressedTrieNaive) ContainsSuffix(suffix string) bool {
	if suffix == "" {
		return false
	}
	node, ok := n.root.nodeByCharacter[rune(suffix[0])]
	if !ok {
		return false
	}
	return node.containsSuffix(suffix, 0)
}

func (n *CompressedTrieNaive) debug() {
	for k, node := range n.root.nodeByCharacter {
		fmt.Println(string(k), " : ", string(node.characters), " : ", node.next.debug())
	}
}

func (n *CompressedNode) contains(pattern string, startAt int) bool {
	node := n
	for node != nil {
		index := startAt
		for nodeCharacterIndex := 0; index < len(pattern) && nodeCharacterIndex < len(node.characters); nodeCharacterIndex, index = nodeCharacterIndex+1, index+1 {
			if node.characters[nodeCharacterIndex] != rune(pattern[index]) {
				return false
			}
		}
		node, startAt = node.next, index
	}
	return true
}

func (n *CompressedNode) containsSuffix(suffix string, startAt int) bool {
	fmt.Println("suffix ", suffix)
	node, previous := n, n
	for node != nil && startAt < len(suffix) {
		index, nodeCharacterIndex := startAt, 0
		for ; index < len(suffix) && nodeCharacterIndex < len(node.characters); nodeCharacterIndex, index = nodeCharacterIndex+1, index+1 {
			if node.characters[nodeCharacterIndex] != rune(suffix[index]) {
				return false
			}
		}
		if nodeCharacterIndex < len(node.characters) {
			return false
		}
		node, previous, startAt = node.next, node, index
	}
	return previous.endOfWord
}

func (n *CompressedNode) debug() string {
	if n == nil {
		return ""
	}
	return string(n.characters) + " - " + n.next.debug()
}
