package p140

import (
	"strings"
)

func wordBreak(s string, wordDict []string) []string {
	var dfs func(s string, wd []string, mem map[string]([]string)) []string
	dfs = func (s string, wd []string, mem map[string]([]string)) []string {
		if res, ok := mem[s]; ok {
			return res
		}

		if len(s) == 0 {
			return []string { "" }
		}

		res := make([]string, 0)
		for _, word := range wd {
			if strings.HasPrefix(s, word) {
				sublist := dfs(s[len(word):], wd, mem)
				for _, sub := range sublist {
					sep := " "
					if len(sub) == 0 {
						sep = ""
					}
					res = append(res, word + sep + sub)
				}
			}
		}
		mem[s] = res
		return res
	}

	return dfs(s, wordDict, map[string]([]string) {})
}
