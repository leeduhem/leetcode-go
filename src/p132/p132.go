package p132

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minCut(s string) int {
	n := len(s)
	cc := make([]int, n+1)	// number of cuts for the first k characters
	for i := 0; i <= n; i++ {
		cc[i] = i-1
	}
	for i := 0; i < n; i++ {
		for j := 0; i-j >= 0 && i+j < n && s[i-j] == s[i+j]; j++ {
			// odd length palindromes
			cc[i+j+1] = min(cc[i+j+1], 1+cc[i-j])
		}

		for j := 1; i-j+1 >= 0 && i+j < n && s[i-j+1] == s[i+j]; j++ {
			// even length palindromes
			cc[i+j+1] = min(cc[i+j+1], 1+cc[i-j+1])
		}
	}

	return cc[n]
}
