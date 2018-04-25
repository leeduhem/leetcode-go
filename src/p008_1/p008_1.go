// https://leetcode.com/problems/string-to-integer-atoi/description/
package p008_1

import (
	"math"
	"strconv"
	"unicode"
	"unicode/utf8"
)

func myAtoi(str string) int {
	if len(str) == 0 {
		return 0
	}

	i := 0

	// Ignore leading white spaces
	for k, c := range str {
		if !unicode.IsSpace(c) {
			i = k
			break
		}
	}

	j := i
	if str[i] == '+' || str[i] == '-' {
		j++
	}

	for {
		r, sz := utf8.DecodeRuneInString(str[j:])
		if (r == utf8.RuneError && (sz == 0 || sz == 1)) ||
			!unicode.IsDigit(r) {
			n, _ := strconv.Atoi(str[i:j])

			if n > math.MaxInt32 {
				return math.MaxInt32
			}
			if n < math.MinInt32 {
				return math.MinInt32
			}
			return n
		}

		j += sz
	}
}
