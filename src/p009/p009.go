// https://leetcode.com/problems/palindrome-number/description/
package p009

import "strconv"

func isSlicePalindrome(s []rune) bool {
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true
}

func isPalindrome(x int) bool {
	xs := strconv.Itoa(x)
	if isSlicePalindrome([]rune(xs)) {
		return true
	}
	return false
}
