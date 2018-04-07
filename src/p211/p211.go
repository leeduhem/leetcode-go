package p211

type WordDictionary struct {
	children map[byte]*WordDictionary
	value    bool
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	wd := new(WordDictionary)
	wd.children = make(map[byte]*WordDictionary)
	return *wd
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	node := this
	i := 0
	for i < len(word) {
		if node1, ok := node.children[word[i]]; ok {
			node = node1
			i++
		} else {
			break
		}
	}

	for i < len(word) {
		chr := word[i]
		node1 := Constructor()
		node.children[chr] = &node1
		node = node.children[chr]
		i++
	}
	node.value = true
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	node := this
	for i := 0; i < len(word); i++ {
		chr := word[i]
		if node1, ok := node.children[chr]; ok {
			node = node1
		} else if chr == '.' {
			for _, wd := range node.children {
				if wd.Search(word[i+1:]) {
					return true
				}
			}
			return false
		} else {
			return false
		}
	}
	return node.value
}
