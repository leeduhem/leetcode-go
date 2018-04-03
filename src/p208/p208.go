package p208

type Trie struct {
	children map[byte]*Trie
	value bool
}


/** Initialize your data structure here. */
func Constructor() Trie {
	trie := new(Trie)
	trie.children = make(map[byte]*Trie)
	return *trie
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
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


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this
	for i := 0; i < len(word); i++ {
		chr := word[i]
		if node1, ok := node.children[chr]; ok {
			node = node1
		} else {
			return false
		}
	}
	return node.value
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for i := 0; i < len(prefix); i++ {
		chr := prefix[i]
		if node1, ok := node.children[chr]; ok {
			node = node1
		} else {
			return false
		}
	}
	return true
}
