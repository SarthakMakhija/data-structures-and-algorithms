package trie

//AllWordsMatching
//str can contain "."
func (t *TrieNode) AllWordsMatching(str string) []string {
	if str == "" {
		return []string{}
	}

	var words []string
	var allWordsMatchingInner func(node *TrieNode, index int, word string)
	allWordsMatchingInner = func(node *TrieNode, index int, word string) {
		if index >= len(str) {
			if node.isEndOfWord == true {
				words = append(words, word)
			}
			return
		}
		character := rune(str[index])

		if character == '.' {
			for ch, trieNode := range node.nodeByCharacter {
				allWordsMatchingInner(trieNode, index+1, word+string(ch))
			}
		} else {
			trieNode, characterExists := node.nodeByCharacter[character]
			if characterExists {
				allWordsMatchingInner(trieNode, index+1, word+string(character))
			} else {
				return
			}
		}
	}
	allWordsMatchingInner(t, 0, "")
	return words
}
