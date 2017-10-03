package p005_1

func max(x,y int) int {
	if (x > y) {
		return x
	}
	return y
}

func expandAroundCenter(s string, left, right int) int {
	L,R := left, right
	for ; L >= 0 && R < len(s) && s[L] == s[R]; L,R = L-1,R+1 {
	}
	return R-L-1
}

func longestPalindrome(s string) string {
	if (len(s) == 0) {
		return s
	}

	start, end := 0,0
	for i,_ := range s {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		len0 := max(len1, len2)
		if (len0 > end - start) {
			start = i - (len0 - 1) / 2;
			end = i + len0 / 2;
		}
	}

	return s[start:end+1]
}
