package p005_1

func intMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func expandAroundCenter(s []rune, left, right int) int {
	L, R := left, right
	for ; L >= 0 && R < len(s) && s[L] == s[R]; L, R = L-1, R+1 {
	}
	return R - L - 1
}

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}

	rs := []rune(s)

	start, end := 0, 0
	for i := range rs {
		len1 := expandAroundCenter(rs, i, i)
		len2 := expandAroundCenter(rs, i, i+1)
		len0 := intMax(len1, len2)
		if len0 > end-start {
			start = i - (len0-1)/2
			end = i + len0/2
		}
	}

	return string(rs[start : end+1])
}
