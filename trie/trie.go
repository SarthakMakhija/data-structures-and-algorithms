package trie

type TrieNode struct {
	nodeByCharacter map[rune]*TrieNode
	isEndOfWord     bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		isEndOfWord:     false,
		nodeByCharacter: make(map[rune]*TrieNode),
	}
}

func (t *TrieNode) Add(word string) {
	if word == "" {
		return
	}
	root := t
	for _, ch := range word {
		trieNode, characterExists := root.nodeByCharacter[ch]
		if !characterExists {
			newNode := NewTrieNode()
			newNode.isEndOfWord = false
			root.nodeByCharacter[ch] = newNode
			root = newNode
		} else {
			root = trieNode
		}
	}
	root.isEndOfWord = true
}

func (t *TrieNode) ExistsCompleteWord(word string) bool {
	if word == "" {
		return false
	}
	root := t
	for _, ch := range word {
		trieNode, characterExists := root.nodeByCharacter[ch]
		if !characterExists {
			return false
		}
		root = trieNode
	}
	if root.isEndOfWord {
		return true
	}
	return false
}

func (t *TrieNode) ExistsPrefix(prefix string) bool {
	if prefix == "" {
		return false
	}
	root := t
	for _, ch := range prefix {
		trieNode, characterExists := root.nodeByCharacter[ch]
		if !characterExists {
			return false
		}
		root = trieNode
	}
	return true
}

func (t *TrieNode) AllWords() []string {
	return t.allWordsPrefixedBy("")
}

func (t *TrieNode) allWordsPrefixedBy(word string) []string {
	var allWords []string

	var allWordsInner func(node *TrieNode, word string)
	allWordsInner = func(node *TrieNode, word string) {
		if node.isEndOfWord {
			allWords = append(allWords, word)
		}
		if !node.isEndOfWord || len(node.nodeByCharacter) != 0 {
			for ch, trieNode := range node.nodeByCharacter {
				allWordsInner(trieNode, word+string(ch))
			}
		}
	}
	allWordsInner(t, word)
	return allWords
}

func (t *TrieNode) AutoCompleteWithPrefix(prefix string) []string {
	node := t
	for _, ch := range prefix {
		trieNode, characterExists := node.nodeByCharacter[ch]
		if !characterExists {
			return []string{}
		}
		node = trieNode
	}
	return node.allWordsPrefixedBy(prefix)
}
