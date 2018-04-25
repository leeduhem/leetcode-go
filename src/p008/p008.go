// https://leetcode.com/problems/string-to-integer-atoi/description/
package p008

import (
	"math"
	"unicode"
	"unicode/utf8"
)

func myAtoi(str string) int {
	if len(str) == 0 {
		return 0
	}

	i := 0

	// Ignore leading white spaces
	for j, c := range str {
		if !unicode.IsSpace(c) {
			i = j
			break
		}
	}

	sig := 1
	if str[i] == '+' || str[i] == '-' {
		if str[i] == '-' {
			sig = -1
		}
		i++
	}

	var n uint64
	for {
		r, sz := utf8.DecodeRuneInString(str[i:])
		if (r == utf8.RuneError && (sz == 0 || sz == 1)) ||
			!unicode.IsDigit(r) {
			return sig * int(n)
		}

		n = n*10 + uint64(r-'0')
		if n > math.MaxInt32 {
			if sig < 0 {
				return math.MinInt32
			}
			return math.MaxInt32
		}

		i += sz
	}
}
