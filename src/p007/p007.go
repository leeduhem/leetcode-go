// https://leetcode.com/problems/reverse-integer/description/
package p007

import (
	"math"
	"strconv"
	"unicode"
)

func reverseSlice(s []rune) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverse(x int) int {
	xs := []rune(strconv.Itoa(x))

	if unicode.IsDigit(xs[0]) {
		reverseSlice(xs[:])
	} else {
		reverseSlice(xs[1:])
	}

	if rx, err := strconv.Atoi(string(xs)); err == nil {
		if math.MinInt32 <= rx && rx <= math.MaxInt32 {
			return rx
		}
	}
	return 0
}
