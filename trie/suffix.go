package trie

type SuffixTrie struct {
	root *SuffixNode
}

type SuffixNode struct {
	nodeByCharacter map[rune]*SuffixNode
	endOfWord       bool
}

func NewSuffixTrie(str string) *SuffixTrie {
	trie := &SuffixTrie{
		root: &SuffixNode{nodeByCharacter: make(map[rune]*SuffixNode)},
	}
	for pass := 1; pass <= len(str); pass++ {
		trie.addSuffix(str, pass-1)
	}
	return trie
}

func (t *SuffixTrie) addSuffix(str string, startingIndex int) {
	currentNode := t.root
	for index := startingIndex; index < len(str); index++ {
		ch := rune(str[index])
		node, ok := currentNode.nodeByCharacter[ch]
		if !ok {
			currentNode.nodeByCharacter[ch] = &SuffixNode{nodeByCharacter: make(map[rune]*SuffixNode)}
			currentNode = currentNode.nodeByCharacter[ch]
		} else {
			currentNode = node
		}
	}
	currentNode.endOfWord = true
}

func (t *SuffixTrie) Contains(pattern string) bool {
	currentNode := t.root
	for index := 0; index < len(pattern); index++ {
		ch := rune(pattern[index])
		node, ok := currentNode.nodeByCharacter[ch]
		if !ok {
			return false
		}
		currentNode = node
	}
	return true
}

func (t *SuffixTrie) ContainsSuffix(suffix string) bool {
	currentNode := t.root
	for index := 0; index < len(suffix); index++ {
		ch := rune(suffix[index])
		node, ok := currentNode.nodeByCharacter[ch]
		if !ok {
			return false
		}
		currentNode = node
	}
	return currentNode.endOfWord == true
}
