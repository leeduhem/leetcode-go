package p126

import (
	"math"
)

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	min := math.MaxInt32
	queue := []string { beginWord }
	minLadder := map[string]([]string) { }

	ladder := map[string]int { }
	for _, word := range wordList {
		ladder[word] = math.MaxInt32
	}
	ladder[beginWord] = 0

	// BFS: Dijisktra search
	for len(queue) != 0 {
		word := queue[0]
		queue = queue[1:]

		step := ladder[word] + 1
		if step > min {
			break
		}

		for i := 0; i < len(word); i++ {
			chs := ([]byte)(word)
			for ch := 'a'; ch <= 'z'; ch++ {
				chs[i] = byte(ch)
				word1 := string(chs)

				if _, ok := ladder[word1]; ok {
					if step > ladder[word1] {
						continue
					} else if step < ladder[word1] {
						queue = append(queue, word1)
						ladder[word1] = step
					}

					if ml, ok := minLadder[word1]; ok {
						minLadder[word1] = append(ml, word)
					} else {
						minLadder[word1] = []string { word }
					}

					if word1 == endWord {
						min = step
					}
				}
			}
		}
	}

	results := [][]string { }

	var backTrace func(word, start string, list []string)
	backTrace = func(word, start string, list []string) {
		if word == start {
			results = append(results, append([]string{start}, list...))
			return
		}
		list = append([]string{word}, list...)
		if ml, ok := minLadder[word]; ok {
			for _, s := range ml {
				backTrace(s, start, list)
			}
		}
		list = list[1:]
	}

	backTrace(endWord, beginWord, []string {})

	return results
}
