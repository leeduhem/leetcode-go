package p139

func wordBreak(s string, wordDict []string) bool {
	wd := map[string]bool {}
	for _, w := range wordDict {
		wd[w] = true
	}

	f := make([]bool, len(s) + 1)
	f[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if f[j] {
				if _, ok := wd[s[j:i]]; ok {
					f[i] = true
					break
				}
			}
		}
	}

	return f[len(s)]
}
