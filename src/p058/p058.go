package p058

import "unicode"

func lengthOfLastWord(s string) int {
	s1 := []rune(s)

	i := len(s1)-1

	// Ignore spaces at end
	for ; i >= 0 && unicode.IsSpace(s1[i]); i-- {}

	start := i
	for ; i >= 0 && ! unicode.IsSpace(s1[i]); i-- {}
	return start - i
}
