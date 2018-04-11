package p214

func reverseCopy(s string) string {
	n := len(s)
	if n == 0 {
		return s
	}
	s1 := make([]byte, n)
	for i := n - 1; i >= 0; i-- {
		s1[i] = s[n-1-i]
	}
	return string(s1)
}

func shortestPalindrome(s string) string {
	j := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == s[j] {
			j++
		}
	}
	if j == len(s) {
		return s
	}
	suffix := s[j:]
	return reverseCopy(suffix) + shortestPalindrome(s[:j]) + suffix
}
