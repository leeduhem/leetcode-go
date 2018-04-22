// https://leetcode.com/problems/longest-palindromic-substring/description/
package p005

func isPalindrome(s []rune) bool {
	n := len(s) / 2
	for i := 0; i < n; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func longestPalindrome(s string) string {
	rs := []rune(s)
	l := len(rs)
	low, high, max := 0, 0, 0

	for i := range rs {
		for j := i + 1; j <= l; j++ {
			if isPalindrome(rs[i:j]) && (j-i+1 > max) {
				max = j - i + 1
				low, high = i, j
			}
		}
	}

	return string(rs[low:high])
}
