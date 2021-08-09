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

func (t *TrieNode) AllWords() []string {
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
	allWordsInner(t, "")
	return allWords
}
