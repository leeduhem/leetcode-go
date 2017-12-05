package p125

import (
	"unicode"
)

func isAlphaNumeric(r rune) bool {
	if unicode.IsLetter(r) || unicode.IsDigit(r) {
		return true
	}
	return false
}

func isPalindrome(s string) bool {
	if s == "" {
		return true
	}

	s1 := []rune(s)
	slen := len(s1)

	for i, j := 0, slen-1; i < j; i, j = i+1, j-1 {
		for ; i < slen && !isAlphaNumeric(s1[i]); i++ {}
		for ; j >= 0 && !isAlphaNumeric(s1[j]); j-- {}
		if i >= j { return true }

		if unicode.ToLower(s1[i]) != unicode.ToLower(s1[j]) {
			return false
		}
	}

	return true
}
