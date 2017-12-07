package p127

func ladderLength(beginWord string, endWord string, wordList1 []string) int {
	wordList := map[string]bool { }
	for _, word := range wordList1 {
		wordList[word] = true
	}

	ll := 1
	beginSet := map[string]bool { beginWord : true }
	endSet := map[string]bool { endWord : true }
	visited := map[string]bool { }

	for len(beginSet) != 0 && len(endSet) != 0 {
		temp := map[string]bool {}
		for word, _ := range beginSet {
			chs := ([]byte)(word)

			for i := 0; i < len(chs); i++ {
				for c := 'a'; c <= 'z'; c++ {
					old := chs[i]
					chs[i] = byte(c)
					target := string(chs)

					_, ok1 := wordList[target]
					_, ok2 := endSet[target]
					if ok2 && ok1 {
						return ll + 1
					}

					_, ok3 := visited[target]
					if !ok3 && ok1 {
						temp[target] = true
						visited[target] = true
					}

					chs[i] = old
				}
			}
		}

		beginSet = temp
		ll++
	}

	return 0
}
