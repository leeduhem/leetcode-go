package p014_1

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// Vertical scanning
	for i,c := range strs[0] {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || []rune(strs[j])[i] != c {
				return strs[0][0:i]
			}
		}
	}

	return strs[0]
}
