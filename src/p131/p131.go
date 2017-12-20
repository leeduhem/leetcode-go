package p131

func isPalindrome(s string) bool {
	l := len(s)
	h := l / 2

	for i,v := range s[0:h] {
		if (v != []rune(s)[l-1-i]) {
			return false
		}
	}

	return true
}

func partition(s string) [][]string {
	if len(s) == 0 {
		return nil
	}

	parts := [][]string {}

	var dfs func(index int, path []string)
	dfs = func(index int, path []string) {
		if index == len(s) {
			parts = append(parts, path)
			return
		}

		for i := index; i < len(s); i++ {
			if isPalindrome(s[index : i+1]) {
				path1 := make([]string, len(path), len(path)+1)
				copy(path1, path)
				dfs(i+1, append(path1, s[index : i+1]))
			}
		}
	}

	dfs(0, []string {})

	return parts
}
