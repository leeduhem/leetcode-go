package p009

import "strconv"

func isStringPalindrome(s string) bool {
	l := len(s)
	h := l / 2

	for i,v := range s[0:h] {
		if (v != []rune(s)[l-1-i]) {
			return false
		}
	}

	return true
}

func isPalindrome(x int) bool {
	xs := strconv.Itoa(x)
	if (isStringPalindrome(xs)) {
		return true
	}
	return false
}
