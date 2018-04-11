package p214_2

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
	rs := reverseCopy(s)
	s1 := s + "#" + rs
	tlb := make([]int, len(s1))

	for i := 1; i < len(s1); i++ {
		j := tlb[i-1]
		for j > 0 && s1[i] != s1[j] {
			j = tlb[j-1]
		}
		if s1[i] == s1[j] {
			j++
		}
		tlb[i] = j
	}

	return rs[:len(s)-tlb[len(tlb)-1]] + s
}
