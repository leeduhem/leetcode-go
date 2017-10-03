package p008

import (
	"math"
	"unicode"
)

func myAtoi(str string) int {
	var n int64 = 0
	s := false

	if (str == "") {
		return 0
	}

	// Ignore leading spaces
	l := len(str)
	i := 0
	for ; i < l && unicode.IsSpace(rune(str[i])); i++ {
	}

	if str[i] == '+' || str[i] == '-' {
		if (str[i] == '-') {
			s = true
		}

		i++
	}

	for ; i < l && unicode.IsDigit(rune(str[i])); i++ {
		n1 := n * 10 + int64(str[i] - '0')
		// overflow
		if n1 < n {
			if s {
				return math.MinInt32
			} else {
				return math.MaxInt32
			}
		}

		n = n1
	}

	if (s) {
		n = -n
	}

	if (n > math.MaxInt32) {
		return math.MaxInt32
	}
	if (n < math.MinInt32) {
		return math.MinInt32
	}
	return int(n)
}
