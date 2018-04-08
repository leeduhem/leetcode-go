package p212

type TrieNode struct {
	next [26]*TrieNode
	word string
}

func buildTrie(words []string) *TrieNode {
	root := new(TrieNode)
	for _, word := range words {
		p := root
		for i := 0; i < len(word); i++ {
			d := word[i] - 'a'
			if p.next[d] == nil {
				p.next[d] = new(TrieNode)
			}
			p = p.next[d]
		}
		p.word = word
	}
	return root
}

func dfs(board [][]byte, i, j int, p *TrieNode, res *[]string) {
	c := board[i][j]

	if c == '#' || p.next[c-'a'] == nil {
		return
	}
	p = p.next[c-'a']
	if p.word != "" { // found one
		*res = append(*res, p.word)
		p.word = "" // de-duplicate
	}

	board[i][j] = '#'
	if i > 0 {
		dfs(board, i-1, j, p, res)
	}
	if j > 0 {
		dfs(board, i, j-1, p, res)
	}
	if i < len(board)-1 {
		dfs(board, i+1, j, p, res)
	}
	if j < len(board[0])-1 {
		dfs(board, i, j+1, p, res)
	}
	board[i][j] = c
}

func findWords(board [][]byte, words []string) []string {
	res := []string{}
	root := buildTrie(words)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			dfs(board, i, j, root, &res)
		}
	}
	return res
}
